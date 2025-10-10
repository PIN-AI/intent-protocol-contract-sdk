package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	intentmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/intentmanager"
	cryptoHelpers "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// SubmitIntentParams 描述 submitIntent 所需参数。
type SubmitIntentParams struct {
	IntentID     [32]byte
	SubnetID     [32]byte
	IntentType   string
	ParamsHash   [32]byte
	Deadline     *big.Int
	PaymentToken common.Address
	Amount       *big.Int
	Value        *big.Int // 可选，ETH 支付时缺省为 Amount
}

// SignedIntent 描述批量提交时的单条 intent 信息与签名。
type SignedIntent struct {
	Data      cryptoHelpers.SignedIntentInput
	Signature []byte
}

// SubmitIntentBatchParams 批量提交参数。
type SubmitIntentBatchParams struct {
	Items         []SignedIntent
	TotalEthValue *big.Int // 可选，未提供时自动统计 PaymentToken==0 的金额
}

// IntentInfo 与合约返回结构一致。
type IntentInfo = intentmanager.DataStructuresIntentInfo

// IntentService 提供 IntentManager 的高层封装。
type IntentService struct {
	backend      bind.ContractBackend
	txManager    *txmgr.Manager
	contract     *intentmanager.IntentManager
	signer       signer.Signer
	chainID      *big.Int
	contractAddr common.Address
}

// NewIntentService 构造 IntentService。
func NewIntentService(backend bind.ContractBackend, txm *txmgr.Manager, contract *intentmanager.IntentManager, sig signer.Signer, chainID *big.Int, contractAddr common.Address) *IntentService {
	return &IntentService{
		backend:      backend,
		txManager:    txm,
		contract:     contract,
		signer:       sig,
		chainID:      new(big.Int).Set(chainID),
		contractAddr: contractAddr,
	}
}

// SubmitIntent 调用合约 submitIntent。
func (s *IntentService) SubmitIntent(ctx context.Context, params SubmitIntentParams) (*types.Transaction, error) {
	if params.Deadline == nil {
		return nil, errors.New("intent: deadline is required")
	}
	if params.Amount == nil {
		return nil, errors.New("intent: amount is required")
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		if params.Value != nil {
			opts.Value = params.Value
		} else if params.PaymentToken == ZeroAddress {
			opts.Value = params.Amount
		}
		return s.contract.SubmitIntent(opts, params.IntentID, params.SubnetID, params.IntentType, params.ParamsHash, params.Deadline, params.PaymentToken, params.Amount)
	})
}

// SubmitIntentsBySignatures 批量提交 Intent。
func (s *IntentService) SubmitIntentsBySignatures(ctx context.Context, params SubmitIntentBatchParams) (*types.Transaction, error) {
	if len(params.Items) == 0 {
		return nil, errors.New("intent: empty batch")
	}
	intents := make([]intentmanager.IIntentManagerIntentData, len(params.Items))
	signatures := make([][]byte, len(params.Items))
	totalEth := new(big.Int)
	for i, item := range params.Items {
		if item.Data.Deadline == nil {
			return nil, errors.New("intent: nil deadline in batch item")
		}
		if item.Data.Amount == nil {
			return nil, errors.New("intent: nil amount in batch item")
		}
		intents[i] = intentmanager.IIntentManagerIntentData{
			IntentId:     item.Data.IntentID,
			SubnetId:     item.Data.SubnetID,
			Requester:    item.Data.Requester,
			IntentType:   item.Data.IntentType,
			ParamsHash:   item.Data.ParamsHash,
			Deadline:     item.Data.Deadline,
			PaymentToken: item.Data.PaymentToken,
			Amount:       item.Data.Amount,
		}
		signatures[i] = item.Signature
		if item.Data.PaymentToken == ZeroAddress {
			totalEth.Add(totalEth, item.Data.Amount)
		}
	}
	if params.TotalEthValue != nil {
		totalEth = new(big.Int).Set(params.TotalEthValue)
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		opts.Value = totalEth
		return s.contract.SubmitIntentsBySignatures(opts, intents, signatures)
	})
}

// FailIntent 调用 failIntent。
func (s *IntentService) FailIntent(ctx context.Context, intentID [32]byte, reason string, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.FailIntent(opts, intentID, reason, validators, signatures)
	})
}

// ProcessExpiredIntent 调用 processExpiredIntent。
func (s *IntentService) ProcessExpiredIntent(ctx context.Context, intentID [32]byte) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ProcessExpiredIntent(opts, intentID)
	})
}

// BatchProcessExpiredIntents 调用 batchProcessExpiredIntents。
func (s *IntentService) BatchProcessExpiredIntents(ctx context.Context, ids [][32]byte) (*types.Transaction, error) {
	if len(ids) == 0 {
		return nil, errors.New("intent: empty ids")
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.BatchProcessExpiredIntents(opts, ids)
	})
}

// GetIntentInfo 读取意图信息。
func (s *IntentService) GetIntentInfo(ctx context.Context, intentID [32]byte) (IntentInfo, error) {
	return s.contract.GetIntentInfo(&bind.CallOpts{Context: ctx}, intentID)
}

// BatchGetIntentInfo 批量读取意图信息。
// 一次性获取多个 Intent 的详细信息，避免多次 RPC 调用。
// 返回数组顺序与输入 ids 顺序一致，不存在的 Intent 返回零值结构体。
func (s *IntentService) BatchGetIntentInfo(ctx context.Context, ids [][32]byte) ([]IntentInfo, error) {
	if len(ids) == 0 {
		return []IntentInfo{}, nil
	}
	return s.contract.BatchGetIntentInfo(&bind.CallOpts{Context: ctx}, ids)
}

// GetIntentsByStatus 根据状态读取 Intent ID 列表。
func (s *IntentService) GetIntentsByStatus(ctx context.Context, status IntentStatus) ([][32]byte, error) {
	return s.contract.GetIntentsByStatus(&bind.CallOpts{Context: ctx}, uint8(status))
}

// GetSubnetIntents 查询某个子网的所有 Intent ID。
func (s *IntentService) GetSubnetIntents(ctx context.Context, subnetID [32]byte) ([][32]byte, error) {
	return s.contract.GetSubnetIntents(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetPendingIntentCount 查询待处理数量。subnetID 为 nil 时查询所有。
func (s *IntentService) GetPendingIntentCount(ctx context.Context, subnetID *[32]byte) (*big.Int, error) {
	var id [32]byte
	if subnetID != nil {
		id = *subnetID
	}
	return s.contract.GetPendingIntentCount(&bind.CallOpts{Context: ctx}, id)
}

// GetTotalIntentCount 查询总 Intent 数。
func (s *IntentService) GetTotalIntentCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetTotalIntentCount(&bind.CallOpts{Context: ctx})
}

// GetRequiredValidatorCount 查询子网阈值。
func (s *IntentService) GetRequiredValidatorCount(ctx context.Context, subnetID [32]byte) (*big.Int, error) {
	return s.contract.GetRequiredValidatorCount(&bind.CallOpts{Context: ctx}, subnetID)
}

// ComputeDigest 计算单条 intent 的 digest，便于批量签名。
func (s *IntentService) ComputeDigest(input cryptoHelpers.SignedIntentInput) ([32]byte, error) {
	return cryptoHelpers.ComputeIntentDigest(input, s.contractAddr, s.chainID)
}

// SignDigest 使用 SDK signer 对 digest 执行 EIP-191 签名。
func (s *IntentService) SignDigest(digest [32]byte) ([]byte, error) {
	return s.signer.SignDigest(digest)
}

// SignIntent 封装 ComputeDigest + SignDigest，一步完成 Intent 签名。
// 简化批量签名流程，适用于 SubmitIntentsBySignatures 场景。
func (s *IntentService) SignIntent(input cryptoHelpers.SignedIntentInput) ([]byte, error) {
	digest, err := s.ComputeDigest(input)
	if err != nil {
		return nil, err
	}
	return s.SignDigest(digest)
}

// ======================== 只读方法：角色与权限 ========================

// DefaultAdminRole 返回默认管理员角色哈希。
func (s *IntentService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GovernanceRole 返回治理角色哈希。
func (s *IntentService) GovernanceRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GOVERNANCEROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin 返回指定角色的管理员角色。
func (s *IntentService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole 检查账户是否拥有指定角色。
func (s *IntentService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== 只读方法：Intent 状态查询 ========================

// IntentExists 检查 Intent 是否存在。
func (s *IntentService) IntentExists(ctx context.Context, intentID [32]byte) (bool, error) {
	return s.contract.IntentExists(&bind.CallOpts{Context: ctx}, intentID)
}

// IsIntentExpired 检查 Intent 是否已过期。
func (s *IntentService) IsIntentExpired(ctx context.Context, intentID [32]byte) (bool, error) {
	return s.contract.IsIntentExpired(&bind.CallOpts{Context: ctx}, intentID)
}

// ======================== 只读方法：全局配置查询 ========================

// DefaultMaxDuration 返回默认最大 Intent 持续时间。
func (s *IntentService) DefaultMaxDuration(ctx context.Context) (*big.Int, error) {
	return s.contract.DEFAULTMAXDURATION(&bind.CallOpts{Context: ctx})
}

// DefaultMinBudget 返回默认最小 Intent 预算。
func (s *IntentService) DefaultMinBudget(ctx context.Context) (*big.Int, error) {
	return s.contract.DEFAULTMINBUDGET(&bind.CallOpts{Context: ctx})
}

// GetMaxIntentDuration 返回当前最大 Intent 持续时间配置。
func (s *IntentService) GetMaxIntentDuration(ctx context.Context) (*big.Int, error) {
	return s.contract.GetMaxIntentDuration(&bind.CallOpts{Context: ctx})
}

// GetMinIntentBudget 返回当前最小 Intent 预算配置。
func (s *IntentService) GetMinIntentBudget(ctx context.Context) (*big.Int, error) {
	return s.contract.GetMinIntentBudget(&bind.CallOpts{Context: ctx})
}

// GetTotalEscrowBalance 返回指定代币的托管总余额。
func (s *IntentService) GetTotalEscrowBalance(ctx context.Context, token common.Address) (*big.Int, error) {
	return s.contract.GetTotalEscrowBalance(&bind.CallOpts{Context: ctx}, token)
}

// Paused 返回 IntentManager 是否处于暂停状态。
func (s *IntentService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface 检查合约是否支持指定接口（ERC165）。
func (s *IntentService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ======================== 写入方法：紧急控制 ========================

// EmergencyPause 紧急暂停 IntentManager（需要 GOVERNANCE_ROLE）。
func (s *IntentService) EmergencyPause(ctx context.Context) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyPause(opts)
	})
}

// EmergencyUnpause 解除 IntentManager 紧急暂停（需要 GOVERNANCE_ROLE）。
func (s *IntentService) EmergencyUnpause(ctx context.Context) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyUnpause(opts)
	})
}

// EmergencyRefundBatch 紧急批量退款（需要 GOVERNANCE_ROLE）。
func (s *IntentService) EmergencyRefundBatch(ctx context.Context, intentIDs [][32]byte, reason string) (*types.Transaction, error) {
	if len(intentIDs) == 0 {
		return nil, errors.New("intent: empty intent IDs")
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyRefundBatch(opts, intentIDs, reason)
	})
}

// ======================== 写入方法：角色管理 ========================

// GrantRole 授予账户指定角色（需要角色管理员权限）。
func (s *IntentService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole 撤销账户的指定角色（需要角色管理员权限）。
func (s *IntentService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole 放弃自己的指定角色。
func (s *IntentService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// ======================== 写入方法：全局配置管理 ========================

// SetMaxIntentDuration 设置最大 Intent 持续时间（需要 GOVERNANCE_ROLE）。
func (s *IntentService) SetMaxIntentDuration(ctx context.Context, maxDuration *big.Int) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetMaxIntentDuration(opts, maxDuration)
	})
}

// SetMinIntentBudget 设置最小 Intent 预算（需要 GOVERNANCE_ROLE）。
func (s *IntentService) SetMinIntentBudget(ctx context.Context, minBudget *big.Int) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetMinIntentBudget(opts, minBudget)
	})
}

// Initialize 初始化 IntentManager 合约（仅在部署后调用一次）。
func (s *IntentService) Initialize(ctx context.Context, admin common.Address, subnetFactory common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, admin, subnetFactory)
	})
}

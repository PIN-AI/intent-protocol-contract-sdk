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

// IntentStatus 枚举，与合约 DataStructures.IntentStatus 对应。
type IntentStatus uint8

const (
	IntentStatusPending  IntentStatus = 0
	IntentStatusComplete IntentStatus = 1
	IntentStatusExpired  IntentStatus = 2
	IntentStatusFailed   IntentStatus = 3
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

// IntentResult 与合约结构体保持一致。
type IntentResult = intentmanager.DataStructuresIntentResult

// FeeDistribution 是 IntentResult 中的费用拆分。
type FeeDistribution = intentmanager.DataStructuresFeeDistribution

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

// CompleteIntent 调用 completeIntent。
func (s *IntentService) CompleteIntent(ctx context.Context, intentID [32]byte, result IntentResult, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.CompleteIntent(opts, intentID, result, validators, signatures)
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

// GetIntentsByStatus 根据状态读取 Intent ID 列表。
func (s *IntentService) GetIntentsByStatus(ctx context.Context, status IntentStatus) ([][32]byte, error) {
	return s.contract.GetIntentsByStatus(&bind.CallOpts{Context: ctx}, uint8(status))
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

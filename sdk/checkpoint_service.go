package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	checkpointmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/checkpointmanager"
	cryptoHelpers "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// CheckpointState alias 合约返回结构。
type CheckpointState = checkpointmanager.DataStructuresCheckpointState

// CheckpointService 提供 CheckpointManager 常用操作（读写）。
type CheckpointService struct {
	contract     *checkpointmanager.CheckpointManager
	backend      bind.ContractBackend
	txMgr        *txmgr.Manager
	signer       signer.Signer
	chainID      *big.Int
	contractAddr common.Address
}

// NewCheckpointService 创建服务。
func NewCheckpointService(
	contract *checkpointmanager.CheckpointManager,
	backend bind.ContractBackend,
	txm *txmgr.Manager,
	sig signer.Signer,
	chainID *big.Int,
	contractAddr common.Address,
) *CheckpointService {
	return &CheckpointService{
		contract:     contract,
		backend:      backend,
		txMgr:        txm,
		signer:       sig,
		chainID:      new(big.Int).Set(chainID),
		contractAddr: contractAddr,
	}
}

// AttachTxManager 允许在运行时更新 TxManager。
func (s *CheckpointService) AttachTxManager(txm *txmgr.Manager) {
	s.txMgr = txm
}

// GetCheckpoint 查询指定子网/epoch 的检查点状态。
func (s *CheckpointService) GetCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (CheckpointState, error) {
	return s.contract.GetCheckpoint(&bind.CallOpts{Context: ctx}, subnetID, epoch)
}

// GetCheckpointProof 返回检查点对应的证明。
func (s *CheckpointService) GetCheckpointProof(ctx context.Context, subnetID [32]byte, epoch uint64) ([]byte, error) {
	return s.contract.GetCheckpointProof(&bind.CallOpts{Context: ctx}, subnetID, epoch)
}

// CheckpointHeader 表示提交检查点所需的头部数据。
type CheckpointHeader struct {
	SubnetID             [32]byte
	Epoch                uint64
	ParentCheckpointHash [32]byte
	Timestamp            *big.Int
	Version              uint32
	ParamsHash           [32]byte
	Roots                cryptoHelpers.CommitmentRoots
	DACommitments        []cryptoHelpers.DACommitment
	EpochSlot            cryptoHelpers.EpochSlot
	Stats                []byte
	AuxHash              [32]byte
	AssignmentsRoot      [32]byte
	ValidationCommitment [32]byte
	PolicyRoot           [32]byte
}

// SubmitCheckpoint 调用合约 submitCheckpoint。
func (s *CheckpointService) SubmitCheckpoint(
	ctx context.Context,
	header CheckpointHeader,
	validators []common.Address,
	signatures [][]byte,
) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	if len(validators) != len(signatures) {
		return nil, errors.New("checkpoint: validators/signatures length mismatch")
	}
	contractHeader := convertCheckpointHeader(header)
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SubmitCheckpoint(opts, contractHeader, validators, signatures)
	})
}

// ComputeCheckpointHash 通过合约计算 checkpoint 哈希。
func (s *CheckpointService) ComputeCheckpointHash(ctx context.Context, header CheckpointHeader) ([32]byte, error) {
	return s.contract.ComputeCheckpointHash(&bind.CallOpts{Context: ctx}, convertCheckpointHeader(header))
}

// ComputeDigest 计算提交所需的签名 digest。
func (s *CheckpointService) ComputeDigest(header CheckpointHeader) ([32]byte, error) {
	return cryptoHelpers.ComputeCheckpointDigest(cryptoHelpers.CheckpointInput{
		SubnetID:             header.SubnetID,
		Epoch:                header.Epoch,
		ParentCheckpointHash: header.ParentCheckpointHash,
		Timestamp:            header.Timestamp,
		Version:              header.Version,
		ParamsHash:           header.ParamsHash,
		Roots:                header.Roots,
		DACommitments:        header.DACommitments,
		EpochSlot:            header.EpochSlot,
		Stats:                header.Stats,
		AuxHash:              header.AuxHash,
		AssignmentsRoot:      header.AssignmentsRoot,
		ValidationCommitment: header.ValidationCommitment,
		PolicyRoot:           header.PolicyRoot,
	}, s.contractAddr, s.chainID)
}

// SignDigest 使用绑定的 signer 对 digest 进行签名。
func (s *CheckpointService) SignDigest(digest [32]byte) ([]byte, error) {
	if s.signer == nil {
		return nil, errors.New("checkpoint: signer not configured")
	}
	return s.signer.SignDigest(digest)
}

func convertCheckpointHeader(header CheckpointHeader) checkpointmanager.DataStructuresCheckpointHeader {
	return checkpointmanager.DataStructuresCheckpointHeader{
		SubnetId:             header.SubnetID,
		Epoch:                header.Epoch,
		ParentCpHash:         header.ParentCheckpointHash,
		Timestamp:            header.Timestamp,
		Version:              header.Version,
		ParamsHash:           header.ParamsHash,
		Roots:                convertCommitmentRoots(header.Roots),
		DaCommitments:        convertDACommitments(header.DACommitments),
		EpochSlot:            convertEpochSlot(header.EpochSlot),
		Stats:                header.Stats,
		AuxHash:              header.AuxHash,
		AssignmentsRoot:      header.AssignmentsRoot,
		ValidationCommitment: header.ValidationCommitment,
		PolicyRoot:           header.PolicyRoot,
	}
}

func convertCommitmentRoots(roots cryptoHelpers.CommitmentRoots) checkpointmanager.DataStructuresCommitmentRoots {
	return checkpointmanager.DataStructuresCommitmentRoots{
		AgentRoot:        roots.AgentRoot,
		AgentServiceRoot: roots.AgentServiceRoot,
		RankRoot:         roots.RankRoot,
		MetricsRoot:      roots.MetricsRoot,
		DataUsageRoot:    roots.DataUsageRoot,
		StateRoot:        roots.StateRoot,
		EventRoot:        roots.EventRoot,
		CrossSubnetRoot:  roots.CrossSubnetRoot,
	}
}

func convertDACommitments(items []cryptoHelpers.DACommitment) []checkpointmanager.DataStructuresDACommitment {
	if len(items) == 0 {
		return nil
	}
	out := make([]checkpointmanager.DataStructuresDACommitment, len(items))
	for i, item := range items {
		out[i] = checkpointmanager.DataStructuresDACommitment{
			Kind:          item.Kind,
			Pointer:       item.Pointer,
			SizeHint:      normalizeBigInt(item.SizeHint),
			SegmentHashes: item.SegmentHashes,
			Expiry:        normalizeBigInt(item.Expiry),
		}
	}
	return out
}

func convertEpochSlot(slot cryptoHelpers.EpochSlot) checkpointmanager.DataStructuresEpochSlot {
	return checkpointmanager.DataStructuresEpochSlot{Epoch: slot.Epoch, Slot: slot.Slot}
}

func normalizeBigInt(v *big.Int) *big.Int {
	if v != nil {
		return v
	}
	return big.NewInt(0)
}

// ======================== 只读方法：角色与权限 ========================

// DefaultAdminRole 返回默认管理员角色哈希。
func (s *CheckpointService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GovernanceRole 返回治理角色哈希。
func (s *CheckpointService) GovernanceRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GOVERNANCEROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin 返回指定角色的管理员角色。
func (s *CheckpointService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole 检查账户是否拥有指定角色。
func (s *CheckpointService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== 只读方法：配置与状态查询 ========================

// DefaultChallengeWindow 返回默认挑战窗口期。
func (s *CheckpointService) DefaultChallengeWindow(ctx context.Context) (*big.Int, error) {
	return s.contract.DEFAULTCHALLENGEWINDOW(&bind.CallOpts{Context: ctx})
}

// GetLatestCheckpoint 返回子网最新的 checkpoint。
func (s *CheckpointService) GetLatestCheckpoint(ctx context.Context, subnetID [32]byte) (CheckpointState, error) {
	return s.contract.GetLatestCheckpoint(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetLatestFinalized 返回子网最新的已最终化 checkpoint。
func (s *CheckpointService) GetLatestFinalized(ctx context.Context, subnetID [32]byte) (CheckpointState, error) {
	return s.contract.GetLatestFinalized(&bind.CallOpts{Context: ctx}, subnetID)
}

// CanFinalizeCheckpoint 检查指定 checkpoint 是否可以最终化。
func (s *CheckpointService) CanFinalizeCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (bool, error) {
	return s.contract.CanFinalizeCheckpoint(&bind.CallOpts{Context: ctx}, subnetID, epoch)
}

// VerifyCheckpoint 验证指定 checkpoint 的完整性与签名。
func (s *CheckpointService) VerifyCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (bool, string, error) {
	result, err := s.contract.VerifyCheckpoint(&bind.CallOpts{Context: ctx}, subnetID, epoch)
	if err != nil {
		return false, "", err
	}
	return result.Valid, result.Reason, nil
}

// VerifyCheckpointContinuity 验证 checkpoint 是否与前一个 checkpoint 连续。
func (s *CheckpointService) VerifyCheckpointContinuity(ctx context.Context, subnetID [32]byte, header CheckpointHeader) (bool, error) {
	return s.contract.VerifyCheckpointContinuity(&bind.CallOpts{Context: ctx}, subnetID, convertCheckpointHeader(header))
}

// Paused 返回 CheckpointManager 是否处于暂停状态。
func (s *CheckpointService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface 检查合约是否支持指定接口（ERC165）。
func (s *CheckpointService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ======================== 写入方法：checkpoint 管理 ========================

// FinalizeCheckpoint 最终化指定 checkpoint（需要 GOVERNANCE_ROLE）。
func (s *CheckpointService) FinalizeCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.FinalizeCheckpoint(opts, subnetID, epoch)
	})
}

// RevertCheckpoint 回滚指定 checkpoint（需要 GOVERNANCE_ROLE）。
func (s *CheckpointService) RevertCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64, reason string) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevertCheckpoint(opts, subnetID, epoch, reason)
	})
}

// ======================== 写入方法：角色管理 ========================

// GrantRole 授予账户指定角色（需要角色管理员权限）。
func (s *CheckpointService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole 撤销账户的指定角色（需要角色管理员权限）。
func (s *CheckpointService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole 放弃自己的指定角色。
func (s *CheckpointService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// SetGovernanceRole 设置治理角色账户（需要 DEFAULT_ADMIN_ROLE）。
func (s *CheckpointService) SetGovernanceRole(ctx context.Context, governance common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetGovernanceRole(opts, governance)
	})
}

// Initialize 初始化 CheckpointManager 合约（仅在部署后调用一次）。
func (s *CheckpointService) Initialize(ctx context.Context, admin common.Address, registry common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, admin, registry)
	})
}

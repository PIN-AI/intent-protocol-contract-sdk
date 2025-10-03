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

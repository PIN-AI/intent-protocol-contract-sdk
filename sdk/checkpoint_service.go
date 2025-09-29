package sdk

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	checkpointmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/checkpointmanager"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// CheckpointState alias 合约返回结构。
type CheckpointState = checkpointmanager.DataStructuresCheckpointState

// CheckpointService 提供 CheckpointManager 常用操作（当前主要为只读）。
type CheckpointService struct {
	contract *checkpointmanager.CheckpointManager
	backend  bind.ContractBackend
	txMgr    *txmgr.Manager
}

// NewCheckpointService 创建服务。
func NewCheckpointService(contract *checkpointmanager.CheckpointManager, backend bind.ContractBackend) *CheckpointService {
	return &CheckpointService{contract: contract, backend: backend}
}

// AttachTxManager 用于未来写操作（如提交检查点）。
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

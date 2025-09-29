package sdk

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	subnetcontract "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnet"
	subnetfactory "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnetfactory"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// SubnetInfoSummary 汇总 SubnetFactory.getSubnetInfo 的返回值。
type SubnetInfoSummary struct {
	Info           subnetfactory.DataStructuresSubnetInfo
	ValidatorCount *big.Int
	AgentCount     *big.Int
	Status         uint8
}

// SubnetFactoryService 封装 SubnetFactory 读写操作。
type SubnetFactoryService struct {
	backend  bind.ContractBackend
	txMgr    *txmgr.Manager
	contract *subnetfactory.SubnetFactory
	signer   signer.Signer
	chainID  *big.Int
}

// NewSubnetFactoryService 构造服务。
func NewSubnetFactoryService(backend bind.ContractBackend, txm *txmgr.Manager, contract *subnetfactory.SubnetFactory, sig signer.Signer, chainID *big.Int) *SubnetFactoryService {
	return &SubnetFactoryService{
		backend:  backend,
		txMgr:    txm,
		contract: contract,
		signer:   sig,
		chainID:  new(big.Int).Set(chainID),
	}
}

// GetSubnetContract 返回子网合约地址。
func (s *SubnetFactoryService) GetSubnetContract(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.SubnetContracts(&bind.CallOpts{Context: ctx}, subnetID)
}

// IsSubnetActive 查询子网是否处于 ACTIVE 状态。
func (s *SubnetFactoryService) IsSubnetActive(ctx context.Context, subnetID [32]byte) (bool, error) {
	return s.contract.IsSubnetActive(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetSubnetInfo 返回子网详细信息与统计。
func (s *SubnetFactoryService) GetSubnetInfo(ctx context.Context, subnetID [32]byte) (SubnetInfoSummary, error) {
	resp, err := s.contract.GetSubnetInfo(&bind.CallOpts{Context: ctx}, subnetID)
	if err != nil {
		return SubnetInfoSummary{}, err
	}
	return SubnetInfoSummary{
		Info:           resp.Info,
		ValidatorCount: resp.ValidatorCount,
		AgentCount:     resp.AgentCount,
		Status:         resp.SubnetStatus,
	}, nil
}

// TotalCreated 返回累计创建的子网数量。
func (s *SubnetFactoryService) TotalCreated(ctx context.Context) (*big.Int, error) {
	return s.contract.TotalCreated(&bind.CallOpts{Context: ctx})
}

// EnumerateSubnets 返回当前所有子网 ID。
func (s *SubnetFactoryService) EnumerateSubnets(ctx context.Context) ([][32]byte, error) {
	total, err := s.contract.TotalCreated(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	count := total.Uint64()
	ids := make([][32]byte, 0, count)
	for i := uint64(0); i < count; i++ {
		id, err := s.contract.AllSubnets(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(i))
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// GetActiveSubnetCount 返回活跃子网数量。
func (s *SubnetFactoryService) GetActiveSubnetCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetActiveSubnetCount(&bind.CallOpts{Context: ctx})
}

// GetCreationStats 返回 totalCreated（兼容旧接口）。
func (s *SubnetFactoryService) GetCreationStats(ctx context.Context) (*big.Int, error) {
	return s.contract.GetCreationStats(&bind.CallOpts{Context: ctx})
}

// PredictSubnetAddress 使用 clone determinism 预测地址。
func (s *SubnetFactoryService) PredictSubnetAddress(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.PredictSubnetAddress(&bind.CallOpts{Context: ctx}, subnetID)
}

// BindSubnet 返回 Subnet 合约实例，便于查询参与者等信息。
func (s *SubnetFactoryService) BindSubnet(subnetAddr common.Address) (*subnetcontract.Subnet, error) {
	return subnetcontract.NewSubnet(subnetAddr, s.backend)
}

// PauseSubnet 调用 pauseSubnet（需要 GUARDIAN_ROLE）。
func (s *SubnetFactoryService) PauseSubnet(ctx context.Context, subnetID [32]byte, reason string) (*types.Transaction, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, errors.New("subnet: pause reason required")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.PauseSubnet(opts, subnetID, reason)
	})
}

// ResumeSubnet 调用 resumeSubnet。
func (s *SubnetFactoryService) ResumeSubnet(ctx context.Context, subnetID [32]byte) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ResumeSubnet(opts, subnetID)
	})
}

// DeprecateSubnet 调用 deprecateSubnet。
func (s *SubnetFactoryService) DeprecateSubnet(ctx context.Context, subnetID [32]byte, reason string) (*types.Transaction, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, errors.New("subnet: deprecate reason required")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.DeprecateSubnet(opts, subnetID, reason)
	})
}

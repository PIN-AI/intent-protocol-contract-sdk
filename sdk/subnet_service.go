package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	subnetcontract "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnet"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// SubnetService 封装 Subnet 合约的常用调用。
type SubnetService struct {
	contract *subnetcontract.Subnet
	txMgr    *txmgr.Manager
}

// NewSubnetService 使用已绑定的 Subnet 合约实例与 TxManager 构造服务。
func NewSubnetService(contract *subnetcontract.Subnet, txm *txmgr.Manager) *SubnetService {
	return &SubnetService{contract: contract, txMgr: txm}
}

// NewSubnetServiceByAddress 通过合约地址绑定 Subnet 并返回服务实例。
func NewSubnetServiceByAddress(backend bind.ContractBackend, addr common.Address, txm *txmgr.Manager) (*SubnetService, error) {
	contract, err := subnetcontract.NewSubnet(addr, backend)
	if err != nil {
		return nil, err
	}
	return &SubnetService{contract: contract, txMgr: txm}, nil
}

// RegisterParticipantParams 描述 RegisterParticipant 调用的可配置项。
type RegisterParticipantParams struct {
	Domain      string
	Endpoint    string
	MetadataURI string
	Value       *big.Int // 可选：注册所需的原生代币数量（如有）
}

func (s *SubnetService) registerParticipant(ctx context.Context, participantType ParticipantType, params RegisterParticipantParams) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		if params.Value != nil {
			opts.Value = params.Value
		}
		return s.contract.RegisterParticipant(opts, uint8(participantType), params.Domain, params.Endpoint, params.MetadataURI)
	})
}

// RegisterValidator 使用 RegisterParticipant 注册验证者。
func (s *SubnetService) RegisterValidator(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantValidator, params)
}

// RegisterAgent 使用 RegisterParticipant 注册代理。
func (s *SubnetService) RegisterAgent(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantAgent, params)
}

// RegisterMatcher 使用 RegisterParticipant 注册匹配器。
func (s *SubnetService) RegisterMatcher(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantMatcher, params)
}

// RegisterParticipantERC20Params 描述 RegisterParticipantERC20 的输入。
type RegisterParticipantERC20Params struct {
	Domain      string
	Endpoint    string
	MetadataURI string
	Amount      *big.Int
}

func (s *SubnetService) registerParticipantERC20(ctx context.Context, participantType ParticipantType, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	if params.Amount == nil {
		return nil, errors.New("subnet: amount is required for ERC20 registration")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RegisterParticipantERC20(opts, uint8(participantType), params.Amount, params.Domain, params.Endpoint, params.MetadataURI)
	})
}

// RegisterValidatorERC20 通过 ERC20 质押注册验证者。
func (s *SubnetService) RegisterValidatorERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantValidator, params)
}

// RegisterAgentERC20 通过 ERC20 质押注册代理。
func (s *SubnetService) RegisterAgentERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantAgent, params)
}

// RegisterMatcherERC20 通过 ERC20 质押注册匹配器。
func (s *SubnetService) RegisterMatcherERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantMatcher, params)
}

// ListActiveParticipants 返回指定类型的活跃参与者列表。
func (s *SubnetService) ListActiveParticipants(ctx context.Context, participantType ParticipantType) ([]subnetcontract.DataStructuresParticipantInfo, error) {
	return s.contract.ListActiveParticipants(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// IsActiveParticipant 查询指定地址在某参与者类型下是否活跃。
func (s *SubnetService) IsActiveParticipant(ctx context.Context, addr common.Address, participantType ParticipantType) (bool, error) {
	return s.contract.IsActiveParticipant(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantInfo 返回参与者详情。
func (s *SubnetService) GetParticipantInfo(ctx context.Context, addr common.Address, participantType ParticipantType) (subnetcontract.DataStructuresParticipantInfo, error) {
	return s.contract.GetParticipantInfo(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantStakeInfo 返回参与者在子网中的质押详情。
func (s *SubnetService) GetParticipantStakeInfo(ctx context.Context, addr common.Address, participantType ParticipantType) (subnetcontract.DataStructuresStakeInfo, error) {
	return s.contract.GetParticipantStakeInfo(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantCount 返回指定类型的参与者数量。
func (s *SubnetService) GetParticipantCount(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.GetParticipantCount(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

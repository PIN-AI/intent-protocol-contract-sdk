package addressbook

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Network represents supported chain network identifiers.
type Network string

const (
	NetworkBase        Network = "base"
	NetworkBaseSepolia Network = "base_sepolia"
	NetworkLocal       Network = "local"
)

// ErrUnsupportedNetwork is returned for unsupported network names or chainId.
var ErrUnsupportedNetwork = errors.New("addressbook: unsupported network")

// Addresses holds a set of core contract addresses.
type Addresses struct {
	IntentManager     common.Address
	SubnetFactory     common.Address
	StakingManager    common.Address
	CheckpointManager common.Address
}

// Merge returns a new address set: prioritizes non-zero addresses from override.
func Merge(base, override Addresses) Addresses {
	result := base
	if override.IntentManager != (common.Address{}) {
		result.IntentManager = override.IntentManager
	}
	if override.SubnetFactory != (common.Address{}) {
		result.SubnetFactory = override.SubnetFactory
	}
	if override.StakingManager != (common.Address{}) {
		result.StakingManager = override.StakingManager
	}
	if override.CheckpointManager != (common.Address{}) {
		result.CheckpointManager = override.CheckpointManager
	}
	return result
}

// Validate checks if all addresses are set (non-zero addresses).
func (a Addresses) Validate() error {
	if a.IntentManager == (common.Address{}) {
		return errors.New("addressbook: intent manager address is zero")
	}
	if a.SubnetFactory == (common.Address{}) {
		return errors.New("addressbook: subnet factory address is zero")
	}
	if a.StakingManager == (common.Address{}) {
		return errors.New("addressbook: staking manager address is zero")
	}
	if a.CheckpointManager == (common.Address{}) {
		return errors.New("addressbook: checkpoint manager address is zero")
	}
	return nil
}

// Resolve returns final contract addresses based on network name and optional override configuration.
// Resolution order: default addresses -> environment variables -> passed overrides.
func Resolve(network Network, overrides *Addresses) (Addresses, error) {
	defaults, ok := defaultAddresses[network]
	if !ok {
		return Addresses{}, fmt.Errorf("addressbook: %w: %s", ErrUnsupportedNetwork, network)
	}
	resolved := defaults
	resolved = Merge(resolved, loadFromEnv(network))
	if overrides != nil {
		resolved = Merge(resolved, *overrides)
	}
	return resolved, resolved.Validate()
}

// NormalizeNetwork converts mixed-case strings to standardized network names.
func NormalizeNetwork(name string) (Network, error) {
	switch Network(strings.ToLower(strings.TrimSpace(name))) {
	case NetworkBase:
		return NetworkBase, nil
	case NetworkBaseSepolia:
		return NetworkBaseSepolia, nil
	case NetworkLocal:
		return NetworkLocal, nil
	default:
		return "", fmt.Errorf("addressbook: %w: %s", ErrUnsupportedNetwork, name)
	}
}

// FromChainID maps network names based on chain ID.
func FromChainID(chainID *big.Int) (Network, error) {
	if chainID == nil {
		return "", errors.New("addressbook: nil chain id")
	}
	switch chainID.Uint64() {
	case 8453:
		return NetworkBase, nil
	case 84532:
		return NetworkBaseSepolia, nil
	case 31337:
		return NetworkLocal, nil
	default:
		return "", fmt.Errorf("addressbook: %w: %s", ErrUnsupportedNetwork, chainID.String())
	}
}

func loadFromEnv(network Network) Addresses {
	vars := envVarMap[network]
	if len(vars) == 0 {
		return Addresses{}
	}
	var out Addresses
	for key, envName := range vars {
		raw := strings.TrimSpace(os.Getenv(envName))
		if raw == "" {
			continue
		}
		addr := common.HexToAddress(raw)
		switch key {
		case "intent_manager":
			out.IntentManager = addr
		case "subnet_factory":
			out.SubnetFactory = addr
		case "staking_manager":
			out.StakingManager = addr
		case "checkpoint_manager":
			out.CheckpointManager = addr
		}
	}
	return out
}

var defaultAddresses = map[Network]Addresses{
	NetworkBase: {
		IntentManager:     common.HexToAddress("0x1111111111111111111111111111111111111111"),
		SubnetFactory:     common.HexToAddress("0x2222222222222222222222222222222222222222"),
		StakingManager:    common.HexToAddress("0x3333333333333333333333333333333333333333"),
		CheckpointManager: common.HexToAddress("0x4444444444444444444444444444444444444444"),
	},
	NetworkBaseSepolia: {
		IntentManager:     common.HexToAddress("0x5555555555555555555555555555555555555555"),
		SubnetFactory:     common.HexToAddress("0x6666666666666666666666666666666666666666"),
		StakingManager:    common.HexToAddress("0x7777777777777777777777777777777777777777"),
		CheckpointManager: common.HexToAddress("0x8888888888888888888888888888888888888888"),
	},
	NetworkLocal: {
		IntentManager:     common.HexToAddress("0x9999999999999999999999999999999999999999"),
		SubnetFactory:     common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
		StakingManager:    common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
		CheckpointManager: common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"),
	},
}

var envVarMap = map[Network]map[string]string{
	NetworkBase: {
		"intent_manager":     "PIN_BASE_INTENT_MANAGER",
		"subnet_factory":     "PIN_BASE_SUBNET_FACTORY",
		"staking_manager":    "PIN_BASE_STAKING_MANAGER",
		"checkpoint_manager": "PIN_BASE_CHECKPOINT_MANAGER",
	},
	NetworkBaseSepolia: {
		"intent_manager":     "PIN_BASE_SEPOLIA_INTENT_MANAGER",
		"subnet_factory":     "PIN_BASE_SEPOLIA_SUBNET_FACTORY",
		"staking_manager":    "PIN_BASE_SEPOLIA_STAKING_MANAGER",
		"checkpoint_manager": "PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER",
	},
	NetworkLocal: {
		"intent_manager":     "PIN_LOCAL_INTENT_MANAGER",
		"subnet_factory":     "PIN_LOCAL_SUBNET_FACTORY",
		"staking_manager":    "PIN_LOCAL_STAKING_MANAGER",
		"checkpoint_manager": "PIN_LOCAL_CHECKPOINT_MANAGER",
	},
}

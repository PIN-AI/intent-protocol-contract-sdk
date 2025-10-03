package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	gethCrypto "github.com/ethereum/go-ethereum/crypto"
)

var (
	checkpointTypeHash = gethCrypto.Keccak256Hash([]byte("PIN_CHECKPOINT_V1(bytes32,uint64,bytes32,uint256,uint32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,address,uint256)"))

	daCommitmentTupleType, _ = abi.NewType("tuple[]", "", []abi.ArgumentMarshaling{
		{Name: "kind", Type: "string"},
		{Name: "pointer", Type: "string"},
		{Name: "size_hint", Type: "uint256"},
		{Name: "segment_hashes", Type: "bytes32[]"},
		{Name: "expiry", Type: "uint256"},
	})
	daCommitmentsArgs = abi.Arguments{{Type: daCommitmentTupleType}}
)

// CommitmentRoots 表示多种承诺的 Merkle 根。
type CommitmentRoots struct {
	AgentRoot        [32]byte
	AgentServiceRoot [32]byte
	RankRoot         [32]byte
	MetricsRoot      [32]byte
	DataUsageRoot    [32]byte
	StateRoot        [32]byte
	EventRoot        [32]byte
	CrossSubnetRoot  [32]byte
}

// DACommitment 描述数据可用性的承诺。
type DACommitment struct {
	Kind          string
	Pointer       string
	SizeHint      *big.Int
	SegmentHashes [][32]byte
	Expiry        *big.Int
}

// EpochSlot 表示检查点对应的纪元/槽位。
type EpochSlot struct {
	Epoch uint64
	Slot  uint64
}

// CheckpointInput 聚合计算摘要所需的全部字段。
type CheckpointInput struct {
	SubnetID             [32]byte
	Epoch                uint64
	ParentCheckpointHash [32]byte
	Timestamp            *big.Int
	Version              uint32
	ParamsHash           [32]byte
	Roots                CommitmentRoots
	DACommitments        []DACommitment
	EpochSlot            EpochSlot
	Stats                []byte
	AuxHash              [32]byte
	AssignmentsRoot      [32]byte
	ValidationCommitment [32]byte
	PolicyRoot           [32]byte
}

// EncodeCommitmentRoots 将 CommitmentRoots 编码为字节数组。
func EncodeCommitmentRoots(roots CommitmentRoots) []byte {
	encoded := make([]byte, 0, 32*8)
	encoded = append(encoded, roots.AgentRoot[:]...)
	encoded = append(encoded, roots.AgentServiceRoot[:]...)
	encoded = append(encoded, roots.RankRoot[:]...)
	encoded = append(encoded, roots.MetricsRoot[:]...)
	encoded = append(encoded, roots.DataUsageRoot[:]...)
	encoded = append(encoded, roots.StateRoot[:]...)
	encoded = append(encoded, roots.EventRoot[:]...)
	encoded = append(encoded, roots.CrossSubnetRoot[:]...)
	return encoded
}

// EncodeDACommitments 将 DACommitment 数组编码为 ABI 字节。
func EncodeDACommitments(commitments []DACommitment) ([]byte, error) {
	if len(commitments) == 0 {
		return daCommitmentsArgs.Pack([]struct {
			Kind          string
			Pointer       string
			SizeHint      *big.Int
			SegmentHashes [][32]byte
			Expiry        *big.Int
		}{})
	}

	converted := make([]struct {
		Kind          string
		Pointer       string
		SizeHint      *big.Int
		SegmentHashes [][32]byte
		Expiry        *big.Int
	}, len(commitments))

	for i, c := range commitments {
		converted[i] = struct {
			Kind          string
			Pointer       string
			SizeHint      *big.Int
			SegmentHashes [][32]byte
			Expiry        *big.Int
		}{
			Kind:          c.Kind,
			Pointer:       c.Pointer,
			SizeHint:      normalizeBigInt(c.SizeHint),
			SegmentHashes: c.SegmentHashes,
			Expiry:        normalizeBigInt(c.Expiry),
		}
	}

	return daCommitmentsArgs.Pack(converted)
}

// EncodeEpochSlot 将 EpochSlot 编码为字节数组。
func EncodeEpochSlot(slot EpochSlot) []byte {
	encoded := make([]byte, 0, 64)
	encoded = append(encoded, common.LeftPadBytes(new(big.Int).SetUint64(slot.Epoch).Bytes(), 32)...)
	encoded = append(encoded, common.LeftPadBytes(new(big.Int).SetUint64(slot.Slot).Bytes(), 32)...)
	return encoded
}

// ComputeCheckpointDigest 计算提交检查点时需要签名的 digest。
func ComputeCheckpointDigest(input CheckpointInput, contract common.Address, chainID *big.Int) ([32]byte, error) {
	var zero [32]byte
	if chainID == nil {
		return zero, errors.New("crypto: nil chain id")
	}
	if input.Timestamp == nil {
		return zero, errors.New("crypto: nil timestamp")
	}

	rootsHash := gethCrypto.Keccak256Hash(EncodeCommitmentRoots(input.Roots))

	daEncoded, err := EncodeDACommitments(input.DACommitments)
	if err != nil {
		return zero, err
	}
	daHash := gethCrypto.Keccak256Hash(daEncoded)

	epochSlotHash := gethCrypto.Keccak256Hash(EncodeEpochSlot(input.EpochSlot))
	statsHash := gethCrypto.Keccak256Hash(input.Stats)

	encoded, err := abi.Arguments{
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeUint64},
		{Type: typeBytes32},
		{Type: typeUint256},
		{Type: typeUint32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeAddress},
		{Type: typeUint256},
	}.Pack(
		checkpointTypeHash,
		input.SubnetID,
		input.Epoch,
		input.ParentCheckpointHash,
		input.Timestamp,
		input.Version,
		input.ParamsHash,
		rootsHash,
		daHash,
		epochSlotHash,
		statsHash,
		input.AuxHash,
		input.AssignmentsRoot,
		input.ValidationCommitment,
		input.PolicyRoot,
		contract,
		chainID,
	)
	if err != nil {
		return zero, err
	}

	return bytesToBytes32(gethCrypto.Keccak256Hash(encoded).Bytes()), nil
}

func normalizeBigInt(v *big.Int) *big.Int {
	if v != nil {
		return v
	}
	return big.NewInt(0)
}

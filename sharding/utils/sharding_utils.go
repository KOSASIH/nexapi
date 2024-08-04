package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ShardIDSize = 32
)

func GenerateShardID(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return hash[:ShardIDSize], nil
}

func ShardIDToString(shardID []byte) string {
	return hex.EncodeToString(shardID)
}

func StringToShardID(shardIDStr string) ([]byte, error) {
	shardID, err := hex.DecodeString(shardIDStr)
	if err!= nil {
		return nil, err
	}
	if len(shardID)!= ShardIDSize {
		return nil, fmt.Errorf("invalid shard ID size")
	}
	return shardID, nil
}

func CalculateShardIndex(shardID []byte, shardCount int) (int, error) {
	if shardCount <= 0 {
		return 0, fmt.Errorf("shard count must be positive")
	}
	shardIndex := new(big.Int)
	shardIndex.SetBytes(shardID)
	shardIndex.Mod(shardIndex, big.NewInt(int64(shardCount)))
	return int(shardIndex.Int64()), nil
}

func SplitDataIntoShards(data []byte, shardSize int) [][]byte {
	shards := make([][]byte, 0)
	for len(data) > shardSize {
		shard := data[:shardSize]
		data = data[shardSize:]
		shards = append(shards, shard)
	}
	if len(data) > 0 {
		shards = append(shards, data)
	}
	return shards
}

func MergeShards(shards [][]byte) []byte {
	data := make([]byte, 0)
	for _, shard := range shards {
		data = append(data, shard...)
	}
	return data
}

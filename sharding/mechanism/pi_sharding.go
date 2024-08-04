package mechanism

import (
	"math"
	"math/big"

	"github.com/consensys/gurvy"
)

const (
	DefaultShardCount = 1024
	DefaultShardSize = 100
)

func ShardPi(ctx context.Context, pi *big.Float, shardCount int, shardSize int) ([]*big.Float, error) {
	if shardCount <= 0 {
		shardCount = DefaultShardCount
	}
	if shardSize <= 0 {
		shardSize = DefaultShardSize
	}
	if shardCount%shardSize != 0 {
		return nil, fmt.Errorf("shardCount must be a multiple of shardSize")
	}
	shardCount = shardCount / shardSize
	shards := make([]*big.Float, shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = new(big.Float)
	}
	cur := new(big.Float)
	cur.SetPrec(64)
	cur.SetString("0")
	for i := 0; i < shardCount; i++ {
		cur.Add(cur, pi)
		cur.Quo(cur, big.NewFloat(float64(shardCount)))
		shards[i].Add(shards[i], cur)
		cur.Mul(cur, big.NewFloat(float64(shardSize)))
	}
	return shards, nil
}

func UnshardPi(ctx context.Context, shards []*big.Float) (*big.Float, error) {
	if len(shards) <= 0 {
		return nil, fmt.Errorf("shards cannot be empty")
	}
	total := new(big.Float)
	for _, shard := range shards {
		total.Add(total, shard)
	}
	return total, nil
}

func VerifyPiShard(ctx context.Context, shard *big.Float, shardIndex int, shardSize int) error {
	if shardIndex < 0 || shardIndex >= len(shards) {
		return fmt.Errorf("shardIndex out of range")
	}
	if shardSize <= 0 {
		return fmt.Errorf("shardSize must be positive")
	}
	if shard.Cmp(new(big.Float).SetInt64(0)) == 0 {
		return fmt.Errorf("shard cannot be zero")
	}
	cur := new(big.Float)
	cur.SetPrec(64)
	cur.SetInt64(int64(shardIndex) * int64(shardSize))
	cur.Quo(cur, big.NewFloat(float64(len(shards))))
	cur.Mul(cur, big.NewFloat(math.Pi))
	if cur.Cmp(shard) != 0 {
		return fmt.Errorf("shard verification failed")
	}
	return nil
}

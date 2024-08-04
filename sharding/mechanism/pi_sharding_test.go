package mechanism

import (
	"math/big"
	"testing"
)

func TestShardPi(t *testing.T) {
	pi := big.NewFloat(math.Pi)
	shards, err := ShardPi(context.Background(), pi, 1024, 100)
	if err!= nil {
		t.Errorf("Expected ShardPi to succeed, but got error: %s", err)
	}
	if len(shards) != 10 {
		t.Errorf("Expected 10 shards, but got %d", len(shards))
	}
}

func TestUnshardPi(t *testing.T) {
	pi := big.NewFloat(math.Pi)
	shards, err := ShardPi(context.Background(), pi, 1024, 100)
	if err!= nil {
		t.Errorf("Expected ShardPi to succeed, but got error: %s", err)
	}
	unshardedPi, err := UnshardPi(context.Background(), shards)
	if err!= nil {
		t.Errorf("Expected UnshardPi to succeed, but got error: %s", err)
	}
	if unshardedPi.Cmp(pi)!= 0 {
		t.Errorf("Expected unsharded Pi to match original Pi")
	}
}

func TestVerifyPiShard(t *testing.T) {
	pi := big.NewFloat(math.Pi)
	shards, err := ShardPi(context.Background(), pi, 1024, 100)
	if err!= nil {
		t.Errorf("Expected ShardPi to succeed, but got error: %s", err)
	}
	for i, shard := range shards {
		err := VerifyPiShard(context.Background(), shard, i, 100)
		if err!= nil {
			t.Errorf("Expected VerifyPiShard to succeed, but got error: %s", err)
		}
	}
}

func BenchmarkShardPi(b *testing.B) {
	pi := big.NewFloat(math.Pi)
	for i := 0; i < b.N; i++ {
		_, err := ShardPi(context.Background(), pi, 1024, 100)
		if err!= nil {
			b.Errorf("Expected ShardPi to succeed, but got error: %s", err)
		}
	}
}

func BenchmarkUnshardPi(b *testing.B) {
	pi := big.NewFloat(math.Pi)
	shards, err := ShardPi(context.Background(), pi, 1024, 100)
	if err!= nil {
		b.Errorf("Expected ShardPi to succeed, but got error: %s", err)
	}
	for i := 0; i < b.N; i++ {
		_, err := UnshardPi(context.Background(), shards)
		if err!= nil {
			b.Errorf("Expected UnshardPi to succeed, but got error: %s", err)
		}
	}
}

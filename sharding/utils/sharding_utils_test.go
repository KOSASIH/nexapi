package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShardID(t *testing.T) {
	data := []byte("hello world")
	shardID, err := GenerateShardID(data)
	assert.NoError(t, err)
	assert.Len(t, shardID, ShardIDSize)
}

func TestShardIDToString(t *testing.T) {
	shardID := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	shardIDStr := ShardIDToString(shardID)
	assert.Equal(t, "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f", shardIDStr)
}

func TestStringToShardID(t *testing.T) {
	shardIDStr := "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
	shardID, err := StringToShardID(shardIDStr)
	assert.NoError(t, err)
	assert.Equal(t, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}, shardID)
}

func TestCalculateShardIndex(t *testing.T) {
	shardID := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	shardCount := 1024
	shardIndex, err := CalculateShardIndex(shardID, shardCount)
	assert.NoError(t, err)
	assert.Equal(t, 123, shardIndex)
}

func TestSplitDataIntoShards(t *testing.T) {
	data := []byte("hello world hello world hello world")
	shardSize := 10
	shards := SplitDataIntoShards(data, shardSize)
	assert.Len(t, shards, 3)
	assert.Equal(t, []byte("hello wor"), shards[0])
	assert.Equal(t, []byte("ld hello w"), shards[1])
	assert.Equal(t, []byte("orld"), shards[2])
}

func TestMergeShards(t *testing.T) {
	shards := [][]byte{
		[]byte("hello wor"),
		[]byte("ld hello w"),
		[]byte("orld"),
	}
	data := MergeShards(shards)
	assert.Equal(t, []byte("hello world hello world"), data)
}

func BenchmarkGenerateShardID(b *testing.B) {
	data := []byte("hello world")
	for i := 0; i < b.N; i++ {
		_, err := GenerateShardID(data)
		if err!= nil {
			b.Errorf("Expected GenerateShardID to succeed, but got error: %s", err)
		}
	}
}

func BenchmarkShardIDToString(b *testing.B) {
	shardID := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	for i := 0; i < b.N; i++ {
		ShardIDToString(shardID)
	}
}

func BenchmarkStringToShardID(b *testing.B) {
	shardIDStr := "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"
	for i := 0; i < b.N; i++ {
		_, err := StringToShardID(shardIDStr)
		if err!= nil {
			b.Errorf("Expected StringToShardID to succeed, but got error: %s", err)
		}
	}
}

func BenchmarkCalculateShardIndex(b *testing.B) {
	shardID := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	shardCount := 1024
	for i := 0; i < b.N; i++ {
		_, err := CalculateShardIndex(shardID, shardCount)
		if err!= nil {
			b.Errorf("Expected CalculateShardIndex to succeed, but got error: %s", err)
		}
	}
}

func BenchmarkSplitDataIntoShards(b *testing.B) {
	data := []byte("hello world hello world hello world")
	shardSize := 10
	for i := 0; i < b.N; i++ {
		SplitDataIntoShards(data, shardSize)
	}
}

func BenchmarkMergeShards(b *testing.B) {
	shards := [][]byte{
		[]byte("hello wor"),
		[]byte("ld hello w"),
		[]byte("orld"),
	}
	for i := 0; i < b.N; i++ {
		MergeShards(shards)
	}
}

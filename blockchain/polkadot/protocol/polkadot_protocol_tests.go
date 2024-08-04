package protocol

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestPolkadotProtocol(t *testing.T) {
	nodeURL := "ws://localhost:9944"
	protocol, err := NewPolkadotProtocol(nodeURL)
	if err!= nil {
		t.Fatal(err)
	}
	defer protocol.client.Close()

	// Test GetBlockHash
	height := uint64(100)
	hash, err := protocol.GetBlockHash(height)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Block hash at height %d: %s\n", height, hash)

	// Test GetBlock
	block, err := protocol.GetBlock(hash)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Block at hash %s: %+v\n", hash, block)

	// Test GetFinalizedHead
	header, err := protocol.GetFinalizedHead()
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Finalized head: %+v\n", header)

	// Test GetBestNumber
	number, err := protocol.GetBestNumber()
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Best block number: %d\n", number)

	// Test GetMetadata
	metadata, err := protocol.GetMetadata()
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Metadata: %+v\n", metadata)

	// Test GetStorage
	key := "0x1234567890abcdef"
	data, err := protocol.GetStorage(key)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Storage value at key %s: %s\n", key, data)

	// Test ExtrinsicSign
	ext := &types.Extrinsic{
		Call: &types.Call{
			Module: "balances",
			Func:   "transfer",
			Args: []types.Arg{
				types.NewArg("amount", types.NewU128(100)),
				types.NewArg("dest", types.NewAccountId("0x1234567890abcdef")),
			},
		},
	}
	signer, err := crypto.GenerateKey()
	if err!= nil {
		t.Fatal(err)
	}
	signedExt, err := protocol.ExtrinsicSign(ext, signer)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Signed extrinsic: %+v\n", signedExt)

	// Test SubmitExtrinsic
	hash, err = protocol.SubmitExtrinsic(signedExt)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Submitted extrinsic with hash %s\n", hash)

	// Test QuerySubscribe
	query := &types.Query{
		Module: "system",
		Func:   "events",
		Args: []types.Arg{
			types.NewArg("phase", types.NewU8(1)),
		},
	}
	subscriber, err := protocol.QuerySubscribe(query)
	if err!= nil {
		t.Fatal(err)
	}
	defer subscriber.Unsubscribe()

	for {
		select {
		case <-subscriber.Done():
			return
		case event := <-subscriber.Chan():
			fmt.Printf("Received event: %+v\n", event)
		}
	}
}

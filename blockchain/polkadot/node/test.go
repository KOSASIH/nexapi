package node

import (
	"testing"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/primitive"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/sr25519"
)

func TestNewPolkadotNode(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()
}

func TestGetBlockHash(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()

	height := uint64(100)
	hash, err := node.GetBlockHash(height)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Block hash at height %d: %s\n", height, hash.Hex())
}

func TestGetBlock(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()

	hash, err := types.HexDecodeString("0x1234567890abcdef")
	if err!= nil {
		t.Fatal(err)
	}
	block, err := node.GetBlock(hash)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Block at hash %s: %+v\n", hash.Hex(), block)
}

func TestGetFinalizedHead(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()

	header, err := node.GetFinalizedHead()
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Finalized head: %+v\n", header)
}

func TestGetBestNumber(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()

	number, err := node.GetBestNumber()
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Best block number: %d\n", number)
}

func TestGetMetadata(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()

	metadata, err := node.GetMetadata()
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Metadata: %+v\n", metadata)
}

func TestGetStorage(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()

	item := types.StorageKey("0x1234567890abcdef")
	data, err := node.GetStorage(item)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Storage value at key %s: %s\n", item.Hex(), data.Hex())
}

func TestExtrinsicSign(t *testing.T) {
	url := "ws://localhost:9944"
	node, err := NewPolkadotNode(url)
	if err!= nil {
		t.Fatal(err)
	}
	defer node.Client.Close()

	ext := extrinsic.NewExtrinsic(primitive.Call{
		Module: "balances",
		Func:   "transfer",
		Args: []primitive.Data{
			primitive.NewU128(100),
			primitive.NewAccountId("0x1234567890abcdef"),
		},
	})
	signer, err := sr25519.NewPairFromSeed("0x1234567890abcdef", nil)
	if err!= nil {
		t.Fatal(err)
	}
	signedExt, err := node.ExtrinsicSign(ext, signer)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Signed extrinsic: %+v\n", signedExt)
}


package node

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestCosmosNode(t *testing.T) {
    // ...

    // Test getting transaction count
    nonce, err := node.GetTransactionCount(address)
    if err!= nil {
        t.Fatal(err)
    }
    fmt.Printf("Transaction count: %d\n", nonce)

    // Test listening for blocks
    time.Sleep(10 * time.Second) // wait for 10 seconds to receive new blocks
    // assert that new blocks were received
    // ...
}

func TestNewCosmosNode(t *testing.T) {
    nodeURL := "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"
    privateKeyHex := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

    _, err := NewCosmosNode(nodeURL, privateKeyHex)
    if err!= nil {
        t.Fatal(err)
    }
}

func TestCosmosNode_SendTransaction(t *testing.T) {
    nodeURL := "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"
    privateKeyHex := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

    node, err := NewCosmosNode(nodeURL, privateKeyHex)
    if err!= nil {
        t.Fatal(err)
    }

    tx := types.NewTransaction(common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"), common.HexToAddress("0x55241586d50469745864804697458648046974586"), big.NewInt(1000000000000000000), 20000, big.NewInt(20000000000))
    err = node.SendTransaction(tx)
    if err!= nil {
        t.Fatal(err)
    }
}

func TestCosmosNode_GetBalance(t *testing.T) {
    nodeURL := "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"
    privateKeyHex := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

    node, err := NewCosmosNode(nodeURL, privateKeyHex)
    if err!= nil {
        t.Fatal(err)
    }

    address := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
    balance, err := node.GetBalance(address)
    if err!= nil {
        t.Fatal(err)
    }
    fmt.Printf("Balance: %s\n", balance)
}

func TestCosmosNode_GetTransactionCount(t *testing.T) {
    nodeURL := "https://mainnet.infura.io/v3/YOUR_PROJECT_ID"
    privateKeyHex := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

    node, err := NewCosmosNode(nodeURL, privateKeyHex)
    if err!= nil {
        t.Fatal(err)
    }

    address := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
    nonce, err := node.GetTransactionCount(address)
    if err!= nil {
        t.Fatal(err)
    }
    fmt.Printf("Transaction count: %d\n", nonce)
}

package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

// CosmosNode represents a Cosmos node
type CosmosNode struct {
	client    *ethclient.Client
	account   accounts.Account
	privateKey *ecdsa.PrivateKey
	nodeURL   string
}

// NewCosmosNode creates a new Cosmos node
func NewCosmosNode(nodeURL string, privateKeyHex string) (*CosmosNode, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err!= nil {
		return nil, err
	}

	account := accounts.Account{
		Address: crypto.PubkeyToAddress(privateKey.PublicKey),
	}

	client, err := ethclient.Dial(nodeURL)
	if err!= nil {
		return nil, err
	}

	return &CosmosNode{
		client:    client,
		account:   account,
		privateKey: privateKey,
		nodeURL:   nodeURL,
	}, nil
}

// Start starts the Cosmos node
func (n *CosmosNode) Start(ctx context.Context) error {
	go n.listenForBlocks(ctx)
	return nil
}

// listenForBlocks listens for new blocks on the Cosmos node
func (n *CosmosNode) listenForBlocks(ctx context.Context) {
	headers := make(chan *types.Header)
	sub, err := n.client.SubscribeNewHead(ctx, headers)
	if err!= nil {
		log.Fatal(err)
	}

	for {
		select {
		case header := <-headers:
			fmt.Printf("Received new block: %s\n", header.Hash().Hex())
			n.processBlock(header)
		case <-ctx.Done():
			sub.Unsubscribe()
			return
		}
	}
}

// processBlock processes a new block on the Cosmos node
func (n *CosmosNode) processBlock(header *types.Header) {
	// Process block logic here
}

// SendTransaction sends a transaction on the Cosmos node
func (n *CosmosNode) SendTransaction(tx *types.Transaction) error {
	return n.client.SendTransaction(context.Background(), tx)
}

// GetBalance gets the balance of an account on the Cosmos node
func (n *CosmosNode) GetBalance(address common.Address) (*big.Int, error) {
	return n.client.BalanceAt(context.Background(), address, nil)
}

// GetTransactionCount gets the transaction count of an account on the Cosmos node
func (n *CosmosNode) GetTransactionCount(address common.Address) (uint64, error) {
	return n.client.PendingNonceAt(context.Background(), address)
}

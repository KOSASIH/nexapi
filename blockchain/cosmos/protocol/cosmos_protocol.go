package protocol

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

// CosmosProtocol represents the Cosmos protocol
type CosmosProtocol struct {
	node    *Node
	client  *ethclient.Client
	account accounts.Account
	privateKey *ecdsa.PrivateKey
}

// NewCosmosProtocol creates a new Cosmos protocol
func NewCosmosProtocol(nodeURL string, privateKeyHex string) (*CosmosProtocol, error) {
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

	node, err := NewNode(nodeURL, privateKeyHex)
	if err!= nil {
		return nil, err
	}

	return &CosmosProtocol{
		node:    node,
		client:  client,
		account: account,
		privateKey: privateKey,
	}, nil
}

// Start starts the Cosmos protocol
func (p *CosmosProtocol) Start(ctx context.Context) error {
	go p.node.listenForBlocks(ctx)
	return nil
}

// SendTransaction sends a transaction on the Cosmos protocol
func (p *CosmosProtocol) SendTransaction(tx *types.Transaction) error {
	return p.client.SendTransaction(context.Background(), tx)
}

// GetBalance gets the balance of an account on the Cosmos protocol
func (p *CosmosProtocol) GetBalance(address common.Address) (*big.Int, error) {
	return p.client.BalanceAt(context.Background(), address, nil)
}

// GetTransactionCount gets the transaction count of an account on the Cosmos protocol
func (p *CosmosProtocol) GetTransactionCount(address common.Address) (uint64, error) {
	return p.client.PendingNonceAt(context.Background(), address)
}

// GetBlockByNumber gets a block by number on the Cosmos protocol
func (p *CosmosProtocol) GetBlockByNumber(number uint64) (*types.Block, error) {
	return p.client.BlockByNumber(context.Background(), big.NewInt(int64(number)))
}

// GetBlockByHash gets a block by hash on the Cosmos protocol
func (p *CosmosProtocol) GetBlockByHash(hash common.Hash) (*types.Block, error) {
	return p.client.BlockByHash(context.Background(), hash)
}

// GetTransactionByHash gets a transaction by hash on the Cosmos protocol
func (p *CosmosProtocol) GetTransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	return p.client.TransactionByHash(context.Background(), hash)
}

// GetTransactionReceipt gets a transaction receipt on the Cosmos protocol
func (p *CosmosProtocol) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	return p.client.TransactionReceipt(context.Background(), hash)
}

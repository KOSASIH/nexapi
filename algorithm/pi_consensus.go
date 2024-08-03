package algorithm

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/pi-network/pi/types"
)

// PiConsensus is the interface for the Pi consensus algorithm
type PiConsensus interface {
	Initialize() error
	Start() error
	Stop() error
	VerifyBlock(*types.Block) (bool, error)
	VerifyTransaction(*types.Transaction) (bool, error)
}

// piConsensus is the implementation of the Pi consensus algorithm
type piConsensus struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    string
	chain      []*types.Block
	transactionPool map[string]*types.Transaction
	mu          sync.RWMutex
}

// NewPiConsensus creates a new instance of the Pi consensus algorithm
func NewPiConsensus(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, address string) PiConsensus {
	return &piConsensus{
		privateKey: privateKey,
		publicKey:  publicKey,
		address:    address,
		chain:      make([]*types.Block, 0),
		transactionPool: make(map[string]*types.Transaction),
	}
}

// Initialize initializes the Pi consensus algorithm
func (pc *piConsensus) Initialize() error {
	// Initialize the chain with the genesis block
	genesisBlock := &types.Block{
		Hash:        "genesis-block",
		PreviousHash: "",
		Transactions: make([]*types.Transaction, 0),
		Timestamp:   time.Now().Unix(),
	}
	pc.chain = append(pc.chain, genesisBlock)
	return nil
}

// Start starts the Pi consensus algorithm
func (pc *piConsensus) Start() error {
	// Start the consensus loop
	go pc.consensusLoop()
	return nil
}

// Stop stops the Pi consensus algorithm
func (pc *piConsensus) Stop() error {
	// Stop the consensus loop
	return nil
}

// VerifyBlock verifies a block
func (pc *piConsensus) VerifyBlock(block *types.Block) (bool, error) {
	// Verify the block hash
	hash, err := pc.calculateBlockHash(block)
	if err != nil {
		return false, err
	}
	if hash != block.Hash {
		return false, fmt.Errorf("invalid block hash")
	}

	// Verify the block transactions
	for _, transaction := range block.Transactions {
		valid, err := pc.VerifyTransaction(transaction)
		if err != nil {
			return false, err
		}
		if !valid {
			return false, fmt.Errorf("invalid transaction")
		}
	}

	return true, nil
}

// VerifyTransaction verifies a transaction
func (pc *piConsensus) VerifyTransaction(transaction *types.Transaction) (bool, error) {
	// Verify the transaction hash
	hash, err := pc.calculateTransactionHash(transaction)
	if err != nil {
		return false, err
	}
	if hash != transaction.Hash {
		return false, fmt.Errorf("invalid transaction hash")
	}

	// Verify the transaction signature
	signatureValid, err := pc.verifyTransactionSignature(transaction)
	if err != nil {
		return false, err
	}
	if !signatureValid {
		return false, fmt.Errorf("invalid transaction signature")
	}

	return true, nil
}

func (pc *piConsensus) consensusLoop() {
	for {
		// Select a random transaction from the transaction pool
		transaction := pc.selectRandomTransaction()
		if transaction == nil {
			continue
		}

		// Create a new block with the transaction
		block := pc.createBlock(transaction)

		// Broadcast the block to the network
		pc.broadcastBlock(block)

		// Add the block to the chain
		pc.addBlockToChain(block)
	}
}

func (pc *piConsensus) calculateBlockHash(block *types.Block) (string, error) {
	hash := sha256.Sum256([]byte(block.Hash + block.PreviousHash + block.Timestamp))
	return hex.EncodeToString(hash[:]), nil
}

func (pc *piConsensus) calculateTransactionHash(transaction *types.Transaction) (string, error) {
	hash := sha256.Sum256([]byte(transaction.ID + transaction.From + transaction.To + fmt.Sprintf("%d", transaction.Amount)))
	return hex.EncodeToString(hash[:]), nil
}

func (pc *piConsensus) verifyTransactionSignature(transaction *types.Transaction) (bool, error) {
	signature, err := hex.DecodeString(transaction.Signature)
	if err != nil {
		return false, err
	}
	r := big.NewInt(0).SetBytes(signature[:len(signature)/2])
	s := big.NewInt(0).SetBytes(signature[len(signature)/2:])
	if ecd

package protocol

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/pi-network/pi-node/types"
	"github.com/pi-network/pi-node/utils"
)

// PiProtocol represents the Pi protocol
type PiProtocol struct {
	// Node private key
	privateKey *ecdsa.PrivateKey

	// Node public key
	publicKey *ecdsa.PublicKey

	// Node address
	address string

	// Node port
	port int

	// Protocol mutex
	mutex sync.Mutex

	// Transaction pool
	transactionPool map[string]*types.Transaction

	// Block chain
	blockChain []*types.Block
}

// NewPiProtocol creates a new Pi protocol
func NewPiProtocol(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, address string, port int) *PiProtocol {
	protocol := &PiProtocol{
		privateKey:     privateKey,
		publicKey:     publicKey,
		address:       address,
		port:          port,
		transactionPool: make(map[string]*types.Transaction),
		blockChain:     make([]*types.Block, 0),
	}

	return protocol
}

// HandleMessage handles a message
func (p *PiProtocol) HandleMessage(message *types.Message) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	switch message.Type {
	case types.MessageTypeNodeInfo:
		return p.handleNodeInfoMessage(message)
	case types.MessageTypeTransaction:
		return p.handleTransactionMessage(message)
	case types.MessageTypeBlock:
		return p.handleBlockMessage(message)
	default:
		return fmt.Errorf("unknown message type: %s", message.Type)
	}
}

func (p *PiProtocol) handleNodeInfoMessage(message *types.Message) error {
	log.Println("Received node info message")

	nodeInfo := &types.NodeInfo{}
	err := json.Unmarshal(message.Data, nodeInfo)
	if err != nil {
		return err
	}

	log.Println("Node info:", nodeInfo)

	return nil
}

func (p *PiProtocol) handleTransactionMessage(message *types.Message) error {
	log.Println("Received transaction message")

	transaction := &types.Transaction{}
	err := json.Unmarshal(message.Data, transaction)
	if err != nil {
		return err
	}

	log.Println("Transaction:", transaction)

	p.transactionPool[transaction.ID] = transaction

	return nil
}

func (p *PiProtocol) handleBlockMessage(message *types.Message) error {
	log.Println("Received block message")

	block := &types.Block{}
	err := json.Unmarshal(message.Data, block)
	if err != nil {
		return err
	}

	log.Println("Block:", block)

	p.blockChain = append(p.blockChain, block)

	return nil
}

// CreateBlock creates a new block
func (p *PiProtocol) CreateBlock() (*types.Block, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	block := &types.Block{
		Timestamp: time.Now().Unix(),
		Transactions: make([]*types.Transaction, 0),
		PreviousBlockHash: p.getPreviousBlockHash(),
	}

	for _, transaction := range p.transactionPool {
		block.Transactions = append(block.Transactions, transaction)
		delete(p.transactionPool, transaction.ID)
	}

	blockHash, err := p.calculateBlockHash(block)
	if err != nil {
		return nil, err
	}

	block.Hash = blockHash

	return block, nil
}

func (p *PiProtocol) getPreviousBlockHash() string {
	if len(p.blockChain) == 0 {
		return ""
	}

	return p.blockChain[len(p.blockChain)-1].Hash
}

func (p *PiProtocol) calculateBlockHash(block *types.Block) (string, error) {
	blockBytes, err := json.Marshal(block)
	if err != nil {
		return "", err
	}

	hash, err := utils.GenerateHash(blockBytes)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// GetBlockChain returns the block chain
func (p *PiProtocol) GetBlockChain() []*types.Block {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.blockChain
}

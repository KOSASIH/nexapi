package protocol

import (
	"testing"

	"github.com/pi-network/pi-node/types"
	"github.com/pi-network/pi-node/utils"
)

func TestNewPiProtocol(t *testing.T) {
	privateKey, err := utils.GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	address, err := utils.GenerateAddress(publicKey)
	if err != nil {
		t.Fatal(err)
	}

	protocol := NewPiProtocol(privateKey, publicKey, address, 8080)
	if protocol == nil {
		t.Errorf("Expected protocol to be created, but got nil")
	}
}

func TestHandleMessage(t *testing.T) {
	privateKey, err := utils.GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	address, err := utils.GenerateAddress(publicKey)
	if err != nil {
		t.Fatal(err)
	}

	protocol := NewPiProtocol(privateKey, publicKey, address, 8080)

	nodeInfoMessage := &types.Message{
		Type: types.MessageTypeNodeInfo,
		Data: []byte(`{"address":"localhost:8080","port":8080}`),
	}

	err = protocol.HandleMessage(nodeInfoMessage)
	if err != nil {
		t.Errorf("Expected node info message to be handled, but got error: %s", err)
	}

	transactionMessage := &types.Message{
		Type: types.MessageTypeTransaction,
		Data: []byte(`{"id":"transaction-id","from":"from-address","to":"to-address","amount":10}`),
	}

	err = protocol.HandleMessage(transactionMessage)
	if err != nil {
		t.Errorf("Expected transaction message to be handled, but got error: %s", err)
	}

	blockMessage := &types.Message{
		Type: types.MessageTypeBlock,
		Data: []byte(`{"hash":"block-hash","previousBlockHash":"previous-block-hash","transactions":[{"id":"transaction-id","from":"from-address","to":"to-address","amount":10}]}`),
	}

	err = protocol.HandleMessage(blockMessage)
	if err != nil {
		t.Errorf("Expected block message to be handled, but got error: %s", err)
	}
}

func TestCreateBlock(t *testing.T) {
	privateKey, err := utils.GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	address, err := utils.GenerateAddress(publicKey)
	if err != nil {
		t.Fatal(err)
	}

	protocol := NewPiProtocol(privateKey, publicKey, address, 8080)

	transaction := &types.Transaction{
		ID:     "transaction-id",
		From:   "from-address",
		To:     "to-address",
		Amount: 10,
	}

	protocol.transactionPool[transaction.ID] = transaction

	block, err := protocol.CreateBlock()
	if err != nil {
		t.Errorf("Expected block to be created, but got error: %s", err)
	}

	if block == nil {
		t.Errorf("Expected block to be created, but got nil")
	}

	if len(block.Transactions) != 1 {
		t.Errorf("Expected block to have 1 transaction, but got %d", len(block.Transactions))
	}
}

func TestGetBlockChain(t *testing.T) {
	privateKey, err := utils.GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	address, err := utils.GenerateAddress(publicKey)
	if err != nil {
		t.Fatal(err)
	}

	protocol := NewPiProtocol(privateKey, publicKey, address, 8080)

	blockChain := protocol.GetBlockChain()
	if blockChain == nil {
		t.Errorf("Expected block chain to be returned, but got nil")
	}

	if len(blockChain) != 0 {
		t.Errorf("Expected block chain to be empty, but got %d blocks", len(blockChain))
	}
}

package algorithm

import (
	"testing"
	"time"

	"github.com/pi-network/pi/types"
)

func TestPiConsensus_Initialize(t *testing.T) {
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

	pc := NewPiConsensus(privateKey, publicKey, address)
	err = pc.Initialize()
	if err != nil {
		t.Errorf("Expected Initialize to succeed, but got error: %s", err)
	}
}

func TestPiConsensus_Start(t *testing.T) {
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

	pc := NewPiConsensus(privateKey, publicKey, address)
	err = pc.Initialize()
	if err != nil {
		t.Fatal(err)
	}
	err = pc.Start()
	if err != nil {
		t.Errorf("Expected Start to succeed, but got error: %s", err)
	}
}

func TestPiConsensus_Stop(t *testing.T) {
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

	pc := NewPiConsensus(privateKey, publicKey, address)
	err = pc.Initialize()
	if err != nil {
		t.Fatal(err)
	}
	err = pc.Start()
	if err != nil {
		t.Fatal(err)
	}
	err = pc.Stop()
	if err != nil {
		t.Errorf("Expected Stop to succeed, but got error: %s", err)
	}
}

func TestPiConsensus_VerifyBlock(t *testing.T) {
	privateKey, err := utils.GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err!= nil {
		t.Fatal(err)
	}
	address, err := utils.GenerateAddress(publicKey)
	if err!= nil {
		t.Fatal(err)
	}

	pc := NewPiConsensus(privateKey, publicKey, address)
	err = pc.Initialize()
	if err!= nil {
		t.Fatal(err)
	}

	block := &types.Block{
		Hash:        "block-hash",
		PreviousHash: "previous-block-hash",
		Transactions: []*types.Transaction{
			{
				ID:     "transaction-id",
				From:   "from-address",
				To:     "to-address",
				Amount: 10,
				Signature: "signature",
			},
		},
		Timestamp: time.Now().Unix(),
	}

	valid, err := pc.VerifyBlock(block)
	if err!= nil {
		t.Errorf("Expected VerifyBlock to succeed, but got error: %s", err)
	}
	if!valid {
		t.Errorf("Expected block to be valid, but got invalid")
	}
}

func TestPiConsensus_VerifyTransaction(t *testing.T) {
	privateKey, err := utils.GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err!= nil {
		t.Fatal(err)
	}
	address, err := utils.GenerateAddress(publicKey)
	if err!= nil {
		t.Fatal(err)
	}

	pc := NewPiConsensus(privateKey, publicKey, address)
	err = pc.Initialize()
	if err!= nil {
		t.Fatal(err)
	}

	transaction := &types.Transaction{
		ID:     "transaction-id",
		From:   "from-address",
		To:     "to-address",
		Amount: 10,
		Signature: "signature",
	}

	valid, err := pc.VerifyTransaction(transaction)
	if err!= nil {
		t.Errorf("Expected VerifyTransaction to succeed, but got error: %s", err)
	}
	if!valid {
		t.Errorf("Expected transaction to be valid, but got invalid")
	}
}

func TestPiConsensus_consensusLoop(t *testing.T) {
	privateKey, err := utils.GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err!= nil {
		t.Fatal(err)
	}
	address, err := utils.GenerateAddress(publicKey)
	if err!= nil {
		t.Fatal(err)
	}

	pc := NewPiConsensus(privateKey, publicKey, address)
	err = pc.Initialize()
	if err!= nil {
		t.Fatal(err)
	}
	err = pc.Start()
	if err!= nil {
		t.Fatal(err)
	}

	// Wait for the consensus loop to run for a few seconds
	time.Sleep(5 * time.Second)

	// Stop the consensus loop
	err = pc.Stop()
	if err!= nil {
		t.Errorf("Expected Stop to succeed, but got error: %s", err)
	}
}

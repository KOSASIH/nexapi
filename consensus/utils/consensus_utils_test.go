package utils

import (
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err!= nil {
		t.Errorf("Expected GeneratePrivateKey to succeed, but got error: %s", err)
	}
	if privateKey == nil {
		t.Errorf("Expected privateKey to be non-nil, but got nil")
	}
}

func TestGeneratePublicKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey, err := GeneratePublicKey(privateKey)
	if err!= nil {
		t.Errorf("Expected GeneratePublicKey to succeed, but got error: %s", err)
	}
	if publicKey == nil {
		t.Errorf("Expected publicKey to be non-nil, but got nil")
	}
}

func TestGenerateAddress(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey, err := GeneratePublicKey(privateKey)
	if err!= nil {
		t.Fatal(err)
	}
	address, err := GenerateAddress(publicKey)
	if err!= nil {
		t.Errorf("Expected GenerateAddress to succeed, but got error: %s", err)
	}
	if address == "" {
		t.Errorf("Expected address to be non-empty, but got empty string")
	}
}

func TestSign(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	message := []byte("hello world")
	signature, err := Sign(privateKey, message)
	if err!= nil {
		t.Errorf("Expected Sign to succeed, but got error: %s", err)
	}
	if signature == "" {
		t.Errorf("Expected signature to be non-empty, but got empty string")
	}
}

func TestVerify(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey, err := GeneratePublicKey(privateKey)
	if err!= nil {
		t.Fatal(err)
	}
	message := []byte("hello world")
	signature, err := Sign(privateKey, message)
	if err!= nil {
		t.Fatal(err)
	}
	valid, err := Verify(publicKey, message, signature)
	if err!= nil {
		t.Errorf("Expected Verify to succeed, but got error: %s", err)
	}
	if!valid {
		t.Errorf("Expected signature to be valid, but got invalid")
	}
}

func TestCalculateBlockHash(t *testing.T) {
	block := &types.Block{
		PreviousHash: "previous-block-hash",
		Transactions: []*types.Transaction{
			{
				ID: "transaction-id",
			},
		},
		Timestamp: time.Now().Unix(),
	}
	hash, err := CalculateBlockHash(block)
	if err!= nil {
		t.Errorf("Expected CalculateBlockHash to succeed, but got error: %s", err)
	}
	if hash == "" {
		t.Errorf("Expected hash to be non-empty, but got empty string")
	}
}

package utils

import (
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Errorf("Expected private key to be generated, but got error: %s", err)
	}

	if privateKey == nil {
		t.Errorf("Expected private key to be generated, but got nil")
	}
}

func TestGeneratePublicKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := GeneratePublicKey(privateKey)
	if err != nil {
		t.Errorf("Expected public key to be generated, but got error: %s", err)
	}

	if publicKey == nil {
		t.Errorf("Expected public key to be generated, but got nil")
	}
}

func TestGenerateAddress(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := GeneratePublicKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	address, err := GenerateAddress(publicKey)
	if err != nil {
		t.Errorf("Expected address to be generated, but got error: %s", err)
	}

	if address == "" {
		t.Errorf("Expected address to be generated, but got empty string")
	}
}

func TestGenerateHash(t *testing.T) {
	data := []byte("hello world")
	hash, err := GenerateHash(data)
	if err != nil {
		t.Errorf("Expected hash to be generated, but got error: %s", err)
	}

	if hash == "" {
		t.Errorf("Expected hash to be generated, but got empty string")
	}
}

func TestSignAndVerify(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	publicKey, err := GeneratePublicKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	data := []byte("hello world")
	signature, err := Sign(privateKey, data)
	if err != nil {
		t.Errorf("Expected signature to be generated, but got error: %s", err)
	}

	valid, err := Verify(publicKey, data, signature)
	if err != nil {
		t.Errorf("Expected signature to be verified, but got error: %s", err)
	}

	if !valid {
		t.Errorf("Expected signature to be valid, but got invalid")
	}
}

package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// GeneratePrivateKey generates a new ECDSA private key
func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

// GeneratePublicKey generates a new ECDSA public key from a private key
func GeneratePublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	return &privateKey.PublicKey, nil
}

// GenerateAddress generates a new address from a public key
func GenerateAddress(publicKey *ecdsa.PublicKey) (string, error) {
	hash := sha256.Sum256(elliptic.Marshal(publicKey.X, publicKey.Y))
	return fmt.Sprintf("0x%x", hash), nil
}

// Sign signs a message with a private key
func Sign(privateKey *ecdsa.PrivateKey, message []byte) (string, error) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, message)
	if err!= nil {
		return "", err
	}
	return fmt.Sprintf("%x%x", r, s), nil
}

// Verify verifies a signature with a public key
func Verify(publicKey *ecdsa.PublicKey, message []byte, signature string) (bool, error) {
	r, s, err := hex.DecodeString(signature)
	if err!= nil {
		return false, err
	}
	return ecdsa.Verify(&publicKey, message, r, s), nil
}

// CalculateBlockHash calculates the hash of a block
func CalculateBlockHash(block *types.Block) (string, error) {
	hash := sha256.Sum256([]byte(block.PreviousHash + block.Transactions[0].ID + fmt.Sprintf("%d", block.Timestamp)))
	return fmt.Sprintf("0x%x", hash), nil
}

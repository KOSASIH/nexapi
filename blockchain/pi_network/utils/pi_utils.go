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
	pubBytes := elliptic.Marshal(publicKey, publicKey.X, publicKey.Y)
	hash := sha256.Sum256(pubBytes)
	address := hex.EncodeToString(hash[:])
	return address, nil
}

// GenerateHash generates a new hash from a byte slice
func GenerateHash(data []byte) (string, error) {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

// Sign signs a byte slice with a private key
func Sign(privateKey *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, data)
	if err != nil {
		return nil, err
	}
	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

// Verify verifies a signature with a public key
func Verify(publicKey *ecdsa.PublicKey, data []byte, signature []byte) (bool, error) {
	r := big.NewInt(0).SetBytes(signature[:len(signature)/2])
	s := big.NewInt(0).SetBytes(signature[len(signature)/2:])
	if ecdsa.Verify(publicKey, data, r, s) {
		return true, nil
	}
	return false, fmt.Errorf("invalid signature")
}

package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GeneratePrivateKey generates a new private key
func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

// GetAccountFromPrivateKey gets an account from a private key
func GetAccountFromPrivateKey(privateKey *ecdsa.PrivateKey) (accounts.Account, error) {
	return accounts.Account{
		Address: crypto.PubkeyToAddress(privateKey.PublicKey),
	}, nil
}

// GetPublicKeyFromPrivateKey gets a public key from a private key
func GetPublicKeyFromPrivateKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	return &privateKey.PublicKey, nil
}

// GetAddressFromPublicKey gets an address from a public key
func GetAddressFromPublicKey(publicKey *ecdsa.PublicKey) (common.Address, error) {
	return crypto.PubkeyToAddress(*publicKey), nil
}

// HexToPrivateKey converts a hex string to a private key
func HexToPrivateKey(hexString string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(hexString)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// HexToAddress converts a hex string to an address
func HexToAddress(hexString string) (common.Address, error) {
	address := common.HexToAddress(hexString)
	return address, nil
}

// AddressToString converts an address to a string
func AddressToString(address common.Address) string {
	return address.Hex()
}

// PrivateKeyToString converts a private key to a string
func PrivateKeyToString(privateKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(privateKey))
}

// PublicKeyToString converts a public key to a string
func PublicKeyToString(publicKey *ecdsa.PublicKey) string {
	return hex.EncodeToString(crypto.FromECDSA(publicKey))
}

// RandomHex generates a random hex string
func RandomHex(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// SplitAddress splits an address into its components
func SplitAddress(address common.Address) (string, string, error) {
	parts := strings.Split(address.Hex(), "0x")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid address format")
	}
	return parts[0], parts[1], nil
}

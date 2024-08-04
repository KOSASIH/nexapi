package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// HexToBytes converts a hex string to a byte slice
func HexToBytes(hexString string) ([]byte, error) {
	return hex.DecodeString(strings.TrimPrefix(hexString, "0x"))
}

// BytesToHex converts a byte slice to a hex string
func BytesToHex(bytes []byte) string {
	return "0x" + hex.EncodeToString(bytes)
}

// U256ToBigInt converts a U256 to a big.Int
func U256ToBigInt(u256 types.U256) *big.Int {
	return big.NewInt(0).SetBytes(u256[:])
}

// BigIntToU256 converts a big.Int to a U256
func BigIntToU256(bigInt *big.Int) types.U256 {
	return types.U256(bigInt.Bytes())
}

// SignMessage signs a message with a private key
func SignMessage(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	return crypto.Sign(privateKey, message)
}

// VerifySignature verifies a signature with a public key
func VerifySignature(publicKey *ecdsa.PublicKey, message []byte, signature []byte) bool {
	return crypto.VerifySignature(publicKey, message, signature)
}

// GenerateKey generates a new ECDSA key pair
func GenerateKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	return crypto.GenerateKey()
}

// GetAccountAddressFromPrivateKey gets the account address from a private key
func GetAccountAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(privateKey.PublicKey)
}

// GetAccountAddressFromPublicKey gets the account address from a public key
func GetAccountAddressFromPublicKey(publicKey *ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(*publicKey)
}

// EncodeExtrinsic encodes an extrinsic to JSON
func EncodeExtrinsic(ext *types.Extrinsic) ([]byte, error) {
	return json.Marshal(ext)
}

// DecodeExtrinsic decodes an extrinsic from JSON
func DecodeExtrinsic(data []byte) (*types.Extrinsic, error) {
	var ext types.Extrinsic
	err := json.Unmarshal(data, &ext)
	return &ext, err
}

// EncodeBlock encodes a block to JSON
func EncodeBlock(block *types.Block) ([]byte, error) {
	return json.Marshal(block)
}

// DecodeBlock decodes a block from JSON
func DecodeBlock(data []byte) (*types.Block, error) {
	var block types.Block
	err := json.Unmarshal(data, &block)
	return &block, err
}

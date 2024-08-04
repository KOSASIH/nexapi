package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestHexToBytes(t *testing.T) {
	hexString := "0x1234567890abcdef"
	bytes, err := HexToBytes(hexString)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Bytes: %x\n", bytes)
}

func TestBytesToHex(t *testing.T) {
	bytes := []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}
	hexString := BytesToHex(bytes)
	fmt.Printf("Hex string: %s\n", hexString)
}

func TestU256ToBigInt(t *testing.T) {
	u256 := types.U256{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}
	bigInt := U256ToBigInt(u256)
	fmt.Printf("Big int: %s\n", bigInt)
}

func TestBigIntToU256(t *testing.T) {
	bigInt := big.NewInt(0x1234567890abcdef)
	u256 := BigIntToU256(bigInt)
	fmt.Printf("U256: %x\n", u256)
}

func TestSignMessage(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err!= nil {
		t.Fatal(err)
	}
	message := []byte("Hello, World!")
	signature, err := SignMessage(privateKey, message)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Signature: %x\n", signature)
}

func TestVerifySignature(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey
	message := []byte("Hello, World!")
	signature, err := SignMessage(privateKey, message)
	if err!= nil {
		t.Fatal(err)
	}
	valid := VerifySignature(publicKey, message, signature)
	if!valid {
		t.Fatal("Invalid signature")
	}
	fmt.Println("Signature is valid")
}

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey, err := GenerateKey()
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Private key: %x\n", privateKey.D.Bytes())
	fmt.Printf("Public key: %x\n", publicKey.X.Bytes())
}

func TestGetAccountAddressFromPrivateKey(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err!= nil {
		t.Fatal(err)
	}
	address := GetAccountAddressFromPrivateKey(privateKey)
	fmt.Printf("Account address: %s\n", address.Hex())
}

func TestGetAccountAddressFromPublicKey(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey
	address := GetAccountAddressFromPublicKey(publicKey)
	fmt.Printf("Account address: %s\n", address.Hex())
}

func TestEncodeExtrinsic(t *testing.T) {
	ext := &types.Extrinsic{
		Call: &types.Call{
			Module: "balances",
			Func:   "transfer",
			Args: []types.Arg{
				types.NewArg("amount", types.NewU128(100)),
				types.NewArg("dest", types.NewAccountId("0x1234567890abcdef")),
			},
		},
	}
	data, err := EncodeExtrinsic(ext)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Encoded extrinsic: %s\n", data)
}

func TestDecodeExtrinsic(t *testing.T) {
	data := []byte(`{"call":{"module":"balances","func":"transfer","args":[{"name":"amount","value":"100"},{"name":"dest","value":"0x1234567890abcdef"}]}}`)
	ext, err := DecodeExtrinsic(data)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Decoded extrinsic: %+v\n", ext)
}

func TestEncodeBlock(t *testing.T) {
	block := &types.Block{
		Header: &types.Header{
			Number: 100,
			Hash:   types.NewHash("0x1234567890abcdef"),
		},
	}
	data, err := EncodeBlock(block)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Encoded block: %s\n", data)
}

func TestDecodeBlock(t *testing.T) {
	data := []byte(`{"header":{"number":100,"hash":"0x1234567890abcdef"}}`)
	block, err := DecodeBlock(data)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Decoded block: %+v\n", block)
}

package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Private key: %s\n", PrivateKeyToString(privateKey))
}

func TestGetAccountFromPrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	account, err := GetAccountFromPrivateKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Account: %v\n", account)
}

func TestGetPublicKeyFromPrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	publicKey, err := GetPublicKeyFromPrivateKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Public key: %s\n", PublicKeyToString(publicKey))
}

func TestGetAddressFromPublicKey(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	publicKey, err := GetPublicKeyFromPrivateKey(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	address, err := GetAddressFromPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Address: %s\n", address.Hex())
}

func TestHexToPrivateKey(t *testing.T) {
	hexString := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
	privateKey, err := HexToPrivateKey(hexString)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Private key: %s\n", PrivateKeyToString(privateKey))
}

func TestHexToAddress(t *testing.T) {
	hexString := "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
	address, err := HexToAddress(hexString)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Address: %s\n", address.Hex())
}

func TestAddressToString(t *testing.T) {
	address := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
	str := AddressToString(address)
	fmt.Printf("Address string: %s\n", str)
}

func TestPrivateKeyToString(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	str := PrivateKeyToString(privateKey)
	fmt.Printf("Private key string: %s\n", str)
}

func TestPublicKeyToString(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err!= nil {
		t.Fatal(err)
	}
	publicKey, err := GetPublicKeyFromPrivateKey(privateKey)
	if err!= nil {
		t.Fatal(err)
	}
	str := PublicKeyToString(publicKey)
	fmt.Printf("Public key string: %s\n", str)
}

func TestRandomHex(t *testing.T) {
	hexString, err := RandomHex(32)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Random hex: %s\n", hexString)
}

func TestSplitAddress(t *testing.T) {
	address := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")
	prefix, hex, err := SplitAddress(address)
	if err!= nil {
		t.Fatal(err)
	}
	fmt.Printf("Prefix: %s, Hex: %s\n", prefix, hex)
}

package utils

import (
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	b, err := GenerateRandomBytes(32)
	if err != nil {
		t.Errorf("Expected GenerateRandomBytes to succeed, but got error: %s", err)
	}
	if len(b) != 32 {
		t.Errorf("Expected GenerateRandomBytes to return 32 bytes, but got %d", len(b))
	}
}

func TestHashSHA256(t *testing.T) {
	data := []byte("hello world")
	hash, err := HashSHA256(data)
	if err != nil {
		t.Errorf("Expected HashSHA256 to succeed, but got error: %s", err)
	}
	if len(hash) != 32 {
		t.Errorf("Expected HashSHA256 to return 32 bytes, but got %d", len(hash))
	}
}

func TestEncryptAES(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err!= nil {
		t.Fatal(err)
	}
	plaintext := []byte("hello world")
	ciphertext, err := EncryptAES(plaintext, key)
	if err!= nil {
		t.Errorf("Expected EncryptAES to succeed, but got error: %s", err)
	}
	if len(ciphertext) < len(plaintext) {
		t.Errorf("Expected ciphertext to be at least as long as plaintext")
	}
}

func TestDecryptAES(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err!= nil {
		t.Fatal(err)
	}
	plaintext := []byte("hello world")
	ciphertext, err := EncryptAES(plaintext, key)
	if err!= nil {
		t.Errorf("Expected EncryptAES to succeed, but got error: %s", err)
	}
	decrypted, err := DecryptAES(ciphertext, key)
	if err!= nil {
		t.Errorf("Expected DecryptAES to succeed, but got error: %s", err)
	}
	if string(decrypted)!= string(plaintext) {
		t.Errorf("Expected decrypted text to match original plaintext, but got %s", decrypted)
	}
}

func TestSignRSA(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err!= nil {
		t.Fatal(err)
	}
	data := []byte("hello world")
	signature, err := SignRSA(data, privateKey)
	if err!= nil {
		t.Errorf("Expected SignRSA to succeed, but got error: %s", err)
	}
	if len(signature) == 0 {
		t.Errorf("Expected signature to be non-empty")
	}
}

func TestVerifyRSA(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err!= nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey
	data := []byte("hello world")
	signature, err := SignRSA(data, privateKey)
	if err!= nil {
		t.Errorf("Expected SignRSA to succeed, but got error: %s", err)
	}
	err = VerifyRSA(data, signature, publicKey)
	if err!= nil {
		t.Errorf("Expected VerifyRSA to succeed, but got error: %s", err)
	}
}

func TestGenerateHMAC(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err!= nil {
		t.Fatal(err)
	}
	data := []byte("hello world")
	hmac, err := GenerateHMAC(data, key)
	if err!= nil {
		t.Errorf("Expected GenerateHMAC to succeed, but got error: %s", err)
	}
	if len(hmac) == 0 {
		t.Errorf("Expected HMAC to be non-empty")
	}
}

func TestVerifyHMAC(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err!= nil {
		t.Fatal(err)
	}
	data := []byte("hello world")
	hmac, err := GenerateHMAC(data, key)
	if err!= nil {
		t.Errorf("Expected GenerateHMAC to succeed, but got error: %s", err)
	}
	err = VerifyHMAC(data, hmac, key)
	if err!= nil {
		t.Errorf("Expected VerifyHMAC to succeed, but got error: %s", err)
	}
}

func TestBase64Encode(t *testing.T) {
	data := []byte("hello world")
	encoded, err := Base64Encode(data)
	if err!= nil {
		t.Errorf("Expected Base64Encode to succeed, but got error: %s", err)
	}
	if len(encoded) == 0 {
		t.Errorf("Expected encoded string to be non-empty")
	}
}

func TestBase64Decode(t *testing.T) {
	encoded := "SGVsbG8gd29ybGQ="
	data, err := Base64Decode(encoded)
	if err!= nil {
		t.Errorf("Expected Base64Decode to succeed, but got error: %s", err)
	}
	if string(data)!= "hello world" {
		t.Errorf("Expected decoded data to match original data, but got %s", data)
	}
}

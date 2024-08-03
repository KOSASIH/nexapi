package encryption

import (
	"testing"
)

func TestNewPiEncryption(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pe, err := NewPiEncryption(key)
	if err != nil {
		t.Errorf("Expected NewPiEncryption to succeed, but got error: %s", err)
	}
	if pe == nil {
		t.Errorf("Expected pe to be non-nil, but got nil")
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pe, err := NewPiEncryption(key)
	if err != nil {
		t.Fatal(err)
	}
	plaintext := []byte("hello world")
	ciphertext, err := pe.Encrypt(plaintext)
	if err != nil {
		t.Errorf("Expected Encrypt to succeed, but got error: %s", err)
	}
	decrypted, err := pe.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("Expected Decrypt to succeed, but got error: %s", err)
	}
	if string(decrypted) != string(plaintext) {
		t.Errorf("Expected decrypted text to match original plaintext, but got %s", decrypted)
	}
}

func TestEncryptBase64DecryptBase64(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pe, err := NewPiEncryption(key)
	if err != nil {
		t.Fatal(err)
	}
	plaintext := []byte("hello world")
	encrypted, err := pe.EncryptBase64(plaintext)
	if err != nil {
		t.Errorf("Expected EncryptBase64 to succeed, but got error: %s", err)
	}
	decrypted, err := pe.DecryptBase64(encrypted)
	if err != nil {
		t.Errorf("Expected DecryptBase64 to succeed, but got error: %s", err)
	}
	if string(decrypted) != string(plaintext) {
		t.Errorf("Expected decrypted text to match original plaintext, but got %s", decrypted)
	}
}

func TestInvalidKeyLength(t *testing.T) {
	key := make([]byte, 16)
	_, err := NewPiEncryption(key)
	if err == nil {
		t.Errorf("Expected NewPiEncryption to fail with invalid key length, but got nil error")
	}
}

func TestInvalidCiphertext(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pe, err := NewPiEncryption(key)
	if err != nil {
		t.Fatal(err)
	}
	ciphertext := []byte("invalid ciphertext")
	_, err = pe.Decrypt(ciphertext)
	if err == nil {
		t.Errorf("Expected Decrypt to fail with invalid ciphertext, but got nil error")
	}
}

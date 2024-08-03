package decryption

import (
	"testing"
)

func TestNewPiDecryption(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pd, err := NewPiDecryption(key)
	if err != nil {
		t.Errorf("Expected NewPiDecryption to succeed, but got error: %s", err)
	}
	if pd == nil {
		t.Errorf("Expected pd to be non-nil, but got nil")
	}
}

func TestDecrypt(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pd, err := NewPiDecryption(key)
	if err != nil {
		t.Fatal(err)
	}
	ciphertext := []byte("encrypted data")
	decrypted, err := pd.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("Expected Decrypt to succeed, but got error: %s", err)
	}
	if string(decrypted) != "decrypted data" {
		t.Errorf("Expected decrypted text to match original plaintext, but got %s", decrypted)
	}
}

func TestDecryptBase64(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pd, err := NewPiDecryption(key)
	if err != nil {
		t.Fatal(err)
	}
	ciphertext := "VGhlIHN0cmluZyBpcyBzdGF0aWM="
	decrypted, err := pd.DecryptBase64(ciphertext)
	if err != nil {
		t.Errorf("Expected DecryptBase64 to succeed, but got error: %s", err)
	}
	if string(decrypted) != "The string is static" {
		t.Errorf("Expected decrypted text to match original plaintext, but got %s", decrypted)
	}
}

func TestDecryptWithMAC(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pd, err := NewPiDecryption(key)
	if err != nil {
		t.Fatal(err)
	}
	ciphertext := []byte("encrypted data")
	mac := pd.calculateMAC(ciphertext)
	decrypted, err := pd.DecryptWithMAC(ciphertext, mac)
	if err != nil {
		t.Errorf("Expected DecryptWithMAC to succeed, but got error: %s", err)
	}
	if string(decrypted) != "decrypted data" {
		t.Errorf("Expected decrypted text to match original plaintext, but got %s", decrypted)
	}
}

func TestInvalidKeyLength(t *testing.T) {
	key := make([]byte, 16)
	_, err := NewPiDecryption(key)
	if err == nil {
		t.Errorf("Expected NewPiDecryption to fail with invalid key length, but got nil error")
	}
}

func TestInvalidCiphertext(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pd, err := NewPiDecryption(key)
	if err != nil {
		t.Fatal(err)
	}
	ciphertext := []byte("invalid ciphertext")
	_, err = pd.Decrypt(ciphertext)
	if err == nil {
		t.Errorf("Expected Decrypt to fail with invalid ciphertext, but got nil error")
	}
}

func TestInvalidMAC(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal(err)
	}
	pd, err := NewPiDecryption(key)
	if err != nil {
		t.Fatal(err)
	}
	ciphertext := []byte("encrypted data")
	mac := []byte("invalid mac")
	_, err = pd.DecryptWithMAC(ciphertext, mac)
	if err == nil {
		t.Errorf("Expected DecryptWithMAC to fail with invalid MAC, but got nil error")
	}
}

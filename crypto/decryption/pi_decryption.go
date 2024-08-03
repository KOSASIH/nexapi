package decryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// PiDecryption is a struct that holds the decryption key and other parameters
type PiDecryption struct {
	key []byte
}

// NewPiDecryption creates a new instance of PiDecryption
func NewPiDecryption(key []byte) (*PiDecryption, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes long")
	}
	return &PiDecryption{key: key}, nil
}

// Decrypt decrypts a ciphertext message using AES-256-CBC
func (pd *PiDecryption) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(pd.key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

// DecryptBase64 decrypts a base64-encoded ciphertext message and returns the result as a plaintext byte slice
func (pd *PiDecryption) DecryptBase64(ciphertext string) ([]byte, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	return pd.Decrypt(ciphertextBytes)
}

// DecryptWithMAC decrypts a ciphertext message using AES-256-CBC and verifies the MAC (Message Authentication Code)
func (pd *PiDecryption) DecryptWithMAC(ciphertext []byte, mac []byte) ([]byte, error) {
	block, err := aes.NewCipher(pd.key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	expectedMAC := pd.calculateMAC(ciphertext)
	if !hmac.Equal(expectedMAC, mac) {
		return nil, errors.New("MAC verification failed")
	}
	return ciphertext, nil
}

func (pd *PiDecryption) calculateMAC(data []byte) []byte {
	mac := hmac.New(sha256.New, pd.key)
	mac.Write(data)
	return mac.Sum(nil)
}

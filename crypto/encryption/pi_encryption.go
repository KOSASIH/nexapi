package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// PiEncryption is a struct that holds the encryption key and other parameters
type PiEncryption struct {
	key []byte
}

// NewPiEncryption creates a new instance of PiEncryption
func NewPiEncryption(key []byte) (*PiEncryption, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes long")
	}
	return &PiEncryption{key: key}, nil
}

// Encrypt encrypts a plaintext message using AES-256-CBC
func (pe *PiEncryption) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(pe.key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// Decrypt decrypts a ciphertext message using AES-256-CBC
func (pe *PiEncryption) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(pe.key)
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

// EncryptBase64 encrypts a plaintext message and returns the result as a base64-encoded string
func (pe *PiEncryption) EncryptBase64(plaintext []byte) (string, error) {
	ciphertext, err := pe.Encrypt(plaintext)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptBase64 decrypts a base64-encoded ciphertext message and returns the result as a plaintext byte slice
func (pe *PiEncryption) DecryptBase64(ciphertext string) ([]byte, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	return pe.Decrypt(ciphertextBytes)
}

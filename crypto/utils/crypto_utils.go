package utils

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

// GenerateRandomBytes generates a random byte slice of the given length
func GenerateRandomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	return b, err
}

// HashSHA256 hashes the given data using SHA-256
func HashSHA256(data []byte) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

// EncryptAES encrypts the given plaintext using AES-256-CBC
func EncryptAES(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	_, err = rand.Read(iv)
	if err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// DecryptAES decrypts the given ciphertext using AES-256-CBC
func DecryptAES(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
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

// SignRSA signs the given data using RSA
func SignRSA(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash, err := HashSHA256(data)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
}

// VerifyRSA verifies the given signature using RSA
func VerifyRSA(data []byte, signature []byte, publicKey *rsa.PublicKey) error {
	hash, err := HashSHA256(data)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash, signature)
}

// GenerateHMAC generates an HMAC (Message Authentication Code) for the given data
func GenerateHMAC(data []byte, key []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(data)
	if err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

// VerifyHMAC verifies the given HMAC for the given data
func VerifyHMAC(data []byte, hmac []byte, key []byte) error {
	expectedHMAC, err := GenerateHMAC(data, key)
	if err != nil {
		return err
	}
	if !hmac.Equal(expectedHMAC, hmac) {
		return errors.New("HMAC verification failed")
	}
	return nil
}

// Base64Encode encodes the given data using Base64
func Base64Encode(data []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(data), nil
}

// Base64Decode decodes the given Base64-encoded string
func Base64Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

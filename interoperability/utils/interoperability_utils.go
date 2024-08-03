package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GenerateRandomString generates a random string of a given length
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// ComputeHMAC computes the HMAC of a given message using a secret key
func ComputeHMAC(message []byte, secretKey []byte) ([]byte, error) {
	h := hmac.New(sha256.New, secretKey)
	_, err := h.Write(message)
	if err!= nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Base64Encode encodes a byte slice to a base64-encoded string
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode decodes a base64-encoded string to a byte slice
func Base64Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// JSONMarshal marshals a struct to a JSON byte slice
func JSONMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JSONUnmarshal unmarshals a JSON byte slice to a struct
func JSONUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// HTTPGet sends an HTTP GET request to a given URL
func HTTPGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err!= nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// URLParse parses a URL string into a URL struct
func URLParse(url string) (*url.URL, error) {
	return url.Parse(url)
}

// URLQueryEscape escapes a URL query string
func URLQueryEscape(query string) string {
	return url.QueryEscape(query)
}

// ErrorString returns a string representation of an error
func ErrorString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// IsError returns true if the error is not nil
func IsError(err error) bool {
	return err!= nil
}

package utils

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	length := 16
	result := GenerateRandomString(length)
	if len(result)!= length*2 {
		t.Errorf("Expected generated string to be %d characters long, but got %d", length*2, len(result))
	}
}

func TestComputeHMAC(t *testing.T) {
	message := []byte("hello")
	secretKey := []byte("secret")
	result, err := ComputeHMAC(message, secretKey)
	if err!= nil {
		t.Errorf("Expected ComputeHMAC to succeed, but got error: %s", err)
	}
	if len(result)!= 32 {
		t.Errorf("Expected HMAC to be 32 bytes long, but got %d", len(result))
	}
}

func TestBase64Encode(t *testing.T) {
	data := []byte("hello")
	encoded := Base64Encode(data)
	if encoded!= "aGVsbG8=" {
		t.Errorf("Expected base64-encoded string to be 'aGVsbG8=', but got %s", encoded)
	}
}

func TestBase64Decode(t *testing.T) {
	encoded := "aGVsbG8="
	data, err := Base64Decode(encoded)
	if err!= nil {
		t.Errorf("Expected Base64Decode to succeed, but got error: %s", err)
	}
	if string(data)!= "hello" {
		t.Errorf("Expected decoded string to be 'hello', but got %s", data)
	}
}

func TestJSONMarshal(t *testing.T) {
	structure := struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}{Foo: "hello", Bar: 42}
	data, err := JSONMarshal(structure)
	if err!= nil {
		t.Errorf("Expected JSONMarshal to succeed, but got error: %s", err)
	}
	if string(data)!= `{"foo":"hello","bar":42}` {
		t.Errorf("Expected marshaled JSON to be '{\"foo\":\"hello\",\"bar\":42}', but got %s", data)
	}
}

func TestJSONUnmarshal(t *testing.T) {
	data := []byte(`{"foo":"hello","bar":42}`)
	structure := struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}{}
	err := JSONUnmarshal(data, &structure)
	if err!= nil {
		t.Errorf("Expected JSONUnmarshal to succeed, but got error: %s", err)
	}
	if structure.Foo!= "hello" {
		t.Errorf("Expected unmarshaled foo to be 'hello', but got %s", structure.Foo)
	}
	if structure.Bar!= 42 {
		t.Errorf("Expected unmarshaled bar to be 42, but got %d", structure.Bar)
	}
}

func TestHTTPGet(t *testing.T) {
	url := "https://example.com"
	data, err := HTTPGet(url)
	if err!= nil {
		t.Errorf("Expected HTTPGet to succeed, but got error: %s", err)
	}
	if len(data) == 0 {
		t.Errorf("Expected HTTP response to be non-empty")
	}
}

func TestURLParse(t *testing.T) {
	urlString := "https://example.com/path?query=string"
	url, err := URLParse(urlString)
	if err!= nil {
		t.Errorf("Expected URLParse to succeed, but got error: %s", err)
	}
	if url.Scheme!= "https" {
		t.Errorf("Expected URL scheme to be 'https', but got %s", url.Scheme)
	}
	if url.Host!= "example.com" {
		t.Errorf("Expected URL host to be 'example.com', but got %s", url.Host)
	}
	if url.Path!= "/path" {
		t.Errorf("Expected URL path to be '/path', but got %s", url.Path)
	}
	if url.Query().Get("query")!= "string" {
		t.Errorf("Expected URL query to be 'query=string', but got %s", url.Query())
	}
}

func TestURLQueryEscape(t *testing.T) {
	query := "hello world"
	escaped := URLQueryEscape(query)
	if escaped!= "hello+world" {
		t.Errorf("Expected URL-escaped query to be 'hello+world', but got %s", escaped)
	}
}

func TestErrorString(t *testing.T) {
	err := errors.New("test error")
	if ErrorString(err)!= "test error" {
		t.Errorf("Expected ErrorString to return 'test error', but got %s", ErrorString(err))
	}
}

func TestIsError(t *testing.T) {
	err := errors.New("test error")
	if!IsError(err) {
		t.Errorf("Expected IsError to return true, but got false")
	}
	if IsError(nil) {
		t.Errorf("Expected IsError to return false, but got true")
	}
}
		

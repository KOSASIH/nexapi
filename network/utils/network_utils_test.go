package utils

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

func TestDialContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := DialContext(ctx, "example.com:80")
	if err!= nil {
		t.Errorf("Expected DialContext to succeed, but got error: %s", err)
	}
	defer conn.Close()
}

func TestConnectContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := ConnectContext(ctx, "example.com:80")
	if err!= nil {
		t.Errorf("Expected ConnectContext to succeed, but got error: %s", err)
	}
	defer conn.Close()
}

func TestNewTLSConfig(t *testing.T) {
	certFile := "testdata/cert.pem"
	keyFile := "testdata/key.pem"
	tlsConfig, err := NewTLSConfig(certFile, keyFile)
	if err!= nil {
		t.Errorf("Expected NewTLSConfig to succeed, but got error: %s", err)
	}
	if tlsConfig.Certificates[0].Leaf.Subject.CommonName!= "example.com" {
		t.Errorf("Expected certificate subject to be example.com, but got %s", tlsConfig.Certificates[0].Leaf.Subject.CommonName)
	}
}

func TestNewLibp2pHost(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	listenAddr := "/ip4/0.0.0.0/tcp/0"
	host, err := NewLibp2pHost(context.Background(), privateKey, listenAddr)
	if err!= nil {
		t.Errorf("Expected NewLibp2pHost to succeed, but got error: %s", err)
	}
	defer host.Close()
	if host.ID().String()!= peer.IDFromPrivateKey(privateKey).String() {
		t.Errorf("Expected host ID to match private key ID")
	}
}

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err!= nil {
		t.Errorf("Expected GetFreePort to succeed, but got error: %s", err)
	}
	if port <= 0 || port >= 65536 {
		t.Errorf("Expected port to be in range 1-65535, but got %d", port)
	}
}

func TestGetExternalIP(t *testing.T) {
	ip, err := GetExternalIP()
	if err!= nil {
		t.Errorf("Expected GetExternalIP to succeed, but got error: %s", err)
	}
	if ip== nil {
		t.Errorf("Expected external IP to be non-nil")
	}
}
	

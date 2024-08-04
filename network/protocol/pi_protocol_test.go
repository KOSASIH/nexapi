package protocol

import (
	"context"
	"crypto/ecdsa"
	"testing"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multibase"
)

func TestNewPiProtocol(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	piProtocol, err := NewPiProtocol(privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiProtocol to succeed, but got error: %s", err)
	}
	if piProtocol.ID == nil {
		t.Errorf("Expected piProtocol.ID to be non-nil")
	}
}

func TestHandleStream(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	piProtocol, err := NewPiProtocol(privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiProtocol to succeed, but got error: %s", err)
	}
	stream, err := piProtocol.Host().NewStream(context.Background(), piProtocol.ID, ProtocolID)
	if err!= nil {
		t.Errorf("Expected NewStream to succeed, but got error: %s", err)
	}
	go piProtocol.HandleStream(stream)
	time.Sleep(100 * time.Millisecond)
	// Verify that the stream was handled correctly...
}

func TestSendMessage(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	piProtocol, err := NewPiProtocol(privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiProtocol to succeed, but got error: %s", err)
	}
	peerID, err := peer.IDFromPrivateKey(privateKey)
	if err!= nil {
		t.Errorf("Expected peer.IDFromPrivateKey to succeed, but got error: %s", err)
	}
	msg := []byte("Hello, world!")
	err = piProtocol.SendMessage(context.Background(), peerID, msg)
	if err!= nil {
		t.Errorf("Expected SendMessage to succeed, but got error: %s", err)
	}
	// Verify that the message was sent correctly...
}

func TestProcessMessage(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	piProtocol, err := NewPiProtocol(privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiProtocol to succeed, but got error: %s", err)
	}
	msg := []byte(`{"type": "hello", "data": "Hello, world!"}`)
	piProtocol.processMessage(msg)
	// Verify that the message was processed correctly...
}

func TestReadMessage(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	piProtocol, err := NewPiProtocol(privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiProtocol to succeed, but got error: %s", err)
	}
	stream, err := piProtocol.Host().NewStream(context.Background(), piProtocol.ID, ProtocolID)
	if err!= nil {
		t.Errorf("Expected NewStream to succeed, but got error: %s", err)
	}
	msg := []byte("Hello, world!")
	err = writeMessage(stream, msg)
	if err!= nil {
		t.Errorf("Expected writeMessage to succeed, but got error: %s", err)
	}
	readMsg, err := readMessage(stream)
	if err!= nil {
		t.Errorf("Expected readMessage to succeed, but got error: %s", err)
	}
	if!bytes.Equal(readMsg, msg) {
		t.Errorf("Expected read message to match written message")
	}
}

func TestWriteMessage(t *testing.T) {
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	piProtocol, err := NewPiProtocol(privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiProtocol to succeed, but got error: %s", err)
	}
	stream, err := piProtocol.Host().NewStream(context.Background(), piProtocol.ID, ProtocolID)
	if err!= nil {
		t.Errorf("Expected NewStream to succeed, but got error: %s", err)
	}
	msg := []byte("Hello, world!")
	err = writeMessage(stream, msg)
	if err!= nil {
		t.Errorf("Expected writeMessage to succeed, but got error: %s", err)
	}
	// Verify that the message was written correctly...
}

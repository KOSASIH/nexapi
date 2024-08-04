package node

import (
	"context"
	"testing"
)

func TestNewPiNode(t *testing.T) {
	ctx := context.Background()
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	node, err := NewPiNode(ctx, privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiNode to succeed, but got error: %s", err)
	}
	if node.Host == nil {
		t.Errorf("Expected node.Host to be non-nil")
	}
	if node.privateKey == nil {
		t.Errorf("Expected node.privateKey to be non-nil")
	}
}

func TestStart(t *testing.T) {
	ctx := context.Background()
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	node, err := NewPiNode(ctx, privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiNode to succeed, but got error: %s", err)
	}
	err = node.Start(ctx)
	if err!= nil {
		t.Errorf("Expected node.Start to succeed, but got error: %s", err)
	}
}

func TestSend(t *testing.T) {
	ctx := context.Background()
	privateKey1, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	node1, err := NewPiNode(ctx, privateKey1)
	if err!= nil {
		t.Errorf("Expected NewPiNode to succeed, but got error: %s", err)
	}
	err = node1.Start(ctx)
	if err!= nil {
		t.Errorf("Expected node1.Start to succeed, but got error: %s", err)
	}
	privateKey2, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	node2, err := NewPiNode(ctx, privateKey2)
	if err!= nil {
		t.Errorf("Expected NewPiNode to succeed, but got error: %s", err)
	}
	err = node2.Start(ctx)
	if err!= nil {
		t.Errorf("Expected node2.Start to succeed, but got error: %s", err)
	}
	message := []byte("Hello, world!")
	err = node1.Send(ctx, node2.peerInfo.ID, message)
	if err!= nil {
		t.Errorf("Expected node1.Send to succeed, but got error: %s", err)
	}
}

func TestMarshalJSON(t *testing.T) {
	ctx := context.Background()
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	node, err := NewPiNode(ctx, privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiNode to succeed, but got error: %s", err)
	}
	jsonData, err := node.MarshalJSON()
	if err!= nil {
		t.Errorf("Expected node.MarshalJSON to succeed, but got error: %s", err)
	}
	var jsonNode struct {
		PublicKey string `json:"public_key"`
		PeerInfo  string `json:"peer_info"`
	}
	err = json.Unmarshal(jsonData, &jsonNode)
	if err!= nil {
		t.Errorf("Expected json.Unmarshal to succeed, but got error: %s", err)
	}
	if jsonNode.PublicKey!= hex.EncodeToString(node.publicKey.X.Bytes()) {
		t.Errorf("Expected jsonNode.PublicKey to match node.publicKey.X")
	}
	if jsonNode.PeerInfo!= node.peerInfo.String() {
		t.Errorf("Expected jsonNode.PeerInfo to match node.peerInfo")
	}
}

func TestUnmarshalJSON(t *testing.T) {
	ctx := context.Background()
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	node, err := NewPiNode(ctx, privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiNode to succeed, but got error: %s", err)
	}
	jsonData, err := node.MarshalJSON()
	if err!= nil {
		t.Errorf("Expected node.MarshalJSON to succeed, but got error: %s", err)
	}
	unmarshaledNode, err := UnmarshalJSON(jsonData)
	if err!= nil {
		t.Errorf("Expected UnmarshalJSON to succeed, but got error: %s", err)
	}
	if unmarshaledNode.publicKey.X.Cmp(node.publicKey.X)!= 0 {
		t.Errorf("Expected unmarshaledNode.publicKey.X to match node.publicKey.X")
	}
	if unmarshaledNode.peerInfo.ID!= node.peerInfo.ID {
		t.Errorf("Expected unmarshaledNode.peerInfo.ID to match node.peerInfo.ID")
	}
}

func TestHandleStream(t *testing.T) {
	ctx := context.Background()
	privateKey, _, err := GenerateKeyPair()
	if err!= nil {
		t.Errorf("Expected GenerateKeyPair to succeed, but got error: %s", err)
	}
	node, err := NewPiNode(ctx, privateKey)
	if err!= nil {
		t.Errorf("Expected NewPiNode to succeed, but got error: %s", err)
	}
	err = node.Start(ctx)
	if err!= nil {
		t.Errorf("Expected node.Start to succeed, but got error: %s", err)
	}
	stream, err := node.Host.NewStream(ctx, node.peerInfo.ID, protocol.ID("/pi/1.0.0"))
	if err!= nil {
		t.Errorf("Expected node.Host.NewStream to succeed, but got error: %s", err)
	}
	go node.handleStream(stream)
	time.Sleep(100 * time.Millisecond)
	// Verify that the stream was handled correctly...
}
	

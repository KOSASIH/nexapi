package node

import (
	"testing"

	"github.com/pi-network/pi-node/types"
	"github.com/pi-network/pi-node/utils"
)

func TestNewPiNode(t *testing.T) {
	config := &types.Config{
		Port: 8080,
	}

	node, err := NewPiNode(config)
	if err != nil {
		t.Fatal(err)
	}

	if node.config.Port != 8080 {
		t.Errorf("Expected port to be 8080, but got %d", node.config.Port)
	}

	if node.privateKey == nil {
		t.Errorf("Expected private key to be generated, but got nil")
	}

	if node.publicKey == nil {
		t.Errorf("Expected public key to be generated, but got nil")
	}

	if node.address == "" {
		t.Errorf("Expected address to be generated, but got empty string")
	}
}

func TestStartPiNode(t *testing.T) {
	config := &types.Config{
		Port: 8080,
	}

	node, err := NewPiNode(config)
	if err != nil {
		t.Fatal(err)
	}

	err = node.Start()
	if err != nil {
		t.Fatal(err)
	}

	// Check if node is listening on port 8080
	resp, err := http.Get("http://localhost:8080/api/v1/node")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code to be 200, but got %d", resp.StatusCode)
	}
}

func TestStopPiNode(t *testing.T) {
	config := &types.Config{
		Port: 8080,
	}

	node, err := NewPiNode(config)
	if err != nil {
		t.Fatal(err)
	}

	err = node.Start()
	if err != nil {
		t.Fatal(err)
	}

	err = node.Stop()
	if err != nil {
		t.Fatal(err)
	}

	// Check if node is no longer listening on port 8080
	resp, err := http.Get("http://localhost:8080/api/v1/node")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	if resp != nil {
		defer resp.Body.Close()
	}
}

func TestHandleNodeRequest(t *testing.T) {
	config := &types.Config{
		Port: 8080,
	}

	node, err := NewPiNode(config)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/node", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	node.handleNodeRequest(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code to be 200, but got %d", w.Code)
	}

	var nodeInfo types.NodeInfo
	err = json.Unmarshal(w.Body.Bytes(), &nodeInfo)
	if err != nil {
		t.Fatal(err)
	}

	if nodeInfo.Address != node.address {
		t.Errorf("Expected address to be %s, but got %s", node.address, nodeInfo.Address)
	}

	if nodeInfo.Port != node.port {
		t.Errorf("Expected port to be %d, but got %d", node.port, nodeInfo.Port)
	}
}

func TestHandleTransactionsRequest(t *testing.T) {
	config := &types.Config{
		Port: 8080,
	}

	node, err := NewPiNode(config)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/v1/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	node.handleTransactionsRequest(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code to be 200, but got %d", w.Code)
	}

	var transactions []*types.Transaction
	err = json.Unmarshal(w.Body.Bytes(), &transactions)
	if err != nil {
		t.Fatal(err)
	}

	if len(transactions) != 0 {
		t.Errorf("Expected transactions to be empty, but got %d transactions", len(transactions))
	}
}

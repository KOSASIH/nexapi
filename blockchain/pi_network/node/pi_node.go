package node

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/pi-network/pi-node/types"
	"github.com/pi-network/pi-node/utils"
)

// PiNode represents a node in the Pi Network
type PiNode struct {
	// Node configuration
	config *types.Config

	// Node private key
	privateKey *ecdsa.PrivateKey

	// Node public key
	publicKey *ecdsa.PublicKey

	// Node address
	address string

	// Node port
	port int

	// Node router
	router *mux.Router

	// Node server
	server *http.Server

	// Node mutex
	mutex sync.Mutex
}

// NewPiNode creates a new Pi Node
func NewPiNode(config *types.Config) (*PiNode, error) {
	privateKey, err := utils.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}

	publicKey, err := utils.GeneratePublicKey(privateKey)
	if err != nil {
		return nil, err
	}

	address, err := utils.GenerateAddress(publicKey)
	if err != nil {
		return nil, err
	}

	node := &PiNode{
		config:     config,
		privateKey: privateKey,
		publicKey:  publicKey,
		address:    address,
		port:       config.Port,
		router:     mux.NewRouter(),
	}

	node.router.HandleFunc("/api/v1/node", node.handleNodeRequest).Methods("GET")
	node.router.HandleFunc("/api/v1/transactions", node.handleTransactionsRequest).Methods("POST")

	node.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", node.port),
		Handler: node.router,
	}

	return node, nil
}

// Start starts the Pi Node
func (n *PiNode) Start() error {
	log.Println("Starting Pi Node...")

	n.mutex.Lock()
	defer n.mutex.Unlock()

	err := n.server.ListenAndServe()
	if err != nil {
		return err
	}

	log.Println("Pi Node started successfully!")

	return nil
}

// Stop stops the Pi Node
func (n *PiNode) Stop() error {
	log.Println("Stopping Pi Node...")

	n.mutex.Lock()
	defer n.mutex.Unlock()

	err := n.server.Close()
	if err != nil {
		return err
	}

	log.Println("Pi Node stopped successfully!")

	return nil
}

// handleNodeRequest handles requests to the /api/v1/node endpoint
func (n *PiNode) handleNodeRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to /api/v1/node")

	nodeInfo := &types.NodeInfo{
		Address: n.address,
		Port:    n.port,
	}

	json.NewEncoder(w).Encode(nodeInfo)
}

// handleTransactionsRequest handles requests to the /api/v1/transactions endpoint
func (n *PiNode) handleTransactionsRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to /api/v1/transactions")

	transactions := make([]*types.Transaction, 0)

	// Process transactions here...

	json.NewEncoder(w).Encode(transactions)
}

package node

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

// PiNode represents a node in the Pi network
type PiNode struct {
	host.Host
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	peerInfo   peer.AddrInfo
}

// NewPiNode creates a new PiNode instance
func NewPiNode(ctx context.Context, privateKey *ecdsa.PrivateKey) (*PiNode, error) {
	h, err := host.New(ctx)
	if err!= nil {
		return nil, err
	}
	publicKey := &privateKey.PublicKey
	peerInfo, err := peer.AddrInfoFromP2pAddr(multiaddr.StringCast("/ip4/127.0.0.1/tcp/0"))
	if err!= nil {
		return nil, err
	}
	return &PiNode{
		Host:      h,
		privateKey: privateKey,
		publicKey:  publicKey,
		peerInfo:  peerInfo,
	}, nil
}

// Start starts the PiNode
func (n *PiNode) Start(ctx context.Context) error {
	return n.Host.SetStreamHandler(protocol.ID("/pi/1.0.0"), n.handleStream)
}

// handleStream handles incoming streams
func (n *PiNode) handleStream(s network.Stream) {
	fmt.Println("Received stream from", s.Conn().RemotePeer())
	// Handle the stream...
}

// Send sends a message to another node
func (n *PiNode) Send(ctx context.Context, to peer.ID, message []byte) error {
	s, err := n.Host.NewStream(ctx, to, protocol.ID("/pi/1.0.0"))
	if err!= nil {
		return err
	}
	_, err = s.Write(message)
	return err
}

// GenerateKeyPair generates a new key pair
func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err!= nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// MarshalJSON marshals the PiNode to JSON
func (n *PiNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PublicKey string `json:"public_key"`
		PeerInfo  string `json:"peer_info"`
	}{
		PublicKey: hex.EncodeToString(n.publicKey.X.Bytes()),
		PeerInfo:  n.peerInfo.String(),
	})
}

// UnmarshalJSON unmarshals a JSON byte slice to a PiNode
func UnmarshalJSON(data []byte) (*PiNode, error) {
	var jsonNode struct {
		PublicKey string `json:"public_key"`
		PeerInfo  string `json:"peer_info"`
	}
	err := json.Unmarshal(data, &jsonNode)
	if err!= nil {
		return nil, err
	}
	publicKeyX, err := hex.DecodeString(jsonNode.PublicKey)
	if err!= nil {
		return nil, err
	}
	publicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     publicKeyX,
		Y:     new(big.Int).SetBytes([]byte{}),
	}
	peerInfo, err := peer.AddrInfoFromP2pAddr(jsonNode.PeerInfo)
	if err!= nil {
		return nil, err
	}
	return &PiNode{
		publicKey: publicKey,
		peerInfo:  peerInfo,
	}, nil
}

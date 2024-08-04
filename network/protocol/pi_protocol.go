package protocol

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multibase"
)

const (
	ProtocolID = "/pi/1.0.0"
)

type PiProtocol struct {
	peer.ID
	privateKey *ecdsa.PrivateKey
}

func NewPiProtocol(privateKey *ecdsa.PrivateKey) (*PiProtocol, error) {
	id, err := peer.IDFromPrivateKey(privateKey)
	if err!= nil {
		return nil, err
	}
	return &PiProtocol{ID: id, privateKey: privateKey}, nil
}

func (p *PiProtocol) HandleStream(s network.Stream) {
	// Handle incoming stream
	fmt.Println("Received stream from", s.Conn().RemotePeer())
	// Read message from stream
	msg, err := readMessage(s)
	if err!= nil {
		fmt.Println("Error reading message:", err)
		return
	}
	// Process message
	p.processMessage(msg)
}

func (p *PiProtocol) processMessage(msg []byte) {
	// Unmarshal message
	var message struct {
		Type string `json:"type"`
		Data []byte `json:"data"`
	}
	err := json.Unmarshal(msg, &message)
	if err!= nil {
		fmt.Println("Error unmarshaling message:", err)
		return
	}
	// Handle message based on type
	switch message.Type {
	case "hello":
		fmt.Println("Received hello message from", p.ID)
	case "data":
		fmt.Println("Received data message from", p.ID)
		// Process data message
		p.processDataMessage(message.Data)
	default:
		fmt.Println("Unknown message type:", message.Type)
	}
}

func (p *PiProtocol) processDataMessage(data []byte) {
	// Process data message
	fmt.Println("Processing data message...")
	//...
}

func readMessage(s network.Stream) ([]byte, error) {
	// Read message length
	var length uint32
	err := binary.Read(s, binary.BigEndian, &length)
	if err!= nil {
		return nil, err
	}
	// Read message data
	msg := make([]byte, length)
	_, err = io.ReadFull(s, msg)
	if err!= nil {
		return nil, err
	}
	return msg, nil
}

func (p *PiProtocol) SendMessage(ctx context.Context, peerID peer.ID, msg []byte) error {
	// Create a new stream to the peer
	s, err := p.Host().NewStream(ctx, peerID, ProtocolID)
	if err!= nil {
		return err
	}
	// Write message to stream
	err = writeMessage(s, msg)
	if err!= nil {
		return err
	}
	return nil
}

func writeMessage(s network.Stream, msg []byte) error {
	// Write message length
	err := binary.Write(s, binary.BigEndian, uint32(len(msg)))
	if err!= nil {
		return err
	}
	// Write message data
	_, err = s.Write(msg)
	if err!= nil {
		return err
	}
	return nil
}

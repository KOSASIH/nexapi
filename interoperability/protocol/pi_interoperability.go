package protocol

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// PIInteroperability represents the PI Interoperability protocol
type PIInteroperability struct {
	conn *websocket.Conn
}

// NewPIInteroperability creates a new PI Interoperability instance
func NewPIInteroperability(conn *websocket.Conn) *PIInteroperability {
	return &PIInteroperability{conn: conn}
}

// SendMessage sends a message to the PI Interoperability endpoint
func (pi *PIInteroperability) SendMessage(message *PIMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return pi.conn.WriteMessage(websocket.TextMessage, data)
}

// ReceiveMessage receives a message from the PI Interoperability endpoint
func (pi *PIInteroperability) ReceiveMessage() (*PIMessage, error) {
	_, data, err := pi.conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	var message PIMessage
	err = json.Unmarshal(data, &message)
	return &message, err
}

// Close closes the PI Interoperability connection
func (pi *PIInteroperability) Close() error {
	return pi.conn.Close()
}

// PIMessage represents a PI Interoperability message
type PIMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// UnmarshalPayload unmarshals the payload of a PI Interoperability message
func (m *PIMessage) UnmarshalPayload(v interface{}) error {
	return json.Unmarshal(m.Payload, v)
}

// MarshalPayload marshals the payload of a PI Interoperability message
func (m *PIMessage) MarshalPayload(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	m.Payload = data
	return nil
}

// PIProtoMessage represents a PI Interoperability protocol buffer message
type PIProtoMessage struct {
	Type    string          `protobuf:"bytes,1,opt,name=type"`
	Payload []byte          `protobuf:"bytes,2,opt,name=payload"`
}

// Marshal marshals a PI Interoperability protocol buffer message
func (m *PIProtoMessage) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

// Unmarshal unmarshals a PI Interoperability protocol buffer message
func (m *PIProtoMessage) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, m)
}

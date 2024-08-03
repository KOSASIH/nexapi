package protocol

import (
	"testing"
)

func TestPIInteroperability_SendMessage(t *testing.T) {
	conn := &websocket.Conn{}
.pi := NewPIInteroperability(conn)
	message := &PIMessage{Type: "hello", Payload: []byte("world")}
	err := pi.SendMessage(message)
	if err != nil {
		t.Errorf("Expected SendMessage to succeed, but got error: %s", err)
	}
}

func TestPIInteroperability_ReceiveMessage(t *testing.T) {
	conn := &websocket.Conn{}
.pi := NewPIInteroperability(conn)
	message := &PIMessage{Type: "hello", Payload: []byte("world")}
	err := pi.conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		t.Errorf("Expected WriteMessage to succeed, but got error: %s", err)
	}
	receivedMessage, err := pi.ReceiveMessage()
	if err != nil {
		t.Errorf("Expected ReceiveMessage to succeed, but got error: %s", err)
	}
	if receivedMessage.Type != message.Type {
		t.Errorf("Expected received message type to match sent message type, but got %s", receivedMessage.Type)
	}
}

func TestPIInteroperability_Close(t *testing.T) {
	conn := &websocket.Conn{}
.pi := NewPIInteroperability(conn)
	err := pi.Close()
	if err != nil {
		t.Errorf("Expected Close to succeed, but got error: %s", err)
	}
}

func TestPIMessage_UnmarshalPayload(t *testing.T) {
	message := &PIMessage{Payload: []byte(`{"key": "value"}`)}
	var payload struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	err := message.UnmarshalPayload(&payload)
	if err != nil {
		t.Errorf("Expected UnmarshalPayload to succeed, but got error: %s", err)
	}
	if payload.Key != "value" {
		t.Errorf("Expected payload key to be 'value', but got %s", payload.Key)
	}
}

func TestPIMessage_MarshalPayload(t *testing.T) {
	message := &PIMessage{}
	payload := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{Key: "value", Value: "hello"}
	err := message.MarshalPayload(payload)
	if err!= nil {
		t.Errorf("Expected MarshalPayload to succeed, but got error: %s", err)
	}
	if string(message.Payload)!= `{"key":"value","value":"hello"}` {
		t.Errorf("Expected payload to be marshaled correctly, but got %s", message.Payload)
	}
}

func TestPIProtoMessage_Marshal(t *testing.T) {
	message := &PIProtoMessage{Type: "hello", Payload: []byte("world")}
	data, err := message.Marshal()
	if err!= nil {
		t.Errorf("Expected Marshal to succeed, but got error: %s", err)
	}
	if len(data) == 0 {
		t.Errorf("Expected marshaled data to be non-empty")
	}
}

func TestPIProtoMessage_Unmarshal(t *testing.T) {
	data := []byte{0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x05, 0x77, 0x6f, 0x72, 0x6c, 0x64}
	message := &PIProtoMessage{}
	err := message.Unmarshal(data)
	if err!= nil {
		t.Errorf("Expected Unmarshal to succeed, but got error: %s", err)
	}
	if message.Type!= "hello" {
		t.Errorf("Expected unmarshaled type to be 'hello', but got %s", message.Type)
	}
	if string(message.Payload)!= "world" {
		t.Errorf("Expected unmarshaled payload to be 'world', but got %s", message.Payload)
	}
}

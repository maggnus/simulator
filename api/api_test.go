package api

import (
	"simulator/configs"
	"testing"
)

func TestNewClient(t *testing.T) {
	config := configs.Config{
		Server: configs.ServerConfig{
			Host: "localhost",
			Port: 8080,
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
	}

	if client.host != "localhost" {
		t.Errorf("Expected host to be 'localhost', but got %v", client.host)
	}

	if client.port != 8080 {
		t.Errorf("Expected port to be 8080, but got %v", client.port)
	}
}

func TestNewMessage(t *testing.T) {
	message := NewMessage("test_sender", "test_instruction")

	if message.Sender != "test_sender" {
		t.Errorf("Expected sender to be 'test_sender', but got %v", message.Sender)
	}

	if message.Instruction != "test_instruction" {
		t.Errorf("Expected instruction to be 'test_instruction', but got %v", message.Instruction)
	}
}

func TestNewMessageData(t *testing.T) {
	jsonData := []byte(`{"sender":"test_sender","instruction":"test_instruction"}`)

	message, err := NewMessageData(jsonData)
	if err != nil {
		t.Errorf("NewMessageData() error = %v", err)
	}

	if message.Sender != "test_sender" {
		t.Errorf("Expected sender to be 'test_sender', but got %v", message.Sender)
	}

	if message.Instruction != "test_instruction" {
		t.Errorf("Expected instruction to be 'test_instruction', but got %v", message.Instruction)
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		name    string
		message Message
		wantErr bool
	}{
		{
			name: "Valid message",
			message: Message{Sender: "test_sender", Instruction: "test_instruction"},
			wantErr: false,
		},
		{
			name: "Empty sender",
			message: Message{Sender: "", Instruction: "test_instruction"},
			wantErr: true,
		},
		{
			name: "Empty instruction",
			message: Message{Sender: "test_sender", Instruction: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.message.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

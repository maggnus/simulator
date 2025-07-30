package client

import (
	"net"
	"simulator/configs"
	"testing"
)

func TestNewClient(t *testing.T) {
	config := configs.Config{}

	client, err := NewClient(config)
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
	}

	if client == nil {
		t.Error("Expected client to be non-nil")
	}
}

func TestConnectAndClose(t *testing.T) {
	// Create a mock server
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	go func() {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
	}()

	config := configs.Config{
		App: configs.AppConfig{
			Host: "localhost",
			Port: l.Addr().(*net.TCPAddr).Port,
		},
	}

	client, err := NewClient(config)
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
	}

	err = client.Connect()
	if err != nil {
		t.Errorf("Connect() error = %v", err)
	}

	err = client.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}
}
package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"simulator/api"
	"simulator/configs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	config := configs.Config{}
	server, err := NewServer(config)
	assert.NoError(t, err)
	assert.NotNil(t, server)
}

func TestHandleClient(t *testing.T) {
	// Create a mock connection
	serverConn, clientConn := net.Pipe()

	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)

	go handleClient(serverConn, 1024)

	// Send a message from the client side
	message := api.NewMessage("test_sender", "test_instruction")
	jsonMessage, _ := message.Marshal()
	_, err := clientConn.Write(jsonMessage)
	assert.NoError(t, err)

	// Close the client connection to signal EOF
	clientConn.Close()

	// Wait for handleClient to finish processing
	time.Sleep(100 * time.Millisecond)

	assert.Contains(t, logBuffer.String(), "received")
	assert.Contains(t, logBuffer.String(), "test_sender")
	assert.Contains(t, logBuffer.String(), "test_instruction")
}

func TestServerRun(t *testing.T) {
	config := configs.Config{
		App: configs.AppConfig{
			Host: "localhost",
			Port: 10001, // Use a different port for testing
			ReciveBuffer: 1024,
		},
	}
	server, err := NewServer(config)
	assert.NoError(t, err)

	go func() {
		server.Run()
	}()

	// Wait for the server to start listening
	time.Sleep(100 * time.Millisecond)

	// Connect a client to the server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.App.Host, config.App.Port))
	assert.NoError(t, err)
	defer conn.Close()

	// Send a message
	message := api.NewMessage("client_sender", "client_instruction")
	jsonMessage, _ := message.Marshal()
	_, err = conn.Write(jsonMessage)
	assert.NoError(t, err)

	// Read response (optional, depending on server behavior)
	// buf := make([]byte, 1024)
	// n, err := conn.Read(buf)
	// assert.NoError(t, err)
	// assert.Contains(t, string(buf[:n]), "OK")

	// Give server time to process
	time.Sleep(100 * time.Millisecond)
}

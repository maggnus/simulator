package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"simulator/configs"
	"syscall"
)

type Client struct {
	config     configs.Config
	connection net.Conn
	ConnectFunc func() error // For testing purposes
}

func NewClient(config configs.Config) (*Client, error) {
	return &Client{
		config: config,
	}, nil
}

func (c *Client) Connect() error {
	if c.ConnectFunc != nil {
		return c.ConnectFunc()
	}
	addr := fmt.Sprintf("%v:%v", c.config.App.Host, c.config.App.Port)
	log.Printf("connect to: %v", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	c.connection = conn

	return nil
}

func (c *Client) Close() error {
	if c.connection == nil {
		return nil
	}

	return c.connection.Close()
}

// https://stackoverflow.com/questions/12741386/how-to-know-tcp-connection-is-closed-in-net-package
func (c *Client) isConnected() bool {
	f, err := c.connection.(*net.TCPConn).File()
	if err != nil {
		return false
	}

	b := []byte{0}
	_, _, err = syscall.Recvfrom(int(f.Fd()), b, syscall.MSG_PEEK|syscall.MSG_DONTWAIT)
	return err != nil
}

func (c *Client) Send(message interface{}) error {
	if c.connection == nil {
		return fmt.Errorf("wrong connection type nil")
	}

	if !c.isConnected() {
		log.Printf("connection failed. reconnect...")
		return c.Connect()
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Set the send buffer size to 8192 bytes (8KB)
	err = c.connection.(*net.TCPConn).SetWriteBuffer(c.config.App.SendBuffer)
	if err != nil {
		return fmt.Errorf("error setting send buffer size: %v", err)
	}

	n, err := c.connection.Write(data)
	if err != nil {
		return fmt.Errorf("error writing data: %v", err)
		os.Exit(1)
	}

	log.Printf("sent %d bytes: %+v", n, message)

	return nil
}
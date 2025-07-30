package server

import (
	"fmt"
	"log"
	"net"
	"simulator/configs"
)

type Server struct {
	config configs.Config
}

func NewServer(config configs.Config) (*Server, error) {
	return &Server{
		config: config,
	}, nil
}

func (c *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", c.config.App.Host, c.config.App.Port)

	log.Printf("start listener: %s", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	defer func() {
		err := listener.Close()
		log.Fatalf("exit: %v", err)
	}()

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error: %v", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn, c.config.App.ReciveBuffer)
	}
}
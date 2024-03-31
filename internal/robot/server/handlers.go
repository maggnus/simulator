package server

import (
	"io"
	"log"
	"net"
	"simulator/api"
)

func handleClient(conn net.Conn, rcvbuf int) {
	defer conn.Close()

	log.Printf("connection from: %+v", conn.RemoteAddr())

	// Buffer to store received data
	buf := make([]byte, rcvbuf)

	// Read data from the connection in a loop
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("connection closed by server")
				break
			} else {
				log.Printf("error reading: %v", err)
				break
			}
		}

		message, err := api.NewMessageData(buf[:n])
		if err != nil {
			log.Printf("error: %v", err.Error())
			break
		}

		log.Printf("received %d bytes: %+v", n, message)

	}
}

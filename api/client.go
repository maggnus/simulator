package api

import (
	"fmt"
	"io"
	"net/http"
	"simulator/configs"
)

type Client struct {
	host string
	port int
}

func NewClient(config configs.Config) (*Client, error) {
	client := &Client{
		host: config.Server.Host,
		port: config.Server.Port,
	}

	return client, nil
}

func (c *Client) Command(cmd Command) error {
	addr := fmt.Sprintf("http://%v:%v/%v", c.host, c.port, cmd)
	req, err := http.NewRequest("POST", addr, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers (if needed)
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP Client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Printf("%v", string(body))

	return nil
}

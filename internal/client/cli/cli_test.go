package cli

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"simulator/api"
	"simulator/configs"
	"testing"
)

func TestInstructionCmd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	config := configs.Config{
		Server: configs.ServerConfig{
			Host: "localhost",
			Port: 8080,
		},
	}

	client, err := api.NewClient(config)
	if err != nil {
		t.Errorf("NewClient() error = %v", err)
	}

	Client = client

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "Valid instruction",
			args:    []string{"start"},
			wantErr: false,
		},
		{
			name:    "Invalid instruction",
			args:    []string{"invalid"},
			wantErr: true,
		},
		{
			name:    "No instruction",
			args:    []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs(append([]string{"instruction"}, tt.args...))
			if err := rootCmd.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebCmd(t *testing.T) {
	// Redirect log output to a buffer
	var buf bytes.Buffer
	log.SetOutput(&buf)

	go func() {
		rootCmd.SetArgs([]string{"web"})
		rootCmd.Execute()
	}()

	// Give the server time to start
	//time.Sleep(1 * time.Second)

	// Check if the server is running
	//resp, err := http.Get("http://localhost:8080")
	//if err != nil {
	//	t.Errorf("http.Get() error = %v", err)
	//}
	//
	//if resp.StatusCode != http.StatusOK {
	//	t.Errorf("Expected status code %v, but got %v", http.StatusOK, resp.StatusCode)
	//}
}

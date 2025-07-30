package web

import (
	"net/http"
	"net/http/httptest"
	"simulator/api"
	"simulator/configs"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(3)

	if len(board.Cells) != 3 {
		t.Errorf("Expected board size to be 3, but got %v", len(board.Cells))
	}

	if board.S != happy {
		t.Errorf("Expected board symbol to be %v, but got %v", happy, board.S)
	}
}

func TestPrint(t *testing.T) {
	board := NewBoard(3)
	board.Print()
}

func TestHandlers(t *testing.T) {
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
		name       string
		path       string
		statusCode int
	}{
		{
			name:       "Root",
			path:       "/",
			statusCode: http.StatusOK,
		},
		{
			name:       "Start",
			path:       "/start",
			statusCode: http.StatusOK,
		},
		{
			name:       "Stop",
			path:       "/stop",
			statusCode: http.StatusOK,
		},
		{
			name:       "Right",
			path:       "/right",
			statusCode: http.StatusOK,
		},
		{
			name:       "Left",
			path:       "/left",
			statusCode: http.StatusOK,
		},
		{
			name:       "Forward",
			path:       "/forward",
			statusCode: http.StatusOK,
		},
		{
			name:       "Back",
			path:       "/back",
			statusCode: http.StatusOK,
		},
	}

	router := SetupRoutes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.statusCode)
			}
		})
	}
}

package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"simulator/configs"
	"simulator/internal/controller/client"
	"simulator/internal/controller/db"
	"simulator/internal/controller/middleware"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"
)

func TestServerRun(t *testing.T) {
	// Create a real db.DB instance for testing with an in-memory SQLite database
	testConfig := configs.Config{
		Database: configs.DatabaseConfig{
			Driver: "sqlite3",
			Source: "file:ent?mode=memory&cache=shared&_fk=1",
		},
	}
	realDB, err := db.NewDB(testConfig)
	assert.NoError(t, err)
	defer realDB.Client.Close()

	// Create a real client.Client instance and mock its Connect method
	clientConfig := configs.Config{
		App: configs.AppConfig{Host: "localhost", Port: 10000},
	}
	realClient, err := client.NewClient(clientConfig)
	assert.NoError(t, err)
	realClient.ConnectFunc = func() error { return nil } // Mock Connect

	// Create a test server instance
	config := configs.Config{
		Server: configs.ServerConfig{
			Host: "localhost",
			Port: 8080,
		},
	}
	server := &Server{
		config: config,
		db:     realDB,
		client: realClient,
	}

	// Create a new Echo instance for testing
	e := echo.New()
	e.HideBanner = true

	// Register handlers and middleware
		dbLogger := middleware.DBLogger(server.db, nil)
	e.POST("/start", startHandler, dbLogger)
	e.POST("/stop", stopHandler, dbLogger)
	e.POST("/left", leftHandler, dbLogger)
	e.POST("/right", rightHandler, dbLogger)
	e.POST("/forward", forwardHandler, dbLogger)
	e.POST("/back", backHandler, dbLogger)

	// Create a test HTTP server
	ts := httptest.NewServer(e)
	defer ts.Close()

	// Test each endpoint
	endpoints := []string{"start", "stop", "left", "right", "forward", "back"}
	for _, endpoint := range endpoints {
		resp, err := http.Post(fmt.Sprintf("%s/%s", ts.URL, endpoint), "application/json", nil)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		resp.Body.Close()
	}
}

package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"simulator/configs"
	"simulator/internal/controller/client"
	"simulator/internal/controller/db"
	"simulator/internal/controller/middleware"
)

type Server struct {
	config configs.Config
	db     *db.DB
	client *client.Client
}

func NewServer(config configs.Config) (*Server, error) {
	log.Printf("load config: %v", config)

	db, err := db.NewDB(config)
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err.Error())
	}

	client, err := client.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("client error: %v", err.Error())
	}

	err = client.Connect()
	if err != nil {
		return nil, fmt.Errorf("client connection error: %v", err.Error())
	}

	return &Server{
		config: config,
		db:     db,
		client: client,
	}, nil
}

func (c *Server) Run() error {
	addr := fmt.Sprintf("%v:%v", c.config.Server.Host, c.config.Server.Port)

	log.Printf("start server: %s", addr)

	e := echo.New()
	e.HideBanner = true

	// DB Logger middleware
	dbLogger := middleware.DBLogger(c.db, c.client)

	e.POST("/start", startHandler, dbLogger)
	e.POST("/stop", stopHandler, dbLogger)
	e.POST("/left", leftHandler, dbLogger)
	e.POST("/right", rightHandler, dbLogger)
	e.POST("/forward", forwardHandler, dbLogger)
	e.POST("/back", backHandler, dbLogger)

	return e.Start(addr)
}

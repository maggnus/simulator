package main

import (
	"log"
	"simulator/configs"
	"simulator/internal/robot/server"
)

// TODO: must be passed as an argument
const (
	CONFIG_FILE = "configs/config.yml"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	srv, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	err = srv.Run()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"log"
	"os"
	"simulator/api"
	"simulator/configs"
	"simulator/internal/client/cli"
	"simulator/internal/client/web"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	cli.Client = client
	web.Client = client

	err = cli.Execute()
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

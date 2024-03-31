package db

import (
	"context"
	"fmt"
	"simulator/configs"
	"simulator/ent"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*ent.Client
}

func NewDB(config configs.Config) (*DB, error) {

	client, err := ent.Open(config.Database.Driver, config.Database.Source)
	if err != nil {
		return nil, err
	}

	db := &DB{client}

	// Generate database schema for files in ./ent/schema directory
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return db, nil
}

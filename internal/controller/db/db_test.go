package db

import (
	"simulator/configs"
	"testing"
)

func TestNewDB(t *testing.T) {
	config := configs.Config{
		Database: struct {
			Driver string `yaml:"driver"`
			Source string `yaml:"source"`
		}{
			Driver: "sqlite3",
			Source: "file:ent?mode=memory&cache=shared&_fk=1",
		},
	}

	db, err := NewDB(config)
	if err != nil {
		t.Errorf("NewDB() error = %v", err)
	}

	if db.Client == nil {
		t.Error("Expected db.Client to be non-nil")
	}
}

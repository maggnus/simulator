package configs

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	ConfigFile = "config_test.yml"

	config, err := LoadConfig()
	if err != nil {
		t.Errorf("LoadConfig() error = %v", err)
	}

	if config.Server.Host != "localhost" {
		t.Errorf("Expected server host to be 'localhost', but got %v", config.Server.Host)
	}

	if config.Server.Port != 8080 {
		t.Errorf("Expected server port to be 8080, but got %v", config.Server.Port)
	}

	if config.Database.Driver != "sqlite3" {
		t.Errorf("Expected database driver to be 'sqlite3', but got %v", config.Database.Driver)
	}

	if config.Database.Source != "file:ent?mode=memory&cache=shared&_fk=1" {
		t.Errorf("Expected database source to be 'file:ent?mode=memory&cache=shared&_fk=1', but got %v", config.Database.Source)
	}

	if config.App.Host != "localhost" {
		t.Errorf("Expected config host to be 'localhost', but got %v", config.App.Host)
	}

	if config.App.Port != 10000 {
		t.Errorf("Expected config port to be 10000, but got %v", config.App.Port)
	}

	if config.App.SendBuffer != 8192 {
		t.Errorf("Expected send buffer to be 8192, but got %v", config.App.SendBuffer)
	}

	if config.App.ReciveBuffer != 8192 {
		t.Errorf("Expected receive buffer to be 8192, but got %v", config.App.ReciveBuffer)
	}
}

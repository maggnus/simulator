package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var ConfigFile = "configs/config.yml"

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

type AppConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	SendBuffer   int    `yaml:"sndbuf"`
	ReciveBuffer int    `yaml:"rcvbuf"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	App      AppConfig      `yaml:"config"`
}

func LoadConfig() (Config, error) {
	var config Config

	f, err := os.ReadFile(ConfigFile)
	if err != nil {
		return config, fmt.Errorf("failed to read configuration file: %v", err.Error())
	}

	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return config, fmt.Errorf("error parsing configuration file: %v", err.Error())
	}

	return config, nil
}
package config

import (
	"encoding/json"
	"io"
	"os"
)

type MongoConfig struct {
	URI      string
	Database string
}

type Config struct {
	MongoConfig   MongoConfig
	ServerAddress string
}

func LoadConfig() (*Config, error) {
	filename := "config.json"
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

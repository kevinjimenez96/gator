package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

var Cfg = Config{}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username

	err := write(*cfg)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", userHome, configFileName), nil
}

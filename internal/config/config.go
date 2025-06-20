package config

import (
	"errors"
	"flag"
	"os"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/ilyakaznacheev/cleanenv"
)

type serverConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	Env    string       `yaml:"env"`
	Server serverConfig `yaml:"server"`
}

func New() (*Config, error) {
	configPath, err := fetchConfigPath()

	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		return nil, err
	}

	var cfg Config

	if err = cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	if cfg.Server.Port <= 0 {
		return nil, models.ErrBadConfigPort
	}

	return &cfg, nil
}

func fetchConfigPath() (string, error) {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")

		if path == "" {
			return "", errors.New("config path is required")
		}
	}

	return path, nil
}

package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrNoInputFile  = errors.New("input-file is empty in config")
	ErrNoOutputFile = errors.New("output-file is empty in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("decode config yaml: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, ErrNoInputFile
	}

	if cfg.OutputFile == "" {
		return nil, ErrNoOutputFile
	}

	return &cfg, nil
}

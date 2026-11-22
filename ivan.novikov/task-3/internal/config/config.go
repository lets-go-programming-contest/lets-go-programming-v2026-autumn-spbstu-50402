package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyConfigPath = errors.New("config file path not specified")
	ErrInvalidConfig   = errors.New("input or output file path not specified in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	if strings.TrimSpace(path) == "" {
		return Config{}, ErrEmptyConfigPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshal yaml config: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Config{}, ErrInvalidConfig
	}

	return cfg, nil
}

package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrMissingField = errors.New("missing required field")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("decode config file: %w", err)
	}

	if cfg.InputFile == "" {
		return Config{}, fmt.Errorf("%w: input-file", ErrMissingField)
	}

	if cfg.OutputFile == "" {
		return Config{}, fmt.Errorf("%w: output-file", ErrMissingField)
	}

	return cfg, nil
}

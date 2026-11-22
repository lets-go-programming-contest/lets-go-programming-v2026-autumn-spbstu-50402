package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyConfigPath = errors.New("config path is empty")
	ErrEmptyInputFile  = errors.New("input-file is required in config")
	ErrEmptyOutputFile = errors.New("output-file is required in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	if strings.TrimSpace(path) == "" {
		return nil, ErrEmptyConfigPath
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config file %q: %w", path, err)
	}
	defer file.Close()

	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode yaml config %q: %w", path, err)
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("validate config %q: %w", path, err)
	}

	return &cfg, nil
}

func (c *Config) validate() error {
	if strings.TrimSpace(c.InputFile) == "" {
		return ErrEmptyInputFile
	}

	if strings.TrimSpace(c.OutputFile) == "" {
		return ErrEmptyOutputFile
	}

	return nil
}

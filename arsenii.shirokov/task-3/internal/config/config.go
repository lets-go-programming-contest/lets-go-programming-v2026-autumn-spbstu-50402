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

	var cfg Config
	if err := decodeFile(path, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("validate config %q: %w", path, err)
	}

	return &cfg, nil
}

func decodeFile(path string, cfg *Config) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open config file %q: %w", path, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close config file %q: %w", path, closeErr)
		}
	}()

	if decodeErr := yaml.NewDecoder(file).Decode(cfg); decodeErr != nil {
		return fmt.Errorf("decode yaml config %q: %w", path, decodeErr)
	}

	return nil
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

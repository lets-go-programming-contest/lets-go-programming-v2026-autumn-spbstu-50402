// Package config reads and validates the application's YAML configuration.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	// ErrEmptyConfigPath means the YAML configuration path is blank.
	ErrEmptyConfigPath = errors.New("config path is empty")
	// ErrEmptyInputFile means input-file is missing from the configuration.
	ErrEmptyInputFile  = errors.New("input-file is required in config")
	// ErrEmptyOutputFile means output-file is missing from the configuration.
	ErrEmptyOutputFile = errors.New("output-file is required in config")
)

// Config contains paths used by the currency conversion application.
type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// Load reads a YAML configuration from path and checks its required fields.
func Load(path string) (Config, error) {
	if strings.TrimSpace(path) == "" {
		return Config{}, fmt.Errorf("validate config path: %w", ErrEmptyConfigPath)
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("open configuration %q: %w", path, err)
	}

	var cfg Config

	decoder := yaml.NewDecoder(file)

	decoder.KnownFields(true)

	decodeErr := decoder.Decode(&cfg)

	closeErr := file.Close()

	if decodeErr != nil {
		return Config{}, fmt.Errorf("decode configuration %q: %w", path, decodeErr)
	}

	if closeErr != nil {
		return Config{}, fmt.Errorf("close configuration %q: %w", path, closeErr)
	}

	if strings.TrimSpace(cfg.InputFile) == "" {
		return Config{}, fmt.Errorf("configuration %q: %w", path, ErrEmptyInputFile)
	}

	if strings.TrimSpace(cfg.OutputFile) == "" {
		return Config{}, fmt.Errorf("configuration %q: %w", path, ErrEmptyOutputFile)
	}

	return cfg, nil
}

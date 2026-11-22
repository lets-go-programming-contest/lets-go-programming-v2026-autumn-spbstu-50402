// Package config reads and validates the application's YAML configuration.
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config contains paths used by the currency conversion application.
type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// Load reads a YAML configuration from path and checks its required fields.
func Load(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("open configuration %q: %w", path, err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("decode configuration %q: %w", path, err)
	}

	if cfg.InputFile == "" {
		return Config{}, fmt.Errorf("configuration %q: input-file is required", path)
	}
	if cfg.OutputFile == "" {
		return Config{}, fmt.Errorf("configuration %q: output-file is required", path)
	}

	return cfg, nil
}

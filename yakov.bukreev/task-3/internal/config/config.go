package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read configuration file %q: %w", path, err)
	}

	var cfg Config

	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal configuration file %q: %w", path, err)
	}

	return cfg, nil
}

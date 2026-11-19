package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	errInputFileNotSet  = errors.New("input file is not specified")
	errOutputFileNotSet = errors.New("output file is not specified")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config file %q: %w", path, err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parse config file %q: %w", path, err)
	}

	err = cfg.Validate()
	if err != nil {
		return Config{}, fmt.Errorf("validate config file %q: %w", path, err)
	}

	return cfg, nil
}

func (cfg Config) Validate() error {
	if cfg.InputFile == "" {
		return errInputFileNotSet
	}

	if cfg.OutputFile == "" {
		return errOutputFileNotSet
	}

	return nil
}

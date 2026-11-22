package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errInputFile  = errors.New("unknown input file")
	errOutputFile = errors.New("unknown output file")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (cfg Config) Check() error {
	inputFile := strings.TrimSpace(cfg.InputFile)
	if inputFile == "" {
		return errInputFile
	}

	outputFile := strings.TrimSpace(cfg.OutputFile)
	if outputFile == "" {
		return errOutputFile
	}

	return nil
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read error %q: %w", path, err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parse error %q: %w", path, err)
	}

	err = cfg.Check()
	if err != nil {
		return Config{}, fmt.Errorf("check error %q: %w", path, err)
	}

	return cfg, nil
}

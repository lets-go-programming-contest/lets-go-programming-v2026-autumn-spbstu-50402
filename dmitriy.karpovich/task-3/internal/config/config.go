package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read configuration file %q: %w", path, err)
	}

	cfg, err := decode(data)
	if err != nil {
		return Config{}, fmt.Errorf("decode configuration file %q: %w", path, err)
	}
	cfg, err = validate(cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func decode(data []byte) (Config, error) {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return Config{}, errors.New("configuration file must contain exactly one YAML document")
		}
		return Config{}, fmt.Errorf("read trailing YAML document: %w", err)
	}

	return cfg, nil
}

func validate(cfg Config) (Config, error) {
	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Config{}, errors.New("required field input-file is missing")
	}
	if cfg.OutputFile == "" {
		return Config{}, errors.New("required field output-file is missing")
	}
	return cfg, nil
}

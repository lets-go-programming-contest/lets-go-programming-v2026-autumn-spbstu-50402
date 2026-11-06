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

var (
	errMultipleYAMLDocuments = errors.New("configuration file must contain exactly one YAML document")
	errMissingInputFile      = errors.New("required field input-file is missing")
	errMissingOutputFile     = errors.New("required field output-file is missing")
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

	err := decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("decode YAML configuration: %w", err)
	}

	var extra any

	err = decoder.Decode(&extra)
	if !errors.Is(err, io.EOF) {
		if err == nil {
			return Config{}, errMultipleYAMLDocuments
		}

		return Config{}, fmt.Errorf("read trailing YAML document: %w", err)
	}

	return cfg, nil
}

func validate(cfg Config) (Config, error) {
	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Config{}, errMissingInputFile
	}

	if cfg.OutputFile == "" {
		return Config{}, errMissingOutputFile
	}

	return cfg, nil
}

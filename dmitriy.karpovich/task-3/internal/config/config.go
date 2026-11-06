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

}

func validate(cfg Config) (Config, error) {

}

package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const dirPerm = 0o755

var (
	ErrInvalidConfigFlag = errors.New("invalid config flag")
	ErrEmptyConfigFiles  = errors.New("input and output files must be set in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ReadConfig() (string, error) {
	configPath := flag.String("config", "", "YAML-config path")
	flag.Parse()

	if *configPath == "" {
		return "", ErrInvalidConfigFlag
	}

	return *configPath, nil
}

func ParseConfig(configPath string) (Config, error) {
	yamlConf, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlConf, &cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshal config: %w", err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Config{}, ErrEmptyConfigFiles
	}

	if dir := filepath.Dir(cfg.OutputFile); dir != "" {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			return Config{}, fmt.Errorf("create output directory: %w", err)
		}
	}

	return cfg, nil
}

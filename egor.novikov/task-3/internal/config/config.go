package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ReadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	if cfg.InputFile == "" {
		panic("input-file is not specified in config")
	}

	if cfg.OutputFile == "" {
		panic("output-file is not specified in config")
	}

	return cfg
}

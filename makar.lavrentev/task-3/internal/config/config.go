package config

import (
	"flag"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const dirPerm = 0o755

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ReadConfig() string {
	configPath := flag.String("config", "", "YAML-config path")
	flag.Parse()

	if *configPath == "" {
		panic("invalid config flag")
	}

	return *configPath
}

func ParseConfig(configPath string) Config {
	yamlConf, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(yamlConf, &cfg); err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("input and output files must be set in config")
	}

	if dir := filepath.Dir(cfg.OutputFile); dir != "" {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			panic(err)
		}
	}

	return cfg
}

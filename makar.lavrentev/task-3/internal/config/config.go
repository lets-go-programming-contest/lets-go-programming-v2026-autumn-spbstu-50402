package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

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

	return cfg
}

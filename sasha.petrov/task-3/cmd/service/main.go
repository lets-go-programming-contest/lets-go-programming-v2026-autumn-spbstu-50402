package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func readFlags() string {
	path := flag.String("config", "config.yaml", "path to yaml config")

	flag.Parse()

	return *path
}

func main() {
	path := readFlags()

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("config: input-file and output-file are required")
	}
}

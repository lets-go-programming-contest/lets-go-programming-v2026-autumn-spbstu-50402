package main

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func main() {
	configPath := flag.String("config", "", "YAML-config path")
	flag.Parse()
	if *configPath == "" {
		log.Fatal("Invalid config flag")
	}

	yamlConf, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Config import error: %v", err)
	}
	
	var cfg Config
	err = yaml.Unmarshal(yamlConf, &cfg)
	if err != nil {
		log.Fatalf("YAML decode error: %v", err)
	}
}

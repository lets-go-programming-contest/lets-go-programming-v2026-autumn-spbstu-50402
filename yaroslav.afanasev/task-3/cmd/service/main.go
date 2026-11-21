package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/Jaroslaw-super1/task-3/internal/config"
	"github.com/Jaroslaw-super1/task-3/internal/must"
)

var errConfigPathNotSet = errors.New("config path is not specified")

func main() {
	configPath, err := readConfigPath()
	must.NoError(err)

	cfg, err := config.Load(configPath)
	must.NoError(err)

	fmt.Printf("Input: %s, Output: %s\n", cfg.InputFile, cfg.OutputFile)
}

func readConfigPath() (string, error) {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		return "", errConfigPathNotSet
	}

	return *configPath, nil
}

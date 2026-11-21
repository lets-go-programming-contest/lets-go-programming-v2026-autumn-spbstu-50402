package main

import (
	"errors"
	"flag"

	"github.com/Jaroslaw-super1/task-3/internal/must"
)

var errConfigPathNotSet = errors.New("config path is not specified")

func main() {
	configPath, err := readConfigPath()
	must.NoError(err)

	println(configPath)
	println(err)
}

func readConfigPath() (string, error) {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()
	if *configPath == "" {
		return "", errConfigPathNotSet
	}
	return *configPath, nil
}

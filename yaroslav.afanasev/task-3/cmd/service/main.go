package main

import (
	"errors"
	"flag"
)

var errConfigPathNotSet = errors.New("config path is not specified")

func main() {
	configPath, err := readConfigPath()

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

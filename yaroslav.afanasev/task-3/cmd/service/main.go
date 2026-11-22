package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/Jaroslaw-super1/task-3/internal/cbr"
	"github.com/Jaroslaw-super1/task-3/internal/config"
	"github.com/Jaroslaw-super1/task-3/internal/currency"
	"github.com/Jaroslaw-super1/task-3/internal/must"
)

var errConfigPathNotSet = errors.New("config path is not specified")

func main() {
	configPath, err := readConfigPath()
	must.NoError(err)

	cfg, err := config.Load(configPath)
	must.NoError(err)

	rates, err := cbr.ParseFile(cfg.InputFile)
	must.NoError(err)

	currency.SortByValueDescending(rates)

	fmt.Printf("Загружено валют: %d\n", len(rates))
	fmt.Printf("Топ-1: %s = %.4f\n", rates[0].CharCode, rates[0].Value)
}

func readConfigPath() (string, error) {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		return "", errConfigPathNotSet
	}

	return *configPath, nil
}

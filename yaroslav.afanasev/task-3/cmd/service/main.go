package main

import (
	"errors"
	"flag"

	"github.com/Jaroslaw-super1/task-3/internal/cbr"
	"github.com/Jaroslaw-super1/task-3/internal/config"
	"github.com/Jaroslaw-super1/task-3/internal/currency"
	"github.com/Jaroslaw-super1/task-3/internal/must"
	"github.com/Jaroslaw-super1/task-3/internal/storage"
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

	err = storage.WriteJSON(cfg.OutputFile, rates)
	must.NoError(err)

	println("Успешно! Результат сохранен в", cfg.OutputFile)
}

func readConfigPath() (string, error) {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		return "", errConfigPathNotSet
	}

	return *configPath, nil
}

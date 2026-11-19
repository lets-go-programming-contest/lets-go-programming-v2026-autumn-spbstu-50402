package main

import (
	"errors"
	"flag"

	"github.com/Alastor20/task-3/internal/cbr"
	"github.com/Alastor20/task-3/internal/config"
	"github.com/Alastor20/task-3/internal/must"
	"github.com/Alastor20/task-3/internal/output"
)

var errConfigPathNotSet = errors.New("config path is not specified")

func main() {
	configPath, err := readConfigPath()
	must.NoError(err)

	cfg, err := config.Load(configPath)
	must.NoError(err)

	currencies, err := cbr.ParseFile(cfg.InputFile)
	must.NoError(err)

	cbr.SortByValueDesc(currencies)

	must.NoError(output.WriteJSON(cfg.OutputFile, currencies))
}

func readConfigPath() (string, error) {
	configPath := flag.String("config", "", "path to config file")

	flag.Parse()

	if *configPath == "" {
		return "", errConfigPathNotSet
	}

	return *configPath, nil
}

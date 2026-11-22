package main

import (
	"flag"
	"fmt"

	"github.com/mikeee322/task-3/internal/config"
	"github.com/mikeee322/task-3/internal/currency"
	"github.com/mikeee322/task-3/internal/jsonwriter"
	"github.com/mikeee322/task-3/internal/xmlparser"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to YAML configuration file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}

func run(configPath string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	currencies, err := xmlparser.ParseFile(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("parse currencies: %w", err)
	}

	currency.SortByValueDesc(currencies)
	if err := jsonwriter.WriteFile(cfg.OutputFile, currencies); err != nil {
		return fmt.Errorf("write currencies: %w", err)
	}
	return nil
}

package main

import (
	"flag"
	"fmt"

	"arsenii.shirokov/task-3/internal/config"
	"arsenii.shirokov/task-3/internal/currency"
	"arsenii.shirokov/task-3/internal/jsonwriter"
	"arsenii.shirokov/task-3/internal/xmlparser"
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

	var valCurs currency.ValCurs
	if err := xmlparser.ParseFile(cfg.InputFile, &valCurs); err != nil {
		return fmt.Errorf("parse currencies: %w", err)
	}

	currency.SortByValueDesc(valCurs.Valutes)

	if err := jsonwriter.WriteFile(cfg.OutputFile, valCurs.Valutes); err != nil {
		return fmt.Errorf("write currencies: %w", err)
	}

	return nil
}

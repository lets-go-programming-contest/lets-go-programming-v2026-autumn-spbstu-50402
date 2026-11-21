package main

import (
	"flag"

	"github.com/R0me0Star/task-3/internal/config"
	"github.com/R0me0Star/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to YAML config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDesc(currencies)

	if err := currency.WriteJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}

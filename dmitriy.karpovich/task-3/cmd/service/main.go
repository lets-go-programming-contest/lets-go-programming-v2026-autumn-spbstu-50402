package main

import (
	"flag"


	"github.com/voyadger01/task-3/internal/cbr"
	"github.com/voyadger01/task-3/internal/config"
	"github.com/voyadger01/task-3/internal/currency"
	"github.com/voyadger01/task-3/internal/storage"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	rates, err := cbr.ReadFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDescending(rates)

	if err := storage.WriteJSON(cfg.OutputFile, rates); err != nil {
		panic(err)
	}
}

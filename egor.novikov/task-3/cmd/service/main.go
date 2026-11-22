package main

import (
	"flag"
	"github.com/EGGorkaa/task-3/internal/config"
	"github.com/EGGorkaa/task-3/internal/currency"
	"github.com/EGGorkaa/task-3/internal/output"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg := config.ReadConfig(*configPath)

	currencies := currency.ReadXML(cfg.InputFile)

	currency.SortByValueDesc(currencies)

	output.WriteJSON(cfg.OutputFile, currencies)
}

package main

import (
	"flag"

	"github.com/Kayv1e/task-3/internal/config"
	"github.com/Kayv1e/task-3/internal/parser"
	"github.com/Kayv1e/task-3/internal/processor"
	"github.com/Kayv1e/task-3/internal/storage"
)

func main() {
	configPath := flag.String("config", "", "YAML config file path")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := parser.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	processor.SortByValueDesc(currencies)

	if err := storage.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}

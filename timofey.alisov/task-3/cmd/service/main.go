package main

import (
	"flag"

	"github.com/timfoi/task-3/internal/config"
	"github.com/timfoi/task-3/internal/parser"
	"github.com/timfoi/task-3/internal/sorter"
	"github.com/timfoi/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	valutes, err := parser.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sorter.ByValueDesc(valutes)

	if err := writer.WriteJSON(cfg.OutputFile, valutes); err != nil {
		panic(err)
	}
}

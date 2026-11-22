package main

import (
	"slices"

	"github.com/Pilipchenok/task-3/internal/config"
	"github.com/Pilipchenok/task-3/internal/exporter"
	"github.com/Pilipchenok/task-3/internal/model"
	"github.com/Pilipchenok/task-3/internal/parser"
)

func main() {
	cfgPath, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	cfg, err := config.ParseConfig(cfgPath)
	if err != nil {
		panic(err)
	}

	curs, err := parser.DecodeXML[model.ValCurs](cfg.InputFile)
	if err != nil {
		panic(err)
	}

	slices.SortFunc(curs.Valutes, model.SortFunc)

	if err := exporter.WriteJSON(curs.Valutes, cfg.OutputFile); err != nil {
		panic(err)
	}
}

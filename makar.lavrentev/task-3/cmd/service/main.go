package main

import (
	"slices"

	"github.com/Pilipchenok/task-3/internal/config"
	"github.com/Pilipchenok/task-3/internal/exporter"
	"github.com/Pilipchenok/task-3/internal/parser"
	"github.com/Pilipchenok/task-3/internal/service"
)

func main() {
	cfgPath := config.ReadConfig()
	cfg := config.ParseConfig(cfgPath)
	curs := parser.DecodeXML(cfg.InputFile)
	slices.SortFunc(curs.Valutes, service.SortFunc)
	exporter.WriteJSON(curs.Valutes, cfg.OutputFile)
}

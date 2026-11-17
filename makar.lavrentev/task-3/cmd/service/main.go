package main

import (
	"slices"

	"github.com/Pilipchenok/task-3/internal/config"
	"github.com/Pilipchenok/task-3/internal/exporter"
	"github.com/Pilipchenok/task-3/internal/parser"
	"github.com/Pilipchenok/task-3/internal/service"
)

func main() {
	cfg := config.ParseConfig()
	curs := parser.DecodeXML(cfg.InputFile)
	valutes := service.SimplifyData(curs)
	slices.SortFunc(valutes, service.SortFunc)
	exporter.WriteJSON(valutes, cfg.OutputFile)
}
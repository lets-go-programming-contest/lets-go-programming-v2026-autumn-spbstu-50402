package main

import (
	"flag"

	"github.com/Alastor20/task-3/internal/cbr"
	"github.com/Alastor20/task-3/internal/config"
	"github.com/Alastor20/task-3/internal/output"
)

func main() {
	confPath := flag.String("config", "", "path to config file")

	flag.Parse()

	cfg, err := config.Load(*confPath)
	if err != nil {
		panic(err)
	}

	currencies, err := cbr.ParseFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	cbr.SortByValueDesc(currencies)

	err = output.WriteJSON(cfg.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}

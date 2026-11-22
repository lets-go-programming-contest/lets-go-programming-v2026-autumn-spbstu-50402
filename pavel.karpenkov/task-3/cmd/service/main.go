package main

import (
	"errors"
	"flag"

	"github.com/xRonnyn/lets-go-programming-v2026-autumn-spbstu-50402/task-3/cbr"
	"github.com/xRonnyn/lets-go-programming-v2026-autumn-spbstu-50402/task-3/config"
	"github.com/xRonnyn/lets-go-programming-v2026-autumn-spbstu-50402/task-3/currency"
)

var errConfigPathNotSet = errors.New("config path is not set")

func main() {
	configPath := flag.String("config", "", "path to yaml config file")
	flag.Parse()

	if *configPath == "" {
		panic(errConfigPathNotSet)
	}

	cfg, err := config.Parse(*configPath)
	if err != nil {
		panic(err)
	}

	valutes, err := cbr.Read(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.Convert(valutes)
	if err != nil {
		panic(err)
	}

	currency.Sort(currencies)

	err = currency.Save(cfg.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"errors"
	"flag"
	"fmt"
	"sort"

	"github.com/InsafM1/task-3/internal/cbr"
	"github.com/InsafM1/task-3/internal/config"
	"github.com/InsafM1/task-3/internal/storage"
)

var ErrMissingConfigFields = errors.New("input-file and output-file are required")

func main() {
	path := flag.String("config", "", "configuration file path")
	flag.Parse()

	cfg, err := config.LoadConfigFile(*path)
	if err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic(fmt.Errorf("config %q: input-file or output-file are required: %w", *path, ErrMissingConfigFields))
	}

	currencies, err := cbr.ReadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	err = storage.WriteJSON(cfg.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}

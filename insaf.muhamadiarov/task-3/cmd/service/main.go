package main

import (
	"flag"

	"github.com/InsafM1/task-3/internal/config"
	"github.com/InsafM1/task-3/internal/cbr"
)

func main() {
	path := flag.String("config", "", "configuration file path");
	flag.Parse()

	cfg, err := config.LoadConfigFile(*path)
	if err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic(fmt.Errorf("config %q: input-file or output-file are required", *path))
	}

	currencies, err := cbr.ReadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}
}

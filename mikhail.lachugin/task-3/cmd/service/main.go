package main

import (
	"flag"
	"fmt"

	"github.com/mikeee322/task-3/internal/config"
	"github.com/mikeee322/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to the YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("configuration path is required: pass it with -config")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.LoadSorted(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded %d currencies from %s\n", len(currencies), cfg.InputFile)
	for _, item := range currencies {
		fmt.Printf("%s (%d): %g\n", item.CharCode, item.NumCode, item.Value)
	}
}

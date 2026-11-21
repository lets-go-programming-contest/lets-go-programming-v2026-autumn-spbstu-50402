package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/KuzRK/task-3/internal/config"
	"github.com/KuzRK/task-3/internal/currency"
	"github.com/KuzRK/task-3/internal/output"
)

func main() {
	configPath := flag.String("config", "", "path to the YAML configuration file")
	flag.Parse()

	if strings.TrimSpace(*configPath) == "" {
		panic("the -config flag is required")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Errorf("load configuration %q: %w", *configPath, err))
	}

	currencies, err := currency.Read(cfg.InputFile)
	if err != nil {
		panic(fmt.Errorf("read currencies from %q: %w", cfg.InputFile, err))
	}

	if err := output.WriteJSON(cfg.OutputFile, currencies); err != nil {
		panic(fmt.Errorf("write JSON to %q: %w", cfg.OutputFile, err))
	}
}

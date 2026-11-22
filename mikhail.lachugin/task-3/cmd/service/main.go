package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mikeee322/task-3/internal/config"
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

	// The remaining processing steps are added in the following project parts.
	fmt.Fprintf(os.Stdout, "Configuration loaded: input=%s output=%s\n", cfg.InputFile, cfg.OutputFile)
}

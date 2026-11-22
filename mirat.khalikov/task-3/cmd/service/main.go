package main

import (
	"flag"
	"fmt"
	"task-3/cmd/service/internal/config"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("config file is required")
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Input File: ", cfg.InputFile)
	fmt.Println("Output File: ", cfg.OutputFile)
}

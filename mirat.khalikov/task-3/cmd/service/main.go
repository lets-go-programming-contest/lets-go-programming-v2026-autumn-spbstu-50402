package main

import (
	"flag"
	"task-3/cmd/service/internal/config"
	"task-3/cmd/service/internal/output"
	"task-3/cmd/service/internal/xml"
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

	currencies, err := xml.ParseFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	xml.SortCur(currencies)

	err = output.WriteFile(cfg.OutputFile, currencies)
	if err != nil {
		panic(err)
	}

}

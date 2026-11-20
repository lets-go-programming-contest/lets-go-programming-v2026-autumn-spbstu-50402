package main

import (
	"flag"

	"github.com/Sp0146001/task-3/internal/config"
	"github.com/Sp0146001/task-3/internal/rate"
	"github.com/Sp0146001/task-3/internal/report"
)

func readFlags() string {
	path := flag.String("config", "config.yaml", "path to yaml config")

	flag.Parse()

	return *path
}

func main() {
	path := readFlags()

	cfg, err := config.Read(path)
	if err != nil {
		panic(err)
	}

	list, err := rate.Read(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	rate.SortByValue(list)

	if err := report.Write(cfg.OutputFile, list); err != nil {
		panic(err)
	}
}

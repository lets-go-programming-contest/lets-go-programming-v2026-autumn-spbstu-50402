package main

import (
	"flag"
	"fmt"

	"github.com/BukreevYas/task-3/internal/config"
	"github.com/BukreevYas/task-3/internal/report"
)

func main() {
	path := flag.String("config", "", "configuration file path")
	flag.Parse()

	cfg, err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}

	data, err := report.ParseReport(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(data[0].Value)
}

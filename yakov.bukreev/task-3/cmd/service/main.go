package main

import (
	"flag"
	"fmt"
	"sort"

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

	data, err := report.ReadXml(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})

	fmt.Println(data[0].Value)
}

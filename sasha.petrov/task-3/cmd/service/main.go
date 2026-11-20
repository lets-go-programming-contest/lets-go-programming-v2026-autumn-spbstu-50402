package main

import (
	"encoding/xml"
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

	var curs struct {
		XMLName xml.Name     `xml:"ValCurs"`
		Items   []rate.Quote `xml:"Valute"`
	}

	if err := rate.Read(cfg.InputFile, &curs); err != nil {
		panic(err)
	}

	rate.SortByValue(curs.Items)

	if err := report.Write(cfg.OutputFile, curs.Items); err != nil {
		panic(err)
	}
}

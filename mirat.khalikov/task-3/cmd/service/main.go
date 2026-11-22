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
	checkError(err)

	currencies, err := xml.ParseFile(cfg.InputFile)
	checkError(err)

	xml.SortCur(currencies)

	err = output.WriteFile(cfg.OutputFile, currencies)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

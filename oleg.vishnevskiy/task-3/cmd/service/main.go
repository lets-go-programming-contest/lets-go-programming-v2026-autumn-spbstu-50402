package main

import (
	"flag"

	"github.com/voleg904/task-3/internal/config"
	"github.com/voleg904/task-3/internal/jsonconverter"
	"github.com/voleg904/task-3/internal/xmlparser"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	path, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}

	curs, err := xmlparser.Parse(path.InputPath)
	if err != nil {
		panic(err)
	}

	err = jsonconverter.Save(curs, path.OutputPath)
	if err != nil {
		panic(err)
	}
}

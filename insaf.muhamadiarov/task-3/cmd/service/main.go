package main

import (
	"flag"

	"github.com/InsafM1/task-3/internal/config"
)

func main() {
	path := flag.String("config", "", "configuration file path");
	flag.Parse()

	cfg, err := config.LoadConfigFile(*path)
	if err != nil {
		panic(err)
	}
}

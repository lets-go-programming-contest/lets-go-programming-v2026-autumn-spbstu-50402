package main

import (
	"flag"

	"github.com/Alastor20/task-3/internal/config"
)

func main() {
	confPath := flag.String("config", "", "path to config file")

	flag.Parse()

	_, err := config.Load(*confPath)
	if err != nil {
		panic(err)
	}
}

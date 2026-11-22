package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/xRonnyn/lets-go-programming-v2026-autumn-spbstu-50402/task-3/config"
)

var errConfigPathNotSet = errors.New("config path is not set")

func main() {
	configPath := flag.String("config", "", "path to yaml config file")
	flag.Parse()

	if *configPath == "" {
		panic(errConfigPathNotSet)
	}

	cfg, err := config.Parse(*configPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}

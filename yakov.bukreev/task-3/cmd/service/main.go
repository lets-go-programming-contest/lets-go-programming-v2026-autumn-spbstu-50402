package main

import (
	"flag"
	"fmt"

	"github.com/BukreevYas/task-3/internal/config"
)

func main() {
	path := flag.String("config", "", "configuration file path")
	flag.Parse()

	cfg, err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.InputFile)
}

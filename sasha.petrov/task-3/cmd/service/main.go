package main

import (
	"flag"
)

func readFlags() string {
	path := flag.String("config", "", "path to yaml config")

	flag.Parse()

	if *path == "" {
		panic("config flag is empty")
	}

	return *path
}

func main() {
	_ = readFlags()
}

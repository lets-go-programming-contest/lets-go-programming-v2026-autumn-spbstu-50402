package main

import (
	"flag"
)

func main() {
	_ = flag.String("config", "", "path to yaml config")

	flag.Parse()
}

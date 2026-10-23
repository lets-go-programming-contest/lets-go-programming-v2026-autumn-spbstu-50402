package main

import (
	"errors"
	"fmt"
)

const (
	minDeps = 1
	maxDeps = 1000

	minWorkers = 1
	maxWorkers = 1000

	minT = 15
	maxT = 30
)

var (
//forerrors
)

type tempBounds struct {
	lowest, highest int
}

func newTempBounds() *tempBounds {}

func scanDeps() {}

func scanWorkers() {}

func scanBounds() {}

func processDeps() {}

func main() {
	var deps int
	deps, err := scanN()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := range deps {
		if err := processDeps(); err != nil {
			fmt.Println("Error: ", err)
		}
}

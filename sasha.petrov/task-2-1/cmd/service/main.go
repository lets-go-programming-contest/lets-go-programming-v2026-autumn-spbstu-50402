package main

import (
	"errors"
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30

	minDepartments = 1
	minWorkers     = 1
	maxDepartments = 1000
	maxWorkers     = 1000

	bolshe = ">="
	menshe = "<="
)

type temperature struct {
	low  int
	high int
}

func newTemperature() *temperature {
	return &temperature{low: minTemperature, high: maxTemperature}
}

func isKnownOperation(operation string) bool {
	return operation == bolshe || operation == menshe
}

var errInvalidInput = errors.New("invalid input")

func runTemperature() error { return nil }

func main() {
	if err := runTemperature(); err != nil {
		fmt.Printf("error is: %v", err)
	}
}

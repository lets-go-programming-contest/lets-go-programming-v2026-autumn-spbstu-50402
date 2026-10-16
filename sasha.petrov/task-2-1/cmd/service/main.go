package main

import (
	"errors"
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30

	minData = 1
	maxData = 1000

	bolshe = ">="
	menshe = "<="
)

var errInvalidInput = errors.New("invalid input of data")

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

func readData(detailInfo string) (int, error) {
	var data int
	if _, err := fmt.Scan(&data); err != nil {
		return 0, fmt.Errorf("read %s: %w", detailInfo, err)
	}

	if minData > data || data > maxData {
		return 0, fmt.Errorf("%w: %s not between 1 and 1000", errInvalidInput, detailInfo)
	}

	return data, nil
}

func runTemperature() error {
	departments, err := readData("departments")
	if err != nil {
		return err
	}

	for range departments {
		workers, err := readData("workers")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := runTemperature(); err != nil {
		fmt.Printf("error is: %v\n", err)
	}
}

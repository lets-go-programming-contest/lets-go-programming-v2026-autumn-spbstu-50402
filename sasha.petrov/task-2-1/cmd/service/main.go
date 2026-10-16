package main

import (
	"errors"
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30
	noTemperature  = -1

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

func (r *temperature) update(operation string, temperature int) error {
	if !isKnownOperation(operation) {
		return fmt.Errorf("%w: unknown operation %q", errInvalidInput, operation)
	}

	if err := checkBounds("temperature", temperature, minTemperature, maxTemperature); err != nil {
		return err
	}

	if operation == bolshe {
		r.low = max(r.low, temperature)
	} else {
		r.high = min(r.high, temperature)
	}

	return nil
}

func isKnownOperation(operation string) bool {
	return operation == bolshe || operation == menshe
}

func checkBounds(detailInfo string, value, minValue, maxValue int) error {
	if value < minValue || value > maxValue {
		return fmt.Errorf("%w: %s %d is not in [%d; %d]", errInvalidInput, detailInfo, value, minValue, maxValue)
	}

	return nil
}

func (r *temperature) result() int {
	if r.low > r.high {
		return noTemperature
	}

	return r.low
}

func readData(detailInfo string) (int, error) {
	var data int
	if _, err := fmt.Scan(&data); err != nil {
		return 0, fmt.Errorf("read %s: %w", detailInfo, err)
	}

	if err := checkBounds(detailInfo, data, minData, maxData); err != nil {
		return 0, err
	}

	return data, nil
}

func readInfoOfTemperature() (string, int, error) {
	var (
		operation           string
		temperatureToChange int
	)

	if _, err := fmt.Scan(&operation); err != nil {
		return "", 0, fmt.Errorf("read operation: %w", err)
	}

	if _, err := fmt.Scan(&temperatureToChange); err != nil {
		return "", 0, fmt.Errorf("read temperature: %w", err)
	}

	return operation, temperatureToChange, nil
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
		limits := newTemperature()
		for range workers {
			operaton, temperature, err := readInfoOfTemperature()
			if err != nil {
				return err
			}
			if err := limits.update(operaton, temperature); err != nil {
				return err
			}
			fmt.Println(limits.result())
		}
	}

	return nil
}

func main() {
	if err := runTemperature(); err != nil {
		fmt.Printf("error is: %v\n", err)
	}
}

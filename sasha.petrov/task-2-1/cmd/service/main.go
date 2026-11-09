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

	atLeast = ">="
	atMost  = "<="
)

var (
	errOutOfRange       = errors.New("value is out of range")
	errUnknownOperation = errors.New("unknown operation")
	errNoTemperature    = errors.New("no correct temperature")
)

type temperature struct {
	low  int
	high int
}

func newTemperature(low, high int) *temperature {
	return &temperature{low: low, high: high}
}

func (temp *temperature) update(operation string, temperature int) error {
	if !isKnownOperation(operation) {
		return fmt.Errorf("update by %q: %w", operation, errUnknownOperation)
	}

	if err := checkBounds(temperature, minTemperature, maxTemperature); err != nil {
		return fmt.Errorf("update to %d: %w", temperature, err)
	}

	if operation == atLeast {
		temp.low = max(temp.low, temperature)
	} else {
		temp.high = min(temp.high, temperature)
	}

	return nil
}

func isKnownOperation(operation string) bool {
	return operation == atLeast || operation == atMost
}

func checkBounds(value, minValue, maxValue int) error {
	if value < minValue || value > maxValue {
		return errOutOfRange
	}

	return nil
}

func (temp *temperature) result() (int, error) {
	if temp.low > temp.high {
		return 0, errNoTemperature
	}

	return temp.low, nil
}

func readData() (int, error) {
	var data int
	if _, err := fmt.Scan(&data); err != nil {
		return 0, fmt.Errorf("scan data: %w", err)
	}

	if err := checkBounds(data, minData, maxData); err != nil {
		return 0, fmt.Errorf("data %d: %w", data, err)
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
	departments, err := readData()
	if err != nil {
		return fmt.Errorf("read departments: %w", err)
	}

	for range departments {
		workers, err := readData()
		if err != nil {
			return fmt.Errorf("read workers: %w", err)
		}

		limits := newTemperature(minTemperature, maxTemperature)

		for range workers {
			operation, temperature, err := readInfoOfTemperature()
			if err != nil {
				return err
			}

			if err := limits.update(operation, temperature); err != nil {
				return err
			}

			answer, err := limits.result()

			switch {
			case errors.Is(err, errNoTemperature):
				answer = noTemperature
			case err != nil:
				return fmt.Errorf("calculate temperature: %w", err)
			}

			fmt.Println(answer)
		}
	}

	return nil
}

func main() {
	if err := runTemperature(); err != nil {
		fmt.Printf("error is: %v\n", err)
	}
}

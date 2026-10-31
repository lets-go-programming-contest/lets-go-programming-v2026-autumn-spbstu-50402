package main

import (
	"errors"
	"fmt"
)

const (
	initialMinTemperature = 15
	initialMaxTemperature = 30
)

var (
	errInvalidOperation         = errors.New("invalid operation: expected >= or <=")
	errNoComfortableTemperature = errors.New("no comfortable temperature")
)

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange(minTemperature, maxTemperature int) TemperatureRange {
	return TemperatureRange{
		min: minTemperature,
		max: maxTemperature,
	}
}

func (r *TemperatureRange) Apply(operation string, temperature int) error {
	switch operation {
	case ">=":
		r.min = max(r.min, temperature)
	case "<=":
		r.max = min(r.max, temperature)
	default:
		return fmt.Errorf("%w: %q", errInvalidOperation, operation)
	}

	return nil
}

func (r *TemperatureRange) ComfortableMin() (int, error) {
	if r.min <= r.max {
		return r.min, nil
	}

	return 0, errNoComfortableTemperature
}

func readInt() (int, error) {
	var value int
	if _, err := fmt.Scan(&value); err != nil {
		return 0, fmt.Errorf("scan int: %w", err)
	}

	return value, nil
}

func readOperation() (string, error) {
	var operation string
	if _, err := fmt.Scan(&operation); err != nil {
		return "", fmt.Errorf("scan operation: %w", err)
	}

	return operation, nil
}

func processDepartment() error {
	people, err := readInt()
	if err != nil {
		return fmt.Errorf("read number of people: %w", err)
	}

	temperatureRange := NewTemperatureRange(initialMinTemperature, initialMaxTemperature)

	for range people {
		operation, err := readOperation()
		if err != nil {
			return fmt.Errorf("read operation: %w", err)
		}

		temperature, err := readInt()
		if err != nil {
			return fmt.Errorf("read temperature: %w", err)
		}

		if err := temperatureRange.Apply(operation, temperature); err != nil {
			return err
		}

		minTemperature, err := temperatureRange.ComfortableMin()
		if err != nil {
			fmt.Println(-1)
		} else {
			fmt.Println(minTemperature)
		}
	}

	return nil
}

func main() {
	departments, err := readInt()
	if err != nil {
		fmt.Printf("read number of departments: %v\n", err)

		return
	}

	for range departments {
		if err := processDepartment(); err != nil {
			fmt.Println(err)

			return
		}
	}
}

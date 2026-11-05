package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

var (
	errInvalidRangeUpdate = errors.New("range update failed")
	errOptimalTemperature = errors.New("temperature measurement error")
	errInvalidOperation   = errors.New("invalid operation")
)

type tempRange struct {
	lowerBound, upperBound int
}

func newTempRange(minTemp, maxTemp int) tempRange {
	return tempRange{lowerBound: minTemp, upperBound: maxTemp}
}

func parseCondition() (string, int, error) {
	var operation string

	_, err := fmt.Scan(&operation)
	if err != nil {
		return "", 0, fmt.Errorf("error: reading operation is failed: %w", err)
	}

	if operation != "<=" && operation != ">=" {
		return "", 0, fmt.Errorf("%w: %q", errInvalidOperation, operation)
	}

	var temperature int

	_, err = fmt.Scan(&temperature)
	if err != nil {
		return "", 0, fmt.Errorf("error: reading number is failed: %w", err)
	}

	return operation, temperature, nil
}

func (limits *tempRange) update(operation string, temperature int) error {
	switch operation {
	case ">=":
		if temperature > limits.lowerBound {
			limits.lowerBound = temperature
		}
	case "<=":
		if temperature < limits.upperBound {
			limits.upperBound = temperature
		}
	default:
		return errInvalidRangeUpdate
	}

	return nil
}

func (limits *tempRange) optimal() (int, error) {
	if limits.lowerBound <= limits.upperBound {
		return limits.lowerBound, nil
	}

	return 0, errOptimalTemperature
}

func main() {
	var department int

	_, err := fmt.Scanln(&department)
	if err != nil {
		fmt.Println("Invalid department argument:", err)

		return
	}

	for range department {
		var employee int

		_, err := fmt.Scanln(&employee)
		if err != nil {
			fmt.Println("Invalid employee argument:", err)

			return
		}

		temperatureRange := newTempRange(minTemp, maxTemp)

		for range employee {
			operation, temperature, err := parseCondition()
			if err != nil {
				fmt.Println("invalid employee data:", err)

				return
			}

			err = temperatureRange.update(operation, temperature)
			if err != nil {
				fmt.Println("update range failed:", err)

				return
			}

			optimalTemperature, err := temperatureRange.optimal()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimalTemperature)
		}
	}
}

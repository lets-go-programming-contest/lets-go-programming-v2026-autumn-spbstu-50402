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
	errInvalidOperation     = errors.New("invalid operation")
	errNoOptimalTemperature = errors.New("cannot get optimal temperature")
)

type temperatureRange struct {
	lower int
	upper int
}

func newTemperatureRange(lower, upper int) temperatureRange {
	return temperatureRange{
		lower: lower,
		upper: upper,
	}
}

func getEmployeeData() (string, int, error) {
	var operation string
	if _, err := fmt.Scan(&operation); err != nil {
		return "", 0, fmt.Errorf("failed to scan operation: %w", err)
	}

	if operation != "<=" && operation != ">=" {
		return "", 0, fmt.Errorf("%w: entered %q", errInvalidOperation, operation)
	}

	var temperature int
	if _, err := fmt.Scan(&temperature); err != nil {
		return "", 0, fmt.Errorf("failed to scan temperature: %w", err)
	}

	return operation, temperature, nil
}

func (tr *temperatureRange) updateBounds(operation string, temperature int) error {
	switch operation {
	case "<=":
		if temperature < tr.upper {
			tr.upper = temperature
		}
	case ">=":
		if temperature > tr.lower {
			tr.lower = temperature
		}
	default:
		return errInvalidOperation
	}

	return nil
}

func (tr *temperatureRange) getOptimalTemperature() (int, error) {
	if tr.lower <= tr.upper {
		return tr.lower, nil
	}

	return 0, errNoOptimalTemperature
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scanln(&departmentsCount); err != nil {
		fmt.Printf("cannot read the number of departments: %v\n", err)

		return
	}

	for range departmentsCount {
		var employees int
		if _, err := fmt.Scanln(&employees); err != nil {
			fmt.Printf("cannot read the number of employees: %v\n", err)

			return
		}

		tempRange := newTemperatureRange(minTemp, maxTemp)

		for range employees {
			operation, temperature, err := getEmployeeData()
			if err != nil {
				fmt.Printf("cannot get employee data: %v\n", err)

				return
			}

			err = tempRange.updateBounds(operation, temperature)
			if err != nil {
				fmt.Printf("cannot update temperature bounds: %v\n", err)

				return
			}

			optimalTemp, err := tempRange.getOptimalTemperature()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimalTemp)
		}
	}
}

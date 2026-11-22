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
	errInvalidComparisonOperator = errors.New("invalid operator")
	errNoOptimalTemperature      = errors.New("optimal temperature range is empty")
)

type TemperatureRange struct {
	lower int
	upper int
}

func NewTemperatureRange(lower, upper int) TemperatureRange {
	return TemperatureRange{
		lower: lower,
		upper: upper,
	}
}

func (tr *TemperatureRange) UpdateBounds(operation string, temperature int) error {
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
		return errInvalidComparisonOperator
	}

	return nil
}

func (tr *TemperatureRange) GetOptimalTemperature() (int, error) {
	if tr.lower <= tr.upper {
		return tr.lower, nil
	}

	return 0, errNoOptimalTemperature
}

func scanEmployee() (string, int, error) {
	var operation string
	if _, err := fmt.Scan(&operation); err != nil {
		return "", 0, fmt.Errorf("failed to scan operation: %w", err)
	}

	if operation != "<=" && operation != ">=" {
		return "", 0, fmt.Errorf("%w: got %q", errInvalidComparisonOperator, operation)
	}

	var temperature int
	if _, err := fmt.Scan(&temperature); err != nil {
		return "", 0, fmt.Errorf("failed to scan temperature: %w", err)
	}

	return operation, temperature, nil
}

func main() {
	var numberOfDepartment int
	if _, err := fmt.Scanln(&numberOfDepartment); err != nil {
		fmt.Printf("error reading the number of departments: %v\n", err)

		return
	}

	for range numberOfDepartment {
		var staff int
		if _, err := fmt.Scanln(&staff); err != nil {
			fmt.Printf("error reading the staff: %v\n", err)

			return
		}

		tempRange := NewTemperatureRange(minTemp, maxTemp)

		for range staff {
			operation, temperature, err := scanEmployee()
			if err != nil {
				fmt.Printf("failed to read employee data: %v\n", err)

				return
			}

			err = tempRange.UpdateBounds(operation, temperature)
			if err != nil {
				fmt.Printf("failed to update temperature bounds: %v\n", err)

				return
			}

			optimalTemp, err := tempRange.GetOptimalTemperature()
			if errors.Is(err, errNoOptimalTemperature) {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimalTemp)
		}
	}
}

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
	ErrInvalidOperation     = errors.New("invalid operation")
	ErrNoOptimalTemperature = errors.New("cannot get optimal temperature")
)

func GetEmployeeData() (string, int, error) {
	var operation string
	if _, err := fmt.Scan(&operation); err != nil {
		return "", 0, fmt.Errorf("failed to scan operation: %w", err)
	}

	if operation != "<=" && operation != ">=" {
		return "", 0, fmt.Errorf("%w: entered %q", ErrInvalidOperation, operation)
	}

	var temperature int
	if _, err := fmt.Scan(&temperature); err != nil {
		return "", 0, fmt.Errorf("failed to scan temperature: %w", err)
	}

	return operation, temperature, nil
}

func UpdateBounds(operation string, temperature int, lower *int, upper *int) error {
	switch operation {
	case "<=":
		if temperature < *upper {
			*upper = temperature
		}
	case ">=":
		if temperature > *lower {
			*lower = temperature
		}
	default:
		return ErrInvalidOperation
	}

	return nil
}

func GetOptimalTemperature(lower *int, upper *int) (int, error) {
	if *lower <= *upper {
		return *lower, nil
	}

	return 0, ErrNoOptimalTemperature
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

		lower := minTemp
		upper := maxTemp

		for range employees {
			operation, temperature, err := GetEmployeeData()
			if err != nil {
				fmt.Printf("cannot get employee data: %v\n", err)

				return
			}

			err = UpdateBounds(operation, temperature, &lower, &upper)
			if err != nil {
				fmt.Printf("cannot update temperature bounds: %v\n", err)

				return
			}

			optimalTemp, err := GetOptimalTemperature(&lower, &upper)
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimalTemp)
		}
	}
}

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
	errInvalidOperator          = errors.New("unsupported operator")
	errRangeTemperatureNotFound = errors.New("the request violares optimal temperature")
)

type TempRange struct {
	lowLevel  int
	highLevel int
}

func toChangeTemp(tempRange *TempRange, operation string, newTemp int) (int, error) {
	switch operation {
	case ">=":
		if newTemp > tempRange.lowLevel {
			tempRange.lowLevel = newTemp
		}

	case "<=":
		if newTemp < tempRange.highLevel {
			tempRange.highLevel = newTemp
		}

	default:
		return 0, fmt.Errorf("%w: %s", errInvalidOperator, operation)
	}

	if tempRange.lowLevel > tempRange.highLevel {
		return 0, fmt.Errorf("%w: %s %d", errRangeTemperatureNotFound, operation, newTemp)
	}

	return tempRange.lowLevel, nil
}

func readLine() (string, int, error) {
	var (
		operation string
		temp      int
	)

	_, err := fmt.Scan(&operation)
	if err != nil {
		return "", 0, fmt.Errorf("faild to read the operaton: %w", err)
	}

	_, err = fmt.Scan(&temp)
	if err != nil {
		return "", 0, fmt.Errorf("faild to read the temperature: %w", err)
	}

	return operation, temp, nil
}

func main() {
	var numberOfDepartments int

	_, err := fmt.Scan(&numberOfDepartments)
	if err != nil {
		fmt.Printf("error reading the number of departments: %v\n", err)

		return
	}

	for range numberOfDepartments {
		var numberOfEmployees int

		_, err := fmt.Scan(&numberOfEmployees)
		if err != nil {
			fmt.Printf("error reading the number of employees: %v\n", err)

			return
		}

		tempRange := TempRange{minTemp, maxTemp}

		for range numberOfEmployees {
			operation, newTemp, err := readLine()
			if err != nil {
				fmt.Printf("error reading the Line: %v\n", err)

				return
			}

			result, err := toChangeTemp(&tempRange, operation, newTemp)
			if errors.Is(err, errInvalidOperator) {
				fmt.Printf("error changing range of temperature: %v\n", err)

				return
			} else if errors.Is(err, errRangeTemperatureNotFound) {
				fmt.Println(-1)

				continue
			}

			fmt.Println(result)
		}
	}
}

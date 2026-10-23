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

func CreateTempRange(lowLevel, highLevel int) TempRange {
	return TempRange{
		lowLevel:  lowLevel,
		highLevel: highLevel,
	}
}

func (tempRange *TempRange) ToChangeTemp(operation string, newTemp int) error {
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
		return errInvalidOperator
	}

	return nil
}

func (tempRange *TempRange) GetOptimalTempRange() (int, error) {
	if tempRange.lowLevel > tempRange.highLevel {
		return 0, errRangeTemperatureNotFound
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

		tempRange := CreateTempRange(minTemp, maxTemp)

		for range numberOfEmployees {
			operation, newTemp, err := readLine()
			if err != nil {
				fmt.Printf("error reading the Line: %v\n", err)

				return
			}

			err = tempRange.ToChangeTemp(operation, newTemp)
			if errors.Is(err, errInvalidOperator) {
				fmt.Printf("error changing range of temperature: %v\n", err)

				return
			}

			result, err := tempRange.GetOptimalTempRange()
			if errors.Is(err, errRangeTemperatureNotFound) {
				fmt.Println(-1)

				continue
			}

			fmt.Println(result)
		}
	}
}

package main

import (
	"errors"
	"fmt"
)

var errInvalidOperator = errors.New("unsupported operator")

type TempRange struct {
	lowLevel  int
	highLevel int
}

func toChangeTemp(tempRange *TempRange, operation string, newTemp int) (int, error) {
	switch operation {
	case ">=":

		if newTemp <= tempRange.highLevel {

			if newTemp > tempRange.lowLevel {
				tempRange.lowLevel = newTemp
			}

			return tempRange.lowLevel, nil
		}
	case "<=":

		if newTemp > tempRange.lowLevel {

			if newTemp < tempRange.highLevel {
				tempRange.highLevel = newTemp
			}

			return tempRange.lowLevel, nil
		}
	default:

		return 0, fmt.Errorf("%w: %s", errInvalidOperator, operation)
	}

	return -1, nil
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

		tempRange := TempRange{15, 30}

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
			}
			fmt.Println(result)
		}
	}
}

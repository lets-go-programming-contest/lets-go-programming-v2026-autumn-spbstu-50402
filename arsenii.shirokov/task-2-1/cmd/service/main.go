package main

import (
	"errors"
	"fmt"
)

const (
	lRange = 14
	rRange = 31
)

var (
	errInvalidComparisonOperator = errors.New("invalid operator")
	errInvalidTemperature        = errors.New("invalid temperature")
)

func scanEmployee() (string, int, error) {
	var operation string
	if _, err := fmt.Scan(&operation); err != nil || !(operation == "<=" || operation == ">=") {
		return "", 0, errInvalidComparisonOperator
	}

	var temperature int
	if _, err := fmt.Scan(&temperature); err != nil || temperature >= rRange || temperature <= lRange {
		return "", 0, errInvalidTemperature
	}

	return operation, temperature, nil
}

func getOptimalTemperature(staff int) error {
	minLess := rRange
	maxLarger := lRange

	for range staff {
		operation, temperature, err := scanEmployee()
		if err != nil {
			return err
		}

		if operation == "<=" && temperature < minLess {
			minLess = temperature
		} else if operation == ">=" && temperature > maxLarger {
			maxLarger = temperature
		}

		if maxLarger <= minLess {
			fmt.Println(maxLarger)
		} else {
			fmt.Println(-1)
		}
	}

	return nil
}

func main() {
	var numberOfDepartment int
	if _, err := fmt.Scanln(&numberOfDepartment); err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range numberOfDepartment {
		var staff int
		if _, err := fmt.Scanln(&staff); err != nil {
			fmt.Println("Invalid number of employees")
		}

		err := getOptimalTemperature(staff)

		switch {
		case err == nil:
		case errors.Is(err, errInvalidComparisonOperator):
			fmt.Println("Invalid operator")

			return
		case errors.Is(err, errInvalidTemperature):
			fmt.Println("Invalid temperature")

			return
		default:
			fmt.Println(err)

			return
		}
	}
}

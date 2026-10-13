package main

import (
	"errors"
	"fmt"
)

const lRange = 14
const rRange = 31

var (
	errInvalidComparisonOperator = errors.New("invalid operator")
	errInvalidTemperature        = errors.New("invalid temperature")
)

func getIntermediateResult(minLess int, maxLarger int) {
	if minLess < maxLarger && minLess != rRange {
		fmt.Println(-1)
	} else {
		if maxLarger != lRange {
			fmt.Println(maxLarger)
		} else {
			fmt.Println(minLess)
		}
	}
}

func getOptimalTemperature(staff int) error {
	var minLess int = rRange
	var maxLarger int = lRange

	for range staff {
		var operation string
		if _, err := fmt.Scan(&operation); err != nil || !(operation == "<=" || operation == ">=") {
			return errInvalidComparisonOperator
		}

		var temperature int
		if _, err := fmt.Scan(&temperature); err != nil || temperature >= rRange || temperature <= lRange {
			return errInvalidTemperature
		}

		if operation == "<=" && temperature < minLess {
			minLess = temperature
		} else if operation == ">=" && temperature > maxLarger {
			maxLarger = temperature
		}

		getIntermediateResult(minLess, maxLarger)
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

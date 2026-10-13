package main

import (
	"errors"
	"fmt"
)

var (
	errInvalidComparisonOperator = errors.New("invalid operator")
	errInvalidTemperature        = errors.New("invalid temperature")
)

func getOptimalTemperature(k int) error {
	var minimumLess int = 31
	var maximumLarger int = 14
	for range k {
		var op string
		if _, err := fmt.Scan(&op); err != nil || !(op == "<=" || op == ">=") {
			return errInvalidComparisonOperator
		}

		var t int
		if _, err := fmt.Scan(&t); err != nil || t > 30 || t < 15 {
			return errInvalidTemperature
		}

		if op == "<=" && t < minimumLess {
			minimumLess = t
		} else if op == ">=" && t > maximumLarger {
			maximumLarger = t
		}

		if minimumLess < maximumLarger && minimumLess != 31 {
			fmt.Println(-1)
		} else {
			if maximumLarger != 14 {
				fmt.Println(maximumLarger)
			} else {
				fmt.Println(minimumLess)
			}
		}
	}
	return nil
}

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for range n {
		var k int
		if _, err := fmt.Scanln(&k); err != nil {
			fmt.Println("Invalid number of employees")
		}

		err := getOptimalTemperature(k)
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

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
	errInvalidOperator    = errors.New("invalid operator")
	errOptimalTemperature = errors.New("cannot get optimal temperature")
)

type TemperatureRange struct {
	low  int
	high int
}

func newTemperatureRange(low, high int) TemperatureRange {
	return TemperatureRange{
		low:  low,
		high: high,
	}
}

func (tr *TemperatureRange) update(operator string, temp int) error {
	switch operator {
	case ">=":
		if temp > tr.low {
			tr.low = temp
		}
	case "<=":
		if temp < tr.high {
			tr.high = temp
		}
	default:
		return errInvalidOperator
	}

	return nil
}

func (tr *TemperatureRange) optimal() (int, error) {
	if tr.low <= tr.high {
		return tr.low, nil
	}

	return 0, errOptimalTemperature
}

func readRequirement() (string, int, error) {
	var operator string
	if _, err := fmt.Scan(&operator); err != nil {
		return "", 0, fmt.Errorf("failed to read operator: %w", err)
	}

	if operator != ">=" && operator != "<=" {
		return "", 0, fmt.Errorf("%w: %s", errInvalidOperator, operator)
	}

	var temp int
	if _, err := fmt.Scan(&temp); err != nil {
		return "", 0, fmt.Errorf("failed to read temperature: %w", err)
	}

	return operator, temp, nil
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("invalid number of departments")

		return
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("invalid number of employees")

			return
		}

		tempRange := newTemperatureRange(minTemp, maxTemp)

		for range employees {
			op, temp, err := readRequirement()
			if err != nil {
				fmt.Println(err)

				return
			}

			if err := tempRange.update(op, temp); err != nil {
				fmt.Println(err)

				return
			}

			optimalTemp, err := tempRange.optimal()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimalTemp)
		}
	}
}

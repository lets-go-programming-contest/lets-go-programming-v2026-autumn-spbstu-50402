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
	errOptimalTemperature = errors.New("cannot get optimal temperature")
	errInvalidOperator    = errors.New("invalid operator")
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

func (tr *TemperatureRange) update(op string, temp int) error {
	switch op {
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
	var op string
	if _, err := fmt.Scan(&op); err != nil {
		return "", 0, fmt.Errorf("failed to read operator: %w", err)
	}

	if op != ">=" && op != "<=" {
		return "", 0, fmt.Errorf("%w: %s", errInvalidOperator, op)
	}

	var temp int
	if _, err := fmt.Scan(&temp); err != nil {
		return "", 0, fmt.Errorf("failed to read temperature: %w", err)
	}

	return op, temp, nil
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

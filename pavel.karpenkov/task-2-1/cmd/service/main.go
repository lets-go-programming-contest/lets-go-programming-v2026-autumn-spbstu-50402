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
	errInvalidRangeUpdate = errors.New("range update failed")
	errOptimalTemperature = errors.New("temperature measurement error")
	errInvalidOperation   = errors.New("invalid operation")
)

type TempRange struct {
	lo, hi int
}

func NewTempRange(minTemp, maxTemp int) TempRange {
	return TempRange{lo: minTemp, hi: maxTemp}
}

func parseCondition() (string, int, error) {
	var operation string

	_, err := fmt.Scan(&operation)
	if err != nil {
		return "", 0, fmt.Errorf("error: reading operation is failed: %w", err)
	}

	if operation != "<=" && operation != ">=" {
		return "", 0, fmt.Errorf("%w: %q", errInvalidOperation, operation)
	}

	var temp int

	_, err1 := fmt.Scan(&temp)
	if err1 != nil {
		return "", 0, fmt.Errorf("error: reading number is failed: %w", err)
	}

	return operation, temp, nil
}

func (r *TempRange) Update(op string, temp int) error {
	switch op {
	case ">=":
		if temp > r.lo {
			r.lo = temp
		}
	case "<=":
		if temp < r.hi {
			r.hi = temp
		}
	default:
		return errInvalidRangeUpdate
	}

	return nil
}

func (r *TempRange) Optimal() (int, error) {
	if r.lo <= r.hi {
		return r.lo, nil
	}

	return 0, errOptimalTemperature
}

func main() {
	var department int

	_, err := fmt.Scanln(&department)
	if err != nil {
		fmt.Println("Invalid department argument")

		return
	}

	for range department {
		var employee int

		_, err := fmt.Scanln(&employee)
		if err != nil {
			fmt.Println("Invalid employee argument")

			return
		}

		temperatureRange := NewTempRange(minTemp, maxTemp)

		for range employee {
			operation, temperature, err := parseCondition()
			if err != nil {
				fmt.Printf("ERROR: invalid employee data")

				return
			}
			err = temperatureRange.Update(operation, temperature)
			if err != nil {
				fmt.Printf("ERROR: update range failed")

				return
			}

			optimalTemperature, err := temperatureRange.Optimal()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimalTemperature)
		}
	}
}

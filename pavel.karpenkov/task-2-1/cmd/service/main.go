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
	errInvalidRangeUpdate = errors.New("Range update failed")
	errOptimalTemperature = errors.New("Temperature measurement error")
)

type TempRange struct {
	lo, hi int
}

func NewTempRange(minTemp, maxTemp int) TempRange {
	return TempRange{lo: minTemp, hi: maxTemp}
}
func parseCondition() (string, int, error) {
	var op string
	_, err := fmt.Scan(&op)
	if err != nil {
		return "", 0, fmt.Errorf("ERROR: reading operation is failed: %w", err)
	}
	if op != "<=" && op != ">=" {
		return "", 0, fmt.Errorf("ERROR: invalid operation: %w", err)
	}
	var temp int
	_, err1 := fmt.Scan(&temp)
	if err1 != nil {
		return "", 0, fmt.Errorf("ERROR: reading number is failed: %w", err)
	}
	return op, temp, nil
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
	var (
		department, employee int
	)

	_, err := fmt.Scanln(&department)
	if err != nil {
		fmt.Println("Invalid department argument")
		return
	}

	for range department {
		_, err := fmt.Scanln(&employee)
		if err != nil {
			fmt.Println("Invalid employee argument")
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

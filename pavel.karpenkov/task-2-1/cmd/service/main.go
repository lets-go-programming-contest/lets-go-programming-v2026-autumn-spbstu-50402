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

type RangeOfTemperature struct {
	min int
	max int
}

func NewRangeOfTemperature(min, max int) RangeOfTemperature {
	return RangeOfTemperature{min: min, max: max}
}
func parseCondition() (string, int, error) {
	var op string
	_, err := fmt.Scan(&op)
	if err != nil {
		return "", 0, fmt.Errorf("ERROR: reading operation is failed", err)
	}
	if op != "<=" && op != ">=" {
		return "", 0, fmt.Errorf("ERROR: invalid operation", err)
	}
	var temp int
	_, err1 := fmt.Scan(&temp)
	if err1 != nil {
		return "", 0, fmt.Errorf("ERROR: readint number is failed", err)
	}
	return op, temp, nil
}
func (rangeTemperature *RangeOfTemperature) updateRange(operation string, temperature int) error {
	switch operation {
	case ">=":
		if temperature > rangeTemperature.min {
			rangeTemperature.min = temperature
		}
	case "<=":
		if temperature < rangeTemperature.max {
			rangeTemperature.max = temperature
		}
	default:
		return errInvalidRangeUpdate
	}
	return nil
}
func (rangeTemperature *RangeOfTemperature) getOptimalTemperature() (int, error) {
	if rangeTemperature.min <= rangeTemperature.max {
		return rangeTemperature.min, nil
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
		temperatureRange := NewRangeOfTemperature(minTemp, maxTemp)
		for range employee {
			operation, temperature, err := parseCondition()
			if err != nil {
				fmt.Printf("ERROR: invalid employee data")
				return
			}
			err = temperatureRange.updateRange(operation, temperature)
			if err != nil {
				fmt.Printf("ERROR: update range failed")
			}
			optimalTemperature, err := temperatureRange.getOptimalTemperature()
			if err != nil {
				fmt.Println(-1)
				return
			}
			fmt.Println(optimalTemperature)
		}
	}
}

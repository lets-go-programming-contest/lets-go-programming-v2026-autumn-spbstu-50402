package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

type RangeOfTemperature struct {
	min int
	max int
}

func NewRangeOfTemperature(min, max int) (RangeOfTemperature, error) {
	if min > max {
		return nil, errors.New("ERROR: min > max temperature")
	}
	return RangeOfTemperature{min: min, max: max}, nil
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
		temperatureRange, err := NewRangeOfTemperature(minTemp, maxTemp)
		for range employee {
			operation, temperature, err := parseCondition()
		}
	}
}

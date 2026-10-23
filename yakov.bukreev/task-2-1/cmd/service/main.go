package main

import (
	"errors"
	"fmt"
)

const (
	minDepartments = 1
	maxDepartments = 1000
	minWorkers     = 1
	maxWorkers     = 1000
	minTemp        = 15
	maxTemp        = 30
)

var (
	errInvalidWorkersNumber    = errors.New("invalid workers number")
	errWorkersNumberOutOfRange = errors.New("workers number out of range")
	errInvalidPreference       = errors.New("invalid preference for worker")
	errInvalidPreferenceSign   = errors.New("invalid preference sign for worker")
	errInvalidTemperature      = errors.New("invalid temperature for worker")
)

type TemperatureBounds struct {
	low, high int
}

func (bounds *TemperatureBounds) applyBound(sign string, temperature int) {
	switch sign {
	case ">=":
		bounds.low = max(bounds.low, temperature)
	case "<=":
		bounds.high = min(bounds.high, temperature)
	}
}

func (bounds *TemperatureBounds) getOptimalTemperature() int {
	if bounds.low <= bounds.high {
		return bounds.low
	} else {
		return -1
	}
}

func processDepartment() error {
	var workers int

	if _, err := fmt.Scan(&workers); err != nil {
		return fmt.Errorf("%w: %w", errInvalidWorkersNumber, err)
	}

	if workers < minWorkers || workers > maxWorkers {
		return errWorkersNumberOutOfRange
	}

	bounds := TemperatureBounds{
		low:  minTemp,
		high: maxTemp,
	}

	for idx := range workers {
		var (
			sign        string
			temperature int
		)

		if _, err := fmt.Scan(&sign, &temperature); err != nil {
			return fmt.Errorf("%w %d: %w", errInvalidPreference, idx+1, err)
		}

		if sign != ">=" && sign != "<=" {
			return fmt.Errorf("%w %d", errInvalidPreferenceSign, idx+1)
		}

		if temperature < minTemp || temperature > maxTemp {
			return fmt.Errorf("%w %d", errInvalidTemperature, idx+1)
		}

		bounds.applyBound(sign, temperature)
		fmt.Println(bounds.getOptimalTemperature())
	}

	return nil
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Printf("Invalid departments number: %v.", err)

		return
	}

	if departments < minDepartments || departments > maxDepartments {
		fmt.Println("Departments number out of range.")

		return
	}

	for i := range departments {
		if err := processDepartment(); err != nil {
			fmt.Printf("Error (department %d): %v.", i+1, err)
		}
	}
}

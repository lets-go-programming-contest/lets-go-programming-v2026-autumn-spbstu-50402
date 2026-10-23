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
	errIncorrectBounds         = errors.New("incorrect temperature bounds")
)

type temperatureBounds struct {
	low, high int
}

func newTemperatureBounds() *temperatureBounds {
	b := temperatureBounds{
		low:  minTemp,
		high: maxTemp,
	}

	return &b
}

func (bounds *temperatureBounds) applyBound(sign string, temperature int) error {
	switch sign {
	case ">=":
		bounds.low = max(bounds.low, temperature)
	case "<=":
		bounds.high = min(bounds.high, temperature)
	default:
		return errInvalidPreferenceSign
	}

	return nil
}

func (bounds *temperatureBounds) getOptimalTemperature() (int, error) {
	if bounds.low > bounds.high {
		return 0, errIncorrectBounds
	}

	return bounds.low, nil
}

func processDepartment() error {
	var workers int

	if _, err := fmt.Scan(&workers); err != nil {
		return fmt.Errorf("%w: %w", errInvalidWorkersNumber, err)
	}

	if workers < minWorkers || workers > maxWorkers {
		return errWorkersNumberOutOfRange
	}

	bounds := newTemperatureBounds()

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

		if err := bounds.applyBound(sign, temperature); err != nil {
			return fmt.Errorf("(worker %d): %w", idx+1, err)
		}

		optimalTemp, err := bounds.getOptimalTemperature()
		if err != nil {
			optimalTemp = -1
		}

		fmt.Println(optimalTemp)
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

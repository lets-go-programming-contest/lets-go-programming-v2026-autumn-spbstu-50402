package main

import (
	"errors"
	"fmt"
)

var (
	errInvalidDepartmentCount = errors.New("invalid number of departments provided")
	errInvalidWorkerCount     = errors.New("invalid number of workers provided")
	errInvalidModifier        = errors.New("invalid format of operation modifier")
	errInvalidTemperature     = errors.New("invalid format of temperature value")
	errTemperatureBounds      = errors.New("temperature boundaries violation")
)

const (
	minTemp = 15
	maxTemp = 30
)

type TemperatureBoundary struct {
	lowerBound int
	upperBound int
}

func NewTemperatureBoundary(lower, upper int) (*TemperatureBoundary, error) {
	if lower > upper {
		return nil, fmt.Errorf("%w: lower bound %d cannot exceed upper bound %d", errTemperatureBounds, lower, upper)
	}

	return &TemperatureBoundary{
		lowerBound: lower,
		upperBound: upper,
	}, nil
}

func (cb *TemperatureBoundary) Adjust(modifier string, temp int) error {
	switch modifier {
	case ">=":
		if temp > cb.lowerBound {
			cb.lowerBound = temp
		}
	case "<=":
		if temp < cb.upperBound {
			cb.upperBound = temp
		}
	default:
		return fmt.Errorf("%w: operator '%s' is not supported", errInvalidModifier, modifier)
	}

	return nil
}

func (cb *TemperatureBoundary) IsValid() bool {
	return cb.lowerBound <= cb.upperBound
}

func (cb *TemperatureBoundary) OptimalTemperature() int {
	if !cb.IsValid() {
		return -1
	}

	return cb.lowerBound
}

func ParseDepartmentsCount() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("invalid departments input: %w", err)
	}

	if departments < 1 || departments > 1000 {
		return 0, errInvalidDepartmentCount
	}

	return departments, nil
}

func ParseWorkersCount() (int, error) {
	var workers int

	_, err := fmt.Scan(&workers)
	if err != nil {
		return 0, fmt.Errorf("invalid workers input: %w", err)
	}

	if workers < 1 || workers > 1000 {
		return 0, errInvalidWorkerCount
	}

	return workers, nil
}

func ParseConstraint() (string, int, error) {
	var modifier string

	_, err := fmt.Scan(&modifier)
	if err != nil {
		return "", 0, fmt.Errorf("invalid modifier input: %w", err)
	}

	if modifier != "<=" && modifier != ">=" {
		return "", 0, errInvalidModifier
	}

	var temperature int

	_, err = fmt.Scan(&temperature)
	if err != nil {
		return "", 0, fmt.Errorf("invalid temperature input: %w", err)
	}

	if temperature < minTemp || temperature > maxTemp {
		return "", 0, errInvalidTemperature
	}

	return modifier, temperature, nil
}

func ExecuteSimulation() error {
	deptCount, err := ParseDepartmentsCount()
	if err != nil {
		return err
	}

	for range deptCount {
		workerCount, err := ParseWorkersCount()
		if err != nil {
			return err
		}

		tempBounds, err := NewTemperatureBoundary(minTemp, maxTemp)
		if err != nil {
			return err
		}

		for range workerCount {
			modifier, temperature, err := ParseConstraint()
			if err != nil {
				return err
			}

			err = tempBounds.Adjust(modifier, temperature)
			if err != nil {
				return err
			}

			fmt.Println(tempBounds.OptimalTemperature())
		}
	}

	return nil
}

func main() {
	if err := ExecuteSimulation(); err != nil {
		fmt.Println(err)
	}
}

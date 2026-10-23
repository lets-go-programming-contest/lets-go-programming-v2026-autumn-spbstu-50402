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
	minDepartments = 1
	maxDepartments = 1000

	minWorkers = 1
	maxWorkers = 1000

	minTemp = 15
	maxTemp = 30
)

type TemperatureBoundary struct {
	lowerBound int
	upperBound int
}

func NewTemperatureBoundary(lower, upper int) (TemperatureBoundary, error) {
	if lower > upper {
		return TemperatureBoundary{}, errTemperatureBounds
	}

	return TemperatureBoundary{
		lowerBound: lower,
		upperBound: upper,
	}, nil
}

func (cb *TemperatureBoundary) Adjust(constraint Constraint) error {
	switch constraint.modifier {
	case ">=":
		if constraint.temperature > cb.lowerBound {
			cb.lowerBound = constraint.temperature
		}
	case "<=":
		if constraint.temperature < cb.upperBound {
			cb.upperBound = constraint.temperature
		}
	default:
		return fmt.Errorf("%w: operator '%s' is not supported", errInvalidModifier, constraint.modifier)
	}

	return nil
}

func (cb *TemperatureBoundary) isValid() bool {
	return cb.lowerBound <= cb.upperBound
}

func (cb *TemperatureBoundary) OptimalTemperature() (int, error) {
	if !cb.isValid() {
		return 0, errTemperatureBounds
	}

	return cb.lowerBound, nil
}

func ParseDepartmentsCount() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("invalid departments input: %w", err)
	}

	if departments < minDepartments || departments > maxDepartments {
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

	if workers < minWorkers || workers > maxWorkers {
		return 0, errInvalidWorkerCount
	}

	return workers, nil
}

type Constraint struct {
	modifier    string
	temperature int
}

func NewConstraint(modifier string, temperature int) Constraint {
	return Constraint{modifier: modifier, temperature: temperature}
}

func ParseConstraint() (Constraint, error) {
	var modifier string

	_, err := fmt.Scan(&modifier)
	if err != nil {
		return NewConstraint("", 0), fmt.Errorf("invalid modifier input: %w", err)
	}

	if modifier != "<=" && modifier != ">=" {
		return NewConstraint("", 0), errInvalidModifier
	}

	var temperature int

	_, err = fmt.Scan(&temperature)
	if err != nil {
		return NewConstraint("", 0), fmt.Errorf("invalid temperature input: %w", err)
	}

	if temperature < minTemp || temperature > maxTemp {
		return NewConstraint("", 0), errInvalidTemperature
	}

	return NewConstraint(modifier, temperature), nil
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
			constraint, err := ParseConstraint()
			if err != nil {
				return err
			}

			err = tempBounds.Adjust(constraint)
			if err != nil {
				return err
			}

			optimalTemp, err := tempBounds.OptimalTemperature()
			if err != nil {
				optimalTemp = -1
			}
			fmt.Println(optimalTemp)
		}
	}

	return nil
}

func main() {
	if err := ExecuteSimulation(); err != nil {
		fmt.Println(err)
	}
}

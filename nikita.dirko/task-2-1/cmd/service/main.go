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

type temperatureBoundary struct {
	lowerBound int
	upperBound int
}

func newTemperatureBoundary(lower, upper int) (temperatureBoundary, error) {
	if lower > upper {
		return temperatureBoundary{}, errTemperatureBounds
	}

	return temperatureBoundary{
		lowerBound: lower,
		upperBound: upper,
	}, nil
}

func (cb *temperatureBoundary) adjust(constraint constraint) error {
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
		return fmt.Errorf("%w: operator '%q' is not supported", errInvalidModifier, constraint.modifier)
	}

	return nil
}

func (cb *temperatureBoundary) isValid() bool {
	return cb.lowerBound <= cb.upperBound
}

func (cb *temperatureBoundary) optimalTemperature() (int, error) {
	if !cb.isValid() {
		return 0, errTemperatureBounds
	}

	return cb.lowerBound, nil
}

func parseDepartmentsCount() (int, error) {
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

func parseWorkersCount() (int, error) {
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

type constraint struct {
	modifier    string
	temperature int
}

func newConstraint(modifier string, temperature int) constraint {
	return constraint{modifier: modifier, temperature: temperature}
}

func parseConstraint() (constraint, error) {
	var modifier string

	_, err := fmt.Scan(&modifier)
	if err != nil {
		return constraint{}, fmt.Errorf("invalid modifier input: %w", err)
	}

	if modifier != "<=" && modifier != ">=" {
		return constraint{}, errInvalidModifier
	}

	var temperature int

	_, err = fmt.Scan(&temperature)
	if err != nil {
		return constraint{}, fmt.Errorf("invalid temperature input: %w", err)
	}

	if temperature < minTemp || temperature > maxTemp {
		return constraint{}, errInvalidTemperature
	}

	return newConstraint(modifier, temperature), nil
}

func executeSimulation() error {
	deptCount, err := parseDepartmentsCount()
	if err != nil {
		return err
	}

	for range deptCount {
		workerCount, err := parseWorkersCount()
		if err != nil {
			return err
		}

		tempBounds, err := newTemperatureBoundary(minTemp, maxTemp)
		if err != nil {
			return err
		}

		for range workerCount {
			constraint, err := parseConstraint()
			if err != nil {
				return err
			}

			err = tempBounds.adjust(constraint)
			if err != nil {
				return err
			}

			optimalTemp, err := tempBounds.optimalTemperature()
			if err != nil {
				fmt.Println(-1)
				continue
			}

			fmt.Println(optimalTemp)
		}
	}

	return nil
}

func main() {
	if err := executeSimulation(); err != nil {
		fmt.Println(err)
	}
}

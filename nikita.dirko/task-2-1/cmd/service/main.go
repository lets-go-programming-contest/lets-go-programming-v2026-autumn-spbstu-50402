package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidDepartmentCount = errors.New("invalid number of departments provided")
	ErrInvalidWorkerCount     = errors.New("invalid number of workers provided")
	ErrInvalidModifier        = errors.New("invalid format of operation modifier")
	ErrInvalidTemperature     = errors.New("invalid format of temperature value")
	ErrTemperatureBounds      = errors.New("temperature boundaries violation")
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
		return nil, fmt.Errorf("%w: lower bound %d cannot exceed upper bound %d", ErrTemperatureBounds, lower, upper)
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
		return fmt.Errorf("%w: operator '%s' is not supported", ErrInvalidModifier, modifier)
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
		return 0, ErrInvalidDepartmentCount
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
		return 0, ErrInvalidWorkerCount
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
		return "", 0, ErrInvalidModifier
	}

	var temperature int

	_, err = fmt.Scan(&temperature)
	if err != nil {
		return "", 0, fmt.Errorf("invalid temperature input: %w", err)
	}

	if temperature < minTemp || temperature > maxTemp {
		return "", 0, ErrInvalidTemperature
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

func ErrorHandler(err error) {
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, ErrInvalidDepartmentCount):
		fmt.Println("Invalid number of departments")
	case errors.Is(err, ErrInvalidWorkerCount):
		fmt.Println("Invalid number of workers")
	case errors.Is(err, ErrInvalidModifier):
		fmt.Println("Invalid format of operation")
	case errors.Is(err, ErrInvalidTemperature):
		fmt.Println("Invalid format of temperature")
	case errors.Is(err, ErrTemperatureBounds):
		fmt.Println("Temperature bounds error occurred")
	default:
		fmt.Printf("An unexpected system error occurred: %v\n", err)
	}
}

func main() {
	if err := ExecuteSimulation(); err != nil {
		ErrorHandler(err)
	}
}

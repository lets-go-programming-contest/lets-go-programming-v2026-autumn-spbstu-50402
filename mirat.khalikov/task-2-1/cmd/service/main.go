package main

import (
	"errors"
	"fmt"
)

var (
	errInvalidCondition   = errors.New("invalid condition")
	errInvalidDepartments = errors.New("input should be between 1 and 1000")
	errInvalidWorkers     = errors.New("input should be between 1 and 1000")
	errInvalidTemperature = errors.New("input temperature should be between 15 and 30")
	errNonStrictCondition = errors.New("input condition should be non-strict")
)

const (
	minD = 1
	maxD = 1000

	minW = 1
	maxW = 1000

	minT = 15
	maxT = 30
)

type Temperature struct {
	minimal int
	maximal int
}

func (t *Temperature) Update(temperature int, condition string) error {
	switch condition {
	case ">=":
		if temperature > t.minimal {
			t.minimal = temperature
		}
	case "<=":
		if temperature < t.maximal {
			t.maximal = temperature
		}
	default:
		return fmt.Errorf("condition %q: %w", condition, errInvalidCondition)
	}

	return nil
}

func readDepartments() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	if departments < minD || departments > maxD {
		return 0, fmt.Errorf("departments: %d: %w", departments, errInvalidDepartments)
	}

	return departments, nil
}

func readWorkers() (int, error) {
	var workers int

	_, err := fmt.Scan(&workers)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	if workers < minW || workers > maxW {
		return 0, fmt.Errorf("workers: %d: %w", workers, errInvalidWorkers)
	}

	return workers, nil
}

func readCondition() (string, error) {
	var condition string

	_, err := fmt.Scan(&condition)
	if err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}

	if condition != ">=" && condition != "<=" {
		return "", errNonStrictCondition
	}

	return condition, nil
}

func readTemperature() (int, error) {
	var temperature int

	_, err := fmt.Scan(&temperature)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}

	if temperature < minT || temperature > maxT {
		return 0, errInvalidTemperature
	}

	return temperature, nil
}

func main() {
	departments, err := readDepartments()
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		workers, err := readWorkers()
		if err != nil {
			fmt.Println(err)

			return
		}

		temp := Temperature{minimal: minT, maximal: maxT}

		for range workers {
			condition, err := readCondition()
			if err != nil {
				fmt.Println(err)

				return
			}

			temperature, err := readTemperature()
			if err != nil {
				fmt.Println(err)

				return
			}

			err = temp.Update(temperature, condition)
			if err != nil {
				fmt.Println(err)

				return
			}

			if temp.minimal <= temp.maximal {
				fmt.Println(temp.minimal)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

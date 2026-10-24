package main

import (
	"errors"
	"fmt"
)

const (
	minDeps = 1
	maxDeps = 1000

	minWorkers = 1
	maxWorkers = 1000

	minT = 15
	maxT = 30
)

var (
	errDepsValueTooLow         = errors.New("deps input is too low")
	errDepsValueTooHigh        = errors.New("deps input is too high")
	errWorkersValueTooLow      = errors.New("workers input is too low")
	errWorkersValueTooHigh     = errors.New("workers input is too high")
	errTemperatureTooLow       = errors.New("temperature input is too low")
	errTemperatureTooHigh      = errors.New("temperature input is too high")
	errUnknownBoundSign        = errors.New("unknown sign for bound")
	errInvalidDepsInput        = errors.New("invalid input for deps")
	errInvalidWorkersInput     = errors.New("invalid input for workers")
	errInvalidBoundInput       = errors.New("invalid input for new bound")
	errInvalidTemperatureInput = errors.New("invalid input for temperature request")
)

type tempBounds struct {
	lowest  int
	highest int
}

func newTempBounds() *tempBounds {
	return &tempBounds{
		lowest:  minT,
		highest: maxT,
	}
}

func scanDeps() (int, error) {
	var deps int

	if _, err := fmt.Scan(&deps); err != nil {
		return 0, errInvalidDepsInput
	}

	if deps < minDeps {
		return 0, errDepsValueTooLow
	}

	if deps > maxDeps {
		return 0, errDepsValueTooHigh
	}

	return deps, nil
}

func scanWorkers() (int, error) {
	var workers int

	if _, err := fmt.Scan(&workers); err != nil {
		return 0, errInvalidWorkersInput
	}

	if workers < minWorkers {
		return 0, errWorkersValueTooLow
	}

	if workers > maxWorkers {
		return 0, errWorkersValueTooHigh
	}

	return workers, nil
}

func scanBound() (string, int, error) {
	var sign string

	if _, err := fmt.Scan(&sign); err != nil {
		return "", 0, errInvalidBoundInput
	}

	if sign != "<=" && sign != ">=" {
		return "", 0, errUnknownBoundSign
	}

	var temperature int

	if _, err := fmt.Scan(&temperature); err != nil {
		return "", 0, errInvalidTemperatureInput
	}

	if temperature < minT {
		return "", 0, errTemperatureTooLow
	}

	if temperature > maxT {
		return "", 0, errTemperatureTooHigh
	}

	return sign, temperature, nil
}

func processDeps(bounds *tempBounds, workers int) error {
	for range workers {
		sign, temperature, err := scanBound()
		if err != nil {
			return err
		}

		switch sign {
		case "<=":
			if temperature < bounds.highest {
				bounds.highest = temperature
			}

		case ">=":
			if temperature > bounds.lowest {
				bounds.lowest = temperature
			}

		default:
			return errUnknownBoundSign
		}

		if bounds.lowest > bounds.highest {
			fmt.Println(-1)
		} else {
			fmt.Println(bounds.lowest)
		}
	}

	return nil
}

func main() {
	deps, err := scanDeps()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	for range deps {
		workers, err := scanWorkers()
		if err != nil {
			fmt.Println("Error:", err)

			return
		}
		if err = processDeps(newTempBounds(), workers); err != nil {
			fmt.Println("Error:", err)

			return
		}
	}
}

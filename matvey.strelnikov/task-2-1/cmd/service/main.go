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
	errDepsValueTooLow     = errors.New("deps input is too low")
	errDepsValueTooHigh    = errors.New("deps input is too high")
	errWorkersValueTooLow  = errors.New("workers input is too low")
	errWorkersValueTooHigh = errors.New("workers input is too high")
	errTemperatureTooLow   = errors.New("temperature input is too low")
	errTemperatureTooHigh  = errors.New("temperature input is too high")
	errUnknownBoundSign    = errors.New("unknown sign for bound")
)

type boundsInterface interface {
	compareBounds(sign string, t int) error
	getResult() int
}

type tempBounds struct {
	lowest  int
	highest int
}

func newTempBounds(low, high int) tempBounds {
	return tempBounds{
		lowest:  low,
		highest: high,
	}
}

func (b *tempBounds) compareBounds(sign string, temperature int) error {
	switch sign {
	case "<=":
		if temperature < b.highest {
			b.highest = temperature
		}
	case ">=":
		if temperature > b.lowest {
			b.lowest = temperature
		}
	default:
		return errUnknownBoundSign
	}

	return nil
}

func (b *tempBounds) getResult() int {
	if b.lowest > b.highest {
		return -1
	}

	return b.lowest
}

func scanDeps() (int, error) {
	var deps int

	if _, err := fmt.Scan(&deps); err != nil {
		return 0, fmt.Errorf("invalid input for deps: %w", err)
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
		return 0, fmt.Errorf("invalid input for workers: %w", err)
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
		return "", 0, fmt.Errorf("invalid input for new bound: %w", err)
	}

	if sign != "<=" && sign != ">=" {
		return "", 0, errUnknownBoundSign
	}

	var temperature int

	if _, err := fmt.Scan(&temperature); err != nil {
		return "", 0, fmt.Errorf("invalid input for temperature: %w", err)
	}

	if temperature < minT {
		return "", 0, errTemperatureTooLow
	}

	if temperature > maxT {
		return "", 0, errTemperatureTooHigh
	}

	return sign, temperature, nil
}

func processDeps(workers int) error {
	tmp := newTempBounds(minT, maxT)

	var bounds boundsInterface = &tmp

	for range workers {
		sign, temperature, err := scanBound()
		if err != nil {
			return err
		}

		err = bounds.compareBounds(sign, temperature)
		if err != nil {
			return err
		}

		fmt.Println(bounds.getResult())
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

		if err = processDeps(workers); err != nil {
			fmt.Println("Error:", err)

			return
		}
	}
}

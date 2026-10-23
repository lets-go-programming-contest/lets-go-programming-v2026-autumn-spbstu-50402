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
	errDepsValueTooLow = errors.New("Deps input is too low")
	errDepsValueTooHigh = errors.New("Deps input is too high")
	errWorkersValueTooLow = errors.New("Workers input is too low")
	errWorkersValueTooHigh = errors.New("Workers input is too high")
	errTemperatureTooLow =errors.New("Temperature input is too low")
	errTemperatureTooHigh =errors.New("Temperature input is too high")
	errUnknownBoundSign = errors.New("Unknown sign for bound")
)

type tempBounds struct {
	lowest, highest int
}

func newTempBounds() *tempBounds {
	return &tempBounds{
		lowest: minT,
		highest: maxT,
	}
}

func scanDeps() (int, error) {
	var deps int
	_, err := fmt.Scan(&deps)
	if err != nil {
		return 0, fmt.Errorf("invalid input for deps")
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
	_, err := fmt.Scan(&workers)
	if err != nil {
		return 0, fmt.Errorf("invalid input for workers")
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
	_, err := fmt.Scan(&sign)
	if err != nil {
		return "", 0, fmt.Errorf("invalid input for new bound")
	}

	if sign != "<=" && sign != ">=" {
		return "", 0, errUnknownBoundSign
	}

	var t int
	_, err = fmt.Scan(&t)
	if err != nil {
		return "", 0, fmt.Errorf("invalid input for temperature request")
	}

	if t < minT {
		return "", 0, errTemperatureTooLow
	}

	if t > maxT {
		return "", 0, errTemperatureTooHigh
	}

	return sign, t, nil
}

func (b*tempBounds) compareBounds(sign string, t int) error {
	switch sign {
	case "<=":
		b.highest = min(b.highest, t)
	case ">=":
		b.lowest = max(b.lowest, t)
	default:
		return errUnknownBoundSign
	}

	return nil
}

func processDeps(b *tempBounds, workers int) error {
	for i := 0; i < workers; i++ {
		sign, t, err := scanBound()
		if err != nil {
			return err
		}

		err = b.compareBounds(sign,t)
		if err != nil {
			return err
		}

		if b.lowest > b.highest {
			fmt.Println(-1)
		} else {
			fmt.Println(b.lowest)
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

	workers, err := scanWorkers()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := 0; i < deps; i++ {
		if err := processDeps(newTempBounds(), workers); err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}

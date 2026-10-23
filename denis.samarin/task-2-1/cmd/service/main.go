package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

var (
	ErrUnknownOperation     = errors.New("unknown operation")
	ErrNoOptimalTemperature = errors.New("no optimal temperature")
)

type temperatureRange struct {
	lower int
	upper int
}

func newTemperatureRange(lower, upper int) temperatureRange {
	return temperatureRange{
		lower: lower,
		upper: upper,
	}
}

func (tr *temperatureRange) update(sign string, temp int) error {
	switch sign {
	case ">=":
		tr.lower = max(tr.lower, temp)
	case "<=":
		tr.upper = min(tr.upper, temp)
	default:
		return fmt.Errorf("%w: %q", ErrUnknownOperation, sign)
	}

	return nil
}

func (tr *temperatureRange) optimal() (int, error) {
	if tr.lower > tr.upper {
		return 0, ErrNoOptimalTemperature
	}

	return tr.lower, nil
}

func readDepartments() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("read departments count: %w", err)
	}

	return departments, nil
}

func readWorkers() (int, error) {
	var workers int

	_, err := fmt.Scan(&workers)
	if err != nil {
		return 0, fmt.Errorf("read workers count: %w", err)
	}

	return workers, nil
}

func readRequest() (string, int, error) {
	var sign string

	_, err := fmt.Scan(&sign)
	if err != nil {
		return "", 0, fmt.Errorf("read operation: %w", err)
	}

	var temp int

	_, err = fmt.Scan(&temp)
	if err != nil {
		return "", 0, fmt.Errorf("read temperature: %w", err)
	}

	return sign, temp, nil
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

		tempRange := newTemperatureRange(minTemp, maxTemp)

		for range workers {
			sign, temp, err := readRequest()
			if err != nil {
				fmt.Println(err)

				return
			}

			err = tempRange.update(sign, temp)
			if err != nil {
				fmt.Println(err)

				return
			}

			optimal, err := tempRange.optimal()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimal)
		}
	}
}

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

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid number of workers")

			return
		}

		tempRange := newTemperatureRange(minTemp, maxTemp)

		for range workers {
			var (
				sign string
				temp int
			)

			_, err = fmt.Scan(&sign, &temp)
			if err != nil {
				fmt.Println("Invalid input")

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

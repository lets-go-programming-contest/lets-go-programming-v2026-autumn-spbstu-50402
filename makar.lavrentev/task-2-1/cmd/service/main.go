package main

import (
	"errors"
	"fmt"
)

const (
	startMinTemp       = 15
	startMaxTemp       = 30
	invalidTemperature = -1
)

var (
	errUnsupportedOperation = errors.New("unsupported operation")
	errInvalidRange         = errors.New("temperatureRange: lower bound exceeds upper bound")
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

func (tr *temperatureRange) updateTemps(operation string, temperature int) error {
	switch operation {
	case ">=":
		tr.lower = max(tr.lower, temperature)
	case "<=":
		tr.upper = min(tr.upper, temperature)
	default:
		return fmt.Errorf("%w: %s", errUnsupportedOperation, operation)
	}

	return nil
}

func (tr *temperatureRange) current() (int, error) {
	if tr.lower <= tr.upper {
		return tr.lower, nil
	}

	return 0, errInvalidRange
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid amount of departments: ", err)

		return
	}

	for range departments {
		var staff int

		if _, err := fmt.Scan(&staff); err != nil {
			fmt.Println("Invalid amount of staff: ", err)

			return
		}

		tempRange := newTemperatureRange(startMinTemp, startMaxTemp)

		for range staff {
			var (
				operation   string
				temperature int
			)

			if _, err := fmt.Scan(&operation, &temperature); err != nil {
				fmt.Println("Invalid input:", err)

				return
			}

			if err := tempRange.updateTemps(operation, temperature); err != nil {
				fmt.Println("Invalid operation:", err)

				return
			}

			temp, err := tempRange.current()
			if err != nil {
				fmt.Println(invalidTemperature)
			} else {
				fmt.Println(temp)
			}
		}
	}
}

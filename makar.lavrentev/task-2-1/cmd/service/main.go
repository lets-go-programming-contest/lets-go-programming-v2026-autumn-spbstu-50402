package main

import (
	"errors"
	"fmt"
)

const (
	startMinTemp = 15
	startMaxTemp = 30
)

var errUnsupportedOperation = errors.New("unsupported operation")

type temperatureRange struct {
	lower int
	upper int
}

func newTempRange(lower, upper int) temperatureRange {
	return temperatureRange{
		lower: lower,
		upper: upper,
	}
}

func (tr *temperatureRange) updateTemps(operation string, temperature int) error {
	if operation != "<=" && operation != ">=" {
		return fmt.Errorf("%w: %s", errUnsupportedOperation, operation)
	}

	if operation == ">=" && temperature > tr.lower {
		tr.lower = temperature
	} else if operation == "<=" && temperature < tr.upper {
		tr.upper = temperature
	}

	return nil
}

func (tr *temperatureRange) current() int {
	if tr.lower <= tr.upper {
		return tr.lower
	}

	return -1
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

		tempRange := newTempRange(startMinTemp, startMaxTemp)

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

			fmt.Println(tempRange.current())
		}
	}
}

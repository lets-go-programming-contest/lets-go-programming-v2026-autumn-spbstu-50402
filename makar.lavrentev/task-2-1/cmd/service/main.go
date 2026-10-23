package main

import (
	"fmt"
)

const (
	startMinTemp = 15
	startMaxTemp = 30
)

type TemperatureRange struct {
	lower int
	upper int
}

func NewTempRange(lower, upper int) TemperatureRange {
	return TemperatureRange{
		lower: lower,
		upper: upper,
	}
}

func (tr *TemperatureRange) UpdateTemps(operation string, temperature int) {
	if operation == ">=" && temperature > tr.lower {
		tr.lower = temperature
	} else if operation == "<=" && temperature < tr.upper {
		tr.upper = temperature
	}
}

func (tr *TemperatureRange) OutputTemps() {
	if tr.lower <= tr.upper {
		fmt.Println(tr.lower)
	} else {
		fmt.Println(-1)
	}
}

func processEmployeeInput(staff int) {
	tempRange := NewTempRange(startMinTemp, startMaxTemp)

	for range staff {
		var (
			operation   string
			temperature int
		)

		if _, err := fmt.Scan(&operation, &temperature); err != nil {
			fmt.Println("Invalid input: ", err)

			return
		}

		if operation != "<=" && operation != ">=" {
			fmt.Println("Invalid operation input")

			return
		}

		tempRange.UpdateTemps(operation, temperature)
		tempRange.OutputTemps()
	}
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid amount of departments")

		return
	}

	for range departments {
		var staff int

		if _, err := fmt.Scan(&staff); err != nil {
			fmt.Println("Invalid amount of staff")

			return
		}

		processEmployeeInput(staff)
	}
}

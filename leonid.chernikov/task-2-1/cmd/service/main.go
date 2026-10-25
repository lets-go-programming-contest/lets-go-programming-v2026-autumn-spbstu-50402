package main

import (
	"errors"
	"fmt"
)

const (
	minTemp   = 15
	maxTemp   = 30
	highCount = 1000
	lowCount  = 1
)

var (
	errUnknownSign = errors.New("unknown comparison sign")
)

type department struct {
	low  int
	high int
}

func (d *department) update(sign string, temp int) error {
	switch sign {
	case ">=":
		if temp > d.low {
			d.low = temp
		}
	case "<=":
		if temp < d.high {
			d.high = temp
		}
	default:
		return errUnknownSign
	}

	return nil
}

func (d *department) optimal() int {
	if d.low > d.high {
		return -1
	}

	return d.low
}

func readInput() (string, int, error) {
	var (
		compSymb string
		temp     int
	)

	_, err := fmt.Scan(&compSymb, &temp)
	if err != nil {
		return "", 0, fmt.Errorf("read employee request: %w", err)
	}

	return compSymb, temp, nil
}

func main() {
	var departNumber int

	_, err := fmt.Scan(&departNumber)
	if err != nil {
		fmt.Println("Input department number error")

		return
	}

	if departNumber > highCount || departNumber < lowCount {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departNumber {
		dep := department{low: minTemp, high: maxTemp}

		var employeeNumber int

		_, err := fmt.Scan(&employeeNumber)
		if err != nil {
			fmt.Println("Input employee number error")

			return
		}

		if employeeNumber > highCount || employeeNumber < lowCount {
			fmt.Println("Invalid number of employees")

			return
		}

		for range employeeNumber {
			compSymb, temp, err := readInput()
			if err != nil {
				fmt.Println("Input employee request error")

				return
			}

			err = dep.update(compSymb, temp)
			if err != nil {
				fmt.Println("Invalid employee request:", err)

				return
			}

			fmt.Println(dep.optimal())
		}
	}
}

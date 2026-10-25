package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
	highMagicCount = 1000
	lowMagicCount = 1
)

type department struct {
	low  int
	high int
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

func (d *department) updateTempBounds(compSymb string, temp int) {
	if d.low > d.high {
		fmt.Println("-1")

		return
	}

	switch compSymb {
	case ">=":
		switch {
		case temp > d.high:
			d.low = temp

			fmt.Println("-1")
		case temp < d.low:
			fmt.Println(d.low)
		default:
			d.low = temp

			fmt.Println(d.low)
		}
	case "<=":
		switch {
		case temp > d.high:
			fmt.Println(d.low)
		case temp < d.low:
			d.high = temp

			fmt.Println("-1")
		default:
			d.high = temp

			fmt.Println(d.low)
		}
	default:
		fmt.Println("invalid comparison symbol")
	}
}

func main() {
	var departNumber int

	_, err := fmt.Scan(&departNumber)
	if err != nil {
		fmt.Println("Input department number error")

		return
	}

	if departNumber > highMagicCount || departNumber < lowMagicCount {
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

		if employeeNumber > highMagicCount || employeeNumber < lowMagicCount {
			fmt.Println("Invalid number of employees")

			return
		}

		for range employeeNumber {
			compSymb, temp, err := readInput()
			if err != nil {
				fmt.Println("Input employee request error")

				return
			}

			dep.updateTempBounds(compSymb, temp)
		}
	}
}

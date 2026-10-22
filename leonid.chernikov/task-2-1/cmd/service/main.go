package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

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

func updateTempBounds(compSymb string, temp int, lmin *int, lmax *int) {
	if *lmin > *lmax {
		fmt.Println("-1")

		return
	}

	switch compSymb {
	case ">=":
		switch {
		case temp > *lmax:
			*lmin = temp

			fmt.Println("-1")
		case temp < *lmin:
			fmt.Println(*lmin)
		default:
			*lmin = temp

			fmt.Println(*lmin)
		}
	case "<=":
		switch {
		case temp > *lmax:
			fmt.Println(*lmin)
		case temp < *lmin:
			*lmax = temp

			fmt.Println("-1")
		default:
			*lmax = temp

			fmt.Println(*lmin)
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

	if departNumber > 1000 || departNumber < 1 {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departNumber {
		localMin := minTemp
		localMax := maxTemp

		var employeeNumber int

		_, err := fmt.Scan(&employeeNumber)
		if err != nil {
			fmt.Println("Input employee number error")

			return
		}

		if employeeNumber > 1000 || employeeNumber < 1 {
			fmt.Println("Invalid number of employees")

			return
		}

		for range employeeNumber {
			compSymb, temp, err := readInput()
			if err != nil {
				fmt.Println("Input employee request error")

				return
			}

			updateTempBounds(compSymb, temp, &localMin, &localMax)
		}
	}
}

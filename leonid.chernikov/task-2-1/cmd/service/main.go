package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func readInput() (compSymb string, temp int, err error) {

	_, err = fmt.Scan(&compSymb, &temp)
	return compSymb, temp, err
}

func updateTempBounds(compSymb string, temp int, lmin *int, lmax *int) {

	if *lmin > *lmax {
		fmt.Println("-1")
		return
	}

	switch compSymb {
	case ">=":
		if temp > *lmax {
			fmt.Println("-1")
			*lmin = temp
		}
		if temp < *lmin {
			fmt.Println(*lmin)
		}
		if temp >= *lmin && temp <= *lmax {
			*lmin = temp
			fmt.Println(*lmin)
		}
	case "<=":
		if temp > *lmax {
			fmt.Println(*lmin)
		}
		if temp < *lmin {
			fmt.Println("-1")
			*lmax = temp
		}
		if temp >= *lmin && temp <= *lmax {
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

	for i := 0; i < departNumber; i++ {

		localMin := minTemp
		localMax := maxTemp

		var employeeNumber int

		_, err := fmt.Scan(&employeeNumber)
		if err != nil {
			fmt.Println("Input employee number error")
			return
		}

		for j := 0; j < employeeNumber; j++ {

			compSymb, temp, err := readInput()
			if err != nil {
				fmt.Println("Input employee request error")
				return
			}

			updateTempBounds(compSymb, temp, &localMin, &localMax)
		}
	}
}

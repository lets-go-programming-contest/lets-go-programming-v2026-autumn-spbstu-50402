package main

import "fmt"

func readInput() (compSymb string, temp int){
	fmt.Scan(&compSymb, &temp)
	return compSymb, temp
}

func updateTempBounds(compSymb string, temp int, lmin *int, lmax *int) {
	switch compSymb {
	case ">=":
		if temp > *lmax {
			fmt.Println("-1")
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
		}
		if temp >= *lmin && temp <= *lmax {
			*lmax = temp
			fmt.Println(*lmin)
		}
	default:
		fmt.Println("invalid comprasion symbol")
	}
}

func main() {
	var constMin int = 15
	var constMax int = 30

	var departNumber int
	fmt.Scan(&departNumber)

	for i := 0; i < departNumber; i++ {
		var localMin int = constMin
		var localMax int = constMax
		var employeeNumber int
		fmt.Scan(&employeeNumber)
		for j := 0; j < employeeNumber; j++ {
			compSymb, temp := readInput()
			updateTempBounds(compSymb, temp , &localMin, &localMax)
		}
	}
}
	
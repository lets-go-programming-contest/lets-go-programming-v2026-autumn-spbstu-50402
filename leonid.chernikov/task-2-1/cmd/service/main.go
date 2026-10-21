package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func readInput() (compSymb string, temp int){
	fmt.Scan(&compSymb, &temp)
	return compSymb, temp
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
		fmt.Println("invalid comprasion symbol")
	}
}

func main() {

	var departNumber int
	fmt.Scan(&departNumber)

	for i := 0; i < departNumber; i++ {
		localMin := minTemp
		localMax := maxTemp
		var employeeNumber int
		fmt.Scan(&employeeNumber)
		for j := 0; j < employeeNumber; j++ {
			compSymb, temp := readInput()
			updateTempBounds(compSymb, temp , &localMin, &localMax)
		}
	}
}
	
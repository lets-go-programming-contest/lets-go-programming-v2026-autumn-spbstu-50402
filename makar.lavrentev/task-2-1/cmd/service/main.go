package main

import (
	"fmt"
)

const startMinTemp = 15
const startMaxTemp = 30

func UpdateTemps(oprtn string, temp int, lower *int, upper *int) {
	if oprtn == ">=" && temp > *lower {
		*lower = temp
	} else if oprtn == "<=" && temp < *upper {
		*upper = temp
	}

	if *lower <= *upper {
		fmt.Println(*lower)
	} else {
		fmt.Println(-1)
	}
}

func processEmployeeInput(staff int) {
	lower := startMinTemp
	upper := startMaxTemp

	for range staff {
		var (
			oprtn string
			temp  int
		)

		if _, err := fmt.Scan(&oprtn, &temp); err != nil {
			fmt.Println("Invalid input: ", err)

			return
		}

		if oprtn != "<=" && oprtn != ">=" {
			fmt.Println("Invalid operation input")

			return
		}

		UpdateTemps(oprtn, temp, &lower, &upper)
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

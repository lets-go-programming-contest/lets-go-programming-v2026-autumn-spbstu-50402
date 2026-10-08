package main

import (
	"fmt"
)

func processStaff(staff int) {
	minTemp := 15
	maxTemp := 30

	for range staff {
		var (
			oprtn string
			temp int
		)

		if _, err := fmt.Scan(&oprtn, &temp); err != nil {
			fmt.Println("Invalid input")

			return
		}

		if (oprtn != "<=" && oprtn != ">=") || temp < 15 || temp > 30 {
			fmt.Println("Invalid input")

			return
		}

		if oprtn == ">=" && temp > minTemp {
			minTemp = temp
		} else if oprtn == "<=" && temp < maxTemp {
			maxTemp = temp
		}

		if minTemp <= maxTemp {
			fmt.Println(minTemp)
		} else {
			fmt.Println(-1)
		}
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

		processStaff(staff)
	}
}

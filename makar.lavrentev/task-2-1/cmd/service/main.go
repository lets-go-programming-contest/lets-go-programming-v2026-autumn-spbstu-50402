package main

import (
	"fmt"
)

func main() {
	var departments, staff int
	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid amount of departments")

		return
	}

	for _ = range departments {
		_, err = fmt.Scan(&staff)
		if err != nil {
			fmt.Println("Invalid amount of staff")

			return
		}

		minTemp := 15
		maxTemp := 30

		for _ = range staff {
			var oprtn string
			_, err = fmt.Scan(&oprtn)
			if err != nil || (oprtn != "<=" && oprtn != ">=") {
				fmt.Println("Invalid operation")

				return
			}

			var temp int
			_, err = fmt.Scan(&temp)
			if err != nil || (temp < 15 || temp > 30) {
				fmt.Println("Invalid temperature")

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
}

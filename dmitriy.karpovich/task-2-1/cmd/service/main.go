package main

import "fmt"

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		var workers int
		if _, err := fmt.Scan(&workers); err != nil {
			fmt.Println("Invalid number of workers")

			return
		}

		low := 15
		high := 30

		for range workers {
			var operation string
			if _, err := fmt.Scan(&operation); err != nil {
				fmt.Println("Invalid format of operation")

				return
			}

			var temperature int
			if _, err := fmt.Scan(&temperature); err != nil {
				fmt.Println("Invalid format of temperature")

				return
			}

			switch operation {
			case ">=":
				if temperature > low {
					low = temperature
				}
			case "<=":
				if temperature < high {
					high = temperature
				}
			}

			if low <= high {
				fmt.Println(low)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

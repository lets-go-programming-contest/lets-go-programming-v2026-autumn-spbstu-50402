package main

import "fmt"

func main() {
	var parts int
	if _, err := fmt.Scan(&parts); err != nil {
		fmt.Println("Invalid number of parts")

		return
	}

	for range parts {
		var workers int
		if _, err := fmt.Scan(&workers); err != nil {
			fmt.Println("Invalid workers input")

			return
		}

		low := 15
		high := 30

		for range workers {
			var operation string
			if _, err := fmt.Scan(&operation); err != nil {
				fmt.Println("Invalid change cmd input")

				return
			}

			var tmp int
			if _, err := fmt.Scan(&tmp); err != nil {
				fmt.Println("Invalid temperature input")

				return
			}

			switch operation {
			case ">=":
				low = max(low, tmp)
			case "<=":
				high = min(high, tmp)
			}

			if low <= high {
				fmt.Println(low)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

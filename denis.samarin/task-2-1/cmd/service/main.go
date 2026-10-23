package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid number of workers")

			return
		}

		low, high := minTemp, maxTemp

		for range workers {
			var (
				sign string
				temp int
			)

			_, err = fmt.Scan(&sign, &temp)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			switch sign {
			case ">=":
				low = max(low, temp)
			case "<=":
				high = min(high, temp)
			default:
				fmt.Println("Invalid operation")

				return
			}

			if low <= high {
				fmt.Println(low)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

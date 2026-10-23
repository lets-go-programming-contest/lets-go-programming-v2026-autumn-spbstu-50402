package main

import "fmt"

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	low, high := 15, 30

	for range departments {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid number of workers")
			return
		}

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

			if sign == ">=" {
				if temp > low {
					low = temp
				}
			} else if sign == "<=" {
				if temp < high {
					high = temp
				}
			} else {
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

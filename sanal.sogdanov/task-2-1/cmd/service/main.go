package main

import "fmt"

func main() {
	var departmentCount int

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("Invalid input the number of departments")

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err = fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("invalid input the number of workers")

			return
		}

		low, high := 15, 30

		for range employeeCount {
			var (
				operator    string
				temperature int
			)

			_, err = fmt.Scan(&operator, &temperature)
			if err != nil {
				fmt.Println("Invalid input the temperature condition")

				return
			}

			switch operator {
			case ">=":
				if temperature > low {
					low = temperature
				}
			case "<=":
				if temperature < high {
					high = temperature
				}
			default:
				fmt.Println("Expected <= or >=")

				return
			}

			if low > high {
				fmt.Println(-1)
			} else {
				fmt.Println(low)
			}
		}
	}
}

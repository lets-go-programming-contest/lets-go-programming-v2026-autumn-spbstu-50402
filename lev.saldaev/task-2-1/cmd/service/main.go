package main

import "fmt"

func main() {
	var (
		departmentCount int
		operator        string
	)

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Println("invalid department amount")

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err = fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Println("invalid employee amount")

			return
		}

		minTemp := 15
		maxTemp := 30

		for range employeeCount {
			_, err = fmt.Scan(&operator)
			if err != nil {
				fmt.Println("invalid input")

				return
			}

			var temperature int

			_, err = fmt.Scan(&temperature)
			if err != nil {
				fmt.Println("invalid input")

				return
			}

			switch operator {
			case ">=":
				minTemp = max(minTemp, temperature)
			case "<=":
				maxTemp = min(maxTemp, temperature)
			default:
				fmt.Println("invalid input")

				return
			}

			if maxTemp < minTemp {
				fmt.Println("-1")
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}

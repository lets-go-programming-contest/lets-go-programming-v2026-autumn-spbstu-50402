package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input the number of departments")
		return
	}

	for i := 0; i < n; i++ {
		var k int
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("invalid input the number of workers")
			return
		}

		low, high := 15, 30

		for j := 0; j < k; j++ {
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

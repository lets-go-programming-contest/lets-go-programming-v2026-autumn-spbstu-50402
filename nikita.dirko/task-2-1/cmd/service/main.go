package main

import "fmt"

func main() {
	var N int
	if _, err := fmt.Scan(&N); err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for range N {
		var K int
		if _, err := fmt.Scan(&K); err != nil {
			fmt.Println("Invalid number of workers")
			return
		}

		for range K {
			var modifier string
			if _, err := fmt.Scan(&modifier); err != nil {
				fmt.Println("Invalid format of operation")
				return
			}

			var temperature int
			if _, err := fmt.Scan(&temperature); err != nil {
				fmt.Println("Invalid format of temperature")
				return
			}
		}
	}
}

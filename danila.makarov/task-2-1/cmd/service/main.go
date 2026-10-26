package main

import "fmt"

func main() {
	var depCount int

	_, err := fmt.Scan(&depCount)
	if err != nil {
		fmt.Println("Invalid departments count")

		return
	}

	for range depCount {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid workers count")

			return
		}

		for range workers {
			var (
				sign string
				temp int
			)

			_, err = fmt.Scan(&sign, &temp)
			if err != nil {
				fmt.Println("Invalid temperature")

				return
			}

			fmt.Println("wish:", sign, temp)
		}
	}
}

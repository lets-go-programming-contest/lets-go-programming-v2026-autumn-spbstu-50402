package main

import (
	"fmt"
)

func main() {
	var parts int

	_, err := fmt.Scan(&parts)
	if err != nil {
		fmt.Println("Invalid parts input")

		return
	}

	for range parts {

		var workers string

		_, err := fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid workers input")

			return
		}

		for range workers {
			var changeCmd string

			_, err := fmt.Scan(&changeCmd)
			if err != nil {
				fmt.Println("Invalid change cmd input")

				return
			}

			var temperature int

			_, err = fmt.Scan(&temperature)
			if err != nil {
				fmt.Println("Invalid temperature input")

				return
			}
		}
	}
}

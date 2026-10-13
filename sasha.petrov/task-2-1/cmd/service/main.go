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

		minTmp := 15
		maxTmp := 30

		var workers int

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

			switch changeCmd {
			case ">=":
				minTmp = max(minTmp, temperature)
			case "<=":
				maxTmp = min(maxTmp, temperature)
			}

			if minTmp > maxTmp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTmp)
			}
		}
	}
}

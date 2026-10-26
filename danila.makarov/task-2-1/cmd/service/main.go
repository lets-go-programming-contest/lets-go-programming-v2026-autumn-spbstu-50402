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

		fmt.Println("department, workers:", workers)
	}
}

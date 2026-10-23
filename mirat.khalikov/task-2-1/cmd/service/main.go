package main

import "fmt"

func main() {
	var (
		departments int
		err         error
	)

	_, err = fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	for range departments {
		var workers int
		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		for range workers {
			var condition string
			_, err = fmt.Scan(&condition)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
			var temperature int
			_, err = fmt.Scan(&temperature)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
		}
	}
}

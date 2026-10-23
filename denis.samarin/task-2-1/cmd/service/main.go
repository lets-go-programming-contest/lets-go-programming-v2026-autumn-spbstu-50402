package main

import "fmt"

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for range departments {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid number of workers")
			return
		}
	}
}

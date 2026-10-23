package main

import "fmt"

func main() {
	var (
		department, employee int
	)

	_, err := fmt.Scanln(&department)
	if err != nil {
		fmt.Println("Invalid department argument")
		return
	}

	for range department {
		_, err := fmt.Scanln(&employee)
		if err != nil {
			fmt.Println("Invalid employee argument")
		}

	}
}

package main

import "fmt"

func getOptimalTemperature(k int) int

func main() {
	var n int
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println("Invalid number of departments")
		return
	}

	for range n {
		var k int
		if _, err := fmt.Scanln(&k); err != nil {
			fmt.Println("Invalid number of employees")
		}
		fmt.Println(getOptimalTemperature(k))
	}
}

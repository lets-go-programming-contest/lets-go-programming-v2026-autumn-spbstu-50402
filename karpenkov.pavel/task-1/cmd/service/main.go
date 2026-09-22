package main

import "fmt"

func main() {
	var a, b int
	var operator string
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("input error -", err)
		return
	}
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("input error -", err)
		return
	}
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("input error -", err)
		return
	}
	if operator == "+" {
		fmt.Println(a + b)
	} else if operator == "-" {
		fmt.Println(a - b)
	} else if operator == "*" {
		fmt.Println(a * b)
	} else if operator == "/" {
		if b == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		fmt.Println(a / b)
	} else {
		fmt.Println("Error: invalid operator")
		return
	}
}

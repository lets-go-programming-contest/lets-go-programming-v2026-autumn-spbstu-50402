package main

import "fmt"

func main() {
	var a, b int
	var operand string
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
	_, err = fmt.Scanln(&operand)
	if err != nil {
		fmt.Println("input error -", err)
		return
	}
	if operand == "+" {
		fmt.Println(a + b)
	}
	if operand == "-" {
		fmt.Println(a - b)
	}

}

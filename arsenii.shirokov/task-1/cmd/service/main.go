package main

import "fmt"

func main() {
	var (
		a, b int
		op   string
	)
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
}

package main

import "fmt"

func main() {
	var (
		a, b   int
		op     string
		result int
	)

	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scanln(&op); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		result = a + b

	case "-":
		result = a - b

	case "*":
		result = a * b

	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = a / b

	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}

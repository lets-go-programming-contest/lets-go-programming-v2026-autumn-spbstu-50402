package main

import "fmt"

func main() {
	var a int
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var b int
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var op string
	if _, err := fmt.Scanln(&op); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	var res int
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		res = a / b
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(res)
}

package main

import "fmt"

func main() {
	var lhs int
	_, err := fmt.Scan(&lhs)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var rhs int
	_, err = fmt.Scan(&rhs)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var operation string
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	var result int
	switch operation {
	case "+":
		result = lhs + rhs
	case "-":
		result = lhs - rhs
	case "*":
		result = lhs * rhs
	case "/":
		if rhs == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = lhs / rhs
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(result)
}

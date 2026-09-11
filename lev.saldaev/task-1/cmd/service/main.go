package main

import "fmt"

func main() {
	var (
		n1 int
		n2 int
		op string
	)

	_, err := fmt.Scan(&n1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&n2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&op)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		fmt.Println(n1 + n2)
	case "-":
		fmt.Println(n1 - n2)
	case "*":
		fmt.Println(n1 * n2)
	case "/":
		if n2 == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(n1 / n2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}

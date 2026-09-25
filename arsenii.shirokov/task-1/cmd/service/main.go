package main

import (
	"errors"
	"fmt"
)

var (
	errDivisionByZero   = errors.New("division by zero")
	errInvalidOperation = errors.New("invalid operation")
)

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errDivisionByZero
		}
		return a / b, nil
	default:
		return 0, errInvalidOperation
	}
}

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

	res, err := calculate(a, b, op)
	switch {
	case err == nil:
	case errors.Is(err, errDivisionByZero):
		fmt.Println("Division by zero")
		return
	case errors.Is(err, errInvalidOperation):
		fmt.Println("Invalid operation")
		return
	default:
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

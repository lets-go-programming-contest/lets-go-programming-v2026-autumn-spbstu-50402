package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero   = errors.New("Division by zero")
	ErrInvalidOperation = errors.New("Invalid operation")
)

func calculate(a int, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidOperation
	}
}

func main() {
	var (
		a, b     int
		operator string
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
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}
	res, err1 := calculate(a, b, operator)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(res)
}

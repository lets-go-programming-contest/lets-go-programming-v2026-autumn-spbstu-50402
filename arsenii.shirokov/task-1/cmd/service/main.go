package main

import (
	"errors"
	"fmt"
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
			return 0, errors.New("Division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("Invalid operation")
	}
}

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

	res, err := calculate(a, b, op)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}

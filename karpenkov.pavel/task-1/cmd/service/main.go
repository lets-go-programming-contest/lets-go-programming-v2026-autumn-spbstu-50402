package main

import (
	"errors"
	"fmt"
)

func calculate(a int, b int, operator string) (int, error) {
	if operator == "+" {
		return a + b, nil
	} else if operator == "-" {
		return a - b, nil
	} else if operator == "*" {
		return a * b, nil
	} else if operator == "/" {
		if b == 0 {
			return 0, errors.New("Division by zero")
		}
		return a / b, nil
	} else {
		return 0, errors.New("Invalid operation")
	}
}

func main() {
	var (
		a, b int
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
	if err1 == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err1)
	}
}

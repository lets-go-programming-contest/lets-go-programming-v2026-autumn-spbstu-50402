package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
		sign string
	)

	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&sign)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	answer := num1

	switch sign {
	case "+":
		answer += num2
	case "-":
		answer -= num2
	case "*":
		answer *= num2
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		answer /= num2
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(answer)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Invalid first operand")
		return
	}
	firstOperandLine := scanner.Text()
	firstOperandNum, err := strconv.Atoi(firstOperandLine)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if !scanner.Scan() {
		fmt.Println("Invalid second operand")
		return
	}
	secondOperandLine := scanner.Text()
	secondOperandNum, err := strconv.Atoi(secondOperandLine)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if !scanner.Scan() {
		fmt.Println("Invalid operation")
		return
	}
	operation := scanner.Text()

	var result int
	switch operation {
	case "/":
		if secondOperandNum == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = firstOperandNum / secondOperandNum
	case "*":
		result = firstOperandNum * secondOperandNum
	case "+":
		result = firstOperandNum + secondOperandNum
	case "-":
		result = firstOperandNum - secondOperandNum
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}

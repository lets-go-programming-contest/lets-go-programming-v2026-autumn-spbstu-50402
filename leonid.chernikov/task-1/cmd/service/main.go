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
	first_operand_line := scanner.Text()
	first_operand_num, err := strconv.Atoi(first_operand_line)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if !scanner.Scan() {
		fmt.Println("Invalid second operand")
		return
	}
	second_operand_line := scanner.Text()
	second_operand_num, err := strconv.Atoi(second_operand_line)
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
		if second_operand_num == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = first_operand_num / second_operand_num
	case "*":
		result = first_operand_num * second_operand_num
	case "+":
		result = first_operand_num + second_operand_num
	case "-":
		result = first_operand_num - second_operand_num
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(result)
}
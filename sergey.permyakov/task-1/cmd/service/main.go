package main

import "fmt"

func main() {
	var frstOp int
	_, err1 := fmt.Scan(&frstOp)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	var scndOp int
	_, err2 := fmt.Scan(&scndOp)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	var operation string
	_, err3 := fmt.Scan(&operation)
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}

	var res int
	switch operation {
	case "+":
		res = frstOp + scndOp
	case "-":
		res = frstOp - scndOp
	case "*":
		res = frstOp * scndOp
	case "/":
		if scndOp == 0 {
			fmt.Println("Division by zero")
			return
		}
		res = frstOp / scndOp
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(res)
}

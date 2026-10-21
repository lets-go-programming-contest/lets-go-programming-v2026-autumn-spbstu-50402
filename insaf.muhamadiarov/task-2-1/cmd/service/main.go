package main

import (
	"fmt"
	"errors"
)

type TempRange struct{
	LowLevel
	HighLevel
}

func readLine() (string, int, error) {
	var (
		operation string
		temp int
	)

	_, err := fmt.Scan(&operation)
	if err != nil {
		return "", 0, fmt.Errorf("faild to read the operaton: %w", err)
	}

	_, err := fmt.Scan(&temp)
	if err != nil {
		return "", 0, fmt.Errorf("faild to read the temperature: %w", err)
	}

	return operation, temp, nil	
}

func main() {
	var numberOfDepartments int
	_, err := fmt.Scan(&numberOfDepartments)
	if err != nil {
		fmt.Printf("Error reading the number of departments: %w\n", err)
		return
	}

	for range numberOfDepartments {
		var numberOfEmployees int
		_, err := fmt.Scan(&numberOfEmployees)
		if err != nil {
			fmt.Printf("Error reading the number of employees: %w\n", err)
			return
		}

		tempRange := TempRange{15, 30}

		for range numberOfEmployees {
			operation, temp, err := readLine()
			if err != nil {
				fmt.Printf("Error reading the Line: %w\n", err)
				return
			}


			fmt.Println()
		}
	}
}





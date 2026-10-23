package main

import "fmt"

const (
	MinTemp = 15
	MaxTemp = 30
)

type TempRange struct {
	minTemp int
	maxTemp int
}

func NewTempRange(minTemp_ int, maxTemp_ int) *TempRange {
	return &TempRange{minTemp: minTemp_, maxTemp: maxTemp_}
}

func (tRange *TempRange) Apply(operator string, temperature int) error {
	switch operator {
	case ">=":
		if temperature > tRange.minTemp {
			tRange.minTemp = temperature
		}
	case "<=":
		if temperature < tRange.maxTemp {
			tRange.maxTemp = temperature
		}
	default:
		return fmt.Errorf("unknown operator %q", operator)
	}
	return nil
}

func (tRange *TempRange) Value() int {
	if tRange.maxTemp < tRange.minTemp {
		return -1
	}
	return tRange.minTemp
}

func main() {
	var (
		departmentCount int
		operator        string
	)

	_, err := fmt.Scan(&departmentCount)
	if err != nil {
		fmt.Printf("invalid department amount: %v\n", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		_, err = fmt.Scan(&employeeCount)
		if err != nil {
			fmt.Printf("invalid employee amount: %v\n", err)

			return
		}

		state := NewTempRange(MinTemp, MaxTemp)

		for range employeeCount {
			_, err = fmt.Scan(&operator)
			if err != nil {
				fmt.Printf("invalid input: %v\n", err)

				return
			}

			var temperature int

			_, err = fmt.Scan(&temperature)
			if err != nil {
				fmt.Printf("invalid input: %v\n", err)

				return
			}

			if err := state.Apply(operator, temperature); err != nil {
				fmt.Printf("apply: %v\n", err)
				return
			}

			fmt.Println(state.Value())
		}
	}
}

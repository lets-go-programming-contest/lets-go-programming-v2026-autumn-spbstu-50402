package main

import "fmt"

const (
	minAllowedTemperature = 15
	maxAllowedTemperature = 30
	noSuitableTemperature = -1
)

type temperatureController interface {
	applyConstraint(operator string, temperature int) error
	optimalTemperature() int
}

type temperatureRange struct {
	low  int
	high int
}

func (state *temperatureRange) applyConstraint(operator string, temperature int) error {
	switch operator {
	case ">=":
		if temperature > state.low {
			state.low = temperature
		}
	case "<=":
		if temperature < state.high {
			state.high = temperature
		}
	default:
		return fmt.Errorf("expected <= or >=, got %q", operator)
	}

	return nil
}

func (state *temperatureRange) optimalTemperature() int {
	if state.low > state.high {
		return noSuitableTemperature
	}

	return state.low
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println("Invalid number of departments:", err)

		return
	}

	for range departmentCount {
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println("Invalid number of employees:", err)

			return
		}

		var controller temperatureController = &temperatureRange{
			low:  minAllowedTemperature,
			high: maxAllowedTemperature,
		}

		for range employeeCount {
			var (
				operator    string
				temperature int
			)

			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				fmt.Println("Invalid temperature condition:", err)

				return
			}

			if err := controller.applyConstraint(operator, temperature); err != nil {
				fmt.Println("Invalid temperature condition:", err)

				return
			}

			fmt.Println(controller.optimalTemperature())
		}
	}
}

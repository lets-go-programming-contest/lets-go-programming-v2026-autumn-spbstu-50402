package main

import "fmt"

const (
	minAllowedTemperature = 15
	maxAllowedTemperature = 30
	noSuitableTemperature = -1
)

var (
	errInvalidOperator       = fmt.Errorf("expected <= or >=")
	errNoSuitableTemperature = fmt.Errorf("no suitable temperature")
)

type temperatureRange struct {
	low  int
	high int
}

func newTemperatureRange() *temperatureRange {
	return &temperatureRange{
		low:  minAllowedTemperature,
		high: maxAllowedTemperature,
	}
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
		return fmt.Errorf("%w, got %q", errInvalidOperator, operator)
	}

	return nil
}

func (state *temperatureRange) optimalTemperature() (int, error) {
	if state.low > state.high {
		return 0, errNoSuitableTemperature
	}

	return state.low, nil
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

		controller := newTemperatureRange()

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

			result, err := controller.optimalTemperature()
			if err != nil {
				fmt.Println(noSuitableTemperature)

				continue
			}

			fmt.Println(result)
		}
	}
}

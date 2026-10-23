package main

import "fmt"

const (
	minDepartment  = 1
	maxDepartment  = 1000
	minStaff       = 1
	maxStaff       = 1000
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var departments int = 0
	_, err := fmt.Scan(&departments)

	if err != nil {
		fmt.Println("Invalid input departments")
		return
	}

	if departments < minDepartment || departments > maxDepartment {
		fmt.Printf("Depatments must be form %d to %d", minDepartment, maxDepartment)
		return
	}

	for range departments {
		var staff int = 0
		_, err = fmt.Scan(&staff)

		if err != nil {
			fmt.Println("Invalid input staff")
			return
		}

		if staff < minStaff || staff > maxStaff {
			fmt.Printf("Staff must be from %d to %d", minStaff, maxStaff)
			return
		}

		var maxTempOfDepartment int = maxTemperature
		var minTempOfDepartment int = minTemperature

		for range staff {
			var condition string = ""
			_, err = fmt.Scan(&condition)

			if err != nil {
				fmt.Println("Invalid input condition")
				return
			}

			if condition != ">=" && condition != "<=" {
				fmt.Println("Condition of temperature must be '<=' or '>='")
				return
			}

			var temperature int = minTemperature
			_, err = fmt.Scan(&temperature)

			if err != nil {
				fmt.Println("Invalid input temperature value")
				return
			}

			if temperature < minTemperature || temperature > maxTemperature {
				fmt.Printf("Temperature must be from %d to %d", minTemperature, maxTemperature)
				return
			}

			if condition == ">=" && temperature > minTempOfDepartment {
				minTempOfDepartment = temperature
			} else if condition == "<=" && temperature < maxTempOfDepartment {
				maxTempOfDepartment = temperature
			}

			if maxTempOfDepartment < minTempOfDepartment {
				fmt.Println("-1")
			} else {
				fmt.Println(minTempOfDepartment)
			}
		}
	}
}

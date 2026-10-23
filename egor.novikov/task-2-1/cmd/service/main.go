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

func parseDepartment() (int, error) {
	var departments int
	_, err := fmt.Scan(&departments)

	if err != nil {
		return 0, fmt.Errorf("Invalid input departments")
	}

	if departments < minDepartment || departments > maxDepartment {
		return 0, fmt.Errorf(
			"Depatments must be form %d to %d",
			minDepartment,
			maxDepartment,
		)
	}

	return departments, nil
}

func parseStaff() (int, error) {
	var staff int
	_, err := fmt.Scan(&staff)

	if err != nil {
		return 0, fmt.Errorf("Invalid input staff")
	}

	if staff < minStaff || staff > maxStaff {
		return 0, fmt.Errorf(
			"Staff must be from %d to %d",
			minStaff,
			maxStaff,
		)
	}

	return staff, nil
}

func parseTemperature() (int, error) {
	var temperature int
	_, err := fmt.Scan(&temperature)

	if err != nil {
		return 0, fmt.Errorf("Invalid input temperature value")
	}

	if temperature < minTemperature || temperature > maxTemperature {
		return 0, fmt.Errorf(
			"Temperature must be from %d to %d",
			minTemperature,
			maxTemperature,
		)
	}

	return temperature, nil
}

func main() {
	departments, err := parseDepartment()
	if err != nil {
		fmt.Println(err)
		return
	}

	for range departments {
		staff, err := parseStaff()
		if err != nil {
			fmt.Println(err)
			return
		}

		maxTempOfDepartment := maxTemperature
		minTempOfDepartment := minTemperature

		for range staff {
			var condition string
			_, err = fmt.Scan(&condition)

			if err != nil {
				fmt.Println("Invalid input condition")
				return
			}

			if condition != ">=" && condition != "<=" {
				fmt.Println("Condition of temperature must be '<=' or '>='")
				return
			}

			temperature, err := parseTemperature()
			if err != nil {
				fmt.Println(err)
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

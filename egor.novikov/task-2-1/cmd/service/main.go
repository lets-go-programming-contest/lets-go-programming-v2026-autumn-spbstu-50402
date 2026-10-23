package main

import (
	"errors"
	"fmt"
)

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
		return 0, errors.New("Invalid input departments")
	}

	if departments < minDepartment || departments > maxDepartment {
		return 0, errors.New("Depatments must be form 1 to 1000")
	}

	return departments, nil
}

func parseStaff() (int, error) {
	var staff int
	_, err := fmt.Scan(&staff)

	if err != nil {
		return 0, errors.New("Invalid input staff")
	}

	if staff < minStaff || staff > maxStaff {
		return 0, errors.New("Staff must be from 1 to 1000")
	}

	return staff, nil
}

func parseTemperature() (int, error) {
	var temperature int
	_, err := fmt.Scan(&temperature)

	if err != nil {
		return 0, errors.New("Invalid input temperature value")
	}

	if temperature < minTemperature || temperature > maxTemperature {
		return 0, errors.New("Temperature must be from 15 to 30")
	}

	return temperature, nil
}

func parseCondition() (string, error) {
	var condition string
	_, err := fmt.Scan(&condition)

	if err != nil {
		return "", errors.New("Invalid input condition")
	}

	if condition != ">=" && condition != "<=" {
		return "", errors.New("Condition of temperature must be '<=' or '>='")
	}

	return condition, nil
}

func processStaff(minTemp *int, maxTemp *int) error {
	condition, err := parseCondition()
	if err != nil {
		return err
	}

	temperature, err := parseTemperature()
	if err != nil {
		return err
	}

	if condition == ">=" && temperature > *minTemp {
		*minTemp = temperature
	} else if condition == "<=" && temperature < *maxTemp {
		*maxTemp = temperature
	}

	if *maxTemp < *minTemp {
		fmt.Println("-1")
	} else {
		fmt.Println(*minTemp)
	}

	return nil
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

		minTemp := minTemperature
		maxTemp := maxTemperature

		for range staff {
			err = processStaff(&minTemp, &maxTemp)
			if err != nil {
				fmt.Println(err)

				return
			}
		}
	}
}

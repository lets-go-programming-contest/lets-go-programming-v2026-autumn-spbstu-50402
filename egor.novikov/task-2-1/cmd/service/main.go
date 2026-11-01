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

var (
	errInvalidDepartments     = errors.New("invalid input departments")
	errInvalidDepartmentCount = errors.New("departments must be form 1 to 1000")
	errInvalidStaff           = errors.New("invalid input staff")
	errInvalidStaffCount      = errors.New("staff must be from 1 to 1000")
	errInvalidTemperature     = errors.New("invalid input temperature value")
	errTemperatureOutOfRange  = errors.New("temperature must be from 15 to 30")
	errInvalidConditionInput  = errors.New("invalid input condition")
	errInvalidCondition       = errors.New("condition of temperature must be '<=' or '>='")
)

func parseDepartment() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errInvalidDepartments, err)
	}

	if departments < minDepartment || departments > maxDepartment {
		return 0, errInvalidDepartmentCount
	}

	return departments, nil
}

func parseStaff() (int, error) {
	var staff int

	_, err := fmt.Scan(&staff)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errInvalidStaff, err)
	}

	if staff < minStaff || staff > maxStaff {
		return 0, errInvalidStaffCount
	}

	return staff, nil
}

func parseTemperature() (int, error) {
	var temperature int

	_, err := fmt.Scan(&temperature)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errInvalidTemperature, err)
	}

	if temperature < minTemperature || temperature > maxTemperature {
		return 0, errTemperatureOutOfRange
	}

	return temperature, nil
}

func parseCondition() (string, error) {
	var condition string

	_, err := fmt.Scan(&condition)
	if err != nil {
		return "", fmt.Errorf("%w: %w", errInvalidConditionInput, err)
	}

	if condition != ">=" && condition != "<=" {
		return "", errInvalidCondition
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

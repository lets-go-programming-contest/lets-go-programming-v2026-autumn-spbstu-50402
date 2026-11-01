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

type DepartmentTemp struct {
	minTemp int
	maxTemp int
}

func NewDepartment() *DepartmentTemp {
	return &DepartmentTemp{
		minTemp: minTemperature,
		maxTemp: maxTemperature,
	}
}

func (d *DepartmentTemp) UpdateBounds(condition string, temperature int) error {
	switch condition {
	case ">=":
		if temperature > d.minTemp {
			d.minTemp = temperature
		}
	case "<=":
		if temperature < d.maxTemp {
			d.maxTemp = temperature
		}
	default:
		return errInvalidCondition
	}

	return nil
}

func (d *DepartmentTemp) GetResult() int {
	if d.maxTemp < d.minTemp {
		return -1
	}

	return d.minTemp
}

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

func processStaff(department *DepartmentTemp) error {
	condition, err := parseCondition()
	if err != nil {
		return err
	}

	temperature, err := parseTemperature()
	if err != nil {
		return err
	}

	err = department.UpdateBounds(condition, temperature)
	if err != nil {
		return err
	}

	fmt.Println(department.GetResult())

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

		department := NewDepartment()

		for range staff {
			err = processStaff(department)
			if err != nil {
				fmt.Println(err)

				return
			}
		}
	}
}

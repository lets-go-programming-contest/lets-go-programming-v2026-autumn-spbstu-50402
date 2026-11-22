package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30

	minDepartments = 1
	maxDepartments = 1000

	minWorkers = 1
	maxWorkers = 1000
)

var (
	errDepartmentsOutOfRange = errors.New("departments count is out of range")
	errWorkersOutOfRange     = errors.New("workers count is out of range")
	errTemperatureOutOfRange = errors.New("temperature is outside the allowed range")
	errUnknownOperation      = errors.New("unknown operation")
	errNoValidTemperature    = errors.New("no valid temperature in bounds")
)

type bounds struct {
	left, right int
}

func newBounds(left, right int) bounds {
	return bounds{
		left:  left,
		right: right,
	}
}

func (b *bounds) change(sign string, temp int) error {
	switch sign {
	case "<=":
		b.right = min(b.right, temp)
	case ">=":
		b.left = max(b.left, temp)
	default:
		return errUnknownOperation
	}

	return nil
}

func (b *bounds) valid() bool {
	return b.left <= b.right
}

func (b *bounds) optimal() (int, error) {
	if !b.valid() {
		return 0, errNoValidTemperature
	}

	return b.left, nil
}

func scanDepartments() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("error reading departments count: %w", err)
	}

	if departments < minDepartments || departments > maxDepartments {
		return 0, errDepartmentsOutOfRange
	}

	return departments, nil
}

func scanWorkers() (int, error) {
	var workers int

	_, err := fmt.Scan(&workers)
	if err != nil {
		return 0, fmt.Errorf("error reading workers count: %w", err)
	}

	if workers < minWorkers || workers > maxWorkers {
		return 0, errWorkersOutOfRange
	}

	return workers, nil
}

func scanRequest() (string, int, error) {
	var sign string

	_, err := fmt.Scan(&sign)
	if err != nil {
		return "", 0, fmt.Errorf("error reading employee's operation: %w", err)
	}

	if sign != "<=" && sign != ">=" {
		return "", 0, errUnknownOperation
	}

	var temp int

	_, err = fmt.Scan(&temp)
	if err != nil {
		return "", 0, fmt.Errorf("error reading employee's temperature: %w", err)
	}

	if temp < minTemp || temp > maxTemp {
		return "", 0, errTemperatureOutOfRange
	}

	return sign, temp, nil
}

func main() {
	departments, err := scanDepartments()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	for range departments {
		workers, err := scanWorkers()
		if err != nil {
			fmt.Println("Error:", err)

			return
		}

		temperatureBounds := newBounds(minTemp, maxTemp)

		for range workers {
			sign, temp, err := scanRequest()
			if err != nil {
				fmt.Println("Error:", err)

				return
			}

			err = temperatureBounds.change(sign, temp)
			if err != nil {
				fmt.Println("Error:", err)

				return
			}

			optimal, err := temperatureBounds.optimal()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimal)
		}
	}
}

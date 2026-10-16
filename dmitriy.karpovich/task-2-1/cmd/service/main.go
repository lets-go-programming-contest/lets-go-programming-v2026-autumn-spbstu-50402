package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

var (
	errDepartmentsOutOfRange = errors.New("departments count is out of range")
	errWorkersOutOfRange     = errors.New("workers count is out of range")
	errTemperatureOutOfRange = errors.New("temperature is outside the allowed range")
	errUnknownOperation      = errors.New("unknown operation")
)

type bounds struct {
	left, right int
}

func (b *bounds) change(sign string, temp int) {
	switch sign {
	case "<=":
		b.right = min(b.right, temp)
	case ">=":
		b.left = max(b.left, temp)
	}
}

func (b *bounds) check() bool {
	return b.left <= b.right
}

func scanDepartments() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("error reading departments count: %w", err)
	}

	if departments < 1 || departments > 1000 {
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

	if workers < 1 || workers > 1000 {
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

		temperatureBounds := bounds{
			left:  minTemp,
			right: maxTemp,
		}

		for range workers {
			sign, temp, err := scanRequest()
			if err != nil {
				fmt.Println("Error:", err)

				return
			}

			temperatureBounds.change(sign, temp)

			if temperatureBounds.check() {
				fmt.Println(temperatureBounds.left)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidDepartmentCount = errors.New("invalid number of departments provided")
	ErrInvalidWorkerCount     = errors.New("invalid number of workers provided")
	ErrInvalidModifier        = errors.New("invalid format of operation modifier")
	ErrInvalidTemperature     = errors.New("invalid format of temperature value")
	ErrTemperatureBounds      = errors.New("temperature boundaries violation")
)
func main() {
const (
	minTemp = 15
	maxTemp = 30
)

type TemperatureBoundary struct {
	lowerBound int
	upperBound int
}

func NewTemperatureBoundary(lower, upper int) (*TemperatureBoundary, error) {
	if lower > upper {
		return nil, fmt.Errorf("%w: lower bound %d cannot exceed upper bound %d", ErrTemperatureBounds, lower, upper)
	}
	return &TemperatureBoundary{
		lowerBound: lower,
		upperBound: upper,
	}, nil
}

func (cb *TemperatureBoundary) Adjust(modifier string, temp int) error {
	switch modifier {
	case ">=":
		if temp > cb.lowerBound {
			cb.lowerBound = temp
		}
	case "<=":
		if temp < cb.upperBound {
			cb.upperBound = temp
		}
	default:
		return fmt.Errorf("%w: operator '%s' is not supported", ErrInvalidModifier, modifier)
	}
	return nil
}

func (cb *TemperatureBoundary) IsValid() bool {
	return cb.lowerBound <= cb.upperBound
}

func (cb *TemperatureBoundary) OptimalTemperature() int {
	if !cb.IsValid() {
		return -1
	}
	return cb.lowerBound
}
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		var workers int
		if _, err := fmt.Scan(&workers); err != nil {
			fmt.Println("Invalid number of workers")

			return
		}

		minTemp := 15
		maxTemp := 30

		for range workers {
			var modifier string
			if _, err := fmt.Scan(&modifier); err != nil {
				fmt.Println("Invalid format of operation")

				return
			}

			var temperature int
			if _, err := fmt.Scan(&temperature); err != nil {
				fmt.Println("Invalid format of temperature")

				return
			}

			switch modifier {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
func ErrorHandler(err error) {
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, ErrInvalidDepartmentCount):
		fmt.Println("Invalid number of departments")
	case errors.Is(err, ErrInvalidWorkerCount):
		fmt.Println("Invalid number of workers")
	case errors.Is(err, ErrInvalidModifier):
		fmt.Println("Invalid format of operation")
	case errors.Is(err, ErrInvalidTemperature):
		fmt.Println("Invalid format of temperature")
	case errors.Is(err, ErrTemperatureBounds):
		fmt.Println("Temperature bounds error occurred")
	default:
		fmt.Printf("An unexpected system error occurred: %v\n", err)
	}
}

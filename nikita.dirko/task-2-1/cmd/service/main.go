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

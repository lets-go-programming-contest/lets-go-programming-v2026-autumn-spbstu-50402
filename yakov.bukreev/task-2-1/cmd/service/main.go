package main

import "fmt"

const (
	MinTemp = 15
	MaxTemp = 30
)

type TemperatureBounds struct {
	low, high int
}

func (bounds *TemperatureBounds) applyBound(sign string, temperature int) {
	switch sign {
	case ">=":
		bounds.low = max(bounds.low, temperature)
	case "<=":
		bounds.high = min(bounds.high, temperature)
	}
}

func (bounds *TemperatureBounds) getOptimalTemperature() int {
	if bounds.low <= bounds.high {
		return bounds.low
	} else {
		return -1
	}
}

func processDepartment() error {
	var workers int

	if _, err := fmt.Scan(&workers); err != nil {
		return fmt.Errorf("invalid workers number")
	}

	if workers < 1 || workers > 1000 {
		return fmt.Errorf("workers number out of range")
	}

	bounds := TemperatureBounds{
		low:  MinTemp,
		high: MaxTemp,
	}

	for idx := range workers {
		var (
			sign        string
			temperature int
		)

		if _, err := fmt.Scan(&sign, &temperature); err != nil {
			return fmt.Errorf("invalid preference for worker %d", idx+1)
		}

		if sign != ">=" && sign != "<=" {
			return fmt.Errorf("invalid preference sign for worker %d", idx+1)
		}

		if temperature < MinTemp || temperature > MaxTemp {
			return fmt.Errorf("invalid temperature for worker %d", idx+1)
		}

		bounds.applyBound(sign, temperature)
		fmt.Println(bounds.getOptimalTemperature())
	}

	return nil
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Printf("Invalid deparments number: %v.", err)
		return
	}

	if departments < 1 || departments > 1000 {
		fmt.Println("Deparments number out of range.")
		return
	}

	for i := range departments {
		if err := processDepartment(); err != nil {
			fmt.Printf("Error (department %d): %v.", i+1, err)
		}
	}
}

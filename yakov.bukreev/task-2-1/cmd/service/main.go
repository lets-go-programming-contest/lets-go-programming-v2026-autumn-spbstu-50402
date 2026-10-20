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

func processDepartment() {
	var workers int

	if _, err := fmt.Scan(&workers); err != nil {
		fmt.Println("Invalid workers number.")
		return
	}

	if workers < 1 || workers > 1000 {
		fmt.Println("Workers number out of range")
		return
	}

	bounds := TemperatureBounds{
		low:  MinTemp,
		high: MaxTemp,
	}

	for range workers {
		var (
			sign        string
			temperature int
		)

		if _, err := fmt.Scan(&sign, &temperature); err != nil {
			fmt.Println("Invalid preference.")
			return
		}

		if sign != ">=" && sign != "<=" {
			fmt.Println("Invalid preference sign.")
			return
		}

		if temperature < MinTemp || temperature > MaxTemp {
			fmt.Println("Invalid temperature.")
			return
		}

		bounds.applyBound(sign, temperature)
		fmt.Println(bounds.getOptimalTemperature())
	}
}

func main() {
	var deparments int

	if _, err := fmt.Scan(&deparments); err != nil {
		fmt.Println("Invalid deparments number.")
		return
	}

	if deparments < 1 || deparments > 1000 {
		fmt.Println("Deparments number out of range.")
		return
	}

	for range deparments {
		processDepartment()
	}
}

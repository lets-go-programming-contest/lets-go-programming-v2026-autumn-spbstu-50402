package main

import "fmt"

const (
	MinTemp = 15
	MaxTemp = 30
)

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

	var tempLow, tempHigh int = MinTemp, MaxTemp

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

		switch sign {
		case ">=":
			tempLow = max(tempLow, temperature)
		case "<=":
			tempHigh = min(tempHigh, temperature)
		}

		if tempLow <= tempHigh {
			fmt.Println(tempLow)
		} else {
			fmt.Println(-1)
		}
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

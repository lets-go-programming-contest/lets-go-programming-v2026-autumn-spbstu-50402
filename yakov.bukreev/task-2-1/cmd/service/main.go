package main

import "fmt"

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

	var tempMin, tempMax int = 15, 30

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

		if temperature < 15 || temperature > 30 {
			fmt.Println("Invalid temperature.")
			return
		}

		switch sign {
		case ">=":
			tempMin = max(tempMin, temperature)
		case "<=":
			tempMax = min(tempMax, temperature)
		}

		if tempMin <= tempMax {
			fmt.Println(tempMin)
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

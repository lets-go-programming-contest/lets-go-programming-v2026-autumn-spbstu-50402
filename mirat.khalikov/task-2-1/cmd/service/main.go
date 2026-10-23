package main

import "fmt"

const (
	MIN_D = 1
	MAX_D = 1000
	MIN_W = 1
	MAX_W = 1000
	MIN_T = 15
	MAX_T = 30
)

func parseTemperature(temperature int, condition string, minTemp int, maxTemp int) (int, int) {
	if condition == ">=" {
		if temperature > minTemp {
			minTemp = temperature
		}
	} else if condition == "<=" {
		if temperature < maxTemp {
			maxTemp = temperature
		}
	}
	return minTemp, maxTemp
}

func main() {
	var (
		departments int
		err         error
	)

	_, err = fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if departments < MIN_D || departments > MAX_D {
		fmt.Println("Input should be between 1 and 1000")
		return
	}

	for range departments {
		var workers int
		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		if workers < MIN_W || workers > MAX_W {
			fmt.Println("Input should be between 1 and 1000")
			return
		}
		minTemp := MIN_T
		maxTemp := MAX_T
		for range workers {
			var (
				condition   string
				temperature int
			)
			_, err = fmt.Scan(&condition)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
			if condition != ">=" && condition != "<=" {
				fmt.Println("Input condition should be non-strict")
				return
			}
			_, err = fmt.Scan(&temperature)
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
			if temperature < MIN_T || temperature > MAX_T {
				fmt.Println("Input temperature should be between 15 and 30")
				return
			}
			minTemp, maxTemp = parseTemperature(temperature, condition, minTemp, maxTemp)
			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

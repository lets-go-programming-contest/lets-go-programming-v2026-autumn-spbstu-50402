package main

import "fmt"

const (
	minD = 1
	maxD = 1000

	minW = 1
	maxW = 1000

	minT = 15
	maxT = 30
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

func readDepartments() (int, bool) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid input")

		return 0, false
	}

	if departments < minD || departments > maxD {
		fmt.Println("Input should be between 1 and 1000")

		return 0, false
	}

	return departments, true
}

func readWorkers() (int, bool) {
	var workers int

	_, err := fmt.Scan(&workers)
	if err != nil {
		fmt.Println("Invalid input")

		return 0, false
	}

	if workers < minW || workers > maxW {
		fmt.Println("Input should be between 1 and 1000")

		return 0, false
	}

	return workers, true
}

func readCondition() (string, bool) {
	var condition string

	_, err := fmt.Scan(&condition)
	if err != nil {
		fmt.Println("Invalid input")

		return "", false
	}

	if condition != ">=" && condition != "<=" {
		fmt.Println("Input condition should be non-strict")

		return "", false
	}

	return condition, true
}

func readTemperature() (int, bool) {
	var temperature int

	_, err := fmt.Scan(&temperature)
	if err != nil {
		fmt.Println("Invalid input")

		return 0, false
	}

	if temperature < minT || temperature > maxT {
		fmt.Println("Input temperature should be between 15 and 30")

		return 0, false
	}

	return temperature, true
}

func main() {
	departments, check := readDepartments()
	if !check {
		return
	}

	for range departments {
		workers, check := readWorkers()
		if !check {
			return
		}

		minTemp, maxTemp := minT, maxT

		for range workers {
			condition, check := readCondition()
			if !check {
				return
			}

			temperature, check := readTemperature()
			if !check {
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

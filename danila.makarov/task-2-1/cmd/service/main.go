package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type TempRange struct {
	low  int
	high int
}

func (r *TempRange) Update(sign string, temp int) {
	if sign == ">=" && temp > r.low {
		r.low = temp
	}

	if sign == "<=" && temp < r.high {
		r.high = temp
	}
}

func (r *TempRange) Best() int {
	if r.low > r.high {
		return -1
	}

	return r.low
}

func main() {
	var depCount int

	_, err := fmt.Scan(&depCount)
	if err != nil {
		fmt.Println("Invalid departments count")

		return
	}

	for range depCount {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid workers count")

			return
		}

		tempRange := TempRange{low: minTemp, high: maxTemp}

		for range workers {
			var (
				sign string
				temp int
			)

			_, err = fmt.Scan(&sign, &temp)
			if err != nil {
				fmt.Println("Invalid temperature")

				return
			}

			if sign != ">=" && sign != "<=" {
				fmt.Println("Invalid sign")

				return
			}

			tempRange.Update(sign, temp)

			fmt.Println(tempRange.Best())
		}
	}
}

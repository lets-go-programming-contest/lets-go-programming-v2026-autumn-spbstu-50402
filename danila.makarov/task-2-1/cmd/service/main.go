package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type Wish struct {
	sign string
	temp int
}

func findTemp(wishes []Wish) int {
	low := minTemp
	high := maxTemp

	for _, wish := range wishes {
		if wish.sign == ">=" && wish.temp > low {
			low = wish.temp
		}

		if wish.sign == "<=" && wish.temp < high {
			high = wish.temp
		}
	}

	if low > high {
		return -1
	}

	return low
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

		wishes := make([]Wish, 0, workers)

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

			wishes = append(wishes, Wish{sign: sign, temp: temp})

			fmt.Println(findTemp(wishes))
		}
	}
}

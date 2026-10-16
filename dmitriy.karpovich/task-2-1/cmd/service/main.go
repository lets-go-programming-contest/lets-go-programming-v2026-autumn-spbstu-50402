package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type bounds struct {
	left, right int
}

func (b *bounds) change(sign string, temp int) error {
	if temp < minTemp || temp > maxTemp {
		return fmt.Errorf(
			"temperature %d is outside the allowed range [%d, %d]",
			temp, minTemp, maxTemp,
		)
	}

	switch sign {
	case "<=":
		b.right = min(b.right, temp)
	case ">=":
		b.left = max(b.left, temp)
	default:
		return fmt.Errorf("unknown operation %q", sign)
	}

	return nil
}

func (b *bounds) check() bool {
	return b.left <= b.right
}

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

		low := 15
		high := 30

		for range workers {
			var operation string
			if _, err := fmt.Scan(&operation); err != nil {
				fmt.Println("Invalid format of operation")

				return
			}

			var temperature int
			if _, err := fmt.Scan(&temperature); err != nil {
				fmt.Println("Invalid format of temperature")

				return
			}

			switch operation {
			case ">=":
				if temperature > low {
					low = temperature
				}
			case "<=":
				if temperature < high {
					high = temperature
				}
			}

			if low <= high {
				fmt.Println(low)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

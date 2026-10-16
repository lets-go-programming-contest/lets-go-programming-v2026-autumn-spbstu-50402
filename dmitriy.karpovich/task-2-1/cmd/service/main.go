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
			"Temperature %d is outside the allowed range [%d, %d]",
			temp, minTemp, maxTemp,
		)
	}

	switch sign {
	case "<=":
		b.right = min(b.right, temp)
	case ">=":
		b.left = max(b.left, temp)
	default:

		return fmt.Errorf("Unknown operation %q", sign)
	}

	return nil
}

func (b *bounds) check() bool {
	return b.left <= b.right
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Error reading departments count:", err)

		return
	}

	if departments < 1 || departments > 1000 {
		fmt.Println("Error: departments count is out of range")

		return
	}

	for range departments {
		var workers int

		if _, err := fmt.Scan(&workers); err != nil {
			fmt.Println("Error reading workers count:", err)

			return
		}

		if workers < 1 || workers > 1000 {
			fmt.Println("Error: workers count is out of range")

			return
		}

		temperatureBounds := bounds{minTemp, maxTemp}

		for range workers {
			var sign string
			var temp int

			if _, err := fmt.Scan(&sign, &temp); err != nil {
				fmt.Println("Error reading employee's request:", err)

				return
			}

			if err := temperatureBounds.change(sign, temp); err != nil {
				fmt.Println("Error:", err)

				return
			}

			if temperatureBounds.check() {
				fmt.Println(temperatureBounds.left)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

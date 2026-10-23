package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type temperatureRange struct {
	lower int
	upper int
}

func (tr *temperatureRange) update(sign string, temp int) {
	switch sign {
	case ">=":
		tr.lower = max(tr.lower, temp)
	case "<=":
		tr.upper = min(tr.upper, temp)
	}
}

func (tr temperatureRange) optimal() int {
	if tr.lower > tr.upper {
		return -1
	}

	return tr.lower
}

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		var workers int

		_, err = fmt.Scan(&workers)
		if err != nil {
			fmt.Println("Invalid number of workers")

			return
		}

		tempRange := temperatureRange{lower: minTemp, upper: maxTemp}

		for range workers {
			var (
				sign string
				temp int
			)

			_, err = fmt.Scan(&sign, &temp)
			if err != nil {
				fmt.Println("Invalid input")

				return
			}

			if sign != ">=" && sign != "<=" {
				fmt.Println("Invalid operation")

				return
			}

			tempRange.update(sign, temp)
			fmt.Println(tempRange.optimal())
		}
	}
}

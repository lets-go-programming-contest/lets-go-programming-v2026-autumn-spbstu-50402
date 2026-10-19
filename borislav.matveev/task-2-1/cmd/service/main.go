package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		departments int
		employees   int
		temperature int
		operation   string
	)

	if _, err := fmt.Fscan(reader, &departments); err != nil {
		return
	}

	for department := 0; department < departments; department++ {
		if _, err := fmt.Fscan(reader, &employees); err != nil {
			return
		}

		minTemperature := 15
		maxTemperature := 30

		for employee := 0; employee < employees; employee++ {
			if _, err := fmt.Fscan(reader, &operation, &temperature); err != nil {
				return
			}

			switch operation {
			case ">=":
				if temperature > minTemperature {
					minTemperature = temperature
				}
			case "<=":
				if temperature < maxTemperature {
					maxTemperature = temperature
				}
			default:
				return
			}

			if minTemperature > maxTemperature {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, minTemperature)
			}
		}
	}
}

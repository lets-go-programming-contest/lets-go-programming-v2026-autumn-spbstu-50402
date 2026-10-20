package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer func() {
		if err := writer.Flush(); err != nil {
			log.Print(err)
		}
	}()

	var (
		departments int
		employees   int
		temperature int
		operation   string
	)

	if _, err := fmt.Fscan(reader, &departments); err != nil {
		return
	}

	for range departments {
		if _, err := fmt.Fscan(reader, &employees); err != nil {
			return
		}

		for range employees {
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

			result := minTemperature
			if minTemperature > maxTemperature {
				result = -1
			}

			if _, err := fmt.Fprintln(writer, result); err != nil {
				return
			}
		}
	}
}

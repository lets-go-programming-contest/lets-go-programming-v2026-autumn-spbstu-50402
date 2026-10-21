package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

var errInvalidOperation = errors.New("invalid operation")

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
	)

	if _, err := fmt.Fscan(reader, &departments); err != nil {
		return
	}

	for range departments {
		if _, err := fmt.Fscan(reader, &employees); err != nil {
			return
		}

		if err := processDepartment(reader, writer, employees); err != nil {
			return
		}
	}
}

func processDepartment(reader *bufio.Reader, writer *bufio.Writer, employees int) error {
	lowerBound := minTemperature
	upperBound := maxTemperature

	var (
		temperature int
		operation   string
	)

	for range employees {
		if _, err := fmt.Fscan(reader, &operation, &temperature); err != nil {
			return err
		}

		switch operation {
		case ">=":
			if temperature > lowerBound {
				lowerBound = temperature
			}
		case "<=":
			if temperature < upperBound {
				upperBound = temperature
			}
		default:
			return errInvalidOperation
		}

		result := lowerBound
		if lowerBound > upperBound {
			result = -1
		}

		if _, err := fmt.Fprintln(writer, result); err != nil {
			return err
		}
	}

	return nil
}

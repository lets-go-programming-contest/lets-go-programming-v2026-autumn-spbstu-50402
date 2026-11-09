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

type temperatureState struct {
	lowerBound int
	upperBound int
}

func newTemperatureState() temperatureState {
	return temperatureState{
		lowerBound: minTemperature,
		upperBound: maxTemperature,
	}
}

func (state *temperatureState) applyRequirement(operation string, temperature int) error {
	switch operation {
	case ">=":
		if temperature > state.lowerBound {
			state.lowerBound = temperature
		}
	case "<=":
		if temperature < state.upperBound {
			state.upperBound = temperature
		}
	default:
		return errInvalidOperation
	}

	return nil
}

func (state *temperatureState) result() int {
	if state.lowerBound > state.upperBound {
		return -1
	}

	return state.lowerBound
}

func main() {
	// Buffer input to reduce reads from stdin.
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
	state := newTemperatureState()

	var (
		temperature int
		operation   string
	)

	for range employees {
		if _, err := fmt.Fscan(reader, &operation, &temperature); err != nil {
			return fmt.Errorf("read employee requirement %w", err)
		}

		if err := state.applyRequirement(operation, temperature); err != nil {
			return err
		}

		if _, err := fmt.Fprintln(writer, state.result()); err != nil {
			return fmt.Errorf("write department temperature %w", err)
		}
	}

	return nil
}

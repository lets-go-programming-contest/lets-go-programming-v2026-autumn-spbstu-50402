package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	minTemp = 15
	maxTemp = 30
	noTemp  = -1

	minCount       = 1
	maxDepartments = 1000
	maxEmployees   = 1000
)

var (
	errUnknownOperator = errors.New("unknown operator")
	errOutOfRange      = errors.New("value out of range")
)

type operator int

const (
	opGreaterEq operator = iota
	opLessEq
)

func newOperator(raw string) (operator, error) {
	switch raw {
	case ">=":
		return opGreaterEq, nil
	case "<=":
		return opLessEq, nil
	default:
		return 0, fmt.Errorf("%w: %q", errUnknownOperator, raw)
	}
}

type tempRange struct {
	low  int
	high int
}

func newTempRange() *tempRange {
	return &tempRange{low: minTemp, high: maxTemp}
}

func (r *tempRange) apply(oper operator, temp int) {
	switch oper {
	case opGreaterEq:
		r.low = max(r.low, temp)
	case opLessEq:
		r.high = min(r.high, temp)
	}
}

func (r *tempRange) optimal() int {
	if r.low > r.high {
		return noTemp
	}

	return r.low
}

func readCount(reader io.Reader, name string, limit int) (int, error) {
	var value int

	if _, err := fmt.Fscan(reader, &value); err != nil {
		return 0, fmt.Errorf("read %s: %w", name, err)
	}

	if value < minCount || value > limit {
		return 0, fmt.Errorf("%w: %s = %d, want %d..%d", errOutOfRange, name, value, minCount, limit)
	}

	return value, nil
}

func readConstraint(reader io.Reader) (operator, int, error) {
	var (
		raw  string
		temp int
	)

	if _, err := fmt.Fscan(reader, &raw, &temp); err != nil {
		return 0, 0, fmt.Errorf("read constraint: %w", err)
	}

	func newTempRange(low, high int) tempRange {
		return tempRange{low: low, high: high}
	}

	if err != nil {
		return 0, 0, err
	}

	return oper, temp, nil
}

func solveDepartment(reader io.Reader, writer io.Writer) error {
	employees, err := readCount(reader, "K", maxEmployees)
	if err != nil {
		return err
	}

	func (r *tempRange) apply(oper operator, temp int) error {
		switch oper {
		case opGreaterEq:
			r.low = max(r.low, temp)
		case opLessEq:
			r.high = min(r.high, temp)
		default:
			return fmt.Errorf("%w: %d", errUnknownOperator, oper)
		}

		return nil
	}

	for employee := range employees {
		oper, temp, err := readConstraint(reader)
		if err != nil {
			return fmt.Errorf("employee %d: %w", employee+1, err)
		}

		if err := bounds.apply(oper, temp); err != nil {
			return fmt.Errorf("employee %d: %w", employee+1, err)
		}

		if _, err := fmt.Fprintln(writer, bounds.optimal()); err != nil {
			return fmt.Errorf("write answer: %w", err)
		}
	}

	return nil
}

func solveDepartments(reader io.Reader, writer io.Writer) error {
	departments, err := readCount(reader, "N", maxDepartments)
	if err != nil {
		return err
	}

	for department := range departments {
		if err := solveDepartment(reader, writer); err != nil {
			return fmt.Errorf("department %d: %w", department+1, err)
		}
	}

	return nil
}

func solve(input io.Reader, output io.Writer) error {
	writer := bufio.NewWriter(output)
	solveErr := solveDepartments(bufio.NewReader(input), writer)

	if err := writer.Flush(); err != nil {
		return errors.Join(solveErr, fmt.Errorf("flush output: %w", err))
	}

	return solveErr
}

func main() {
	if err := solve(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

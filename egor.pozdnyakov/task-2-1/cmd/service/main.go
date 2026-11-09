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
	errNoOptimalTemp   = errors.New("no optimal temperature")
)

type countError struct {
	name  string
	value int
	lower int
	upper int
}

func (e *countError) Error() string {
	return fmt.Sprintf("invalid %s: expected %d..%d, actual %d", e.name, e.lower, e.upper, e.value)
}

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

func newTempRange(low, high int) tempRange {
	return tempRange{low: low, high: high}
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

func (r *tempRange) optimal() (int, error) {
	if r.low > r.high {
		return 0, errNoOptimalTemp
	}

	return r.low, nil
}

func readCount(reader io.Reader, name string, limit int) (int, error) {
	var value int

	if _, err := fmt.Fscan(reader, &value); err != nil {
		return 0, fmt.Errorf("read %s: %w", name, err)
	}

	if value < minCount || value > limit {
		return 0, &countError{name: name, value: value, lower: minCount, upper: limit}
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

	oper, err := newOperator(raw)
	if err != nil {
		return 0, 0, err
	}

	return oper, temp, nil
}

func solveDepartment(reader io.Reader, writer io.Writer) error {
	employees, err := readCount(reader, "employees count", maxEmployees)
	if err != nil {
		return err
	}

	bounds := newTempRange(minTemp, maxTemp)

	for range employees {
		oper, temp, err := readConstraint(reader)
		if err != nil {
			return err
		}

		if err := bounds.apply(oper, temp); err != nil {
			return err
		}

		answer, err := bounds.optimal()
		if errors.Is(err, errNoOptimalTemp) {
			answer = noTemp
		}

		if _, err := fmt.Fprintln(writer, answer); err != nil {
			return fmt.Errorf("write answer: %w", err)
		}
	}

	return nil
}

func solveDepartments(reader io.Reader, writer io.Writer) error {
	departments, err := readCount(reader, "departments count", maxDepartments)
	if err != nil {
		return err
	}

	for range departments {
		if err := solveDepartment(reader, writer); err != nil {
			return err
		}
	}

	return nil
}

func solve(input io.Reader, output io.Writer) (err error) {
	writer := bufio.NewWriter(output)

	defer func() {
		if flushErr := writer.Flush(); flushErr != nil {
			err = errors.Join(err, fmt.Errorf("flush output: %w", flushErr))
		}
	}()

	return solveDepartments(bufio.NewReader(input), writer)
}

func main() {
	if err := solve(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

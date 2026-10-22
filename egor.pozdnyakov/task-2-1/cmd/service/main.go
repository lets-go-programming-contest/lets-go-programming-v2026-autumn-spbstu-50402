package main

import (
	"fmt"
)

func main() {}

const (
	minTemp = 15
	maxTemp = 30
	noTemp  = -1
)

const (
	minCount       = 1
	maxDepartments = 1000
	maxEmployees   = 1000
)

type operator int

const (
	opGreaterEq operator = iota
	opLessEq
)

func (op operator) String() string {
	switch op {
	case opGreaterEq:
		return ">="
	case opLessEq:
		return "<="
	default:
		return "?"
	}
}

func parseOperator(raw string) (operator, error) {
	switch raw {
	case ">=":
		return opGreaterEq, nil
	case "<=":
		return opLessEq, nil
	default:
		return 0, fmt.Errorf("%w %q", errUnknownOperator, raw)
	}
}

package main

import (
	"errors"
	"fmt"
)

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

var (
	errUnknownOperator = errors.New("unknown operator")
	errOutOfRange      = errors.New("value out of range")
)

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

func main() {
	if err := solve(os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

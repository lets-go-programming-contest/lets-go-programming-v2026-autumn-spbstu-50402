package main

import (
	"errors"
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

var (
	errUnknownOperation     = errors.New("unknown operation")
	errNoOptimalTemperature = errors.New("no optimal temperature")
)

type request struct {
	sign string
	temp int
}

type temperatureRange struct {
	lower int
	upper int
}

func newTemperatureRange(lower, upper int) temperatureRange {
	return temperatureRange{
		lower: lower,
		upper: upper,
	}
}

func (tr *temperatureRange) update(req request) error {
	switch req.sign {
	case ">=":
		tr.lower = max(tr.lower, req.temp)
	case "<=":
		tr.upper = min(tr.upper, req.temp)
	default:
		return fmt.Errorf("%w: %q", errUnknownOperation, req.sign)
	}

	return nil
}

func (tr *temperatureRange) optimal() (int, error) {
	if tr.lower > tr.upper {
		return 0, errNoOptimalTemperature
	}

	return tr.lower, nil
}

func readDepartments() (int, error) {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		return 0, fmt.Errorf("read departments count: %w", err)
	}

	return departments, nil
}

func readWorkers() (int, error) {
	var workers int

	_, err := fmt.Scan(&workers)
	if err != nil {
		return 0, fmt.Errorf("read workers count: %w", err)
	}

	return workers, nil
}

func readRequest() (request, error) {
	var req request

	_, err := fmt.Scan(&req.sign)
	if err != nil {
		return req, fmt.Errorf("read operation: %w", err)
	}

	_, err = fmt.Scan(&req.temp)
	if err != nil {
		return req, fmt.Errorf("read temperature: %w", err)
	}

	return req, nil
}

func printOptimalTemperatures(workers int) error {
	tempRange := newTemperatureRange(minTemp, maxTemp)

	for range workers {
		req, err := readRequest()
		if err != nil {
			return err
		}

		err = tempRange.update(req)
		if err != nil {
			return err
		}

		optimal, err := tempRange.optimal()

		switch {
		case err == nil:
			fmt.Println(optimal)
		case errors.Is(err, errNoOptimalTemperature):
			fmt.Println(-1)
		default:
			return err
		}
	}
	return nil
}

func main() {
	departments, err := readDepartments()
	if err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		workers, err := readWorkers()
		if err != nil {
			fmt.Println(err)

			return
		}

		err = printOptimalTemperatures(workers)
		if err != nil {
			fmt.Println(err)

			return
		}
	}
}

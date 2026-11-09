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

func newRequest(sign string, temp int) request {
	return request{
		sign: sign,
		temp: temp,
	}
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

func readCount() (int, error) {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("scan count: %w", err)
	}

	return count, nil
}

func readRequest() (request, error) {
	var (
		sign string
		temp int
	)

	_, err := fmt.Scan(&sign)
	if err != nil {
		return request{}, fmt.Errorf("read operation: %w", err)
	}

	_, err = fmt.Scan(&temp)
	if err != nil {
		return request{}, fmt.Errorf("read temperature: %w", err)
	}

	return newRequest(sign, temp), nil
}

func processDepartment(workers int) error {
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

func run() error {
	departments, err := readCount()
	if err != nil {
		return fmt.Errorf("read departments: %w", err)
	}

	for range departments {
		workers, err := readCount()
		if err != nil {
			return fmt.Errorf("read workers: %w", err)
		}

		err = processDepartment(workers)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

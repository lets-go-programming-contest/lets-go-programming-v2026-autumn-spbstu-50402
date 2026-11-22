package sttodec

import (
	"fmt"
	"strconv"
	"strings"
)

func StToDec(value string) (float64, error) {
	value = strings.ReplaceAll(value, ",", ".")
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("parse failed: %w", err)
	}
	return result, nil
}

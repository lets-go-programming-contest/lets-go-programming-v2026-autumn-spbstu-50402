package decimal

import (
	"fmt"
	"strconv"
	"strings"
)

type Value float64

func (value *Value) UnmarshalText(text []byte) error {
	normalized := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse decimal value %q: %w", text, err)
	}

	*value = Value(parsed)

	return nil
}

package currency

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrInvalidNumCode = errors.New("invalid NumCode")
	ErrInvalidValue   = errors.New("invalid currency Value")
)

type NumCode int

func (code *NumCode) UnmarshalText(text []byte) error {
	value := strings.TrimSpace(string(text))
	if value == "" {
		*code = 0

		return nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%w: %q", ErrInvalidNumCode, value)
	}

	*code = NumCode(parsed)

	return nil
}

type RateValue float64

func (value *RateValue) UnmarshalText(text []byte) error {
	original := strings.TrimSpace(string(text))
	if original == "" {
		return fmt.Errorf("%w: value is empty", ErrInvalidValue)
	}

	normalized := strings.ReplaceAll(original, ",", ".")

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil || math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return fmt.Errorf("%w: %q", ErrInvalidValue, original)
	}

	*value = RateValue(parsed)

	return nil
}

type Rate struct {
	NumCode  NumCode   `json:"num_code"  xml:"NumCode"`
	CharCode string    `json:"char_code" xml:"CharCode"`
	Value    RateValue `json:"value"     xml:"Value"`
}

func SortByValueDescending(rates []Rate) {
	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})
}

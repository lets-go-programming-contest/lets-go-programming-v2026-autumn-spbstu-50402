package currency

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	errEmptyCurrencyValue = errors.New("currency Value is empty")
	ErrInvalidNumCode     = errors.New("invalid NumCode")
)

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func NewRate(numCodeText, charCode, valueText string) (Rate, error) {
	var numCode int

	numCodeText = strings.TrimSpace(numCodeText)

	if numCodeText != "" {
		parsedNumCode, err := strconv.Atoi(numCodeText)
		if err != nil {
			return Rate{}, fmt.Errorf(
				"%w: %q: %w",
				ErrInvalidNumCode,
				numCodeText,
				err,
			)
		}

		numCode = parsedNumCode
	}

	valueText = strings.ReplaceAll(strings.TrimSpace(valueText), ",", ".")

	if valueText == "" {
		return Rate{}, errEmptyCurrencyValue
	}

	parsedValue, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return Rate{}, fmt.Errorf("invalid Value %q: %w", valueText, err)
	}

	return Rate{
		NumCode:  numCode,
		CharCode: strings.TrimSpace(charCode),
		Value:    parsedValue,
	}, nil
}

func SortByValueDescending(rates []Rate) {
	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})
}

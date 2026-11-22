package currency

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/xRonnyn/lets-go-programming-v2026-autumn-spbstu-50402/task-3/cbr"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Convert(valutes []cbr.Valute) ([]Currency, error) {
	currencies := make([]Currency, 0, len(valutes))

	for _, item := range valutes {
		text := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value of %s: %w", item.CharCode, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

func Sort(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

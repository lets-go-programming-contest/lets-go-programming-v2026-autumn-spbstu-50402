package currency

import "sort"

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func SortByValueDescending(rates []Rate) {
	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})
}

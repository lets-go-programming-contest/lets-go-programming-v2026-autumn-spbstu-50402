package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CurrencyValue float64

func (v *CurrencyValue) UnmarshalText(text []byte) error {
	normalized := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse currency value %q: %w", string(text), err)
	}

	*v = CurrencyValue(parsed)

	return nil
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    CurrencyValue `json:"value"     xml:"Value"`
}

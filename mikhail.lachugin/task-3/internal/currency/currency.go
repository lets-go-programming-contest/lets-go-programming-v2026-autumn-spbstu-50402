// Package currency contains currency data and sorting helpers.
package currency

import (
	"cmp"
	"encoding/xml"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const floatBitSize = 64

// ValCurs is the root element of the Bank of Russia currency XML document.
type ValCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

// Currency contains the fields used in the JSON result.
type Currency struct {
	NumCode  int     `json:"num_code" xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    Decimal `json:"value" xml:"Value"`
}

// Decimal accepts the comma decimal separator used in the Bank of Russia XML.
type Decimal float64

// UnmarshalXML parses a decimal value from an XML element.
func (d *Decimal) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode decimal element %q: %w", start.Name.Local, err)
	}

	value, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(raw), ",", ".", 1), floatBitSize)
	if err != nil {
		return fmt.Errorf("parse decimal value %q: %w", raw, err)
	}
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return fmt.Errorf("decimal value must be finite: %q", raw)
	}

	*d = Decimal(value)
	return nil
}

// SortByValueDesc orders currencies by value from highest to lowest.
func SortByValueDesc(currencies []Currency) {
	slices.SortStableFunc(currencies, func(a, b Currency) int {
		return cmp.Compare(float64(b.Value), float64(a.Value))
	})
}

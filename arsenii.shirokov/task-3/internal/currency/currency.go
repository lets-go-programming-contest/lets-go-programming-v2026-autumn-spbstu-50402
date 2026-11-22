package currency

import (
	"cmp"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const (
	floatBitSize     = 64
	replaceCommaOnce = 1
)

var ErrInvalidDecimal = errors.New("decimal value must be a finite number")

type ValCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    Decimal `json:"value"     xml:"Value"`
}

type Decimal float64

func (d *Decimal) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := dec.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode decimal element %q: %w", start.Name.Local, err)
	}

	normalized := strings.Replace(strings.TrimSpace(raw), ",", ".", replaceCommaOnce)

	val, err := strconv.ParseFloat(normalized, floatBitSize)
	if err != nil {
		return fmt.Errorf("parse decimal value %q: %w", raw, err)
	}

	if math.IsNaN(val) || math.IsInf(val, 0) {
		return fmt.Errorf("%w: %q", ErrInvalidDecimal, raw)
	}

	*d = Decimal(val)

	return nil
}

func SortByValueDesc(currencies []Currency) {
	slices.SortStableFunc(currencies, func(a, b Currency) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

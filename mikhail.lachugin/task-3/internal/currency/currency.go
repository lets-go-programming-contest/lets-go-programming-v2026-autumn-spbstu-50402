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

const floatBitSize = 64

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
		return fmt.Errorf("%w: %q", ErrInvalidDecimal, raw)
	}

	*d = Decimal(value)

	return nil
}

func SortByValueDesc(currencies []Currency) {
	slices.SortStableFunc(currencies, func(a, b Currency) int {
		return cmp.Compare(float64(b.Value), float64(a.Value))
	})
}

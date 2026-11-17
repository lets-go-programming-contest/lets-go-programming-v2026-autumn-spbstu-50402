package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CommaFloat float64

func (c *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	s = strings.Replace(s, ",", ".", 1)

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("parse float: %w", err)
	}

	*c = CommaFloat(val)

	return nil
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

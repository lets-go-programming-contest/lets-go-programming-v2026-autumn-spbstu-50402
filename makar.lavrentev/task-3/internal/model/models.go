package model

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type CommaFloat float64

func (c *CommaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	s = strings.Replace(s, ",", ".", 1)
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*c = CommaFloat(val)
	return nil
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `xml:"NumCode" json:"num_code"`
	CharCode string     `xml:"CharCode" json:"char_code"`
	Value    CommaFloat `xml:"Value" json:"value"`
}

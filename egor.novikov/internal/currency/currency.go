package currency

import "encoding/xml"

type XMLCurrency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type XMLDocument struct {
	XMLName  xml.Name      `xml:"ValCurs"`
	Currencies []XMLCurrency `xml:"Valute"`
}
package xml_parser

import (
	"encoding/xml"
    "fmt"
    "strings"
)

type Valute struct {
    NumCode int `xml:"NumCode"`
    CharCode string `xml:"CharCode"`
    Value string `xml:"Value"`
}

type Curs struct {
	Valutes []Valute `xml:"Valute"`
}

func Parse(data []byte) (Curs, error) {
    res := Curs{}
    err := xml.Unmarshal(data, &res)
    if err != nil {
        return Curs{}, err
    }
    return res, nil
}

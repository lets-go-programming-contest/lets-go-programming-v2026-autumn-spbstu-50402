package xml_parser

import (
	"encoding/xml"
    "fmt"
    "os"
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

func Parse(path string) (Curs, error) {
	data, err := os.ReadFile(path)
    if err != nil {
        return Curs{}, fmt.Errorf("failed to read file %q: %w", path, err)
    }
    res := Curs{}
    err := xml.Unmarshal(data, &res)
    if err != nil {
        return Curs{}, fmt.Errorf("failed to unmarshal %q: %w", err)
    }
    return res, nil
}

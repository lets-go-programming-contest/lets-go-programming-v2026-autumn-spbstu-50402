package cbr

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ReadXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml from the path %q: %w", path, err)
	}

	var doc ValCurs
	err = xml.Unmarshal(data, &doc)
	if err != nil {
		return nil, fmt.Errorf("unmarshal from the path %q: %w", path, err)
	}

	if len(doc.Valutes) == 0 {
		return nil, fmt.Errorf("the file from the path %q has not elements", path)
	}

	result := make([]Currency, 0, len(doc.Valutes))
	for _, v := range doc.Valutes {
		num, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
		if err != nil {
			return nil, fmt.Errorf("bad NumCode %q: %w", v.NumCode, err)
		}

		val, err := strconv.ParseFloat(strings.ReplaceAll(v.Value, ",", "."), 64)
		if err != nil {
			return nil, fmt.Errorf("bad value %q: %w", v.Value, err)
		}

		result = append(result, Currency{
			NumCode:  num,
			CharCode: strings.TrimSpace(v.CharCode),
			Value:    val,
		})
	}

	return result, nil
}

package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

var ErrNoElements = errors.New("file has no elements")

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

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode xml from the path %q: %w", path, err)
	}

	if len(doc.Valutes) == 0 {
		return nil, fmt.Errorf("the file from the path %q has not elements: %w", path, ErrNoElements)
	}

	result := make([]Currency, 0, len(doc.Valutes))

	for _, valute := range doc.Valutes {
		numCodeStr := strings.TrimSpace(valute.NumCode)

		var num int
		if numCodeStr == "" {
			num = 0
		} else {
			num, err = strconv.Atoi(numCodeStr)
			if err != nil {
				return nil, fmt.Errorf("bad NumCode %q: %w", valute.NumCode, err)
			}
		}

		val, err := strconv.ParseFloat(strings.ReplaceAll(valute.Value, ",", "."), 64)
		if err != nil {
			return nil, fmt.Errorf("bad value %q: %w", valute.Value, err)
		}

		result = append(result, Currency{
			NumCode:  num,
			CharCode: strings.TrimSpace(valute.CharCode),
			Value:    val,
		})
	}

	return result, nil
}

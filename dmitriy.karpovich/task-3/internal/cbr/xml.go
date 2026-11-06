package cbr

import (
	"encoding/xml"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/voyadger01/task-3/internal/currency"
)

type xmlDocument struct {
	XMLName xml.Name
	Rates   []xmlRate `xml:"Valute"`
}

type xmlRate struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func ReadFile(path string) ([]currency.Rate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", path, err)
	}

	rates, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("process XML file %q: %w", path, err)
	}
	return rates, nil
}

func Parse(data []byte) ([]currency.Rate, error) {
	var document xmlDocument
	if err := xml.Unmarshal(data, &document); err != nil {
		return nil, fmt.Errorf("decode XML document: %w", err)
	}
	if document.XMLName.Local != "ValCurs" {
		return nil, fmt.Errorf("invalid XML root element: expected ValCurs, got %q", document.XMLName.Local)
	}
	if len(document.Rates) == 0 {
		return nil, fmt.Errorf("XML document does not contain Valute elements")
	}

	rates := make([]currency.Rate, 0, len(document.Rates))
	for index, item := range document.Rates {
		rate, err := parseRate(item, index+1)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func parseRate(item xmlRate, position int) (currency.Rate, error) {
	numCode, err := strconv.Atoi(strings.TrimSpace(item.NumCode))
	if err != nil {
		return currency.Rate{}, fmt.Errorf("currency #%d: invalid NumCode %q: %w", position, item.NumCode, err)
	}

	charCode := strings.TrimSpace(item.CharCode)
	if charCode == "" {
		return currency.Rate{}, fmt.Errorf("currency #%d: CharCode is missing", position)
	}

	valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return currency.Rate{}, fmt.Errorf("currency %s: invalid Value %q: %w", charCode, item.Value, err)
	}
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return currency.Rate{}, fmt.Errorf("currency %s: Value must be a finite number", charCode)
	}

	return currency.Rate{
		NumCode:  numCode,
		CharCode: charCode,
		Value:    value,
	}, nil
}

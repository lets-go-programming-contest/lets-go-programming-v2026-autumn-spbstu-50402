package report

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type CurrencyInfo struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    float32 `xml:"Value"`
}

type reportData struct {
	Currecies []CurrencyInfo `xml:"Valute"`
}

func ParseReport(path string) ([]CurrencyInfo, error) {
	report, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read report file %q: %w", path, err)
	}

	report = bytes.ReplaceAll(report, []byte(","), []byte("."))

	var data reportData

	decoder := xml.NewDecoder(bytes.NewReader(report))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal report file %q: %w", path, err)
	}

	return data.Currecies, nil
}

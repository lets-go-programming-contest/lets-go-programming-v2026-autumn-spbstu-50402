package report

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type CurrencyInfo struct {
	NumCode  int     `xml:"NumCode"   json:"num_code"`
	CharCode string  `xml:"CharCode"  json:"char_code"`
	Value    float32 `xml:"Value"     json:"value"`
}

type reportData struct {
	Currecies []CurrencyInfo `xml:"Valute"`
}

func ReadXml(path string) ([]CurrencyInfo, error) {
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

func WriteJson(path string, mode os.FileMode, data []CurrencyInfo) error {
	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("Failed to marshal json data for file %q: %w", path, err)
	}

	if err := os.WriteFile(path, jsonBytes, mode); err != nil {
		return fmt.Errorf("Failed to write json file %q: %w", path, err)
	}

	return nil
}

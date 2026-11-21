package report

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/net/html/charset"
)

type CurrencyInfo struct {
	NumCode  int     `json:"num_code"   xml:"NumCode"`
	CharCode string  `json:"char_code"  xml:"CharCode"`
	Value    float32 `json:"value"      xml:"Value"`
}

type reportData struct {
	Currecies []CurrencyInfo `xml:"Valute"`
}

func ReadXML(path string) ([]CurrencyInfo, error) {
	report, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read report file %q: %w", path, err)
	}

	report = bytes.ReplaceAll(report, []byte(","), []byte("."))

	var data reportData

	decoder := xml.NewDecoder(bytes.NewReader(report))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal report file %q: %w", path, err)
	}

	return data.Currecies, nil
}

func WriteJSON(path string, fileMode, dirMode os.FileMode, data []CurrencyInfo) error {
	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal json data for file %q: %w", path, err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirMode); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", dir, err)
	}

	if err := os.WriteFile(path, jsonBytes, fileMode); err != nil {
		return fmt.Errorf("failed to write json file %q: %w", path, err)
	}

	return nil
}

package cbr

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/voyadger01/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

type xmlDocument struct {
	Rates []currency.Rate `xml:"Valute"`
}

func ParseFile(path string) ([]currency.Rate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", path, err)
	}

	rates, err := parse(data)
	if err != nil {
		return nil, fmt.Errorf("parse XML file %q: %w", path, err)
	}

	return rates, nil
}

func parse(data []byte) ([]currency.Rate, error) {
	var document xmlDocument

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err := decoder.Decode(&document)
	if err != nil {
		return nil, fmt.Errorf("decode XML document: %w", err)
	}

	return document.Rates, nil
}

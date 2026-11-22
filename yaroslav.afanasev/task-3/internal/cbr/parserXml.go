package cbr

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Jaroslaw-super1/task-3/internal/currency"

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

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var document xmlDocument
	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("decode XML document: %w", err)
	}

	return document.Rates, nil
}

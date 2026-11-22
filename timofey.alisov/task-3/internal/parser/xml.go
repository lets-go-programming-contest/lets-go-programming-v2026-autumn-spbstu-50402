package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/timfoi/task-3/internal/models"
	"golang.org/x/net/html/charset"
)

func ParseXML(path string) ([]models.Valute, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml file: %w", err)
	}

	var valCurs models.ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if valCurs.Valutes == nil {
		return []models.Valute{}, nil
	}

	return valCurs.Valutes, nil
}

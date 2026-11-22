// Package xmlparser reads Bank of Russia currency XML documents.
package xmlparser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mikeee322/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

// ErrEmptyInputPath means the XML source path is blank.
var ErrEmptyInputPath = errors.New("XML input path is empty")

// ParseFile opens and parses a currency XML file.
func ParseFile(path string) ([]currency.Currency, error) {
	if strings.TrimSpace(path) == "" {
		return nil, ErrEmptyInputPath
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML file %q: %w", path, err)
	}

	currencies, parseErr := Parse(file)

	closeErr := file.Close()

	if parseErr != nil {
		return nil, fmt.Errorf("parse XML file %q: %w", path, parseErr)
	}

	if closeErr != nil {
		return nil, fmt.Errorf("close XML file %q: %w", path, closeErr)
	}

	return currencies, nil
}

// Parse decodes a currency XML document from reader.
func Parse(reader io.Reader) ([]currency.Currency, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	var document currency.ValCurs
	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	if document.Valutes == nil {
		return make([]currency.Currency, 0), nil
	}

	return document.Valutes, nil
}

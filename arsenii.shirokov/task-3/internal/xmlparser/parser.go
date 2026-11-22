package xmlparser

import (
	"converter/internal/currency"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html/charset"
)

var ErrEmptyInputPath = errors.New("xml input path is empty")

func ParseFile(path string) ([]currency.Currency, error) {
	if strings.TrimSpace(path) == "" {
		return nil, ErrEmptyInputPath
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open xml file %q: %w", path, err)
	}

	currencies, parseErr := Parse(file)
	closeErr := file.Close()

	if parseErr != nil {
		return nil, fmt.Errorf("parse xml file %q: %w", path, parseErr)
	}

	if closeErr != nil {
		return nil, fmt.Errorf("close xml file %q: %w", path, closeErr)
	}

	return currencies, nil
}

func Parse(reader io.Reader) ([]currency.Currency, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs currency.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if valCurs.Valutes == nil {
		return make([]currency.Currency, 0), nil
	}

	return valCurs.Valutes, nil
}

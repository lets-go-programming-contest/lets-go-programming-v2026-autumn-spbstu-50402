package cbr

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Alastor20/task-3/internal/decimal"
	"github.com/Alastor20/task-3/internal/must"
	"golang.org/x/text/encoding/charmap"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valCurs struct {
	Currencies []Currency `xml:"Valute"`
}

func ParseFile(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file %q: %w", path, err)
	}

	defer func() {
		must.NoError(file.Close())
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var document valCurs

	err = decoder.Decode(&document)
	if err != nil {
		return nil, fmt.Errorf("decode xml document: %w", err)
	}

	if document.Currencies == nil {
		document.Currencies = []Currency{}
	}

	return document.Currencies, nil
}

func (currency *Currency) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var payload struct {
		NumCode  string        `xml:"NumCode"`
		CharCode string        `xml:"CharCode"`
		Value    decimal.Value `xml:"Value"`
	}

	if err := decoder.DecodeElement(&payload, &start); err != nil {
		return fmt.Errorf("decode valute element: %w", err)
	}

	numCode, err := parseNumCode(payload.NumCode)
	if err != nil {
		return err
	}

	currency.NumCode = numCode
	currency.CharCode = strings.TrimSpace(payload.CharCode)
	currency.Value = float64(payload.Value)

	return nil
}

func parseNumCode(raw string) (int, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 0, nil
	}

	numCode, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, fmt.Errorf("parse num code %q: %w", raw, err)
	}

	return numCode, nil
}

package cbr

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/voyadger01/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

var (
	errInvalidXMLRoot       = errors.New("invalid XML root element")
	errEmptyXMLDocument     = errors.New("XML document does not contain Valute elements")
	errNonFiniteCurrencyVal = errors.New("value must be a finite number")
)

type xmlDocument struct {
	XMLName xml.Name  `xml:"ValCurs"`
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

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err := decoder.Decode(&document)
	if err != nil {
		return nil, fmt.Errorf("decode XML document: %w", err)
	}

	if document.XMLName.Local != "ValCurs" {
		return nil, fmt.Errorf("%w: got %q", errInvalidXMLRoot, document.XMLName.Local)
	}

	if len(document.Rates) == 0 {
		return nil, errEmptyXMLDocument
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
	numCode := 0
	numCodeText := strings.TrimSpace(item.NumCode)

	if numCodeText != "" {
		parsedNumCode, parseErr := strconv.Atoi(numCodeText)
		if parseErr != nil {
			return currency.Rate{}, fmt.Errorf(
				"currency #%d: invalid NumCode %q: %w",
				position,
				item.NumCode,
				parseErr,
			)
		}

		numCode = parsedNumCode
	}

	charCode := strings.TrimSpace(item.CharCode)

	value := 0.0
	valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

	if valueText != "" {
		parsedValue, parseErr := strconv.ParseFloat(valueText, 64)
		if parseErr != nil {
			return currency.Rate{}, fmt.Errorf(
				"currency #%d: invalid Value %q: %w",
				position,
				item.Value,
				parseErr,
			)
		}

		if math.IsNaN(parsedValue) || math.IsInf(parsedValue, 0) {
			return currency.Rate{}, fmt.Errorf(
				"currency #%d: %w",
				position,
				errNonFiniteCurrencyVal,
			)
		}

		value = parsedValue
	}

	return currency.Rate{
		NumCode:  numCode,
		CharCode: charCode,
		Value:    value,
	}, nil
}

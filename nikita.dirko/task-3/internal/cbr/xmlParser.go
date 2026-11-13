package cbr

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Values  []valute `xml:"Valute"`
}

type valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func ParseFile(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file %q: %w", path, err)
	}
	defer func() { _ = file.Close() }()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var document valCurs
	err = decoder.Decode(&document)

	if err != nil {
		return nil, fmt.Errorf("decode xml document: %w", err)
	}

	currencies := make([]Currency, 0, len(document.Values))

	for _, rawValute := range document.Values {
		currency, err := rawValute.convert()
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (v valute) convert() (Currency, error) {
	numCode, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
	if err != nil {
		return Currency{}, fmt.Errorf("parse code %q: %w", v.NumCode, err)
	}

	normalized := strings.ReplaceAll(strings.TrimSpace(v.Value), ",", ".")
	value, err := strconv.ParseFloat(normalized, 64)

	if err != nil {
		return Currency{}, fmt.Errorf("parse currency value %q: %w", v.Value, err)
	}

	return Currency{
		NumCode:  numCode,
		CharCode: strings.TrimSpace(v.CharCode),
		Value:    value,
	}, nil
}

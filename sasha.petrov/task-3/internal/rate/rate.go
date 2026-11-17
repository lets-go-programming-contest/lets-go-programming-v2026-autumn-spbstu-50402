package rate

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type Quote struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func (quote *Quote) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode valute: %w", err)
	}

	quote.CharCode = strings.TrimSpace(raw.CharCode)

	if code := strings.TrimSpace(raw.NumCode); code != "" {
		numCode, err := strconv.Atoi(code)
		if err != nil {
			return fmt.Errorf("num code %q: %w", raw.NumCode, err)
		}

		quote.NumCode = numCode
	}

	value, err := strconv.ParseFloat(strings.Replace(raw.Value, ",", ".", 1), 64)
	if err != nil {
		return fmt.Errorf("value %q: %w", raw.Value, err)
	}

	quote.Value = value

	return nil
}

func Read(path string) ([]Quote, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml %q: %w", path, err)
	}

	var curs struct {
		XMLName xml.Name `xml:"ValCurs"`
		Items   []Quote  `xml:"Valute"`
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("parse xml %q: %w", path, err)
	}

	return curs.Items, nil
}

func charsetReader(_ string, input io.Reader) (io.Reader, error) {
	return charmap.Windows1251.NewDecoder().Reader(input), nil
}

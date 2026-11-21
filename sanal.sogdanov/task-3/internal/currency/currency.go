package currency

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var (
	errUnsupportedEncoding = errors.New("unsupported XML encoding")
	errTrailingXML         = errors.New("unexpected data after </ValCurs>")
	errMissingCharCode     = errors.New("missing CharCode")
	errInvalidValue        = errors.New("invalid Value")
)

type currencyXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ratesXML struct {
	XMLName    xml.Name      `xml:"ValCurs"`
	Currencies []currencyXML `xml:"Valute"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Read(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = readCharset

	var rates ratesXML
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	if err := checkXMLTail(decoder); err != nil {
		return nil, err
	}

	currencies := make([]Currency, 0, len(rates.Currencies))
	for index, item := range rates.Currencies {
		rate, err := parseCurrency(item)
		if err != nil {
			return nil, fmt.Errorf("currency %d: %w", index+1, err)
		}

		currencies = append(currencies, rate)
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies, nil
}

func readCharset(charset string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(charset, "windows-1251") {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("%w: %q", errUnsupportedEncoding, charset)
}

func checkXMLTail(decoder *xml.Decoder) error {
	for {
		token, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("decode XML: %w", err)
		}

		switch token := token.(type) {
		case xml.Comment, xml.ProcInst:
			continue
		case xml.CharData:
			if len(bytes.TrimSpace(token)) == 0 {
				continue
			}
		}

		return errTrailingXML
	}
}

func parseCurrency(item currencyXML) (Currency, error) {
	numCodeText := strings.TrimSpace(item.NumCode)
	if numCodeText == "" {
		numCodeText = "0"
	}

	numCode, err := strconv.Atoi(numCodeText)
	if err != nil {
		return Currency{}, fmt.Errorf("invalid NumCode %q: %w", item.NumCode, err)
	}

	charCode := strings.TrimSpace(item.CharCode)
	if charCode == "" {
		return Currency{}, errMissingCharCode
	}

	valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return Currency{}, fmt.Errorf("%w %q: %w", errInvalidValue, item.Value, err)
	}

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return Currency{}, fmt.Errorf("%w %q", errInvalidValue, item.Value)
	}

	return Currency{
		NumCode:  numCode,
		CharCode: charCode,
		Value:    value,
	}, nil
}

package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
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
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported XML encoding %q", charset)
	}

	var rates ratesXML
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decode XML: %w", err)
		}

		switch token := token.(type) {
		case xml.Comment, xml.ProcInst:
			continue
		case xml.CharData:
			if len(bytes.TrimSpace(token)) == 0 {
				continue
			}
		}
		return nil, fmt.Errorf("unexpected data after </ValCurs>")
	}

	currencies := make([]Currency, 0, len(rates.Currencies))
	for i, item := range rates.Currencies {
		numCode, err := strconv.Atoi(strings.TrimSpace(item.NumCode))
		if err != nil {
			return nil, fmt.Errorf("currency %d: invalid NumCode %q", i+1, item.NumCode)
		}

		charCode := strings.TrimSpace(item.CharCode)
		if charCode == "" {
			return nil, fmt.Errorf("currency %d: missing CharCode", i+1)
		}

		valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")
		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil || math.IsNaN(value) || math.IsInf(value, 0) {
			return nil, fmt.Errorf("currency %s: invalid Value %q", charCode, item.Value)
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: charCode,
			Value:    value,
		})
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies, nil
}

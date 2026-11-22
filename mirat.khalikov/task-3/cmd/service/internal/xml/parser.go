package xml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/khalii21/task-3/cmd/service/internal/codetoint"
	"github.com/khalii21/task-3/cmd/service/internal/sttodec"
	"golang.org/x/text/encoding/charmap"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value"     xml:"Value"`
}

type xmlCurrency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlData struct {
	Currencies []xmlCurrency `xml:"Valute"`
}

func ParseFile(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file failed %q: %w", path, err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("close XML file failed:", err)
		}
	}()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var doc xmlData

	err = decoder.Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("decode XML file failed: %w", err)
	}

	return parseCurrencies(doc.Currencies)
}

func parseCurrencies(items []xmlCurrency) ([]Currency, error) {
	currencies := make([]Currency, 0, len(items))

	for _, item := range items {
		code, err := codetoint.CodeToInt(item.NumCode)
		if err != nil {
			return nil, fmt.Errorf("convert currency code failed: %w", err)
		}

		value, parseErr := sttodec.StToDec(item.Value)
		if parseErr != nil {
			return nil, fmt.Errorf("parse currency value failed: %w", parseErr)
		}

		currencies = append(currencies, Currency{
			NumCode:  code,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

package xml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"task-3/cmd/service/internal/codetoint"
	"task-3/cmd/service/internal/sttodec"

	"golang.org/x/text/encoding/charmap"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value" xml:"Value"`
}

type xmlData struct {
	Currencies []struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

func ParseFile(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file failed %q: %w", path, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
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

	currencies := make([]Currency, 0, len(doc.Currencies))
	for _, item := range doc.Currencies {

		code, err := codetoint.CodeToInt(item.NumCode)
		if err != nil {
			return nil, fmt.Errorf("convert currency code failed: %w", err)
		}

		value, err := sttodec.StToDec(item.Value)
		if err != nil {
			return nil, fmt.Errorf("parse currency value failed: %w", err)
		}

		currencies = append(currencies, Currency{
			NumCode:  code,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

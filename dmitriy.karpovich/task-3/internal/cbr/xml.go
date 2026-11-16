package cbr

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/voyadger01/task-3/internal/currency"
	"golang.org/x/net/html/charset"
)

type xmlDocument struct {
	XMLName xml.Name  `xml:"ValCurs"`
	Rates   []xmlRate `xml:"Valute"`
}

type xmlRate struct {
	rate currency.Rate
}

func (r *xmlRate) UnmarshalXML(
	decoder *xml.Decoder,
	start xml.StartElement,
) error {
	var item struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	err := decoder.DecodeElement(&item, &start)
	if err != nil {
		return fmt.Errorf("decode Valute element: %w", err)
	}

	rate, err := currency.NewRate(
		item.NumCode,
		item.CharCode,
		item.Value,
	)
	if err != nil {
		return fmt.Errorf("parse currency rate: %w", err)
	}

	r.rate = rate

	return nil
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseFile(path string) ([]currency.Rate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML file %q: %w", path, err)
	}

	defer func() {
		panicIfErr(file.Close())
	}()

	var document xmlDocument

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&document)
	if err != nil {
		return nil, fmt.Errorf("decode XML document: %w", err)
	}

	rates := make([]currency.Rate, 0, len(document.Rates))

	for _, item := range document.Rates {
		rates = append(rates, item.rate)
	}

	return rates, nil
}

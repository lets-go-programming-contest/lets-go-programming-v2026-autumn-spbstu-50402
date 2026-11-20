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

type Decimal float64

func (value *Decimal) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var text string

	if err := decoder.DecodeElement(&text, &start); err != nil {
		return fmt.Errorf("decode value: %w", err)
	}

	number, err := strconv.ParseFloat(strings.Replace(text, ",", ".", 1), 64)
	if err != nil {
		return fmt.Errorf("parse value %q: %w", text, err)
	}

	*value = Decimal(number)

	return nil
}

type Quote struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    Decimal `xml:"Value" json:"value"`
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

func SortByValue(list []Quote) {
	for i := 1; i < len(list); i++ {
		for j := i; j > 0 && list[j].Value > list[j-1].Value; j-- {
			list[j], list[j-1] = list[j-1], list[j]
		}
	}
}

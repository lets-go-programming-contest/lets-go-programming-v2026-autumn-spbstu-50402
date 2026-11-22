package parser

import (
	"encoding/xml"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/Kayv1e/task-3/internal/model"

	"golang.org/x/net/html/charset"
)

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []valute `xml:"Valute"`
}

type valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func ParseXML(path string) ([]model.Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var data valCurs
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Valutes) == 0 {
		return nil, errors.New("no valutes data in input file")
	}

	currencies := make([]model.Currency, 0, len(data.Valutes))
	for _, v := range data.Valutes {
		numCode, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
		if err != nil {
			return nil, err
		}

		charCode := strings.TrimSpace(v.CharCode)
		if charCode == "" {
			return nil, errors.New("empty CharCode field in input file")
		}

		valueStr := strings.ReplaceAll(strings.TrimSpace(v.Value), ",", ".")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, model.Currency{
			NumCode:  numCode,
			CharCode: charCode,
			Value:    value,
		})
	}

	return currencies, nil
}

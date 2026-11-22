package parser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"

	"github.com/Kayv1e/task-3/internal/model"
)

var (
	ErrNoValutes     = errors.New("no valutes data in input file")
	ErrEmptyCharCode = errors.New("empty CharCode field in input file")
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
		return nil, fmt.Errorf("open xml file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var data valCurs

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if len(data.Valutes) == 0 {
		return nil, ErrNoValutes
	}

	currencies := make([]model.Currency, 0, len(data.Valutes))
	for _, item := range data.Valutes {
		numCode, err := strconv.Atoi(strings.TrimSpace(item.NumCode))
		if err != nil {
			return nil, fmt.Errorf("parse NumCode: %w", err)
		}

		charCode := strings.TrimSpace(item.CharCode)
		if charCode == "" {
			return nil, ErrEmptyCharCode
		}

		valueStr := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("parse Value: %w", err)
		}

		currencies = append(currencies, model.Currency{
			NumCode:  numCode,
			CharCode: charCode,
			Value:    value,
		})
	}

	return currencies, nil
}

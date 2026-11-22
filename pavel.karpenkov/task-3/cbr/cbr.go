package cbr

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type valCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func Read(path string) ([]Valute, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs valCurs

	err = decoder.Decode(&curs)
	if err != nil {
		return nil, fmt.Errorf("decode input file: %w", err)
	}

	return curs.Valutes, nil
}

package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func winDecoder(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}

func panicIfErr(fn func() error) {
	if err := fn(); err != nil {
		panic(err)
	}
}

func DecodeXML[T any](inputFile string) (*T, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("open xml file: %w", err)
	}

	defer panicIfErr(file.Close)

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = winDecoder

	var result T

	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return &result, nil
}

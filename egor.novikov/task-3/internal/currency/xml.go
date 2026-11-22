package currency

import (
	"encoding/xml"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

func ReadXML(path string) []Currency {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer closeFile(file)

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var document XMLDocument

	err = decoder.Decode(&document)
	if err != nil {
		panic(err)
	}

	currencies := make([]Currency, 0, len(document.Currencies))

	for _, item := range document.Currencies {
		numCode := 0

		if item.NumCode != "" {
			numCode, err = strconv.Atoi(item.NumCode)
			if err != nil {
				panic(err)
			}
		}

		valueString := strings.Replace(item.Value, ",", ".", 1)

		value, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			panic(err)
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return currencies
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}

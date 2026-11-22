
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

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var document XMLDocument

	err = decoder.Decode(&document)
	if err != nil {
		panic(err)
	}

	currencies := make([]Currency, 0, len(document.Currencies))

	for _, item := range document.Currencies {
		valueString := strings.Replace(item.Value, ",", ".", 1)

		value, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			panic(err)
		}

		currencies = append(currencies, Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return currencies
}
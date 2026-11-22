package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/text/encoding/charmap"
)

type Valute struct {
    NumCode  int    `xml:"NumCode"`
    CharCode string `xml:"CharCode"`
    Value    string `xml:"Value"`
}

type Curs struct {
	Valutes []Valute `xml:"Valute"`
}

func Parse(path string) (Curs, error) {
	data, err := os.ReadFile(path)
    if err != nil {
        return Curs{}, fmt.Errorf("failed to read file %q: %w", path, err)
    }

    data, err = charmap.Windows1251.NewDecoder().Bytes(data)
	if err != nil {
		return Curs{}, fmt.Errorf("failed to decode windows-1251 %q: %w", path, err)
	}

	data = bytes.Replace(data, []byte(`encoding="windows-1251"`), []byte(`encoding="UTF-8"`), 1)

    var res Curs

    err = xml.Unmarshal(data, &res)
    if err != nil {
        return Curs{}, fmt.Errorf("failed to unmarshal %q: %w", path, err)
    }

    return res, nil
}

package parser

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/Pilipchenok/task-3/internal/model"
	"golang.org/x/text/encoding/charmap"
)

func winDecoder(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}

func DecodeXML(inputFile string) *model.ValCurs {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = winDecoder

	var curs model.ValCurs

	if err := decoder.Decode(&curs); err != nil {
		panic(err)
	}

	return &curs
}

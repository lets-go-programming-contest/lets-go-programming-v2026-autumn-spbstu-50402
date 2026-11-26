package xmlparser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html/charset"
)

var ErrEmptyInputPath = errors.New("xml input path is empty")

func ParseFile(path string, target any) (err error) {
	if strings.TrimSpace(path) == "" {
		return ErrEmptyInputPath
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open xml file %q: %w", path, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close xml file %q: %w", path, closeErr)
		}
	}()

	if parseErr := Parse(file, target); parseErr != nil {
		return fmt.Errorf("parse xml file %q: %w", path, parseErr)
	}

	return nil
}

func Parse(reader io.Reader, target any) error {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("decode xml: %w", err)
	}

	return nil
}

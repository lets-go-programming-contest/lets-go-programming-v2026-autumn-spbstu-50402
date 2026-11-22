package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

type Value float64

func (v *Value) UnmarshalText(text []byte) error {
	normalized := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse value %q: %w", text, err)
	}

	*v = Value(parsed)

	return nil
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Value  `json:"value"     xml:"Value"`
}

type rates struct {
	Currencies []Currency `xml:"Valute"`
}

func ReadXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var parsed rates

	if err := decoder.Decode(&parsed); err != nil {
		return nil, fmt.Errorf("decode input file: %w", err)
	}

	return parsed.Currencies, nil
}

func SortByValueDesc(currencies []Currency) {
	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func WriteJSON(path string, currencies []Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encode currencies: %w", err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}

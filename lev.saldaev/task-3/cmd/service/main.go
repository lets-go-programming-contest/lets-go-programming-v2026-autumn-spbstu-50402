package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

const filePermissions = 0o600

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type xmlRate struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlDocument struct {
	Rates []xmlRate `xml:"Valute"`
}

type rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to YAML configuration file")

	flag.Parse()

	cfg, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	rates, err := loadRates(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sortRates(rates)

	err = writeJSON(cfg.OutputFile, rates)
	if err != nil {
		panic(err)
	}
}

func loadConfig(path string) (config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("read config file %q: %w", path, err)
	}

	var cfg config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config{}, fmt.Errorf("decode config file %q: %w", path, err)
	}

	return cfg, nil
}

func loadRates(path string) ([]rate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var document xmlDocument

	err = decoder.Decode(&document)
	if err != nil {
		return nil, fmt.Errorf("decode input XML %q: %w", path, err)
	}

	rates := make([]rate, 0, len(document.Rates))

	for _, item := range document.Rates {
		value, err := parseValue(item.Value)
		if err != nil {
			return nil, fmt.Errorf("parse rate %q: %w", item.CharCode, err)
		}

		rates = append(rates, rate{
			NumCode:  item.NumCode,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    value,
		})
	}

	return rates, nil
}

func parseValue(raw string) (float64, error) {
	normalized := strings.ReplaceAll(strings.TrimSpace(raw), ",", ".")
	if normalized == "" {
		return 0, nil
	}

	value, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, fmt.Errorf("parse decimal value %q: %w", raw, err)
	}

	return value, nil
}

func sortRates(rates []rate) {
	sort.SliceStable(rates, func(left, right int) bool {
		return rates[left].Value > rates[right].Value
	})
}

func writeJSON(path string, rates []rate) error {
	data, err := json.MarshalIndent(rates, "", "  ")
	if err != nil {
		return fmt.Errorf("encode rates as JSON: %w", err)
	}

	err = os.WriteFile(path, data, filePermissions)
	if err != nil {
		return fmt.Errorf("write output file %q: %w", path, err)
	}

	return nil
}

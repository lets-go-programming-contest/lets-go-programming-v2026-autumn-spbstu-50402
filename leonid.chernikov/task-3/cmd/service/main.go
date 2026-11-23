package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func readXML(path string) []Valute {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var currencies ValCurs

	err = xml.Unmarshal(data, &currencies)
	if err != nil {
		panic(err)
	}

	return currencies.Valutes
}

func convertCurrencies(valutes []Valute) []Currency {
	currencies := make([]Currency, 0)

	for _, valute := range valutes {
		value := strings.ReplaceAll(valute.Value, ",", ".")

		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}

		currencies = append(currencies, Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    number,
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}

func saveJSON(path string, currencies []Currency) {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, data, 0o600)
	if err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	valutes := readXML(config.InputFile)

	currencies := convertCurrencies(valutes)

	saveJSON(config.OutputFile, currencies)
}

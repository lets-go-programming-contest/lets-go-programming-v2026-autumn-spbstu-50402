package main

import (
	"cmp"
	"encoding/json"
	"encoding/xml"
	"flag"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID        string `xml:"ID,attr"`
	NumCode   string `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

type JSONValute struct {
	NumCode  string  `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func winDecoder(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}

func toFloat(s string) float64 {
	s = strings.Replace(s, ",", ".", 1)
	val, _ := strconv.ParseFloat(s, 64)

	return val
}

func sortFunc(a, b JSONValute) int {
	return cmp.Compare(a.Value, b.Value)
}

func parseConfig() Config {
	configPath := flag.String("config", "", "YAML-config path")
	flag.Parse()

	if *configPath == "" {
		panic("invalid config flag")
	}

	yamlConf, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var cfg Config

	if err := yaml.Unmarshal(yamlConf, &cfg); err != nil {
		panic(err)
	}

	return cfg
}

func decodeXML(inputFile string) (*ValCurs, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = winDecoder

	var curs ValCurs
	err = decoder.Decode(&curs)

	return &curs, err
}

func simplifyData(curs *ValCurs) []JSONValute {
	valutes := make([]JSONValute, 0, len(curs.Valutes))

	for _, v := range curs.Valutes {
		val := toFloat(v.Value)
		valutes = append(valutes, JSONValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    val,
		})
	}

	return valutes
}

func writeJSON(valutes []JSONValute, outputFile string) {
	jsonData, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		panic(err)
	}

	if dir := filepath.Dir(outputFile); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}

	if err := os.WriteFile(outputFile, jsonData, 0600); err != nil {
		panic(err)
	}
}

func main() {
	cfg := parseConfig()

	curs, err := decodeXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	valutes := simplifyData(curs)

	slices.SortFunc(valutes, sortFunc)

	writeJSON(valutes, cfg.OutputFile)
}

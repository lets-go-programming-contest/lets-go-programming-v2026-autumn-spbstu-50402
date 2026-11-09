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

const (
	dirPerm  = 0o755
	filePerm = 0o600
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
	NumCode   int    `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   int    `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

type JSONValute struct {
	NumCode  int     `json:"num_code"`
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
	return cmp.Compare(b.Value, a.Value)
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

func decodeXML(inputFile string) *ValCurs {
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

	if err := decoder.Decode(&curs); err != nil {
		panic(err)
	}

	return &curs
}

func simplifyData(curs *ValCurs) []JSONValute {
	valutes := make([]JSONValute, 0, len(curs.Valutes))

	for _, v := range curs.Valutes {
		valutes = append(valutes, JSONValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    toFloat(v.Value),
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
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			panic(err)
		}
	}

	if err := os.WriteFile(outputFile, jsonData, filePerm); err != nil {
		panic(err)
	}
}

func main() {
	cfg := parseConfig()
	curs := decodeXML(cfg.InputFile)
	valutes := simplifyData(curs)

	slices.SortFunc(valutes, sortFunc)

	writeJSON(valutes, cfg.OutputFile)
}

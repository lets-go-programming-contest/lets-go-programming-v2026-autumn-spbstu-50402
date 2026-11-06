package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
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
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Value    float64 `json:"Value"`
}

func winDecoder(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}
	return input, nil
}

func toFloat(s string) (float64) {
	s = strings.Replace(s, ",", ".", 1)
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

func main() {
	configPath := flag.String("config", "", "YAML-config path")
	flag.Parse()
	if *configPath == "" {
		panic("Invalid config flag")
	}

	yamlConf, err := os.ReadFile(*configPath)
	if err != nil {
		panic("Config import error")
	}
	
	var cfg Config
	err = yaml.Unmarshal(yamlConf, &cfg)
	if err != nil {
		panic("YAML decode error")
	}

	data, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		panic("XML read error")
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = winDecoder

	var curs ValCurs
	err = decoder.Decode(&curs)
	if err != nil {
		fmt.Printf("XML parce error: %v", err)
	}

	var valutes []JSONValute
	for _, v := range(curs.Valutes) {
		val := toFloat(v.Value)
		valutes = append(valutes, JSONValute{
			NumCode: v.NumCode,
			CharCode: v.CharCode,
			Value: val,
		})
	}

	
}

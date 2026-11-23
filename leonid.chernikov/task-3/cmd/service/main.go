package main

import (
	"encoding/xml"
	"flag"
	"os"

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

	_ = valutes
}

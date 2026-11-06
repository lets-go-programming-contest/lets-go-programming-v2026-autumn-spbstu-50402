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
}

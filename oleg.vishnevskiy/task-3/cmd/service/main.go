package main

import (
    "flag"

    "github.com/voleg904/task-3/internal/config"
    "github.com/voleg904/task-3/internal/json_converter"
    "github.com/voleg904/task-3/internal/xml_parser"
)

func main() {
    cfgPath := flag.String("config", "config.yaml", "path to config")
    flag.Parse()

    path, err := config.Load(*cfgPath)
    if err != nil {
        panic(err)
    }
    curs, err := xml_parser.Parse(path.InputPath)
    if err != nil {
        panic(err)
    }
    err = json_converter.Save(curs, path.OutputPath)
    if err != nil {
        panic(err)
    }
}

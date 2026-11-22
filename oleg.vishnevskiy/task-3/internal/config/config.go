package config

import (
    "errors"
    "fmt"
    "os"
    "strings"

    "gopkg.in/yaml.v3"
)

type Paths struct {
    InputPath string `yaml:"input-file"`
    OutputPath string `yaml:"input-file"`
}

func Parse(data []byte) (Paths, error) {
    resPaths := Paths{}
    err := yaml.Unmarshal(data, &resPaths)
    if err != nil {
        return Paths{}, err
    }
    return resPaths, nil
}

func Check(p Paths) error {
    if p.InputPath == "" {
        return errors.New("Empty input path!")
    }
    if p.OutputPath == "" {
        return errors.New("Empty output path!")
    }
}

func Load(cfgPath string) (Paths, error) {
    data, err := os.ReadFile(cfgPath)
    if err != nil {
        return Paths{}, fmt.Errorf("failed to open file %q: %w", cfgPath, err)
    }

    var p Paths

    p, err = Parse(data)
    if err != nil {
        return Paths{}, fmt.Errorf("failed to parse file %q: %w", cfgPath, err)
    }

    p, err = Check(p)
    if err != nil {
        return Paths{}, err
    }

    return p, nil
}

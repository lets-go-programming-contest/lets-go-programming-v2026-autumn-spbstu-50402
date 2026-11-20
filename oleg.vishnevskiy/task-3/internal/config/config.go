package config

import (
    "errors"
    "fmt"
    "os"
    "strings"

    "gopkg.in/yaml.v3"
)

type Paths struct {
    InputPath string 'yaml:"input-file"'
    OutputPath string 'yaml:"input-file"'
}

func parse(data []byte) (Paths, error) {
    res_paths := Paths{}
    err := yaml.Unmarshal(data, &res_paths)
    if err != nil {
        return Paths{}, err
    }
    return res_paths, nil
}

func check(p Paths) error {
    if p.InputPath == "" {
        return errors.New("Empty input path!")
    }
    if p.OutputPath == "" {
        return errors.New("Empty output path!")
    }
}

func load(cfg_path string) (Paths, error) {
    data, err := os.ReadFile(cfg_path)
    if err != nil {
        return Paths{}, fmt.Errorf("failed to open file %q: %w", cfg_path, err)
    }

    var p Paths

    p, err = parse(data)
    if err != nil {
        return Paths{}, fmt.Errorf("failed to parse file %q: %w", cfg_path, err)
    }

    p, err = check(p)
    if err != nil {
        return Paths{}, err
    }

    return p, nil
}

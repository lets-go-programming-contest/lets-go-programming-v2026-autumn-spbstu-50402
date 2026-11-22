package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyInput  = errors.New("empty input path")
	ErrEmptyOutput = errors.New("empty output path")
)

type Paths struct {
	InputPath  string `yaml:"input-file"`
	OutputPath string `yaml:"output-file"`
}

func Parse(data []byte) (Paths, error) {
    var resPaths Paths

    err := yaml.Unmarshal(data, &resPaths)

    if err != nil {
        return Paths{}, fmt.Errorf("failed to unmarshal yaml: %w", err)
    }

    return resPaths, nil
}

func Check(p Paths) error {
    if p.InputPath == "" {
        return ErrEmptyInput
    }

    if p.OutputPath == "" {
        return ErrEmptyOutput
    }

    return nil
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

    err = Check(p)
    if err != nil {
        return Paths{}, err
    }

    return p, nil
}

package jsonconverter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/voleg904/task-3/internal/xml_parser"
)

const (
	dirPerm  = 0755
	filePerm = 0600
)

type ValuteJSON struct {
    NumCode int `json:"num_code"`
    CharCode string `json:"char_code"`
    Value float64 `json:"value"`
}

func Save(data xmlparser.Curs, path string) (error) {
    err := os.MkdirAll(filepath.Dir(path), dirPerm)
    if err != nil {
        return fmt.Errorf("failed to create directory %q: %w", path, err)
    }
    res := make([]ValuteJSON, 0, len(data.Valutes))

    for i := range data.Valutes {
        val, err := ChangeFormat(data.Valutes[i])

        if err != nil {
            return err
        }

        res = append(res, val)
    }

    sort.SliceStable(res, func(i, j int) bool {
		return res[i].Value > res[j].Value
	})

    toWrite, err := json.Marshal(res)
    if err != nil {
        return fmt.Errorf("failed to marshal json: %w", err)
    }
    err = os.WriteFile(path, toWrite, filePerm)
    if err != nil {
        return fmt.Errorf("failed to write json %q: %w", path, err)
    }

    return nil
}

func ChangeFormat(val xmlparser.Valute) (ValuteJSON, error) {
	valStr := strings.ReplaceAll(strings.TrimSpace(val.Value), ",", ".")
	value, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return ValuteJSON{}, fmt.Errorf("failed to convert val %q: %w", val.Value, err)
	}

    return ValuteJSON{NumCode: val.NumCode, CharCode: val.CharCode, Value: value}, nil
}

package json_converter

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

type ValuteJson struct {
    NumCode int `json:"num_code"`
    CharCode string `json:"char_code"`
    Value float64 `json:"value"`
}

func Save(data xml_parser.Curs, path string) (error) {
    err := os.MkdirAll(filepath.Dir(path), 0755)
    if err != nil {
        return fmt.Errorf("failed to create directory %q: %w", path, err)
    }
    var res []ValuteJson
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
    err = os.WriteFile(path, toWrite, 0644)
    if err != nil {
        return fmt.Errorf("failed to write json %q: %w", path, err)
    }
    return nil
}

func ChangeFormat(val xml_parser.Valute) (ValuteJson, error) {
	valStr := strings.ReplaceAll(strings.TrimSpace(val.Value), ",", ".")
	value, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return ValuteJson{}, fmt.Errorf("failed to convert val %q: %w", val.Value, err)
	}
    return ValuteJson{NumCode: val.NumCode, CharCode: val.CharCode, Value: value}, nil
}

package exporter

import (
	"encoding/json"
	"fmt"
	"os"
)

const filePerm = 0o600

func WriteJSON[T any](data T, outputFile string) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, filePerm); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}

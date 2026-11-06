package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/voyadger01/task-3/internal/currency"
)

func WriteJSON(path string, rates []currency.Rate) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return fmt.Errorf("output file path is empty")
	}

	data, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON data: %w", err)
	}
	data = append(data, '\n')

	if err := createParentDirectory(path); err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}
	return nil
}

func createParentDirectory(path string) error {
	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create output directory %q: %w", directory, err)
	}
	return nil
}

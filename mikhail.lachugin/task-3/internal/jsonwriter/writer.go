// Package jsonwriter writes currencies as JSON.
package jsonwriter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mikeee322/task-3/internal/currency"
)

const jsonIndent = "    "

// WriteFile creates the output directory when needed and writes currencies to path.
func WriteFile(path string, currencies []currency.Currency) error {
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("JSON output path is empty")
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, 0o750); err != nil {
		return fmt.Errorf("create output directory %q: %w", directory, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", path, err)
	}

	writeErr := Write(file, currencies)
	closeErr := file.Close()
	if writeErr != nil {
		return fmt.Errorf("write JSON to %q: %w", path, writeErr)
	}
	if closeErr != nil {
		return fmt.Errorf("close output file %q: %w", path, closeErr)
	}
	return nil
}

// Write encodes currencies as indented JSON to writer.
func Write(writer io.Writer, currencies []currency.Currency) error {
	if currencies == nil {
		currencies = make([]currency.Currency, 0)
	}

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", jsonIndent)
	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}
	return nil
}

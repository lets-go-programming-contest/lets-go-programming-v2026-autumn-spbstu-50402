package jsonwriter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"arsenii.shirokov/task-3/internal/currency"
)

const (
	dirPerm    = 0o750
	jsonIndent = "    "
)

var ErrEmptyOutputPath = errors.New("json output path is empty")

func WriteFile(path string, currencies []currency.Currency) error {
	if strings.TrimSpace(path) == "" {
		return ErrEmptyOutputPath
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("create output directory %q: %w", dir, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", path, err)
	}

	writeErr := Write(file, currencies)
	closeErr := file.Close()

	if writeErr != nil {
		return fmt.Errorf("write json to %q: %w", path, writeErr)
	}

	if closeErr != nil {
		return fmt.Errorf("close output file %q: %w", path, closeErr)
	}

	return nil
}

func Write(writer io.Writer, currencies []currency.Currency) error {
	if currencies == nil {
		currencies = make([]currency.Currency, 0)
	}

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", jsonIndent)

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

package jsonwriter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	dirPerm    = 0o750
	jsonIndent = "    "
)

var ErrEmptyOutputPath = errors.New("json output path is empty")

func WriteFile(path string, data any) (err error) {
	if strings.TrimSpace(path) == "" {
		return ErrEmptyOutputPath
	}

	dir := filepath.Dir(path)
	if mkdirErr := os.MkdirAll(dir, dirPerm); mkdirErr != nil {
		return fmt.Errorf("create output directory %q: %w", dir, mkdirErr)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", path, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close output file %q: %w", path, closeErr)
		}
	}()

	if writeErr := Write(file, data); writeErr != nil {
		return fmt.Errorf("write json to %q: %w", path, writeErr)
	}

	return nil
}

func Write(writer io.Writer, data any) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", jsonIndent)

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

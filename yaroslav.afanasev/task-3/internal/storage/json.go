package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var errEmptyOutputPath = errors.New("output file path is empty")

const (
	filePerms = 0o600
	dirPerms  = 0o750
)

func WriteJSON[T any](path string, data T) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return errEmptyOutputPath
	}

	encoded, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON data: %w", err)
	}

	if err := createParentDirectory(path); err != nil {
		return err
	}

	if err := os.WriteFile(path, encoded, filePerms); err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}

	return nil
}

func createParentDirectory(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerms); err != nil {
		return fmt.Errorf("create output directory %q: %w", dir, err)
	}

	return nil
}

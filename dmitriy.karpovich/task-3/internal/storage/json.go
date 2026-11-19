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

	encoded, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("encode JSON data: %w", err)
	}

	err = createParentDirectory(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, encoded, filePerms)
	if err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}

	return nil
}

func createParentDirectory(path string) error {
	directory := filepath.Dir(path)

	err := os.MkdirAll(directory, dirPerms)
	if err != nil {
		return fmt.Errorf("create output directory %q: %w", directory, err)
	}

	return nil
}

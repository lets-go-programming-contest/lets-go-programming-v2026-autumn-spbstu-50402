package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/voyadger01/task-3/internal/currency"
)

var errEmptyOutputPath = errors.New("output file path is empty")

const (
	filePerms = 0e600
	dirPerms  = 0e750
)

func WriteJSON(path string, rates []currency.Rate) error {
	path = strings.TrimSpace(path)
	if path == "" {
		return errEmptyOutputPath
	}

	data, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON data: %w", err)
	}

	data = append(data, '\n')

	err = createParentDirectory(path)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, filePerms)
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

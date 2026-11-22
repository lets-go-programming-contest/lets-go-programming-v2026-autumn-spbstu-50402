package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/timfoi/task-3/internal/models"
)

const (
	dirPerm  = 0o750
	filePerm = 0o600
)

func WriteJSON(path string, valutes []models.Valute) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}

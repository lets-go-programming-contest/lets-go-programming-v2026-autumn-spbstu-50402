package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Kayv1e/task-3/internal/model"
)

const outputDirPerm = 0o755

func SaveJSON(path string, currencies []model.Currency) error {
	outputDir := filepath.Dir(path)
	if err := os.MkdirAll(outputDir, outputDirPerm); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

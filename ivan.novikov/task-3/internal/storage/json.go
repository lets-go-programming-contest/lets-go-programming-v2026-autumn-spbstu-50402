package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Kayv1e/task-3/internal/model"
)

func SaveJSON(path string, currencies []model.Currency) error {
	outputDir := filepath.Dir(path)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(currencies)
}

package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Save(path string, currencies []Currency) error {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(currencies)
	if err != nil {
		return fmt.Errorf("encode output file: %w", err)
	}

	return nil
}

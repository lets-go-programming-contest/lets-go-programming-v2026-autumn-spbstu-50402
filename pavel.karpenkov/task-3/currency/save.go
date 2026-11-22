package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const dirPermissions = 0o755

func Save(path string, currencies []Currency) error {
	err := os.MkdirAll(filepath.Dir(path), dirPermissions)
	if err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(currencies)
	if err != nil {
		return fmt.Errorf("encode output file: %w", err)
	}

	return nil
}

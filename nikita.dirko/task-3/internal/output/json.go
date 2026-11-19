package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	filePermissions = 0o600
	dirPermissions  = 0o750
)

func WriteJSON(path string, data any) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	encoded, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("encode data to json: %w", err)
	}

	if err := os.WriteFile(path, encoded, filePermissions); err != nil {
		return fmt.Errorf("write output file %q: %w", path, err)
	}

	return nil
}

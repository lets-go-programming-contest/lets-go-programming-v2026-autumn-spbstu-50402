package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	filePerms = 0o644
	dirPerms  = 0o750
)

func WriteJSON(path string, data any) error {
	encoded, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("encode JSON data: %w", err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerms); err != nil {
		return fmt.Errorf("create output directory %q: %w", dir, err)
	}

	err = os.WriteFile(path, encoded, filePerms)
	if err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}

	return nil
}

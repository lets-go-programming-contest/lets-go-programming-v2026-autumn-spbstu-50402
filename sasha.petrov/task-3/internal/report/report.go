package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirMode  = 0o755
	fileMode = 0o600
)

func Write(path string, data any) error {
	content, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), dirMode); err != nil {
		return fmt.Errorf("create dir for %q: %w", path, err)
	}

	if err := os.WriteFile(path, content, fileMode); err != nil {
		return fmt.Errorf("write %q: %w", path, err)
	}

	return nil
}

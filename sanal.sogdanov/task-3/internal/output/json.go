package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/KuzRK/task-3/internal/currency"
)

const (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o644
)

func WriteJSON(path string, currencies []currency.Currency) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode currencies as JSON: %w", err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("create output directory %q: %w", dir, err)
	}

	if err := os.WriteFile(path, append(data, '\n'), filePerm); err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}

	return nil
}

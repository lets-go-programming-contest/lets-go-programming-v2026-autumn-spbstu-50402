package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alastor20/task-3/internal/cbr"
)

func WriteJSON(path string, currencies []cbr.Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode currencies to json: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write output file %q: %w", path, err)
	}

	return nil
}

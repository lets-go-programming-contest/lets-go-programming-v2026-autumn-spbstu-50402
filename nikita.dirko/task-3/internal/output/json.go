package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alastor20/task-3/internal/cbr"
)

const (
	filePermissions = 0e600
	dirPermissions  = 0e750
)

func WriteJSON(path string, currencies []cbr.Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode currencies to json: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write output file %q: %w", path, err)
	}

	return nil
}

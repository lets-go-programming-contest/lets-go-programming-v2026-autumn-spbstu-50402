package output

import (
	"encoding/json"
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
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return err
	}

	return os.WriteFile(path, append(data, '\n'), filePerm)
}

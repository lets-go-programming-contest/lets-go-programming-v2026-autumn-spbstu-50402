package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/khalii21/task-3/cmd/service/internal/xml"
)

const (
	filePermissions = 0o600
	dirPermissions  = 0o750
)

func WriteFile(path string, currencies []xml.Currency) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("make JSON failed: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), dirPermissions)
	if err != nil {
		return fmt.Errorf("create output directory failed: %w", err)
	}

	err = os.WriteFile(path, data, filePermissions)
	if err != nil {
		return fmt.Errorf("write JSON file failed: %w", err)
	}

	return nil
}

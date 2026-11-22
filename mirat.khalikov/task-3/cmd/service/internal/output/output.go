package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"task-3/cmd/service/internal/xml"
)

func WriteFile(path string, currencies []xml.Currency) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("make JSON failed: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("create output directory failed: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("write JSON file failed: %w", err)
	}

	return nil
}

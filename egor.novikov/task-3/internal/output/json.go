
package output

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/EGGorkaa/task-3/internal/currency"
)

const (
	dirPermissions  os.FileMode = 0o755
	filePermissions os.FileMode = 0o600
)

func WriteJSON(path string, currencies []currency.Currency) {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		panic(err)
	}
}

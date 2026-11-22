package output

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/EGGorkaa/task-3/internal/currency"
)

func WriteJSON(path string, currencies []currency.Currency) {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, data, 0o600); err != nil {
		panic(err)
	}
}

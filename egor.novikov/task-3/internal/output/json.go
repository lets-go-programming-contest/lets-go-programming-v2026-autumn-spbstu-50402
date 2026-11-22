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

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, data, 0600)
	if err != nil {
		panic(err)
	}
}

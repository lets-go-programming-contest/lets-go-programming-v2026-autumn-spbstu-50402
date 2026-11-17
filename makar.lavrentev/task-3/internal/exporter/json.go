package exporter

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Pilipchenok/task-3/internal/model"
)

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

func WriteJSON(valutes []model.JSONValute, outputFile string) {
	jsonData, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		panic(err)
	}
	if dir := filepath.Dir(outputFile); dir != "" {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			panic(err)
		}
	}
	if err := os.WriteFile(outputFile, jsonData, filePerm); err != nil {
		panic(err)
	}
}

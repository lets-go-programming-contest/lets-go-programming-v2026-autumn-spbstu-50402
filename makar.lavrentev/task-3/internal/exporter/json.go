package exporter

import (
	"encoding/json"
	"os"

	"github.com/Pilipchenok/task-3/internal/model"
)

const filePerm = 0o600

func WriteJSON(valutes []model.Valute, outputFile string) {
	jsonData, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(outputFile, jsonData, filePerm); err != nil {
		panic(err)
	}
}

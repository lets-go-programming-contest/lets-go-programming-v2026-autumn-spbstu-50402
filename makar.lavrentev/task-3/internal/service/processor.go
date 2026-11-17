package service

import (
	"cmp"

	"github.com/Pilipchenok/task-3/internal/model"
)

func SortFunc(a, b model.Valute) int {
	return cmp.Compare(b.Value, a.Value)
}

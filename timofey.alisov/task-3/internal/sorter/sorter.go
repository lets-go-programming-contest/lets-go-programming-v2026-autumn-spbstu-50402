package sorter

import (
	"cmp"
	"slices"

	"github.com/timfoi/task-3/internal/models"
)

func ByValueDesc(valutes []models.Valute) {
	slices.SortStableFunc(valutes, func(a, b models.Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

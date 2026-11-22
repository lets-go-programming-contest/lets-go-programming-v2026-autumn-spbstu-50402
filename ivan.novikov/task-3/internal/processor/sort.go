package processor

import (
	"sort"

	"github.com/Kayv1e/task-3/internal/model"
)

func SortByValueDesc(currencies []model.Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

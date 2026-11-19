package cbr

import (
	"cmp"
	"slices"
)

func SortByValueDesc(currencies []Currency) {
	slices.SortFunc(currencies, func(left, right Currency) int {
		return cmp.Compare(right.Value, left.Value)
	})
}

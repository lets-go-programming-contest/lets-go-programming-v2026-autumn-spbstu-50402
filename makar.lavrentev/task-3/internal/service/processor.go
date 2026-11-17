package service

import (
	"cmp"
	"strconv"
	"strings"

	"github.com/Pilipchenok/task-3/internal/model"
)

func toFloat(s string) float64 {
	s = strings.Replace(s, ",", ".", 1)
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

func SortFunc(a, b model.JSONValute) int {
	return cmp.Compare(b.Value, a.Value)
}

func SimplifyData(curs *model.ValCurs) []model.JSONValute {
	valutes := make([]model.JSONValute, 0, len(curs.Valutes))
	for _, v := range curs.Valutes {
		valutes = append(valutes, model.JSONValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    toFloat(v.Value),
		})
	}
	return valutes
}

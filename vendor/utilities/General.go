package utilities

import (
	"html/template"
	"math"
)

func IsUnset(f interface{}) bool {
	if f == nil || f == "" {
		return true
	}
	return false
}

func Float64ToCents(amount float64) int64 {
	a := int64(math.Round(amount * 100))
	return a
}

func CentsToFloat64(amount int64) float64 {
	a := float64(amount) / 100
	return a
}
func IntToPercent(i int) int {
	p := (i * 2) * 10
	return p
}

func HTML(t string) template.HTML {
	return template.HTML(t)
}

package transform

import (
	"strings"
)

func CamelToSnake(s string) string {
	return strings.ToLower(strings.Join(strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	}), "_"))
}

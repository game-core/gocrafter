package transform

import (
	"strings"
)

// CamelToSnake キャメルケースからスネークケースに変換
func CamelToSnake(s string) string {
	return strings.ToLower(strings.Join(strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	}), "_"))
}

// KebabToCamel ケバブケースからキャメルケースに変換
func KebabToCamel(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

// SingularToPlural 単数形から複数形に変換
func SingularToPlural(s string) string {
	if s == "" {
		return s
	}

	// 例外的な変換ルール
	irregularForms := map[string]string{
		"person":    "people",
		"child":     "children",
		"ox":        "oxen",
		"man":       "men",
		"woman":     "women",
		"tooth":     "teeth",
		"foot":      "feet",
		"goose":     "geese",
		"cactus":    "cacti",
		"fungus":    "fungi",
		"focus":     "foci",
		"datum":     "data",
		"medium":    "media",
		"analysis":  "analyses",
		"basis":     "bases",
		"diagnosis": "diagnoses",
		"ellipsis":  "ellipses",
	}
	if val, ok := irregularForms[s]; ok {
		return val
	}

	// 通常の変換ルール
	if strings.HasSuffix(s, "y") && len(s) > 1 && !strings.ContainsAny(string(s[len(s)-2]), "aeiouy") {
		return s[:len(s)-1] + "ies"
	}

	return s + "s"
}

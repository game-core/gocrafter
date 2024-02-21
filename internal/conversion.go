package internal

import (
	"strings"
	"unicode"
)

// SnakeToUpperCamel スネークケースからアッパーキャメルケースに変換
func SnakeToUpperCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}

// SnakeToCamel スネークケースからキャメルケースに変換
func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	result := strings.Join(parts, "")
	result = strings.ToLower(string(result[0])) + result[1:]

	return result
}

// UpperCamelToSnake アッパーキャメルケースからスネークケースに変換
func UpperCamelToSnake(s string) string {
	var result strings.Builder
	result.WriteRune(unicode.ToLower(rune(s[0])))

	for _, char := range s[1:] {
		if unicode.IsUpper(char) {
			result.WriteRune('_')
			result.WriteRune(unicode.ToLower(char))
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

// SingularToPlural 単数形から複数形に変換
func SingularToPlural(s string) string {
	if s == "" {
		return s
	}

	// 単語ごとに分割
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})

	// 最後の単語を複数形に変換
	words[len(words)-1] = convertSingularToPlural(words[len(words)-1])

	// 変換した単語を結合して返す
	return strings.Join(words, "")
}

// convertSingularToPlural 単語ごとに単数形から複数形に変換
func convertSingularToPlural(s string) string {
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

	// 例外的な変換があればそれを返す
	if val, ok := irregularForms[s]; ok {
		return val
	}

	// 最後の文字を判定して変換
	lastChar := s[len(s)-1:]
	if lastChar == "y" && len(s) > 1 && !strings.ContainsAny(string(s[len(s)-2]), "aeiouy") {
		return s[:len(s)-1] + "ies"
	} else if lastChar == "s" {
		return s + "es"
	}

	return s + "s"
}

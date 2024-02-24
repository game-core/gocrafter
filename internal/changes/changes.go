package changes

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

// UpperCamelToCamel アッパーキャメルケースからキャメルケースに変換
func UpperCamelToCamel(s string) string {
	if len(s) == 0 {
		return s
	}
	var result strings.Builder
	result.WriteRune(unicode.ToLower(rune(s[0])))
	result.WriteString(s[1:])
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

// PluralToSingular 複数形から単数形に変換
func PluralToSingular(s string) string {
	if s == "" {
		return s
	}

	// 単語ごとに分割
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})

	// 最後の単語を単数形に変換
	words[len(words)-1] = convertPluralToSingular(words[len(words)-1])

	// 変換した単語を結合して返す
	return returnToOriginalCase(s, strings.Join(words, "_"))
}

// convertPluralToSingular 単語ごとに複数形から単数形に変換
func convertPluralToSingular(s string) string {
	// 例外的な変換ルールの逆
	irregularFormsReverse := map[string]string{
		"people":    "person",
		"children":  "child",
		"oxen":      "ox",
		"men":       "man",
		"women":     "woman",
		"teeth":     "tooth",
		"feet":      "foot",
		"geese":     "goose",
		"cacti":     "cactus",
		"fungi":     "fungus",
		"foci":      "focus",
		"data":      "datum",
		"media":     "medium",
		"analysis":  "analysis",
		"basis":     "basis",
		"diagnosis": "diagnosis",
		"ellipsis":  "ellipsis",
		"bonuses":   "bonus",
		"schedules": "schedule",
	}

	// すでに単数形の場合はそのまま返す
	if val, ok := irregularFormsReverse[s]; ok {
		return val
	}

	// 通常の変換ルールの逆
	if strings.HasSuffix(s, "ies") && len(s) > 3 && !strings.ContainsAny(string(s[len(s)-4]), "aeiouy") {
		return s[:len(s)-3] + "y"
	} else if strings.HasSuffix(s, "es") && len(s) > 2 && !strings.ContainsAny(string(s[len(s)-3]), "aeiouy") {
		return s[:len(s)-2]
	} else if strings.HasSuffix(s, "s") && len(s) > 1 {
		return s[:len(s)-1]
	}

	return s
}

// returnToOriginalCase 文字列を元のケースに戻す
func returnToOriginalCase(original string, converted string) string {
	if IsCamelCase(original) {
		return SnakeToCamel(converted)
	} else if IsSnakeCase(original) {
		return converted
	}

	return converted
}

// IsCamelCase キャメルケースかどうかを判定
func IsCamelCase(s string) bool {
	return strings.IndexFunc(s, unicode.IsUpper) != -1
}

// IsSnakeCase スネークケースかどうかを判定
func IsSnakeCase(s string) bool {
	return strings.Contains(s, "_")
}

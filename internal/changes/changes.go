package changes

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// StringToInt32 stringからint32に変換
func StringToInt32(s string) (int32, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return int32(num), nil
}

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

// CamelToSnake キャメルケースからスネークケースに変換
func CamelToSnake(s string) string {
	return strings.ToLower(regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(s, "${1}_${2}"))
}

// CamelToUpperCamel アッパーキャメルケースからアッパーキャメルケースに変換
func CamelToUpperCamel(s string) string {
	if s == "" {
		return s
	}
	return string(s[0]-'a'+'A') + s[1:]
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

	// 最後の単語を複数形に変換
	words := splitWords(s)
	words[len(words)-1] = convertSingularToPlural(words[len(words)-1])

	return returnToOriginalCase(s, strings.Join(words, "_"))
}

// splitWords 文字列を単語に分割
func splitWords(s string) []string {
	return strings.Split(s, "_")
}

// convertSingularToPlural 単語ごとに単数形から複数形に変換
func convertSingularToPlural(s string) string {
	// 例外的な変換があればそれを返す
	if val, ok := GetConvertMap()[strings.ToLower(s)]; ok {
		if unicode.IsUpper(rune(s[0])) {
			return strings.Title(val)
		}
		return val
	}

	// 最後の文字を判定して変換
	lastChar := s[len(s)-1:]
	if lastChar == "y" && len(s) > 1 && !strings.ContainsAny(string(s[len(s)-2]), "aeiouyAEIOUY") {
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

	// 最後の単語を単数形に変換
	words := splitWords(s)
	words[len(words)-1] = convertPluralToSingular(words[len(words)-1])

	// 変換した単語を結合して返す
	return returnToOriginalCase(s, strings.Join(words, "_"))
}

// convertPluralToSingular 単語ごとに複数形から単数形に変換
func convertPluralToSingular(s string) string {
	// 例外的な変換があればそれを返す
	if val, ok := GetInverseConvertMap()[strings.ToLower(s)]; ok {
		if unicode.IsUpper(rune(s[0])) {
			return strings.Title(val)
		}
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

// Extraction 抽出する
func Extraction(s, es string, p int32) string {
	return strings.Split(s, es)[p]
}

package changes

// GetConvertMap 変換用のMapを取得
func GetConvertMap() map[string]string {
	return map[string]string{
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
		"bonus":     "bonuses",
		"schedule":  "schedules",
		"box":       "boxes",
	}
}

// GetInverseConvertMap 反転された変換用のMapを取得
func GetInverseConvertMap() map[string]string {
	convertMap := GetConvertMap()
	inverseMap := make(map[string]string)

	for key, value := range convertMap {
		inverseMap[value] = key
	}

	return inverseMap
}

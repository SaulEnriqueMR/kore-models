package helpers

import "strings"

func ResolveSatBoolean(booleanSat *string) bool {
	trimmedBoolean := strings.TrimSpace(*booleanSat)

	if strings.ToUpper(trimmedBoolean) == "SI" {
		return true
	}

	if strings.ToUpper(trimmedBoolean) == "NO" {
		return false
	}

	return false
}

package helpers

import (
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func NormalizeSatBoolean(satBoolean string) string {
	trimmedBoolean := strings.TrimSpace(satBoolean)

	// Normalize the string to decomposed form
	t := norm.NFD.String(trimmedBoolean)

	// Use a builder to collect non-mark characters
	var sb strings.Builder
	for _, r := range t {
		// Ignore marks (accents, etc.)
		if !unicode.Is(unicode.Mn, r) {
			sb.WriteRune(r)
		}
	}

	return strings.ToUpper(sb.String())
}

func ResolveSatBoolean(booleanSat string) bool {
	cleanBoolean := NormalizeSatBoolean(booleanSat)

	if strings.ToUpper(cleanBoolean) == "SI" {
		return true
	}

	if strings.ToUpper(cleanBoolean) == "NO" {
		return false
	}

	return false
}

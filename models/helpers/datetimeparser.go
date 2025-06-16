package helpers

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	IsoDatetimeLayout     = "2006-01-02T15:04:05"
	Rfc3339DatetimeLayout = "2006-01-02"
	CustomLayout8601      = "2006-01-02T15:04:05.99"
	IsoDateRegex          = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`
	Rfc3339Regex          = `^\d{4}-\d{2}-\d{2}$`
	XsDateRegex           = `^\d{4}-\d{2}-\d{2}(Z|([+-]\d{2}:\d{2}))?$`
	Iso8601               = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{1,3}$`
)

var MexicoLocation, errLocation = time.LoadLocation("Mexico/General")
var UTC = time.UTC

// ParseDatetime Función encargada de la conversión de los formatos de fechas utilizados por el SAT.
// Aún queda como incognita si el SAT utiliza más formatos pero de momento son los identificados
func ParseDatetime(s string) (time.Time, error) {
	// Hacemos un trim, debido a que muchas veces los campos vienen con espacios al inicio y al final.
	trimmedString := strings.TrimSpace(s)
	if trimmedString == "" {
		return time.Time{}, nil
	}
	// Primero checamos si coincide con el layout de ISO.
	patternISO := regexp.MustCompile(IsoDateRegex)
	if patternISO.MatchString(s) {
		isoDate, err := time.ParseInLocation(IsoDatetimeLayout, trimmedString, UTC)
		return isoDate, err
	}

	// En caso de que no, probamos con el RFC3339.
	patternRFC := regexp.MustCompile(Rfc3339Regex)
	if patternRFC.MatchString(s) {
		isoDate, err := time.ParseInLocation(Rfc3339DatetimeLayout, trimmedString, UTC)
		return isoDate, err
	}

	// xs:date
	patternXsDate := regexp.MustCompile(XsDateRegex)
	if patternXsDate.MatchString(trimmedString) {
		// Parse as RFC3339 date with midnight time added
		parsedTime, err := time.ParseInLocation(time.RFC3339, trimmedString[:10]+"T00:00:00"+trimmedString[10:], UTC)
		return parsedTime, err
	}

	pathIso8601 := regexp.MustCompile(Iso8601)
	if pathIso8601.MatchString(trimmedString) {
		parsedTime, err := time.ParseInLocation(CustomLayout8601, trimmedString, UTC)
		return parsedTime, err
	}

	// Caso extraordinario Retenciones 1.0. Ejemplo "2019-10-20T16:35:28-06:00")
	layout := time.RFC3339
	parsedTime, err := time.ParseInLocation(layout, trimmedString, UTC)
	if err == nil {
		return parsedTime, err
	}

	// En caso de que sea rfc3339 nativo de golang
	return time.Time{}, errors.New("error parsing datetime. String does not match supported patterns")
}

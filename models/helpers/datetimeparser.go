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
	IsoDateRegex          = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`
	Rfc3339Regex          = `^\d{4}-\d{2}-\d{2}$`
)

// ParseDatetime Función encargada de la conversión de los formatos de fechas utilizados por el SAT.
// Aún queda como incognita si el SAT utiliza más formatos pero de momento son los identificados
func ParseDatetime(s string) (time.Time, error) {
	// Hacemos un trim, debido a que muchas veces los campos vienen con espacios al inicio y al final.
	trimmedString := strings.TrimSpace(s)
	if trimmedString == "" {
		return time.Time{}, errors.New("error parsing datetime. String is empty")
	}
	// Primero checamos si coincide con el layout de ISO.
	patternISO := regexp.MustCompile(IsoDateRegex)
	if patternISO.MatchString(s) {
		isoDate, err := time.Parse(IsoDatetimeLayout, trimmedString)
		return isoDate, err
	}
	// En caso de que no, probamos con el RFC3339.
	patternRFC := regexp.MustCompile(Rfc3339Regex)
	if patternRFC.MatchString(s) {
		isoDate, err := time.Parse(Rfc3339DatetimeLayout, trimmedString)
		return isoDate, err
	}

	// Caso extraordinario Retenciones 1.0. Ejemplo "2019-10-20T16:35:28-06:00"
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, trimmedString)
	if err == nil {
		// return parsedTime, err
		// Load the Mexico City location
		location, locErr := time.LoadLocation("America/Mexico_City")
		if locErr != nil {
			return time.Time{}, locErr
		}
		// Convert the parsed time to the Mexico City location
		localTime := parsedTime.In(location)
		return localTime, nil
	}

	// En caso de que sea rfc3339 nativo de golang
	return time.Time{}, errors.New("error parsing datetime. String does not match supported patterns")
}

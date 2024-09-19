package models

import (
	"errors"
	"log"
	"strings"
	"time"
)

var (
	IsoDatetimeLayout     = "2006-01-02T15:04:05"
	Rfc3339DatetimeLayout = "2006-01-02"
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
	if isoDate, err := time.Parse(IsoDatetimeLayout, trimmedString); err == nil {
		log.Println("Entro iso")
		return isoDate, nil
	}
	// En caso de que no, probamos con el RFC3339.
	if rfc3339, err := time.Parse(Rfc3339DatetimeLayout, trimmedString); err == nil {
		log.Println("Entro rfc")
		return rfc3339, nil
	}
	// En caso de que sea rfc3339 nativo de golang
	if rfc3339, err := time.Parse(time.RFC3339, trimmedString); err == nil {
		log.Println("Entro rfc 33")
		return rfc3339, nil
	}
	// Regresamos error.
	return time.Time{}, errors.New("error parsing datetime. String does not match supported patterns")
}

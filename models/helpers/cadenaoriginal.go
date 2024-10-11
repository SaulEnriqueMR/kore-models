package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

func CreateCadenaOriginal(v interface{}) string {
	var sb strings.Builder
	val := reflect.ValueOf(v)

	processValue(val, &sb)

	// Remove the trailing "|" if present
	str := sb.String()
	if len(str) > 0 && str[len(str)-1] == '|' {
		str = str[:len(str)-1]
	}

	aux := sb.String()
	sb.Reset()
	sb.WriteString("||")
	sb.WriteString(aux)
	sb.WriteString("|")
	return sb.String()
}

var ignoreMap = map[string]struct{}{
	"Sello":                  {},
	"Certificado":            {},
	"NoCertificado":          {},
	"XMLName":                {},
	"SelloCfd":               {},
	"NoCertificadoSat":       {},
	"SelloSat":               {},
	"DocumentoFiscalDigital": {},
}

func processValue(val reflect.Value, sb *strings.Builder) {
	if !val.IsValid() || isEmptyValue(val) {
		return // Skip invalid or empty values
	}

	switch val.Kind() {
	case reflect.Ptr:
		if !val.IsNil() {
			processValue(val.Elem(), sb) // Dereference and process non-nil pointers
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := val.Type().Field(i)

			if val.Type().Field(i).PkgPath == "" { // Only process exported fields
				// Accessing field name
				fieldName := fieldType.Name
				_, shouldSkip := ignoreMap[fieldName]
				if !shouldSkip {
					processValue(field, sb)
				}
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			processValue(val.Index(i), sb)
		}
	case reflect.String:
		if val.String() != "" { // Only append non-empty strings
			sb.WriteString(val.String() + "|")
		}
	case reflect.Int, reflect.Int64, reflect.Int32:
		sb.WriteString(fmt.Sprintf("%d|", val.Int()))
	case reflect.Float64, reflect.Float32:
		sb.WriteString(fmt.Sprintf("%f|", val.Float()))
	case reflect.Bool:
		// sb.WriteString(fmt.Sprintf("%t|", val.Bool()))
		return
	default:
		sb.WriteString(fmt.Sprintf("%v|", val.Interface()))
	}
}

// Helper function to determine if a value is empty (e.g., empty string, zero int, nil pointer)
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Slice:
		return v.Len() == 0
	case reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Struct:
		return false // Structs are never considered "empty"
	}
	return false
}

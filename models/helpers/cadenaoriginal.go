package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

func CreateCadenaOriginal(s interface{}) string {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Struct {
		return ""
	}

	ignoreMap := map[string]struct{}{
		"Sello":                  {},
		"Certificado":            {},
		"NoCertificado":          {},
		"XMLName":                {},
		"SelloCfd":               {},
		"NoCertificadoSat":       {},
		"SelloSat":               {},
		"DocumentoFiscalDigital": {},
	}

	var sb strings.Builder
	sb.WriteString("||")

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		// Check if the field name is in the ignore map
		if _, found := ignoreMap[fieldType.Name]; found {
			continue // Skip this field
		}

		var value string

		if field.CanInterface() { // Check if the field can be accessed
			switch field.Kind() {
			case reflect.Bool:
				continue
			case reflect.Ptr:
				// Handle pointer types
				if field.IsNil() {
					value = ""
				} else {
					value = fmt.Sprintf("%v", CreateCadenaOriginal(field.Elem().Interface()))
				}
			case reflect.Struct:
				// Handle nested structs
				value = CreateCadenaOriginal(field.Interface())
			case reflect.Slice, reflect.Array:
				// Handle arrays and slices
				var sliceValues []string
				for j := 0; j < field.Len(); j++ {
					sliceValues = append(sliceValues, fmt.Sprintf("%v", CreateCadenaOriginal(field.Index(j).Interface())))
				}
				value = strings.Join(sliceValues, ",")
			default:
				value = fmt.Sprintf("%v", field.Interface())
			}
		} else {
			value = "" // If the field cannot be accessed, assign an empty string
		}

		sb.WriteString(value)
		sb.WriteString("|")
	}

	sb.WriteString("|")
	return sb.String()
}

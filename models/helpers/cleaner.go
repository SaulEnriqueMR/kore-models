package helpers

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func TrimXml(inputXml []byte) ([]byte, error) {
	// Regex to match attribute values
	re := regexp.MustCompile(`(\w+:\w+|[\w-]+)\s*=\s*"([^"]*)"`)

	// Replace function to trim attribute values
	trimmedXML := re.ReplaceAllStringFunc(string(inputXml), func(match string) string {
		// Split the match into key and value
		parts := strings.SplitN(match, "=", 2)
		if len(parts) != 2 {
			return match // Return the original match if it doesn't contain "="
		}

		// Trim the value (remove leading/trailing spaces and newlines)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(strings.Trim(parts[1], `"`))

		// Reconstruct the attribute with the trimmed value
		return fmt.Sprintf(`%s="%s"`, key, value)
	})

	// Remove extra spaces and newlines between attributes
	trimmedXML = regexp.MustCompile(`>\s+<`).ReplaceAllString(trimmedXML, "><")
	trimmedXML = regexp.MustCompile(`\s+`).ReplaceAllString(trimmedXML, " ")
	trimmedXML = strings.TrimSpace(trimmedXML)
	return []byte(trimmedXML), nil
}

// TrimStringAttributes Función encargada de timmear todos los atributos que sean strings.
// Es muy común que facturas en especial 3.3 o 3.2 vengan con espacios al inicio o al final.
func TrimStringAttributes(s interface{}) {
	v := reflect.ValueOf(s)

	// Ensure we are working with a pointer to a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure it's a struct before proceeding
	if v.Kind() != reflect.Struct {
		return
	}

	// Loop through all the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := field.Type()
		switch field.Kind() {
		case reflect.String:
			// Trim string fields
			field.SetString(strings.TrimSpace(field.String()))
		case reflect.Struct:
			// Recursively handle nested structs
			TrimStringAttributes(field.Addr().Interface())
		case reflect.Slice:
			// Handle slices of structs
			if fieldType.Elem().Kind() == reflect.Struct {
				for j := 0; j < field.Len(); j++ {
					TrimStringAttributes(field.Index(j).Addr().Interface())
				}
			}
		case reflect.Ptr:
			// If the field is a pointer to a struct, trim it recursively
			if field.Elem().Kind() == reflect.Struct {
				TrimStringAttributes(field.Interface())
			}
		default:
			// Default case: do nothing for other types
		}
	}
}

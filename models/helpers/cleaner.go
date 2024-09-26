package helpers

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"strings"
)

func TrimXml(inputXml []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(inputXml))
	var output bytes.Buffer
	encoder := xml.NewEncoder(&output)

	// Iterate through XML tokens
	for {
		tok, err := decoder.Token()
		if err != nil {
			// End of file or error
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			// Trim all attribute values
			for i := range t.Attr {
				t.Attr[i].Value = strings.TrimSpace(t.Attr[i].Value)
			}
			encoder.EncodeToken(t)

		case xml.EndElement:
			// Write end element
			encoder.EncodeToken(t)

		case xml.CharData:
			// Trim text content
			trimmedText := strings.TrimSpace(string(t))
			if trimmedText != "" {
				encoder.EncodeToken(xml.CharData(trimmedText))
			}

		default:
			// For other token types, pass them as they are
			encoder.EncodeToken(tok)
		}
	}

	// Close the encoder and flush the output
	encoder.Flush()

	return output.Bytes(), nil
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

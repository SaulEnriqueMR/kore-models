package comprobante

import (
	"reflect"
	"strings"
	"time"
)

type ComplementoConcepto struct {
	ACuentaTerceros *[]ACuentaTerceros40 `xml:"PorCuentadeTerceros" bson:"ACuentaTerceros,omitempty"`
}

type Complemento struct {
}

type InformacionAdicional struct {
	Comentario *string `bson:"Comentario,omitempty"`
	Desgloce   *string `bson:"Desgloce,omitempty"`
	NotaPie    *string `bson:"NotaPie,omitempty"`
	IdProyecto *string `bson:"IdProyecto,omitempty"`
	Categoria  *string `bson:"Categoria,omitempty"`
}

type Cancelacion struct {
	MotivoCancelacion  *string    `bson:"MotivoCancelacion,omitempty"`
	FolioSustitucion   *string    `bson:"FolioSustitucion,omitempty"`
	EstatusCancelacion *string    `bson:"EstatusCancelacion,omitempty"`
	EsCancelable       *string    `bson:"EsCancelable,omitempty"`
	FechaCancelacion   *time.Time `bson:"FechaCancelacion,omitempty"`
}

// TrimStringAttributes Función encargada de timmear todos los atributos que sean strings.
// Es muy común que facturas en especial 3.3 vengan con espacios al inicio o al final.
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

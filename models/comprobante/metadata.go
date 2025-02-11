package comprobante

import (
	"reflect"
	"time"

	documentofiscaldigital "github.com/SaulEnriqueMR/kore-models/models/documentofiscaldigital"
	"github.com/SaulEnriqueMR/kore-models/utils"
)

type ComprobanteMetadata struct {
	Uuid               string                                   `bson:"Uuid" json:"Uuid"`
	Fecha              time.Time                                `bson:"FechaEmision" json:"FechaEmision"`
	Vigente            bool                                     `bson:"Vigente" json:"Estatus"`
	FechaTimbrado      time.Time                                `bson:"FechaTimbrado" json:"FechaTimbrado"`
	Total              float64                                  `bson:"Total" json:"Total"`
	TipoComprobante    string                                   `bson:"TipoComprobante" json:"TipoComprobante"`
	Emisor             RfcEmisorReceptor                        `bson:"Emisor" json:"Emisor"`
	Receptor           RfcEmisorReceptor                        `bson:"Receptor" json:"Receptor"`
	RfcProvCertif      string                                   `bson:"RfcProvCertif" json:"RfcProvCertif"`
	Metadata           bool                                     `bson:"Metadata" json:"Metadata"`
	Cancelacion        *Cancelacion                             `xml:"Cancelacion" bson:"Cancelacion,omitempty" json:"Cancelacion,omitempty"`
	Transaccion        string                                   `bson:"Transaccion"`
	TotalesMonedaLocal TotalesMonedaLocal                       `bson:"TotalesMonedaLocal" json:"TotalesMonedaLocal"`
	ProcessorMetadata  documentofiscaldigital.ProcessorMetadata `bson:"ProcessorMetadata" json:"ProcessorMetadata"`
}

type Cancelacion struct {
	CanceledByKuantik *bool      `bson:"CanceledByKuantik,omitempty" json:"CanceledByKuantik,omitempty"`
	FechaCancelacion  *time.Time `bson:"FechaCancelacion,omitempty" json:"FechaCancelacion,omitempty"`
}
type RfcEmisorReceptor struct {
	Rfc    string `bson:"Rfc" json:"Rfc"`
	Nombre string `bson:"Nombre" json:"Nombre"`
}

type PagoTercerosMetadata struct {
	Uuid          string `bson:"Uuid" json:"UuidComprobante"`
	RfcTercero    string `bson:"RfcTercero" json:"RfcTercero"`
	NombreTercero string `bson:"NombreTercero" json:"NombreTercero"`
}

type ComprobanteTerceros struct {
	Uuid     string              `bson:"Uuid" json:"UuidComprobante"`
	Terceros []RfcEmisorReceptor `bson:"Terceros" json:"Terceros"`
}

type TotalesMonedaLocal struct {
	Total    float64 `bson:"Total"`
	Subtotal float64 `bson:"Subtotal"`
}

func (c ComprobanteMetadata) GetWholeStruct() []string {
	utils.TourStruct(c)

	// data := reflect.ValueOf(c)
	// fields := make([]string, data.NumField())
	// for i := 0; i < data.NumField(); i++ {
	// 	name := data.Type().Field(i).Name
	// 	log.Println(name)
	// 	log.Println(".....")
	// 	switch data.Field(i).Kind() {
	// 	case reflect.Struct:
	// 		fields = append(fields, getFields(name, data.Field(i).Interface()))
	// 	case reflect.Slice:
	// 		fields = append(fields, getFields(name, data.Field(i).Index(0).Interface()))
	// 	// pinters null too:
	// 	case reflect.Ptr:
	// 		if data.Field(i).IsNil() {
	// 			fields[i] = name
	// 		} else {
	// 			fields = append(fields, getFields(name, data.Field(i).Elem().Interface()))
	// 		}
	// 	default:
	// 		fields[i] = name
	// 	}
	// }
	// return fields
	return nil
}

func getFields(name string, data interface{}) string {
	fields := reflect.ValueOf(data)
	for i := 0; i < fields.NumField(); i++ {
		name = fields.Type().Field(i).Name
		switch fields.Field(i).Kind() {
		case reflect.Struct:
			name = getFields(name, fields.Field(i).Interface())
		case reflect.Slice:
			name = getFields(name, fields.Field(i).Index(0).Interface())
		case reflect.Ptr:
			if fields.Field(i).IsNil() {
				name = name
			} else {
				name = getFields(name, fields.Field(i).Elem().Interface())
			}
		default:
			name = name
		}
	}
	return name
}

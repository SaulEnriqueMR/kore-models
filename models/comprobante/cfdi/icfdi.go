package cfdi

import (
	"encoding/json"
	"reflect"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type ICfdi interface {
	ToCFDI() *CFDI
}

type BaseComprobante struct {
	Value ICfdi
}

func getConcreteType(version string) reflect.Type {
	switch version {
	case "3.2":
		return reflect.TypeOf(Cfdi32{})
	case "3.3":
		return reflect.TypeOf(Cfdi33{})
	case "4.0":
		return reflect.TypeOf(Cfdi40{})
	default:
		return reflect.TypeOf(CfdiMetadata{})
	}
}

func (c *BaseComprobante) UnmarshalJSON(data []byte) error {
	var aux struct {
		Version string `json:"Version"`
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	concreteType := getConcreteType(aux.Version)
	err = helpers.DecodeInterface(&c.Value, data, concreteType)
	if err != nil {
		return err
	}
	return nil
}

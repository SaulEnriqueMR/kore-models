package retenciones

import "github.com/SaulEnriqueMR/kore-models/app/timbrefiscaldigital"

type Shared struct {
	TimbreFiscalDigital *timbrefiscaldigital.TimbreFiscalDigital `xml:"TimbreFiscalDigital" bson:"TimbreFiscalDigital,omitempty"`
}

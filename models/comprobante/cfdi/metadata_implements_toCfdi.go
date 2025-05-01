package cfdi

import (
	"github.com/SaulEnriqueMR/kore-models/models/comprobante"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type CfdiMetadata comprobante.ComprobanteMetadata

func (c CfdiMetadata) ToCFDI() *CFDI {
	return &CFDI{
		Uuid:             c.Uuid,
		Vigencia:         c.Vigente,
		Tipo:             c.TipoComprobante,
		FechaCancelacion: helpers.SafeUnwrap(c.Cancelacion).FechaCancelacion,
		FechaTimbrado:    &c.FechaTimbrado,
		Emisor: Emisor{
			Contribuyente: Contribuyente{
				Rfc:    c.Emisor.Rfc,
				Nombre: c.Emisor.Nombre,
			},
		},
		Receptor: Receptor{
			Contribuyente: Contribuyente{
				Rfc:    c.Receptor.Rfc,
				Nombre: c.Receptor.Nombre,
			},
		},
		Total: c.Total,
	}
}

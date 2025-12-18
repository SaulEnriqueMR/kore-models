package cfdi

import (
	"github.com/SaulEnriqueMR/kore-models/models/comprobante"
	cfdicomplementos "github.com/SaulEnriqueMR/kore-models/models/comprobante/cfdi/cfdi_complementos"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Cfdi30 comprobante.Comprobante30

func (c Cfdi30) ToCFDI() *CFDI {

	return &CFDI{
		Uuid:             c.Uuid,
		Serie:            helpers.SafeUnwrap(c.Serie),
		Folio:            helpers.SafeUnwrap(c.Folio),
		SerieFolio:       c.KuantikMetadata.SerieFolio,
		Version:          c.Version,
		Vigencia:         helpers.SafeUnwrap(c.Vigente),
		Tipo:             c.TipoComprobante,
		MetodoPago:       helpers.SafeUnwrap(c.MetodoPago),
		FechaEmision:     &c.FechaEmision,
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
				Nombre: helpers.SafeUnwrap(c.Receptor.Nombre),
			},
			DomicilioFiscal:  helpers.SafeUnwrap(c.Receptor.Domicilio.Calle),
			ResidenciaFiscal: helpers.SafeUnwrap(c.Receptor.Domicilio.Pais),
		},
		Moneda:           helpers.SafeUnwrap(c.Moneda),
		FormaPago:        c.FormaPago,
		CondicionesPago:  helpers.SafeUnwrap(c.CondicionesPago),
		Descuento:        c.Descuento,
		TipoCambio:       helpers.SafeUnwrap(c.TipoCambio),
		Subtotal:         c.Subtotal,
		Total:            c.Total,
		Complemento:      cfdicomplementos.FromComplemento(helpers.SafeUnwrap(c.Complemento)),
		CfdiRelacionados: nil,
	}
}

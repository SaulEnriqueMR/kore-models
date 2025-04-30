package comprobante

import "github.com/SaulEnriqueMR/kore-models/models/helpers"

func (c Comprobante40) ToCFDI() *CFDI {
	return &CFDI{
		Uuid:             c.Uuid,
		Serie:            helpers.SafeUnwrap(c.Serie),
		Folio:            helpers.SafeUnwrap(c.Folio, "FolioDefault"),
		SerieFolio:       c.KuantikMetadata.SerieFolio,
		Version:          c.Version,
		Vigencia:         helpers.SafeUnwrap(c.Vigente),
		Tipo:             c.TipoComprobante,
		MetodoPago:       helpers.SafeUnwrap(c.MetodoPago),
		FechaEmision:     &c.FechaEmision,
		FechaCancelacion: c.Cancelacion.FechaCancelacion,
		FechaTimbrado:    &c.FechaTimbrado,
		Emisor: Emisor{
			Contribuyente: Contribuyente{
				Rfc:           c.Emisor.Rfc,
				Nombre:        c.Emisor.Nombre,
				RegimenFiscal: c.Emisor.RegimenFiscal,
			},
		},
		Receptor: Receptor{
			Contribuyente: Contribuyente{
				Rfc:    c.Receptor.Rfc,
				Nombre: c.Receptor.Nombre,
			},
			ResidenciaFiscal: helpers.SafeUnwrap(c.Receptor.ResidenciaFiscal),
			DomicilioFiscal:  c.Receptor.DomicilioFiscal,
			UsoCFDI:          c.Receptor.UsoCFDI,
		},
		Moneda:          c.Moneda,
		FormaPago:       helpers.SafeUnwrap(c.FormaPago),
		CondicionesPago: helpers.SafeUnwrap(c.CondicionesPago),
		Descuento:       c.Descuento,
		TipoCambio:      helpers.SafeUnwrap(c.TipoCambio, 1),
		Subtotal:        c.Subtotal,
		Total:           c.Total,
	}
}

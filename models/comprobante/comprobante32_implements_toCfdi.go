package comprobante

import (
	"strings"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

func (c Comprobante32) ToCFDI() *CFDI {
	regimenesFiscales := make([]string, 0, len(c.Emisor.RegimenesFiscales))
	for _, regimen := range c.Emisor.RegimenesFiscales {
		regimenesFiscales = append(regimenesFiscales, regimen.RegimenFiscal)
	}

	return &CFDI{
		Uuid:             c.Uuid,
		Serie:            helpers.SafeUnwrap(c.Serie),
		Folio:            helpers.SafeUnwrap(c.Folio),
		SerieFolio:       c.KuantikMetadata.SerieFolio,
		Version:          c.Version,
		Vigencia:         *c.Vigente,
		Tipo:             c.TipoComprobante,
		MetodoPago:       c.MetodoPago,
		FechaEmision:     &c.FechaEmision,
		FechaCancelacion: c.Cancelacion.FechaCancelacion,
		FechaTimbrado:    &c.FechaTimbrado,
		Emisor: Emisor{
			Contribuyente: Contribuyente{
				Rfc:           c.Emisor.Rfc,
				Nombre:        helpers.SafeUnwrap(c.Emisor.Nombre),
				RegimenFiscal: strings.Join(regimenesFiscales, ", "), // TODO: check if this is correct
			},
		},
		Receptor: Receptor{
			Contribuyente: Contribuyente{
				Rfc:    c.Receptor.Rfc,
				Nombre: helpers.SafeUnwrap(c.Receptor.Nombre),
			},
			DomicilioFiscal:  helpers.SafeUnwrap(&c.Receptor.Domicilio.Calle),
			ResidenciaFiscal: helpers.SafeUnwrap(&c.Receptor.Domicilio.Pais),
		},
		Moneda:          helpers.SafeUnwrap(c.Moneda),
		FormaPago:       c.FormaPago,
		CondicionesPago: helpers.SafeUnwrap(c.CondicionesPago),
		Descuento:       c.Descuento,
		TipoCambio:      c.TipoCambio,
		Subtotal:        c.Subtotal,
		Total:           c.Total,
	}
}

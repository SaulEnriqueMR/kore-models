package comprobante

func (c ComprobanteMetadata) ToCFDI() *CFDI {
	return &CFDI{
		Uuid:             c.Uuid,
		Vigencia:         c.Vigente,
		Tipo:             c.TipoComprobante,
		FechaCancelacion: c.Cancelacion.FechaCancelacion,
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

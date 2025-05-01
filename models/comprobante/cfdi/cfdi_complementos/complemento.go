package cfdicomplementos

import (
	"github.com/SaulEnriqueMR/kore-models/models/comprobante"
	complementopagos "github.com/SaulEnriqueMR/kore-models/models/comprobante/cfdi/cfdi_complementos/complemento_pagos"
)

type ComplementoCfdi struct {
	Pagos []complementopagos.IComplementoPagos
}

func FromComplemento(complemento comprobante.Complemento) (res ComplementoCfdi) {
	// Pagos
	if complemento.Pagos != nil {
		res.Pagos = make([]complementopagos.IComplementoPagos, 0)
		if complemento.Pagos.Pagos10 != nil {
			for _, pago := range *complemento.Pagos.Pagos10 {
				res.Pagos = append(res.Pagos, complementopagos.Pagos10(pago))
			}
		}
		if complemento.Pagos.Pagos20 != nil {
			for _, pago := range *complemento.Pagos.Pagos20 {
				res.Pagos = append(res.Pagos, complementopagos.Pagos20(pago))
			}
		}
	}
	return
}

package complementopagos

import (
	"github.com/AngelTheTwin/slicesutils"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pagos"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Pagos10 pagos.Pagos10

func (p Pagos10) ToPagoCfdi() *ComplementoPagos {
	return &ComplementoPagos{
		Version: p.Version,
		Pagos: slicesutils.Map(p.Pago, func(pagoItem pagos.PagoPagos10) Pago {

			var detalleImpuestos DetalleImpuestos
			impuestosPorDocumento := make([][]Impuesto, len(*pagoItem.DocumentosRelacionados))

			if pagoItem.Impuestos != nil {
				detalleImpuestos.TotalImpuestosRetenidos = new(float64)
				detalleImpuestos.TotalImpuestosTrasladados = new(float64)
				for i, detalle := range *pagoItem.Impuestos {
					*detalleImpuestos.TotalImpuestosRetenidos += helpers.SafeUnwrap(detalle.TotalImpuestosRetenidos)
					*detalleImpuestos.TotalImpuestosTrasladados += helpers.SafeUnwrap(detalle.TotalImpuestosTrasladados)

					newImpuestos := []Impuesto{}
					retenciones := slicesutils.Map(helpers.SafeUnwrap(detalle.Retenciones), func(retencion pagos.RetencionPagos10) Impuesto {
						return Impuesto{
							Tipo:     "Retencion",
							Impuesto: retencion.Impuesto,
							Importe:  retencion.Importe,
						}
					})

					traslados := slicesutils.Map(helpers.SafeUnwrap(detalle.Traslados), func(traslado pagos.TrasladoPagos10) Impuesto {
						return Impuesto{
							Tipo:       "Traslado",
							Impuesto:   traslado.Impuesto,
							TipoFactor: traslado.TipoFactor,
							TasaOCuota: traslado.TasaOCuota,
							Importe:    traslado.Importe,
						}
					})

					newImpuestos = append(newImpuestos, retenciones...)
					newImpuestos = append(newImpuestos, traslados...)
					detalleImpuestos.Impuestos = append(detalleImpuestos.Impuestos, newImpuestos...)
					impuestosPorDocumento[i] = newImpuestos
				}
			}

			if pagoItem.DocumentosRelacionados == nil {
				pagoItem.DocumentosRelacionados = &[]pagos.DoctoRelacionadoPagos10{}
			}

			documentosRelacionados := slicesutils.Map(*pagoItem.DocumentosRelacionados, func(docRelacionado pagos.DoctoRelacionadoPagos10) DocumentoRelacionado {
				return DocumentoRelacionado{
					IdDocumento:          docRelacionado.IdDocumento,
					IdDocumentoString:    docRelacionado.IdDocumentoString,
					Serie:                docRelacionado.Serie,
					Folio:                docRelacionado.Folio,
					Moneda:               docRelacionado.Moneda,
					TipoCambio:           docRelacionado.TipoCambio,
					MetodoPago:           docRelacionado.MetodoPago,
					NoParcialidad:        helpers.SafeUnwrap(docRelacionado.NoParcialidad),
					ImporteSaldoAnterior: helpers.SafeUnwrap(docRelacionado.ImporteSaldoAnterior),
					ImportePagado:        helpers.SafeUnwrap(docRelacionado.ImportePagado),
					ImporteSaldoInsoluto: helpers.SafeUnwrap(docRelacionado.ImporteSaldoInsoluto),
				}
			})

			for i, impuestosDR := range impuestosPorDocumento {
				documentosRelacionados[i].Impuestos = impuestosDR
			}

			return Pago{
				FechaPago:   pagoItem.FechaPago,
				FormaPago:   pagoItem.FormaPago,
				Moneda:      pagoItem.Moneda,
				TipoCambio:  pagoItem.TipoCambio,
				Monto:       pagoItem.Monto,
				NoOperacion: pagoItem.NoOperacion,
				CuentaOrdenante: &CuentaOrdenante{
					RfcEmisorCuentaOrdenante:       pagoItem.RfcEmisorCuentaOrdenante,
					NombreBancoOrdenanteExtranjero: pagoItem.NombreBancoOrdenanteExtranjero,
					CuentaOrdenante:                pagoItem.CuentaOrdenante,
				},
				CuentaBeneficiario: &CuentaBeneficiario{
					RfcEmisorCuentaBeneficiario: pagoItem.RfcEmisorCuentaBeneficiario,
					CuentaBeneficiario:          pagoItem.CuentaBeneficiario,
				},
				TipoCadenaPago:         pagoItem.TipoCadenaPago,
				CertificadoPago:        pagoItem.CertificadoPago,
				CadenaPago:             pagoItem.CadenaPago,
				SelloPago:              pagoItem.SelloPago,
				DocumentosRelacionados: documentosRelacionados,
				DetalleImpuestos:       &detalleImpuestos,
			}
		}),
	}
}

package complementopagos

import (
	"github.com/AngelTheTwin/slicesutils"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pagos"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Pagos20 pagos.Pagos20

func (p Pagos20) ToPagoCfdi() *ComplementoPagos {
	return &ComplementoPagos{
		Version: p.Version,
		Pagos: slicesutils.Map(p.Pago, func(pagoItem pagos.PagoPagos20) Pago {
			detalleImpuestos := slicesutils.Reduce(helpers.SafeUnwrap(helpers.SafeUnwrap(pagoItem.Impuestos).Retenciones), func(acum DetalleImpuestos, curr pagos.RetencionPagos20) DetalleImpuestos {
				*acum.TotalImpuestosRetenidos += curr.Importe
				acum.Impuestos = append(acum.Impuestos, Impuesto{
					Tipo:     "Retencion",
					Impuesto: curr.Impuesto,
					Importe:  curr.Importe,
				})
				return acum
			}, DetalleImpuestos{
				TotalImpuestosRetenidos:   new(float64),
				TotalImpuestosTrasladados: new(float64),
				Impuestos:                 make([]Impuesto, 0),
			})

			detalleImpuestos = slicesutils.Reduce(helpers.SafeUnwrap(helpers.SafeUnwrap(pagoItem.Impuestos).Traslados), func(acum DetalleImpuestos, curr pagos.TrasladoPagos20) DetalleImpuestos {
				*acum.TotalImpuestosTrasladados += helpers.SafeUnwrap(curr.Importe)
				acum.Impuestos = append(acum.Impuestos, Impuesto{
					Tipo:       "Traslado",
					Base:       curr.Base,
					Impuesto:   curr.Impuesto,
					TipoFactor: curr.TipoFactor,
					TasaOCuota: helpers.SafeUnwrap(curr.TasaOCuota),
					Importe:    helpers.SafeUnwrap(curr.Importe),
				})
				return acum
			}, detalleImpuestos)

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
				TipoCadenaPago:  pagoItem.TipoCadenaPago,
				CertificadoPago: pagoItem.CertificadoPago,
				CadenaPago:      pagoItem.CadenaPago,
				SelloPago:       pagoItem.SelloPago,
				DocumentosRelacionados: slicesutils.Map(pagoItem.DocumentosRelacionados, func(dr pagos.DoctoRelacionadoPagos20) DocumentoRelacionado {
					return DocumentoRelacionado{
						IdDocumento:          dr.IdDocumento,
						IdDocumentoString:    dr.IdDocumentoString,
						Serie:                dr.Serie,
						Folio:                dr.Folio,
						Moneda:               dr.Moneda,
						Equivalencia:         dr.Equivalencia,
						NoParcialidad:        dr.NoParcialidad,
						ImporteSaldoAnterior: dr.ImporteSaldoAnterior,
						ImportePagado:        dr.ImportePagado,
						ImporteSaldoInsoluto: dr.ImporteSaldoInsoluto,
						ObjetoImpuesto:       dr.ObjetoImpuesto,
						Impuestos: func(detalle *pagos.ImpuestosDRPagos20) []Impuesto {

							if detalle == nil {
								return make([]Impuesto, 0)
							}

							impuestos := slicesutils.Reduce(helpers.SafeUnwrap(detalle.Retenciones), func(acum []Impuesto, curr pagos.RetencionDRPagos20) []Impuesto {
								return append(acum, Impuesto{
									Tipo:       "Retencion",
									Base:       curr.Base,
									Impuesto:   curr.Impuesto,
									TipoFactor: curr.TipoFactor,
									TasaOCuota: curr.TasaOCuota,
									Importe:    curr.Importe,
								})
							}, make([]Impuesto, 0))
							impuestos = slicesutils.Reduce(helpers.SafeUnwrap(detalle.Traslados), func(acum []Impuesto, curr pagos.TrasladoDRPagos20) []Impuesto {
								return append(acum, Impuesto{
									Tipo:       "Traslado",
									Base:       curr.Base,
									Impuesto:   curr.Impuesto,
									TipoFactor: curr.TipoFactor,
									TasaOCuota: helpers.SafeUnwrap(curr.TasaOCuota),
									Importe:    helpers.SafeUnwrap(curr.Importe),
								})
							}, impuestos)

							return impuestos
						}(dr.Impuestos),
					}
				}),
				DetalleImpuestos: &detalleImpuestos,
			}
		}),
	}
}

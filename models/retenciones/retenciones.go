package retenciones

import (
	"encoding/xml"
	"strings"
)

type Retenciones struct {
	Retenciones10 *Retenciones10 `bson:"Retenciones10,omitempty"`
	Retenciones20 *Retenciones20 `bson:"Retenciones20,omitempty"`
}

func (r *Retenciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	//PATHS atributos
	// 	"uuid": "5E6AB80E-D663-484A-BF87-D73FC7E86D0D",
	// 	"json_content": {
	// 		"retencion": "1",
	// 		"clave_retencion": "01",
	// 		"folio": "ABC",
	// 		"codigo_postal": "91098",
	// 		"fecha": "2024-02-19T12:14:28",
	// 		"version": "2.0",
	// 		"retenciones_relacionadas": {},
	// 		"emisor": {
	// 			"regimen_fiscal_emisor": "624",
	// 			"regimen_fiscal": "624",
	// 			"nombre": "ESCUELA KEMPER URGATE",
	// 			"rfc": "EKU9003173C9"
	// 		},
	// 		"receptor": {
	// 			"nacionalidad": "Extranjero",
	// 			"extranjero": {
	// 				"nombre": "SW SAPIEN",
	// 				"num_reg_id_trib": "ASDASADS",
	// 				"tax_id": "DQWDDQWDQ"
	// 			}
	// 		},
	// 		"periodo": {
	// 			"ejercicio": "2022",
	// 			"mes_fin": "09",
	// 			"mes_inicio": "08"
	// 		},
	// 		"totales": {
	// 			"total_operacion": "2000.00",
	// 			"total_exento": "0",
	// 			"total_retencion": "580.00",
	// 			"total_gravado": "2000.00",
	// 			"impuestos_retenidos": [
	// 				{
	// 					"monto": "580.00",
	// 					"tipo_pago": "03",
	// 					"base": "2000.00",
	// 					"impuesto": "001"
	// 				}
	// 			]
	// 		}
	// 	},
	// 	"xml": "",
	// 	"docType": "comprobante",
	// 	"actionType": "timbrado"
	// }

	// RETENCIONES 10
	// Retenciones10.CveRetenc
	// Retenciones10.Emisor.RFCEmisor
	// Retenciones10.Version
	// // para UUID de ambas versiones
	// Retenciones10.Complemento.TimbreFiscalDigital10.UUID
	// Retenciones10.Complemento.TimbreFiscalDigital11.UUID
	// Retenciones10.FolioInt
	// Retenciones10.FechaExp
	// Retenciones10.Totales.MontoTotOperacion
	// Retenciones10.Totales.MontoTotExent
	// Retenciones10.Totales.MontoTotRet
	// Retenciones10.Totales.MontoTotGrav
	// Retenciones10.Receptor.Nacionalidad
	// Retenciones10.Receptor.Nacional.RFCRecep
	// // para nombre receptor
	// Retenciones10.Receptor.Nacional.NomDenRazSocR
	// Retenciones10.Receptor.Extranjero.NomDenRazSocR

	// Retenciones10.Receptor.Extranjero.NumRegIdTrib

	// // RETENCIONES 20
	// Retenciones20.CveRetenc
	// Retenciones20.FechaExp
	// Retenciones20.Version
	// // para UUID de ambas versiones
	// Retenciones20.Complemento.TimbreFiscalDigital10.UUID
	// Retenciones20.Complemento.TimbreFiscalDigital11.UUID
	// Retenciones20.FolioInt
	// Retenciones20.FechaExp
	// Retenciones20.Totales.MontoTotOperacion
	// Retenciones20.Totales.MontoTotExent
	// Retenciones20.Totales.MontoTotRet
	// Retenciones20.Totales.MontoTotGrav
	// Retenciones20.Receptor.NacionalidadR
	// Retenciones20.Receptor.Nacional.RfcR
	// Retenciones20.Receptor.Nacional.NomDenRazSocR
	// Retenciones20.Receptor.Extranjero.NomDenRazSocR
	// Retenciones20.Receptor.Extranjero.NumRegIdTribR

	var Version string
	for _, attributes := range start.Attr {
		if attributes.Name.Local == "Version" {
			Version = attributes.Value
			Version = strings.TrimSpace(Version)
			break
		}
	}

	if Version == "1.0" {
		var reten10 Retenciones10
		if err := d.DecodeElement(&reten10, &start); err != nil {
			return err
		}

		if reten10.Complemento.TimbreFiscalDigital != nil {
			if reten10.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10 != nil {
				uuid := reten10.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10.Uuid
				reten10.Uuid = strings.ToUpper(uuid)
				reten10.FechaTimbrado = reten10.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10.FechaTimbrado
			}
			if reten10.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11 != nil {
				uuid := reten10.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11.Uuid
				reten10.Uuid = strings.ToUpper(uuid)
				reten10.FechaTimbrado = reten10.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11.FechaTimbrado
			}
		}

		r.Retenciones10 = &reten10
	}

	if Version == "2.0" {
		var reten20 Retenciones20
		if err := d.DecodeElement(&reten20, &start); err != nil {
			return err
		}

		if reten20.Complemento.TimbreFiscalDigital != nil {
			if reten20.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10 != nil {
				uuid := reten20.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10.Uuid
				reten20.Uuid = strings.ToUpper(uuid)
				reten20.FechaTimbrado = reten20.Complemento.TimbreFiscalDigital.TimbreFiscalDigital10.FechaTimbrado
			}
			if reten20.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11 != nil {
				uuid := reten20.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11.Uuid
				reten20.Uuid = strings.ToUpper(uuid)
				reten20.FechaTimbrado = reten20.Complemento.TimbreFiscalDigital.TimbreFiscalDigital11.FechaTimbrado
			}
		}

		r.Retenciones20 = &reten20
	}
	return nil
}

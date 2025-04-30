package gastohidrocarburos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type GastoHidrocarburos10 struct {
	Version         string                     `xml:"Version,attr" bson:"Version" json:"Version"`
	NumeroContrato  string                     `xml:"NumeroContrato,attr" bson:"NoContrato" json:"NoContrato"`
	AreaContractual *string                    `xml:"AreaContractual,attr" bson:"AreaContractual,omitempty" json:"AreaContractual,omitempty"`
	Erogacion       []ErogacionGastoHidrocar10 `xml:"Erogacion" bson:"Erogacion,omitempty" json:"Erogacion,omitempty"`
}

type ErogacionGastoHidrocar10 struct {
	TipoErogacion        string                                `xml:"TipoErogacion,attr" bson:"Tipo" json:"Tipo"`
	MontoCuErogacion     string                                `xml:"MontocuErogacion,attr" bson:"MontoCuErogacion" json:"MontoCuErogacion"`
	Porcentaje           string                                `xml:"Porcentaje,attr" bson:"Porcentaje" json:"Porcentaje"`
	DocumentoRelacionado []DocumentoRelacionadoGastoHidrocar10 `xml:"DocumentoRelacionado" bson:"DocumentoRelacionado,omitempty" json:"DocumentoRelacionado,omitempty"`
	Actividades          *[]ActividadesGastoHidrocar10         `xml:"Actividades" bson:"Actividades,omitempty" json:"Actividades,omitempty"`
	CentroCostos         *[]CentroCostosGastoHidrocar10        `xml:"CentroCostos" bson:"CentroCostos,omitempty" json:"CentroCostos,omitempty"`
}

type DocumentoRelacionadoGastoHidrocar10 struct {
	OrigenErogacion                 string    `xml:"OrigenErogacion,attr" bson:"OrigenErogacion" json:"OrigenErogacion"`
	FolioFiscalVinculado            *string   `xml:"FolioFiscalVinculado,attr" bson:"FolioFiscalVinculado,omitempty" json:"FolioFiscalVinculado,omitempty"`
	RfcProveedor                    *string   `xml:"RFCProveedor,attr" bson:"RfcProveedor,omitempty" json:"RfcProveedor,omitempty"`
	MontoTotalIva                   *float64  `xml:"MontoTotalIVA,attr" bson:"MontoTotalIva,omitempty" json:"MontoTotalIva,omitempty"`
	MontoRetencionIsr               *float64  `xml:"MontoRetencionISR,attr" bson:"MontoRetencionIsr,omitempty" json:"MontoRetencionIsr,omitempty"`
	MontoRetencionIva               *float64  `xml:"MontoRetencionIVA,attr" bson:"MontoRetencionIva,omitempty" json:"MontoRetencionIva,omitempty"`
	MontoRetencionOtrosImpuestos    *float64  `xml:"MontoRetencionOtrosImpuestos,attr" bson:"MontoRetencionOtrosImpuestos,omitempty" json:"MontoRetencionOtrosImpuestos,omitempty"`
	NumeroPedimentoVinculado        *string   `xml:"NumeroPedimentoVinculado,attr" bson:"NumeroPedimentoVinculado,omitempty" json:"NumeroPedimentoVinculado,omitempty"`
	ClavePedimentoVinculado         *string   `xml:"ClavePedimentoVinculado,attr" bson:"ClavePedimentoVinculado,omitempty" json:"ClavePedimentoVinculado,omitempty"`
	ClavePagoPedimentoVinculado     *string   `xml:"ClavePagoPedimentoVinculado,attr" bson:"ClavePagoPedimentoVinculado,omitempty" json:"ClavePagoPedimentoVinculado,omitempty"`
	MontoIvaPedimento               *float64  `xml:"MontoIVAPedimento,attr" bson:"MontoIvaPedimento,omitempty" json:"MontoIvaPedimento,omitempty"`
	OtrosImpuestosPagadosPedimento  *float64  `xml:"OtrosImpuestosPagadosPedimento,attr" bson:"OtrosImpuestosPagadosPedimento,omitempty" json:"OtrosImpuestosPagadosPedimento,omitempty"`
	FechaFolioFiscalVinculadoString *string   `xml:"FechaFolioFiscalVinculado,attr" bson:"FechaFolioFiscalVinculadoString" json:"FechaFolioFiscalVinculadoString"`
	FechaFolioFiscalVinculado       time.Time `bson:"FechaFolioFiscalVinculado,omitempty" json:"FechaFolioFiscalVinculado,omitempty"`
	Mes                             string    `xml:"Mes,attr" bson:"Mes" json:"Mes"`
	MontoTotalErogaciones           float64   `xml:"MontoTotalErogaciones,attr" bson:"MontoTotalErogaciones,omitempty" json:"MontoTotalErogaciones,omitempty"`
}

type ActividadesGastoHidrocar10 struct {
	ActividadRelacionada *string                          `xml:"ActividadRelacionada,attr" bson:"ActividadRelacionada,omitempty" json:"ActividadRelacionada,omitempty"`
	Subactividades       *[]SubActividadesGastoHidrocar10 `xml:"SubActividades" bson:"Subactividades,omitempty" json:"Subactividades,omitempty"`
}

type SubActividadesGastoHidrocar10 struct {
	SubactividadRelacionada *string                  `xml:"SubActividadRelacionada,attr" bson:"SubactividadRelacionada,omitempty" json:"SubactividadRelacionada,omitempty"`
	Tareas                  *[]TareasGastoHidrocar10 `xml:"Tareas" bson:"Tareas,omitempty" json:"Tareas,omitempty"`
}

type TareasGastoHidrocar10 struct {
	TareaRelacionada *string `xml:"TareaRelacionada,attr" bson:"TareaRelacionada,omitempty" json:"TareaRelacionada,omitempty"`
}

type CentroCostosGastoHidrocar10 struct {
	Campo       *string                       `xml:"Campo,attr" bson:"Campo,omitempty" json:"Campo,omitempty"`
	Yacimientos *[]YacimientosGastoHidrocar10 `xml:"Yacimientos" bson:"Yacimientos,omitempty" json:"Yacimientos,omitempty"`
}

type YacimientosGastoHidrocar10 struct {
	Yacimiento *string                 `xml:"Yacimiento,attr" bson:"Yacimiento,omitempty" json:"Yacimiento,omitempty"`
	Pozos      *[]PozosGastoHidrocar10 `xml:"Pozos" bson:"Pozos,omitempty" json:"Pozos,omitempty"`
}

type PozosGastoHidrocar10 struct {
	Pozo *string `xml:"Pozo,attr" bson:"Pozo,omitempty" json:"Pozo,omitempty"`
}

func (iavv10 *DocumentoRelacionadoGastoHidrocar10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DocumentoRelacionadoGastoHidrocar10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*iavv10 = DocumentoRelacionadoGastoHidrocar10(aux)

	if iavv10 != nil {
		if iavv10.FechaFolioFiscalVinculadoString != nil {
			fecha, err := helpers.ParseDatetime(*iavv10.FechaFolioFiscalVinculadoString)
			if err != nil {
				return err
			}
			iavv10.FechaFolioFiscalVinculado = fecha
		}
	}

	return nil
}

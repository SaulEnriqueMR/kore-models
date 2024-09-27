package gastohidrocarburos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type GastoHidrocarburos10 struct {
	Version         string                     `xml:"Version,attr" bson:"Version"`
	NumeroContrato  string                     `xml:"NumeroContrato,attr" bson:"NoContrato"`
	AreaContractual *string                    `xml:"AreaContractual,attr" bson:"AreaContractual,omitempty"`
	Erogacion       []ErogacionGastoHidrocar10 `xml:"Erogacion" bson:"Erogacion,omitempty"`
}

type ErogacionGastoHidrocar10 struct {
	TipoErogacion        string                                `xml:"TipoErogacion,attr" bson:"Tipo"`
	MontoCuErogacion     string                                `xml:"MontocuErogacion,attr" bson:"MontoCuErogacion"`
	Porcentaje           string                                `xml:"Porcentaje,attr" bson:"Porcentaje"`
	DocumentoRelacionado []DocumentoRelacionadoGastoHidrocar10 `xml:"DocumentoRelacionado" bson:"DocumentoRelacionado,omitempty"`
	Actividades          *[]ActividadesGastoHidrocar10         `xml:"Actividades" bson:"Actividades,omitempty"`
	CentroCostos         *[]CentroCostosGastoHidrocar10        `xml:"CentroCostos" bson:"CentroCostos,omitempty"`
}

type DocumentoRelacionadoGastoHidrocar10 struct {
	OrigenErogacion                 string    `xml:"OrigenErogacion,attr" bson:"OrigenErogacion"`
	FolioFiscalVinculado            *string   `xml:"FolioFiscalVinculado,attr" bson:"FolioFiscalVinculado,omitempty"`
	RfcProveedor                    *string   `xml:"RFCProveedor,attr" bson:"RfcProveedor,omitempty"`
	MontoTotalIva                   *float64  `xml:"MontoTotalIVA,attr" bson:"MontoTotalIva,omitempty"`
	MontoRetencionIsr               *float64  `xml:"MontoRetencionISR,attr" bson:"MontoRetencionIsr,omitempty"`
	MontoRetencionIva               *float64  `xml:"MontoRetencionIVA,attr" bson:"MontoRetencionIva,omitempty"`
	MontoRetencionOtrosImpuestos    *float64  `xml:"MontoRetencionOtrosImpuestos,attr" bson:"MontoRetencionOtrosImpuestos,omitempty"`
	NumeroPedimentoVinculado        *string   `xml:"NumeroPedimentoVinculado,attr" bson:"NumeroPedimentoVinculado,omitempty"`
	ClavePedimentoVinculado         *string   `xml:"ClavePedimentoVinculado,attr" bson:"ClavePedimentoVinculado,omitempty"`
	ClavePagoPedimentoVinculado     *string   `xml:"ClavePagoPedimentoVinculado,attr" bson:"ClavePagoPedimentoVinculado,omitempty"`
	MontoIvaPedimento               *float64  `xml:"MontoIVAPedimento,attr" bson:"MontoIvaPedimento,omitempty"`
	OtrosImpuestosPagadosPedimento  *float64  `xml:"OtrosImpuestosPagadosPedimento,attr" bson:"OtrosImpuestosPagadosPedimento,omitempty"`
	FechaFolioFiscalVinculadoString *string   `xml:"FechaFolioFiscalVinculado,attr"`
	FechaFolioFiscalVinculado       time.Time `bson:"FechaFolioFiscalVinculado,omitempty"`
	Mes                             string    `xml:"Mes,attr" bson:"Mes"`
	MontoTotalErogaciones           float64   `xml:"MontoTotalErogaciones,attr" bson:"MontoTotalErogaciones,omitempty"`
}

type ActividadesGastoHidrocar10 struct {
	ActividadRelacionada *string                          `xml:"ActividadRelacionada,attr" bson:"ActividadRelacionada,omitempty"`
	Subactividades       *[]SubActividadesGastoHidrocar10 `xml:"SubActividades" bson:"Subactividades,omitempty"`
}

type SubActividadesGastoHidrocar10 struct {
	SubactividadRelacionada *string                  `xml:"SubActividadRelacionada,attr" bson:"SubactividadRelacionada,omitempty"`
	Tareas                  *[]TareasGastoHidrocar10 `xml:"Tareas" bson:"Tareas,omitempty"`
}

type TareasGastoHidrocar10 struct {
	TareaRelacionada *string `xml:"TareaRelacionada,attr" bson:"TareaRelacionada,omitempty"`
}

type CentroCostosGastoHidrocar10 struct {
	Campo       *string                       `xml:"Campo,attr" bson:"Campo,omitempty"`
	Yacimientos *[]YacimientosGastoHidrocar10 `xml:"Yacimientos" bson:"Yacimientos,omitempty"`
}

type YacimientosGastoHidrocar10 struct {
	Yacimiento *string                 `xml:"Yacimiento,attr" bson:"Yacimiento,omitempty"`
	Pozos      *[]PozosGastoHidrocar10 `xml:"Pozos" bson:"Pozos,omitempty"`
}

type PozosGastoHidrocar10 struct {
	Pozo *string `xml:"Pozo,attr" bson:"Pozo,omitempty"`
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
		fecha, err := helpers.ParseDatetime(*iavv10.FechaFolioFiscalVinculadoString)
		if err != nil {
			return err
		}
		iavv10.FechaFolioFiscalVinculado = fecha
	}

	return nil
}

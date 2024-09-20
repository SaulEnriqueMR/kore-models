package complementos

type GastoHidrocarburos10 struct {
	Version         string                     `xml:"Version,attr" bson:"Version"`
	NumeroContrato  string                     `xml:"NumeroContrato,attr" bson:"NumeroContrato"`
	AreaContractual *string                    `xml:"AreaContractual,attr" bson:"AreaContractual,omitempty"`
	Erogacion       []ErogacionGastoHidrocar10 `xml:"Erogacion" bson:"Erogacion,omitempty"`
}

type ErogacionGastoHidrocar10 struct {
	TipoErogacion        string                                `xml:"TipoErogacion,attr" bson:"TipoErogacion"`
	MontoCuErogacion     string                                `xml:"MontocuErogacion,attr" bson:"MontoCuErogacion"`
	Porcentaje           string                                `xml:"Porcentaje,attr" bson:"Porcentaje"`
	DocumentoRelacionado []DocumentoRelacionadoGastoHidrocar10 `xml:"DocumentoRelacionado" bson:"DocumentoRelacionado,omitempty"`
	Actividades          *[]ActividadesGastoHidrocar10         `xml:"Actividades" bson:"Actividades,omitempty"`
	CentroCostos         *[]CentroCostosGastoHidrocar10        `xml:"CentroCostos" bson:"CentroCostos,omitempty"`
}

type DocumentoRelacionadoGastoHidrocar10 struct {
	OrigenErogacion                string   `xml:"OrigenErogacion,attr" bson:"OrigenErogacion"`
	FolioFiscalVinculado           *string  `xml:"FolioFiscalVinculado,attr" bson:"FolioFiscalVinculado,omitempty"`
	RFCProveedor                   *string  `xml:"RFCProveedor,attr" bson:"RFCProveedor,omitempty"`
	MontoTotalIVA                  *float64 `xml:"MontoTotalIVA,attr" bson:"MontoTotalIVA,omitempty"`
	MontoRetencionISR              *float64 `xml:"MontoRetencionISR,attr" bson:"MontoRetencionISR,omitempty"`
	MontoRetencionIVA              *float64 `xml:"MontoRetencionIVA,attr" bson:"MontoRetencionIVA,omitempty"`
	MontoRetencionOtrosImpuestos   *float64 `xml:"MontoRetencionOtrosImpuestos,attr" bson:"MontoRetencionOtrosImpuestos,omitempty"`
	NumeroPedimentoVinculado       *string  `xml:"NumeroPedimentoVinculado,attr" bson:"NumeroPedimentoVinculado,omitempty"`
	ClavePedimentoVinculado        *string  `xml:"ClavePedimentoVinculado,attr" bson:"ClavePedimentoVinculado,omitempty"`
	ClavePagoPedimentoVinculado    *string  `xml:"ClavePagoPedimentoVinculado,attr" bson:"ClavePagoPedimentoVinculado,omitempty"`
	MontoIVAPedimento              *float64 `xml:"MontoIVAPedimento,attr" bson:",omitempty"`
	OtrosImpuestosPagadosPedimento *float64 `xml:"OtrosImpuestosPagadosPedimento,attr" bson:"OtrosImpuestosPagadosPedimento,omitempty"`
	FechaFolioFiscalVinculado      *string  `xml:"FechaFolioFiscalVinculado,attr" bson:"FechaFolioFiscalVinculado,omitempty"`
	Mes                            string   `xml:"Mes,attr" bson:"Mes"`
	MontoTotalErogaciones          float64  `xml:"MontoTotalErogaciones,attr" bson:"MontoTotalErogaciones,omitempty"`
}

type ActividadesGastoHidrocar10 struct {
	ActividadRelacionada *string                          `xml:"ActividadRelacionada,attr" bson:"ActividadRelacionada,omitempty"`
	SubActividades       *[]SubActividadesGastoHidrocar10 `xml:"SubActividades" bson:"SubActividades,omitempty"`
}

type SubActividadesGastoHidrocar10 struct {
	SubActividadRelacionada *string                  `xml:"SubActividadRelacionada,attr" bson:"SubActividadRelacionada,omitempty"`
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

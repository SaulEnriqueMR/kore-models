package fideicomisonoempresarial

type FideicomisoEmpresarial10 struct {
	Version                          string                           `xml:"Version,attr" bson:"Version" json:"Version"`
	IngresosOEntradas                IngresosOEntradas                `xml:"IngresosOEntradas" bson:"IngresosOEntradas" json:"IngresosOEntradas"`
	DeduccionesOSalidas              DeduccOSalidas                   `xml:"DeduccOSalidas" bson:"DeduccionesOSalidas" json:"DeduccionesOSalidas"`
	RetencionesEfectuadasFideicomiso RetencionesEfectuadasFideicomiso `xml:"RetEfectFideicomiso" bson:"RetencionesEfectuadasFideicomiso" json:"RetencionesEfectuadasFideicomiso"`
}

type IngresosOEntradas struct {
	MontoTotalEntradasPeriodo              float64             `xml:"MontTotEntradasPeriodo,attr" bson:"MontoTotalEntradasPeriodo" json:"MontoTotalEntradasPeriodo"`
	ParteProporcionalAcumulableFideicomiso float64             `xml:"PartPropAcumDelFideicom,attr" bson:"ParteProporcionalAcumulableFideicomiso" json:"ParteProporcionalAcumulableFideicomiso"`
	ProporcionMontoTotal                   float64             `xml:"PropDelMontTot,attr" bson:"ProporcionMontoTotal" json:"ProporcionMontoTotal"`
	IntegracionIngresos                    IntegracionIngresos `xml:"IntegracIngresos" bson:"IntegracionIngresos" json:"IntegracionIngresos"`
}

type IntegracionIngresos struct {
	Concepto string `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
}

type DeduccOSalidas struct {
	MontoTotalEgresosPeriodo     float64            `xml:"MontTotEgresPeriodo,attr" bson:"MontoTotalEgresosPeriodo" json:"MontoTotalEgresosPeriodo"`
	ParteProporcionalFideicomiso float64            `xml:"PartPropDelFideicom,attr" bson:"ParteProporcionalFideicomiso" json:"ParteProporcionalFideicomiso"`
	ProporcionMontoTotal         float64            `xml:"PropDelMontTot,attr" bson:"ProporcionMontoTotal" json:"ProporcionMontoTotal"`
	IntegracionEgresos           IntegracionEgresos `xml:"IntegracEgresos" bson:"IntegracionEgresos" json:"IntegracionEgresos"`
}

type IntegracionEgresos struct {
	Concepto string `xml:"ConceptoS,attr" bson:"Concepto" json:"Concepto"`
}

type RetencionesEfectuadasFideicomiso struct {
	MontoRetencionRelacionadaFideicomiso       float64 `xml:"MontRetRelPagFideic,attr" bson:"MontoRetencionRelacionadaFideicomiso" json:"MontoRetencionRelacionadaFideicomiso"`
	DescripcionRetencionRelacionadaFideicomiso string  `xml:"DescRetRelPagFideic,attr" bson:"DescripcionRetencionRelacionadaFideicomiso" json:"DescripcionRetencionRelacionadaFideicomiso"`
}

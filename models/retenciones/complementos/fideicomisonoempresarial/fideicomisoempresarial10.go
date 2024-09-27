package fideicomisonoempresarial

type FideicomisoEmpresarial10 struct {
	Version                          string                           `xml:"Version,attr" bson:"Version"`
	IngresosOEntradas                IngresosOEntradas                `xml:"IngresosOEntradas" bson:"IngresosOEntradas"`
	DeduccionesOSalidas              DeduccOSalidas                   `xml:"DeduccOSalidas" bson:"DeduccionesOSalidas"`
	RetencionesEfectuadasFideicomiso RetencionesEfectuadasFideicomiso `xml:"RetEfectFideicomiso" bson:"RetencionesEfectuadasFideicomiso"`
}

type IngresosOEntradas struct {
	MontoTotalEntradasPeriodo              float64             `xml:"MontTotEntradasPeriodo,attr" bson:"MontoTotalEntradasPeriodo"`
	ParteProporcionalAcumulableFideicomiso float64             `xml:"PartPropAcumDelFideicom,attr" bson:"ParteProporcionalAcumulableFideicomiso"`
	ProporcionMontoTotal                   float64             `xml:"PropDelMontTot,attr" bson:"ProporcionMontoTotal"`
	IntegracionIngresos                    IntegracionIngresos `xml:"IntegracIngresos" bson:"IntegracionIngresos"`
}

type IntegracionIngresos struct {
	Concepto string `xml:"Concepto,attr" bson:"Concepto"`
}

type DeduccOSalidas struct {
	MontoTotalEgresosPeriodo     float64            `xml:"MontTotEgresPeriodo,attr" bson:"MontoTotalEgresosPeriodo"`
	ParteProporcionalFideicomiso float64            `xml:"PartPropDelFideicom,attr" bson:"ParteProporcionalFideicomiso"`
	ProporcionMontoTotal         float64            `xml:"PropDelMontTot,attr" bson:"ProporcionMontoTotal"`
	IntegracionEgresos           IntegracionEgresos `xml:"IntegracEgresos" bson:"IntegracionEgresos"`
}

type IntegracionEgresos struct {
	Concepto string `xml:"ConceptoS,attr" bson:"Concepto"`
}

type RetencionesEfectuadasFideicomiso struct {
	MontoRetencionRelacionadaFideicomiso       float64 `xml:"MontRetRelPagFideic,attr" bson:"MontoRetencionRelacionadaFideicomiso"`
	DescripcionRetencionRelacionadaFideicomiso string  `xml:"DescRetRelPagFideic,attr" bson:"DescripcionRetencionRelacionadaFideicomiso"`
}

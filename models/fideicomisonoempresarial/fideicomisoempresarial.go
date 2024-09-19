package fideicomisoempresarial

type FideicomisoEmpresarial10 struct {
	Version           string            `xml:"Version,attr" bson:"Version"`
	IngresosOEntradas IngresosOEntradas `xml:"IngresosOEntradas" bson:"IngresosOEntradas"`
	DeduccOSalidas    DeduccOSalidas    `xml:"DeduccOSalidas" bson:"DeduccOSalidas"`
	Retenciones       Retenciones       `xml:"RetEfectFideicomiso" bson:"Retenciones"`
}

type IngresosOEntradas struct {
	MontTotEntradasPeriodo  float64          `xml:"MontTotEntradasPeriodo,attr" bson:"MontTotEntradasPeriodo"`
	PartPropAcumDelFideicom float64          `xml:"PartPropAcumDelFideicom,attr" bson:"PartPropAcumDelFideicom"`
	PropDelMontTot          float64          `xml:"PropDelMontTot,attr" bson:"PropDelMontTot"`
	IntegracIngresos        IntegracIngresos `xml:"IntegracIngresos" bson:"IntegracIngresos"`
}

type IntegracIngresos struct {
	Concepto string `xml:"Concepto,attr" bson:"Concepto"`
}

type DeduccOSalidas struct {
	MontTotEgresPeriodo float64 `xml:"MontTotEgresPeriodo,attr" bson:"MontTotEgresPeriodo"`
	PartPropDelFideicom float64 `xml:"PartPropDelFideicom,attr" bson:"PartPropDelFideicom"`
	PropDelMontTot      float64 `xml:"PropDelMontTot,attr" bson:"PropDelMontTot"`
}

type Retenciones struct {
	MontRetRelPagFideic float64 `xml:"MontRetRelPagFideic,attr" bson:"MontRetRelPagFideic"`
	DescRetRelPagFideic string  `xml:"DescRetRelPagFideic,attr" bson:"DescRetRelPagFideic"`
}

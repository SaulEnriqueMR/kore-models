package arriendamientosfideicomisos

type ArriendamientosFideicomisos10 struct {
	Version           string  `xml:"Version,attr" bson:"Version"`
	Pago              float64 `xml:"PagProvEfecPorFiduc,attr" bson:"Pago"`
	Rendimiento       float64 `xml:"RendimFideicom,attr" bson:"Rendimiento"`
	Deduccion         float64 `xml:"DeduccCorresp,attr" bson:"Deduccion"`
	MontoTotal        float64 `xml:"MontTotRet,attr" bson:"MontoTotal,omitempty"`
	MontoResultado    float64 `xml:"MontResFiscDistFibras,attr" bson:"MontoResultado,omitempty"`
	MontoOtros        float64 `xml:"MontOtrosConceptDistr,attr" bson:"MontoOtros,omitempty"`
	DescripcionMontos string  `xml:"DescrMontOtrosConceptDistr,attr" bson:"DescripcionMontos,omitempty"`
}

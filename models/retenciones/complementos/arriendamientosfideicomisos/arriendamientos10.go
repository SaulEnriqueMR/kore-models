package arriendamientosfideicomisos

type ArriendamientosFideicomisos10 struct {
	Version                                  string  `xml:"Version,attr" bson:"Version" json:"Version"`
	PagoEfectuadoPorFiduciario               float64 `xml:"PagProvEfecPorFiduc,attr" bson:"PagoEfectuadoPorFiduciario" json:"PagoEfectuadoPorFiduciario"`
	RendimientoFideicomiso                   float64 `xml:"RendimFideicom,attr" bson:"RendimientoFideicomiso" json:"RendimientoFideicomiso"`
	DeduccionCorrespondiente                 float64 `xml:"DeduccCorresp,attr" bson:"DeduccionCorrespondiente" json:"DeduccionCorrespondiente"`
	MontoTotalRetenido                       float64 `xml:"MontTotRet,attr" bson:"MontoTotalRetenido,omitempty" json:"MontoTotalRetenido,omitempty"`
	MontoResultadoFiscalDistribuidoPorFibras float64 `xml:"MontResFiscDistFibras,attr" bson:"MontoResultadoFiscalDistribuidoPorFibras,omitempty" json:"MontoResultadoFiscalDistribuidoPorFibras,omitempty"`
	MontoOtrosConceptosDistribuidos          float64 `xml:"MontOtrosConceptDistr,attr" bson:"MontoOtrosConceptosDistribuidos,omitempty" json:"MontoOtrosConceptosDistribuidos,omitempty"`
	DescripcionMontos                        string  `xml:"DescrMontOtrosConceptDistr,attr" bson:"DescripcionMontos,omitempty" json:"DescripcionMontos,omitempty"`
}

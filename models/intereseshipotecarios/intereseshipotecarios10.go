package intereseshipotecarios

type InteresesHipotecarios10 struct {
	Version                    string   `xml:"Version,attr" bson:"Version"`
	CreditoDeInstFinanc        string   `xml:"CreditoDeInstFinanc,attr" bson:"CreditoDeInstFinanc"`
	SaldoInsoluto              float64  `xml:"SaldoInsoluto,attr" bson:"SaldoInsoluto"`
	PropDeducDelCredit         *float64 `xml:"PropDeducDelCredit,attr" bson:"PropDeducDelCredit,omitempty"`
	MontTotIntNominalesDev     *float64 `xml:"MontTotIntNominalesDev,attr" bson:"MontTotIntNominalesDev,omitempty"`
	MontTotIntNominalesDevYPag *float64 `xml:"MontTotIntNominalesDevYPag,attr" bson:"MontTotIntNominalesDevYPag,omitempty"`
	MontTotIntRealPagDeduc     *float64 `xml:"MontTotIntRealPagDeduc,attr" bson:"MontTotIntRealPagDeduc,omitempty"`
	NumContrato                *string  `xml:"NumContrato,attr" bson:"NumContrato,omitempty"`
}

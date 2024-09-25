package dividendos

type Dividendos10 struct {
	Version    string                  `xml:"Version,attr" bson:"Version"`
	DividOUtil *DividOUtilDividendos10 `xml:"DividOUtil" bson:"DividOUtil,omitempty"`
	Remanente  *RemanenteDividendos10  `xml:"Remanente" bson:"Remanente,omitempty"`
}

type DividOUtilDividendos10 struct {
	CveTipDivOUtil            string   `xml:"CveTipDivOUtil,attr" bson:"CveTipDivOUtil"`
	MontISRAcredRetMexico     float64  `xml:"MontISRAcredRetMexico,attr" bson:"MontISRAcredRetMexico"`
	MontISRAcredRetExtranjero float64  `xml:"MontISRAcredRetExtranjero,attr" bson:"MontISRAcredRetExtranjero"`
	MontRetExtDivExt          *float64 `xml:"MontRetExtDivExt,attr" bson:"MontRetExtDivExt,omitempty"`
	TipoSocDistrDiv           string   `xml:"TipoSocDistrDiv,attr" bson:"TipoSocDistrDiv"`
	MontISRAcredNal           *float64 `xml:"MontISRAcredNal,attr" bson:"MontISRAcredNal,omitempty"`
	MontDivAcumNal            *float64 `xml:"MontDivAcumNal,attr" bson:"MontDivAcumNal,omitempty"`
	MontDivAcumExt            *float64 `xml:"MontDivAcumExt,attr" bson:"MontDivAcumExt,omitempty"`
}

type RemanenteDividendos10 struct {
	ProporcionRem *float64 `xml:"ProporcionRem,attr" bson:"ProporcionRem,omitempty"`
}

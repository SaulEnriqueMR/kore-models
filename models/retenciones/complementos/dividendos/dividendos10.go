package dividendos

type Dividendos10 struct {
	Version            string                  `xml:"Version,attr" bson:"Version"`
	DividendoOUtilidad *DividOUtilDividendos10 `xml:"DividOUtil" bson:"DividendoOUtilidad,omitempty"`
	Remanente          *RemanenteDividendos10  `xml:"Remanente" bson:"Remanente,omitempty"`
}

type DividOUtilDividendos10 struct {
	ClaveTipoDivendoOUtilidad          string   `xml:"CveTipDivOUtil,attr" bson:"ClaveTipoDivendoOUtilidad"`
	MontoIsrRetenidoMexico             float64  `xml:"MontISRAcredRetMexico,attr" bson:"MontoIsrRetenidoMexico"`
	MontoIsrRetenidoExtranjero         float64  `xml:"MontISRAcredRetExtranjero,attr" bson:"MontoIsrRetenidoExtranjero"`
	MontoRetencionDividendoExtranjero  *float64 `xml:"MontRetExtDivExt,attr" bson:"MontoRetencionDividendoExtranjero,omitempty"`
	TipoSociedadDistribucionDividendo  string   `xml:"TipoSocDistrDiv,attr" bson:"TipoSociedadDistribucionDividendo"`
	MontoIsrAcreditableNacional        *float64 `xml:"MontISRAcredNal,attr" bson:"MontoIsrAcreditableNacional,omitempty"`
	MontoDividendoAcumulableNacional   *float64 `xml:"MontDivAcumNal,attr" bson:"MontoDividendoAcumulableNacional,omitempty"`
	MontoDividendoAcumulableExtranjero *float64 `xml:"MontDivAcumExt,attr" bson:"MontoDividendoAcumulableExtranjero,omitempty"`
}

type RemanenteDividendos10 struct {
	ProporcionRemanente *float64 `xml:"ProporcionRem,attr" bson:"ProporcionRemanente,omitempty"`
}

package pagosaextranjeros

type Pagosaextranjeros10 struct {
	Version              string                       `xml:"Version,attr" bson:"Version"`
	EsBenefEfectDelCobro string                       `xml:"EsBenefEfectDelCobro,attr" bson:"EsBenefEfectDelCobro"`
	NoBeneficiario       *NoBeneficiarioPagosExtran10 `xml:"NoBeneficiario" bson:"NoBeneficiario,omitempty"`
	Beneficiario         *BeneficiarioPagosExtran10   `xml:"Beneficiario" bson:"Beneficiario,omitempty"`
}

type NoBeneficiarioPagosExtran10 struct {
	PaisDeResidParaEfecFisc string `xml:"PaisDeResidParaEfecFisc,attr" bson:"PaisDeResidParaEfecFisc"`
	ConceptoPago            string `xml:"ConceptoPago,attr" bson:"ConceptoPago"`
	DescripcionConcepto     string `xml:"DescripcionConcepto,attr" bson:"DescripcionConcepto"`
}

type BeneficiarioPagosExtran10 struct {
	Rfc                 string `xml:"RFC,attr" bson:"Rfc"`
	Curp                string `xml:"CURP,attr" bson:"Curp"`
	NomDenRazSocB       string `xml:"NomDenRazSocB,attr" bson:"NomDenRazSocB"`
	ConceptoPago        string `xml:"ConceptoPago,attr" bson:"ConceptoPago"`
	DescripcionConcepto string `xml:"DescripcionConcepto,attr" bson:"DescripcionConcepto"`
}

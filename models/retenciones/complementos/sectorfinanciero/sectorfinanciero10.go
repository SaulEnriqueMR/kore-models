package sectorfinanciero

type SectorFinanciero10 struct {
	Version         string  `xml:"Version,attr" bson:"Version"`
	IdFideicom      string  `xml:"IdFideicom,attr" bson:"IdFideicom"`
	NomFideicom     *string `xml:"NomFideicom,attr" bson:"NomFideicom,omitempty"`
	DescripFideicom string  `xml:"DescripFideicom,attr" bson:"DescripFideicom,omitempty"`
}

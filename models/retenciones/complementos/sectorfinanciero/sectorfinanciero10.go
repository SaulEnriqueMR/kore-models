package sectorfinanciero

type SectorFinanciero10 struct {
	Version                string  `xml:"Version,attr" bson:"Version"`
	IdFideicomiso          string  `xml:"IdFideicom,attr" bson:"IdFideicomiso"`
	NombreFideicomiso      *string `xml:"NomFideicom,attr" bson:"NombreFideicomiso,omitempty"`
	DescripcionFideicomiso string  `xml:"DescripFideicom,attr" bson:"DescripcionFideicomiso,omitempty"`
}

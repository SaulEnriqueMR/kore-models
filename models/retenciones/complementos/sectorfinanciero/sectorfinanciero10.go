package sectorfinanciero

type SectorFinanciero10 struct {
	Version                string  `xml:"Version,attr" bson:"Version" json:"Version"`
	IdFideicomiso          string  `xml:"IdFideicom,attr" bson:"IdFideicomiso" json:"IdFideicomiso"`
	NombreFideicomiso      *string `xml:"NomFideicom,attr" bson:"NombreFideicomiso,omitempty" json:"NombreFideicomiso,omitempty"`
	DescripcionFideicomiso string  `xml:"DescripFideicom,attr" bson:"DescripcionFideicomiso,omitempty" json:"DescripcionFideicomiso,omitempty"`
}

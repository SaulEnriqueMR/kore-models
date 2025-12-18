package comprobante

type UniversalSearch struct {
	Uuid           string `json:"Uuid" bson:"Uuid"`
	CadenaOriginal string `json:"CadenaOriginal" bson:"CadenaOriginal"`
}

func (c *Comprobante30) ToUniversalSearch() *UniversalSearch {
	return &UniversalSearch{
		Uuid:           c.Uuid,
		CadenaOriginal: c.CadenaOriginal,
	}
}

func (c *Comprobante32) ToUniversalSearch() *UniversalSearch {
	return &UniversalSearch{
		Uuid:           c.Uuid,
		CadenaOriginal: c.CadenaOriginal,
	}
}

func (c *Comprobante33) ToUniversalSearch() *UniversalSearch {
	return &UniversalSearch{
		Uuid:           c.Uuid,
		CadenaOriginal: c.CadenaOriginal,
	}
}

func (c *Comprobante40) ToUniversalSearch() *UniversalSearch {
	return &UniversalSearch{
		Uuid:           c.Uuid,
		CadenaOriginal: c.CadenaOriginal,
	}
}

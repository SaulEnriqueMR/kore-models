package acuentaterceros

type ACuentaTerceros11 struct {
	Version                  string                              `xml:"version,attr" bson:"Version"`
	Rfc                      string                              `xml:"rfc,attr" bson:"Rfc"`
	Nombre                   *string                             `xml:"nombre,attr" bson:"Nombre,omitempty"`
	InformacionFiscalTercero *InformacionFiscalTerceroTerceros11 `xml:"InformacionFiscalTercero" bson:"InformacionFiscalTercero,omitempty"`
	InformacionAduanera      *InformacionAduaneraTerceros11      `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
	Parte                    *[]ParteTerceros11                  `xml:"Parte" bson:"Parte,omitempty"`
	CuentaPredial            *CuentaPredialTerceros11            `xml:"CuentaPredial" bson:"CuentaPredial,omitempty"`
	Impuestos                ImpuestosTerceros11                 `xml:"Impuestos" bson:"Impuestos"`
}

type InformacionFiscalTerceroTerceros11 struct {
	Calle        string  `xml:"calle,attr" bson:"Calle"`
	NoExterior   *string `xml:"noExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"noInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"colonia,attr" bson:"Colonia,omitempty"`
	Localidad    *string `xml:"localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"referencia,attr" bson:"Referencia,omitempty"`
	Municipio    string  `xml:"municipio,attr" bson:"Municipio"`
	Estado       string  `xml:"estado,attr" bson:"Estado"`
	Pais         string  `xml:"pais,attr" bson:"Pais"`
	CodigoPostal string  `xml:"codigoPostal,attr" bson:"CodigoPostal"`
}

type InformacionAduaneraTerceros11 struct {
	Numero string  `xml:"numero,attr" bson:"Numero"`
	Fecha  string  `xml:"fecha,attr" bson:"Fecha"`
	Aduana *string `xml:"aduana,attr" bson:"Aduana,omitempty"`
}

type ParteTerceros11 struct {
	Cantidad            float64                          `xml:"cantidad,attr" bson:"Cantidad"`
	Unidad              *string                          `xml:"unidad,attr" bson:"Unidad,omitempty"`
	NoIdentificacion    *string                          `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Descripcion         string                           `xml:"descripcion,attr" bson:"Descripcion"`
	ValorUnitario       *float64                         `xml:"valorUnitario,attr" bson:"ValorUnitario,omitempty"`
	Importe             *float64                         `xml:"importe,attr" bson:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduaneraTerceros11 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
}

type CuentaPredialTerceros11 struct {
	Numero string `xml:"numero,attr" bson:"Numero"`
}

type ImpuestosTerceros11 struct {
	Retenciones *[]RetencionTerceros11 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty"`
	Traslados   *[]TrasladoTerceros11  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
}

type RetencionTerceros11 struct {
	Impuesto string  `xml:"impuesto,attr" bson:"Impuesto"`
	Importe  float64 `xml:"importe,attr" bson:"Importe"`
}

type TrasladoTerceros11 struct {
	Impuesto string  `xml:"impuesto,attr" bson:"Impuesto"`
	Tasa     float64 `xml:"tasa,attr" bson:"Tasa"`
	Importe  float64 `xml:"importe,attr" bson:"Importe"`
}

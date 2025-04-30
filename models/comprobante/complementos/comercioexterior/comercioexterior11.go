package comercioexterior

type ComercioExterior11 struct {
	Version               string                            `xml:"Version,attr" bson:"Version" json:"Version"`
	MotivoTraslado        *string                           `xml:"MotivoTraslado,attr" bson:"MotivoTraslado,omitempty" json:"MotivoTraslado,omitempty"`
	TipoOperacion         string                            `xml:"TipoOperacion,attr" bson:"TipoOperacion" json:"TipoOperacion"`
	ClavePedimento        *string                           `xml:"ClaveDePedimento,attr" bson:"ClavePedimento,omitempty" json:"ClavePedimento,omitempty"`
	CertificadoOrigen     *float64                          `xml:"CertificadoOrigen,attr" bson:"CertificadoOrigen,omitempty" json:"CertificadoOrigen,omitempty"`
	NoCertificadoOrigen   *string                           `xml:"NumCertificadoOrigen,attr" bson:"NoCertificadoOrigen,omitempty" json:"NoCertificadoOrigen,omitempty"`
	NoExportadorConfiable *string                           `xml:"NumeroExportadorConfiable,attr" bson:"NoExportadorConfiable,omitempty" json:"NoExportadorConfiable,omitempty"`
	Incoterm              *string                           `xml:"Incoterm,attr" bson:"Incoterm,omitempty" json:"Incoterm,omitempty"`
	Subdivision           *float64                          `xml:"Subdivision,attr" bson:"Subdivision,omitempty" json:"Subdivision,omitempty"`
	Observaciones         *string                           `xml:"Observaciones,attr" bson:"Observaciones,omitempty" json:"Observaciones,omitempty"`
	TipoCambioUsd         *string                           `xml:"TipoCambioUSD,attr" bson:"TipoCambioUsd,omitempty" json:"TipoCambioUsd,omitempty"`
	TotalUsd              *float64                          `xml:"TotalUSD,attr" bson:"TotalUsd,omitempty" json:"TotalUsd,omitempty"`
	Emisor                *EmisorComercioExterior11         `xml:"Emisor" bson:"Emisor,omitempty" json:"Emisor,omitempty"`
	Propietario           *[]PropietarioComercioExterior11  `xml:"Propietario" bson:"Propietario,omitempty" json:"Propietario,omitempty"`
	Receptor              *ReceptorComercioExterior11       `xml:"Receptor" bson:"Receptor,omitempty" json:"Receptor,omitempty"`
	Destinatario          *[]DestinatarioComercioExterior11 `xml:"Destinatario" bson:"Destinatario,omitempty" json:"Destinatario,omitempty"`
	Mercancias            *[]MercanciaComercioExterior11    `xml:"Mercancias>Mercancia" bson:"Mercancias,omitempty" json:"Mercancias,omitempty"`
}

type EmisorComercioExterior11 struct {
	Curp      *string                      `xml:"Curp,attr" bson:"Curp,omitempty" json:"Curp,omitempty"`
	Domicilio *DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type DomicilioComercioExterior11 struct {
	Calle        string  `xml:"Calle,attr" bson:"Calle" json:"Calle"`
	NoExterior   *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitemtpy" json:"Colonia,omitemtpy"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`
	Municipio    *string `xml:"Municipio,attr" bson:"Municipio,omitemtpy" json:"Municipio,omitemtpy"`
	Estado       string  `xml:"Estado,attr" bson:"Estado" json:"Estado"`
	Pais         string  `xml:"Pais,attr" bson:"Pais" json:"Pais"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal" json:"CodigoPostal"`
}

type PropietarioComercioExterior11 struct {
	NumRegIdTrib     string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib" json:"NumRegIdTrib"`
	ResidenciaFiscal string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal" json:"ResidenciaFiscal"`
}

type ReceptorComercioExterior11 struct {
	NumRegIdTrib *string                      `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	Domicilio    *DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type DestinatarioComercioExterior11 struct {
	NumRegIdTrib *string                       `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	Nombre       *string                       `xml:"Nombre,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	Domicilio    []DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio" json:"Domicilio"`
}

type MercanciaComercioExterior11 struct {
	NoIdentificacion         string                                        `xml:"NoIdentificacion,attr" bson:"NoIdentificacion" json:"NoIdentificacion"`
	FraccionArancelaria      *string                                       `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty" json:"FraccionArancelaria,omitempty"`
	CantidadAduana           *float64                                      `xml:"CantidadAduana,attr" bson:"CantidadAduana,omitempty" json:"CantidadAduana,omitempty"`
	UnidadAduana             *string                                       `xml:"UnidadAduana,attr" bson:"UnidadAduana,omitempty" json:"UnidadAduana,omitempty"`
	ValorUnitarioAduana      *float64                                      `xml:"ValorUnitarioAduana,attr" bson:"ValorUnitarioAduana,omitempty" json:"ValorUnitarioAduana,omitempty"`
	ValorDolares             float64                                       `xml:"ValorDolares,attr" bson:"ValorDolares" json:"ValorDolares"`
	DescripcionesEspecificas *[]DescripcionesEspecificasComercioExterior11 `xml:"DescripcionesEspecificas" bson:"DescripcionesEspecificas,omitempty" json:"DescripcionesEspecificas,omitempty"`
}

type DescripcionesEspecificasComercioExterior11 struct {
	Marca     string  `xml:"Marca,attr" bson:"Marca" json:"Marca"`
	Modelo    *string `xml:"Modelo,attr" bson:"Modelo,omitempty" json:"Modelo,omitempty"`
	Submodelo *string `xml:"SubModelo,attr" bson:"Submodelo,omitempty" json:"Submodelo,omitempty"`
	NoSerie   *string `xml:"NumeroSerie,attr" bson:"NoSerie,omitempty" json:"NoSerie,omitempty"`
}

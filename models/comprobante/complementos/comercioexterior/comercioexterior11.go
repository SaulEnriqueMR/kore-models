package comercioexterior

type ComercioExterior11 struct {
	Version               string                            `xml:"Version,attr" bson:"Version"`
	MotivoTraslado        *string                           `xml:"MotivoTraslado,attr" bson:"MotivoTraslado,omitempty"`
	TipoOperacion         string                            `xml:"TipoOperacion,attr" bson:"TipoOperacion"`
	ClavePedimento        *string                           `xml:"ClaveDePedimento,attr" bson:"ClavePedimento,omitempty"`
	CertificadoOrigen     *float64                          `xml:"CertificadoOrigen,attr" bson:"CertificadoOrigen,omitempty"`
	NoCertificadoOrigen   *string                           `xml:"NumCertificadoOrigen,attr" bson:"NoCertificadoOrigen,omitempty"`
	NoExportadorConfiable *string                           `xml:"NumeroExportadorConfiable,attr" bson:"NoExportadorConfiable,omitempty"`
	Incoterm              *string                           `xml:"Incoterm,attr" bson:"Incoterm,omitempty"`
	Subdivision           *float64                          `xml:"Subdivision,attr" bson:"Subdivision,omitempty"`
	Observaciones         *string                           `xml:"Observaciones,attr" bson:"Observaciones,omitempty"`
	TipoCambioUsd         *string                           `xml:"TipoCambioUSD,attr" bson:"TipoCambioUsd,omitempty"`
	TotalUsd              *float64                          `xml:"TotalUSD,attr" bson:"TotalUsd,omitempty"`
	Emisor                *EmisorComercioExterior11         `xml:"Emisor" bson:"Emisor,omitempty"`
	Propietario           *[]PropietarioComercioExterior11  `xml:"Propietario" bson:"Propietario,omitempty"`
	Receptor              *ReceptorComercioExterior11       `xml:"Receptor" bson:"Receptor,omitempty"`
	Destinatario          *[]DestinatarioComercioExterior11 `xml:"Destinatario" bson:"Destinatario,omitempty"`
	Mercancias            *[]MercanciaComercioExterior11    `xml:"Mercancias>Mercancia" bson:"Mercancias,omitempty"`
}

type EmisorComercioExterior11 struct {
	Curp      *string                      `xml:"Curp,attr" bson:"Curp,omitempty"`
	Domicilio *DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DomicilioComercioExterior11 struct {
	Calle        string  `xml:"Calle,attr" bson:"Calle"`
	NoExterior   *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitemtpy"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`
	Municipio    *string `xml:"Municipio,attr" bson:"Municipio,omitemtpy"`
	Estado       string  `xml:"Estado,attr" bson:"Estado"`
	Pais         string  `xml:"Pais,attr" bson:"Pais"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal"`
}

type PropietarioComercioExterior11 struct {
	NumRegIdTrib     string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib"`
	ResidenciaFiscal string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal"`
}

type ReceptorComercioExterior11 struct {
	NumRegIdTrib *string                      `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	Domicilio    *DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DestinatarioComercioExterior11 struct {
	NumRegIdTrib *string                       `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	Nombre       *string                       `xml:"Nombre,attr" bson:"Nombre,omitempty"`
	Domicilio    []DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio"`
}

type MercanciaComercioExterior11 struct {
	NoIdentificacion         string                                        `xml:"NoIdentificacion,attr" bson:"NoIdentificacion"`
	FraccionArancelaria      *string                                       `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	CantidadAduana           *float64                                      `xml:"CantidadAduana,attr" bson:"CantidadAduana,omitempty"`
	UnidadAduana             *string                                       `xml:"UnidadAduana,attr" bson:"UnidadAduana,omitempty"`
	ValorUnitarioAduana      *float64                                      `xml:"ValorUnitarioAduana,attr" bson:"ValorUnitarioAduana,omitempty"`
	ValorDolares             float64                                       `xml:"ValorDolares,attr" bson:"ValorDolares"`
	DescripcionesEspecificas *[]DescripcionesEspecificasComercioExterior11 `xml:"DescripcionesEspecificas" bson:"DescripcionesEspecificas,omitempty"`
}

type DescripcionesEspecificasComercioExterior11 struct {
	Marca     string  `xml:"Marca,attr" bson:"Marca"`
	Modelo    *string `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	Submodelo *string `xml:"SubModelo,attr" bson:"Submodelo,omitempty"`
	NoSerie   *string `xml:"NumeroSerie,attr" bson:"NoSerie,omitempty"`
}

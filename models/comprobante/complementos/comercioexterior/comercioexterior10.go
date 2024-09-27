package comercioexterior

type ComercioExterior10 struct {
	Version               string                          `xml:"Version,attr" bson:"Version"`
	TipoOperacion         string                          `xml:"TipoOperacion,attr" bson:"TipoOperacion"`
	ClavePedimento        *string                         `xml:"ClaveDePedimento,attr" bson:"ClavePedimento,omitempty"`
	CertificadoOrigen     *float64                        `xml:"CertificadoOrigen,attr" bson:"CertificadoOrigen,omitempty"`
	NoCertificadoOrigen   *string                         `xml:"NumCertificadoOrigen,attr" bson:"NoCertificadoOrigen,omitempty"`
	NoExportadorConfiable *string                         `xml:"NumeroExportadorConfiable,attr" bson:"NoExportadorConfiable,omitempty"`
	Incoterm              *string                         `xml:"Incoterm,attr" bson:"Incoterm,omitempty"`
	Subdivision           *float64                        `xml:"Subdivision,attr" bson:"Subdivision,omitempty"`
	Observaciones         *string                         `xml:"Observaciones,attr" bson:"Observaciones,omitempty"`
	TipoCambioUsd         *float64                        `xml:"TipoCambioUSD,attr" bson:"TipoCambioUsd,omitempty"`
	TotalUsd              *float64                        `xml:"TotalUSD,attr" bson:"TotalUsd,omitempty"`
	Emisor                *EmisorComercioExterior10       `xml:"Emisor" bson:"Emisor,omitempty"`
	Receptor              ReceptorComercioExterior10      `xml:"Receptor" bson:"Receptor"`
	Destinatario          *DestinatarioComercioExterior10 `xml:"Destinatario" bson:"Destinatario,omitempty"`
	Mercancias            *[]MercanciaComercioExterior10  `xml:"Mercancias>Mercancia" bson:"Mercancias,omitempty"`
}

type EmisorComercioExterior10 struct {
	Curp *string `xml:"Curp,attr" bson:"Curp,omitempty"` // Cifrado
}

type ReceptorComercioExterior10 struct {
	Curp         *string `xml:"Curp,attr" bson:"Curp,omitempty"`
	NumRegIdTrib string  `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib"`
}

type DestinatarioComercioExterior10 struct {
	NumRegIdTrib *string                     `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	Rfc          *string                     `xml:"Rfc,attr" bson:"Rfc,omitempty"`
	Curp         *string                     `xml:"Curp,attr" bson:"Curp,omitempty"`
	Nombre       *string                     `xml:"Nombre,attr" bson:"Nombre,omitempty"`
	Domicilio    DomicilioComercioExterior10 `xml:"Domicilio" bson:"Domicilio"`
}

type DomicilioComercioExterior10 struct {
	Calle          string  `xml:"Calle,attr" bson:"Calle"`
	NumeroExterior *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty"`
	NumeroInterior *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty"`
	Colonia        *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`
	Localidad      *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`
	Referencia     *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`
	Municipio      *string `xml:"Municipio,attr" bson:"Municipio,omitempty"`
	Estado         string  `xml:"Estado,attr" bson:"Estado"`
	Pais           string  `xml:"Pais,attr" bson:"Pais"`
	CodigoPostal   string  `xml:"CodigoPostal,attr" bson:"CodigoPostal"`
}

type MercanciaComercioExterior10 struct {
	NoIdentificacion         string                                        `xml:"NoIdentificacion,attr" bson:"NoIdentificacion"`
	FraccionArancelaria      *string                                       `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	CantidadAduana           *float64                                      `xml:"CantidadAduana,attr" bson:"CantidadAduana,omitempty"`
	UnidadAduana             *string                                       `xml:"UnidadAduana,attr" bson:"UnidadAduana,omitempty"`
	ValorUnitarioAduana      *float64                                      `xml:"ValorUnitarioAduana,attr" bson:"ValorUnitarioAduana,omitempty"`
	ValorDolares             float64                                       `xml:"ValorDolares,attr" bson:"ValorDolares"`
	DescripcionesEspecificas *[]DescripcionesEspecificasComercioExterior10 `xml:"DescripcionesEspecificas" bson:"DescripcionesEspecificas,omitempty"`
}

type DescripcionesEspecificasComercioExterior10 struct {
	Marca     string  `xml:"Marca,attr" bson:"Marca"`
	Modelo    *string `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	Submodelo *string `xml:"SubModelo,attr" bson:"Submodelo,omitempty"`
	NoSerie   *string `xml:"NumeroSerie,attr" bson:"NoSerie,omitempty"`
}

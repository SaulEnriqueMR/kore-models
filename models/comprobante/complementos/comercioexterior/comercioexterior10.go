package comercioexterior

type ComercioExterior10 struct {
	Version               string                          `xml:"Version,attr" bson:"Version" json:"Version"`
	TipoOperacion         string                          `xml:"TipoOperacion,attr" bson:"TipoOperacion" json:"TipoOperacion"`
	ClavePedimento        *string                         `xml:"ClaveDePedimento,attr" bson:"ClavePedimento,omitempty" json:"ClavePedimento,omitempty"`
	CertificadoOrigen     *float64                        `xml:"CertificadoOrigen,attr" bson:"CertificadoOrigen,omitempty" json:"CertificadoOrigen,omitempty"`
	NoCertificadoOrigen   *string                         `xml:"NumCertificadoOrigen,attr" bson:"NoCertificadoOrigen,omitempty" json:"NoCertificadoOrigen,omitempty"`
	NoExportadorConfiable *string                         `xml:"NumeroExportadorConfiable,attr" bson:"NoExportadorConfiable,omitempty" json:"NoExportadorConfiable,omitempty"`
	Incoterm              *string                         `xml:"Incoterm,attr" bson:"Incoterm,omitempty" json:"Incoterm,omitempty"`
	Subdivision           *float64                        `xml:"Subdivision,attr" bson:"Subdivision,omitempty" json:"Subdivision,omitempty"`
	Observaciones         *string                         `xml:"Observaciones,attr" bson:"Observaciones,omitempty" json:"Observaciones,omitempty"`
	TipoCambioUsd         *float64                        `xml:"TipoCambioUSD,attr" bson:"TipoCambioUsd,omitempty" json:"TipoCambioUsd,omitempty"`
	TotalUsd              *float64                        `xml:"TotalUSD,attr" bson:"TotalUsd,omitempty" json:"TotalUsd,omitempty"`
	Emisor                *EmisorComercioExterior10       `xml:"Emisor" bson:"Emisor,omitempty" json:"Emisor,omitempty"`
	Receptor              ReceptorComercioExterior10      `xml:"Receptor" bson:"Receptor" json:"Receptor"`
	Destinatario          *DestinatarioComercioExterior10 `xml:"Destinatario" bson:"Destinatario,omitempty" json:"Destinatario,omitempty"`
	Mercancias            *[]MercanciaComercioExterior10  `xml:"Mercancias>Mercancia" bson:"Mercancias,omitempty" json:"Mercancias,omitempty"`
}

type EmisorComercioExterior10 struct {
	Curp *string `xml:"Curp,attr" bson:"Curp,omitempty" json:"Curp,omitempty"` // Cifrado
}

type ReceptorComercioExterior10 struct {
	Curp         *string `xml:"Curp,attr" bson:"Curp,omitempty" json:"Curp,omitempty"`
	NumRegIdTrib string  `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib" json:"NumRegIdTrib"`
}

type DestinatarioComercioExterior10 struct {
	NumRegIdTrib *string                     `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	Rfc          *string                     `xml:"Rfc,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
	Curp         *string                     `xml:"Curp,attr" bson:"Curp,omitempty" json:"Curp,omitempty"`
	Nombre       *string                     `xml:"Nombre,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	Domicilio    DomicilioComercioExterior10 `xml:"Domicilio" bson:"Domicilio" json:"Domicilio"`
}

type DomicilioComercioExterior10 struct {
	Calle          string  `xml:"Calle,attr" bson:"Calle" json:"Calle"`
	NumeroExterior *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NumeroInterior *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia        *string `xml:"Colonia,attr" bson:"Colonia,omitempty" json:"Colonia,omitempty"`
	Localidad      *string `xml:"Localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`
	Referencia     *string `xml:"Referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`
	Municipio      *string `xml:"Municipio,attr" bson:"Municipio,omitempty" json:"Municipio,omitempty"`
	Estado         string  `xml:"Estado,attr" bson:"Estado" json:"Estado"`
	Pais           string  `xml:"Pais,attr" bson:"Pais" json:"Pais"`
	CodigoPostal   string  `xml:"CodigoPostal,attr" bson:"CodigoPostal" json:"CodigoPostal"`
}

type MercanciaComercioExterior10 struct {
	NoIdentificacion         string                                        `xml:"NoIdentificacion,attr" bson:"NoIdentificacion" json:"NoIdentificacion"`
	FraccionArancelaria      *string                                       `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty" json:"FraccionArancelaria,omitempty"`
	CantidadAduana           *float64                                      `xml:"CantidadAduana,attr" bson:"CantidadAduana,omitempty" json:"CantidadAduana,omitempty"`
	UnidadAduana             *string                                       `xml:"UnidadAduana,attr" bson:"UnidadAduana,omitempty" json:"UnidadAduana,omitempty"`
	ValorUnitarioAduana      *float64                                      `xml:"ValorUnitarioAduana,attr" bson:"ValorUnitarioAduana,omitempty" json:"ValorUnitarioAduana,omitempty"`
	ValorDolares             float64                                       `xml:"ValorDolares,attr" bson:"ValorDolares" json:"ValorDolares"`
	DescripcionesEspecificas *[]DescripcionesEspecificasComercioExterior10 `xml:"DescripcionesEspecificas" bson:"DescripcionesEspecificas,omitempty" json:"DescripcionesEspecificas,omitempty"`
}

type DescripcionesEspecificasComercioExterior10 struct {
	Marca     string  `xml:"Marca,attr" bson:"Marca" json:"Marca"`
	Modelo    *string `xml:"Modelo,attr" bson:"Modelo,omitempty" json:"Modelo,omitempty"`
	Submodelo *string `xml:"SubModelo,attr" bson:"Submodelo,omitempty" json:"Submodelo,omitempty"`
	NoSerie   *string `xml:"NumeroSerie,attr" bson:"NoSerie,omitempty" json:"NoSerie,omitempty"`
}

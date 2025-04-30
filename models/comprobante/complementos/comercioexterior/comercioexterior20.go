package comercioexterior

type ComercioExterior20 struct {
	Version               string                            `xml:"Version,attr" bson:"Version" json:"Version"`
	MotivoTraslado        *string                           `xml:"MotivoTraslado,attr" bson:"MotivoTraslado,omitempty" json:"MotivoTraslado,omitempty"`
	ClavePedimento        string                            `xml:"ClaveDePedimento,attr" bson:"ClavePedimento" json:"ClavePedimento"`
	CertificadoOrigen     float64                           `xml:"CertificadoOrigen,attr" bson:"CertificadoOrigen" json:"CertificadoOrigen"`
	NoCertificadoOrigen   *string                           `xml:"NumCertificadoOrigen,attr" bson:"NoCertificadoOrigen,omitempty" json:"NoCertificadoOrigen,omitempty"`
	NoExportadorConfiable *string                           `xml:"NumeroExportadorConfiable,attr" bson:"NoExportadorConfiable,omitempty" json:"NoExportadorConfiable,omitempty"`
	Incoterm              *string                           `xml:"Incoterm,attr" bson:"Incoterm,omitempty" json:"Incoterm,omitempty"`
	Observaciones         *string                           `xml:"Observaciones,attr" bson:"Obsrvaciones,omitempty" json:"Obsrvaciones,omitempty"`
	TipoCambioUsd         float64                           `xml:"TipoCambioUSD,attr" bson:"TipoCambioUsd" json:"TipoCambioUsd"`
	TotalUsd              float64                           `xml:"TotalUSD,attr" bson:"TotalUsd" json:"TotalUsd"`
	Emisor                *EmisorComercioExterior20         `xml:"Emisor" bson:"Emisor,omitempty" json:"Emisor,omitempty"`
	Propietario           *[]PropietarioComercioExterior20  `xml:"Propietario" bson:"Propietario,omitempty" json:"Propietario,omitempty"`
	Receptor              *ReceptorComercioExterior20       `xml:"Receptor" bson:"Receptor,omitempty" json:"Receptor,omitempty"`
	Destinatario          *[]DestinatarioComercioExterior20 `xml:"Destinatario" bson:"Destinatario,omitempty" json:"Destinatario,omitempty"`
	Mercancias            []MercanciaComercioExterior20     `xml:"Mercancias>Mercancia" bson:"Mercancias" json:"Mercancias"`
}

type EmisorComercioExterior20 struct {
	Curp      *string                     `xml:"Curp,attr" bson:"Curp,omitempty" json:"Curp,omitempty"` // Cifrado
	Domicilio DomicilioComercioExterior20 `xml:"Domicilio" bson:"Domicilio" json:"Domicilio"`
}

type PropietarioComercioExterior20 struct {
	NumRegIdTrib     string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib" json:"NumRegIdTrib"`         // Cifrado
	ResidenciaFiscal string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal" json:"ResidenciaFiscal"` // Cifrado
}

type ReceptorComercioExterior20 struct {
	NumRegIdTrib *string                      `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"` // Cifrado
	Domicilio    *DomicilioComercioExterior20 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type DestinatarioComercioExterior20 struct {
	NumRegIdTrib *string                       `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"` // Cifrado
	Nombre       *string                       `xml:"Nombre,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`             // Cifrado
	Domicilio    []DomicilioComercioExterior20 `xml:"Domicilio" bson:"Domicilio" json:"Domicilio"`
}

type MercanciaComercioExterior20 struct {
	NoIdentificacion         string                                        `xml:"NoIdentificacion,attr" bson:"NoIdentificacion" json:"NoIdentificacion"`
	FraccionArancelaria      *string                                       `xml:"FraccionArancelaria,attr" bson:"FraccionArancelario,omitempty" json:"FraccionArancelario,omitempty"`
	CantidadAduana           *float64                                      `xml:"CantidadAduana,attr" bson:"CantidadAduana,omitempty" json:"CantidadAduana,omitempty"`
	UnidadAduana             *string                                       `xml:"UnidadAduana,attr" bson:"UnidadAduana,omitempty" json:"UnidadAduana,omitempty"`
	ValorUnitarioAduana      *float64                                      `xml:"ValorUnitarioAduana,attr" bson:"ValorUnitarioAduana,omitempty" json:"ValorUnitarioAduana,omitempty"`
	ValorDolares             float64                                       `xml:"ValorDolares,attr" bson:"ValorDolares" json:"ValorDolares"`
	DescripcionesEspecificas *[]DescripcionesEspecificasComercioExterior20 `xml:"DescripcionesEspecificas" bson:"DescripcionesEspecificas,omitempty" json:"DescripcionesEspecificas,omitempty"`
}

type DomicilioComercioExterior20 struct {
	Calle        string  `xml:"Calle,attr" bson:"Calle" json:"Calle"`
	NoExterior   *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"` // Cifrado
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty" json:"Colonia,omitempty"`           // Cifrado
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`       // Cifrado
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`     // Cifrado
	Municipio    *string `xml:"Municipio,attr" bson:"Municipio,omitempty" json:"Municipio,omitempty"`       // Cifrado
	Estado       string  `xml:"Estado,attr" bson:"Estado" json:"Estado"`                       // Cifrado
	Pais         string  `xml:"Pais,attr" bson:"Pais" json:"Pais"`                           // Cifrado
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal" json:"CodigoPostal"`           // Cifrado
}

type DescripcionesEspecificasComercioExterior20 struct {
	Marca     string  `xml:"Marca,attr" bson:"Marca" json:"Marca"`
	Modelo    *string `xml:"Modelo,attr" bson:"Modelo,omitempty" json:"Modelo,omitempty"`
	Submodelo *string `xml:"SubModelo,attr" bson:"Submodelo,omitempty" json:"Submodelo,omitempty"`
	NoSerie   *string `xml:"NumeroSerie,attr" bson:"NoSerie,omitempty" json:"NoSerie,omitempty"`
}

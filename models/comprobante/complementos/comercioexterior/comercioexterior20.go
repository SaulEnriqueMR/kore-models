package comercioexterior

type ComercioExterior20 struct {
	Version                   string                            `xml:"Version,attr" bson:"Version"`
	MotivoTraslado            *string                           `xml:"MotivoTraslado,attr" bson:"MotivoTraslado,omitempty"`
	ClavePedimento            string                            `xml:"ClaveDePedimento,attr" bson:"ClavePedimento"`
	CertificadoOrigen         float64                           `xml:"CertificadoOrigen,attr" bson:"CertificadoOrigen"`
	NumCertificadoOrigen      *string                           `xml:"NumCertificadoOrigen,attr" bson:"NumCertificadoOrigen,omitempty"`
	NumeroExportadorConfiable *string                           `xml:"NumeroExportadorConfiable,attr" bson:"NumeroExportadorConfiable,omitempty"`
	Incoterm                  *string                           `xml:"Incoterm,attr" bson:"Incoterm,omitempty"`
	Observaciones             *string                           `xml:"Observaciones,attr" bson:"Obsrvaciones,omitempty"`
	TipoCambioUSD             float64                           `xml:"TipoCambioUSD,attr" bson:"TipoCambioUSD"`
	TotalUSD                  float64                           `xml:"TotalUSD,attr" bson:"TotalUSD"`
	Emisor                    *EmisorComercioExterior20         `xml:"Emisor" bson:"Emisor,omitempty"`
	Propietario               *[]PropietarioComercioExterior20  `xml:"Propietario" bson:"Propietario,omitempty"`
	Receptor                  *ReceptorComercioExterior20       `xml:"Receptor" bson:"Receptor,omitempty"`
	Destinatario              *[]DestinatarioComercioExterior20 `xml:"Destinatario" bson:"Destinatario,omitempty"`
	Mercancias                []MercanciaComercioExterior20     `xml:"Mercancias>Mercancia" bson:"Mercancias"`
}

type EmisorComercioExterior20 struct {
	Curp      *string                     `xml:"Curp,attr" bson:"Curp,omitempty"` // Cifrado
	Domicilio DomicilioComercioExterior20 `xml:"Domicilio" bson:"Domicilio"`
}

type PropietarioComercioExterior20 struct {
	NumRegIdTrib     string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib"`         // Cifrado
	ResidenciaFiscal string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal"` // Cifrado
}

type ReceptorComercioExterior20 struct {
	NumRegIdTrib *string                      `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"` // Cifrado
	Domicilio    *DomicilioComercioExterior20 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DestinatarioComercioExterior20 struct {
	NumRegIdTrib *string                       `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"` // Cifrado
	Nombre       *string                       `xml:"Nombre,attr" bson:"Nombre,omitempty"`             // Cifrado
	Domicilio    []DomicilioComercioExterior20 `xml:"Domicilio" bson:"Domicilio"`
}

type MercanciaComercioExterior20 struct {
	NoIdentificacion         string                                        `xml:"NoIdentificacion,attr" bson:"NoIdentificacion"`
	FraccionArancelaria      *string                                       `xml:"FraccionArancelaria,attr" bson:"FraccionArancelario,omitempty"`
	CantidadAduana           *float64                                      `xml:"CantidadAduana,attr" bson:"CantidadAduana,omitempty"`
	UnidadAduana             *string                                       `xml:"UnidadAduana,attr" bson:"UnidadAduana,omitempty"`
	ValorUnitarioAduana      *float64                                      `xml:"ValorUnitarioAduana,attr" bson:"ValorUnitarioAduana,omitempty"`
	ValorDolares             float64                                       `xml:"ValorDolares,attr" bson:"ValorDolares"`
	DescripcionesEspecificas *[]DescripcionesEspecificasComercioExterior20 `xml:"DescripcionesEspecificas" bson:"DescripcionesEspecificas,omitempty"`
}

type DomicilioComercioExterior20 struct {
	Calle          string  `xml:"Calle,attr" bson:"Calle"`                             // Cifrado
	NumeroExterior *string `xml:"NumeroExterior,attr" bson:"NumeroExterior,omitempty"` // Cifrado
	NumeroInterior *string `xml:"NumeroInterior,attr" bson:"NumeroInterior,omitempty"` // Cifrado
	Colonia        *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`               // Cifrado
	Localidad      *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`           // Cifrado
	Referencia     *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`         // Cifrado
	Municipio      *string `xml:"Municipio,attr" bson:"Municipio,omitempty"`           // Cifrado
	Estado         string  `xml:"Estado,attr" bson:"Estado"`                           // Cifrado
	Pais           string  `xml:"Pais,attr" bson:"Pais"`                               // Cifrado
	CodigoPostal   string  `xml:"CodigoPostal,attr" bson:"CodigoPostal"`               // Cifrado
}

type DescripcionesEspecificasComercioExterior20 struct {
	Marca       string  `xml:"Marca,attr" bson:"Marca"`
	Modelo      *string `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	SubModelo   *string `xml:"SubModelo,attr" bson:"SubModelo,omitempty"`
	NumeroSerie *string `xml:"NumeroSerie,attr" bson:"NumeroSerie,omitempty"`
}

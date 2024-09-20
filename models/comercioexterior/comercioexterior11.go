package comercioexterior

type ComercioExterior11 struct {
	Version                   string                            `xml:"Version,attr" bson:"Version"`
	MotivoTraslado            *string                           `xml:"MotivoTraslado,attr" bson:"MotivoTraslado,omitempty"`
	TipoOperacion             string                            `xml:"TipoOperacion,attr" bson:"TipoOperacion"`
	ClavePedimento            *string                           `xml:"ClaveDePedimento,attr" bson:"ClavePedimento,omitempty"`
	CertificadoOrigen         *float64                          `xml:"CertificadoOrigen,attr" bson:"CertificadoOrigen,omitempty"`
	NumCertificadoOrigen      *string                           `xml:"NumCertificadoOrigen,attr" bson:"NumCertificadoOrigen,omitempty"`
	NumeroExportadorConfiable *string                           `xml:"NumeroExportadorConfiable,attr" bson:"NumeroExportadorConfiable,omitempty"`
	Incoterm                  *string                           `xml:"Incoterm,attr" bson:"Incoterm,omitempty"`
	Subdivision               *float64                          `xml:"Subdivision,attr" bson:"Subdivision,omitempty"`
	Observaciones             *string                           `xml:"Observaciones,attr" bson:"Observaciones,omitempty"`
	TipoCambioUSD             *string                           `xml:"TipoCambioUSD,attr" bson:"TipoCambioUSD,omitempty"`
	TotalUSD                  *float64                          `xml:"TotalUSD,attr" bson:"TotalUSD,omitempty"`
	Emisor                    *EmisorComercioExterior11         `xml:"Emisor" bson:"Emisor,omitempty"`
	Propietario               *[]PropietarioComercioExterior11  `xml:"Propietario" bson:"Propietario,omitempty"`
	Receptor                  *ReceptorComercioExterior11       `xml:"Receptor" bson:"Receptor,omitempty"`
	Destinatario              *[]DestinatarioComercioExterior11 `xml:"Destinatario" bson:"Destinatario,omitempty"`
	Mercancias                *[]MercanciaComercioExterior11    `xml:"Mercancias>Mercancia" bson:"Mercancias,omitempty"`
}

type EmisorComercioExterior11 struct {
	Curp      *string                      `xml:"Curp,attr" bson:"Curp,omitempty"` // Cifrado
	Domicilio *DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DomicilioComercioExterior11 struct {
	Calle          string  `xml:"Calle,attr" bson:"Calle"`                             // Cifrado
	NumeroExterior *string `xml:"NumeroExterior,attr" bson:"NumeroExterior,omitempty"` // Cifrado
	NumeroInterior *string `xml:"NumeroInterior,attr" bson:"NumeroInterior,omitempty"` // Cifrado
	Colonia        *string `xml:"Colonia,attr" bson:"Colonia,omitemtpy"`               // Cifrado
	Localidad      *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`           // Cifrado
	Referencia     *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`         // Cifrado
	Municipio      *string `xml:"Municipio,attr" bson:"Municipio,omitemtpy"`           // Cifrado
	Estado         string  `xml:"Estado,attr" bson:"Estado"`                           // Cifrado
	Pais           string  `xml:"Pais,attr" bson:"Pais"`                               // Cifrado
	CodigoPostal   string  `xml:"CodigoPostal,attr" bson:"CodigoPostal"`               // Cifrado
}

type PropietarioComercioExterior11 struct {
	NumRegIdTrib     string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib"`         // Cifrado
	ResidenciaFiscal string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal"` // Cifrado
}

type ReceptorComercioExterior11 struct {
	NumRegIdTrib *string                      `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"` // Cifrado
	Domicilio    *DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DestinatarioComercioExterior11 struct {
	NumRegIdTrib *string                       `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"` // Cifrado
	Nombre       *string                       `xml:"Nombre,attr" bson:"Nombre,omitempty"`             // Cifrado
	Domicilio    []DomicilioComercioExterior11 `xml:"Domicilio" bson:"Domicilio"`
}

type MercanciaComercioExterior11 struct {
	NoIdentificacion                         string                                        `xml:"NoIdentificacion,attr" bson:"NoIdentificacion"`
	FraccionArancelaria                      *string                                       `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	CantidadAduana                           *float64                                      `xml:"CantidadAduana,attr" bson:"CantidadAduana,omitempty"`
	UnidadAduana                             *string                                       `xml:"UnidadAduana,attr" bson:"UnidadAduana,omitempty"`
	ValorUnitarioAduana                      *float64                                      `xml:"ValorUnitarioAduana,attr" bson:"ValorUnitarioAduana,omitempty"`
	ValorDolares                             float64                                       `xml:"ValorDolares,attr" bson:"ValorDolares"`
	DescripcionesEspecificasComercioExterior *[]DescripcionesEspecificasComercioExterior11 `xml:"DescripcionesEspecificas" bson:"DescripcionesEspecificasComercioExterior,omitempty"`
}

type DescripcionesEspecificasComercioExterior11 struct {
	Marca       string  `xml:"Marca,attr" bson:"Marca"`
	Modelo      *string `xml:"Modelo,attr" bson:"Modelo,omitempty"`
	SubModelo   *string `xml:"SubModelo,attr" bson:"SubModelo,omitempty"`
	NumeroSerie *string `xml:"NumeroSerie,attr" bson:"NumeroSerie,omitempty"`
}

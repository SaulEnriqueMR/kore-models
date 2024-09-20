package cartaporte

import "encoding/xml"

type CartaPorte10 struct {
	XMLName           xml.Name                      `xml:"CartaPorte"`
	Version           string                        `xml:"Version,attr" bson:"Version"`
	TranspInternac    string                        `xml:"TranspInternac,attr" bson:"TranspInternac"`
	EntradaSalidaMerc *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMerc,omitempty"`
	ViaEntradaSalida  *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistRec      *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistRec,omitempty"`
	Ubicaciones       []UbicacionCartaPorte10       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias        MercanciasCartaPorte10        `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte  *FiguraTransporteCartaPorte10 `xml:"FiguraTransporte" bson:"FiguraTransporte,omitempty"`
}

type UbicacionCartaPorte10 struct {
	TipoEstacion       *string                `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty"`
	DistanciaRecorrida *float64               `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty"`
	Origen             *OrigenCartaPorte10    `xml:"Origen" bson:"Origen,omitempty"`
	Destino            *DestinoCartaPorte10   `xml:"Destino" bson:"Destino,omitempty"`
	Domicilio          *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type OrigenCartaPorte10 struct {
	IDOrigen          *string `xml:"IDOrigen,attr" bson:"IDOrigen,omitempty"`
	RFCRemitente      *string `xml:"RFCRemitente,attr" bson:"RFCRemitente,omitempty"`         // Cifrado
	NombreRemitente   *string `xml:"NombreRemitente,attr" bson:"NombreRemitente,omitempty"`   // Cifrado
	NumRegIdTrib      *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`         // Cifrado
	ResidenciaFiscal  *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"` // Cifrado
	NumEstacion       *string `xml:"NumEstacion,attr" bson:"NumEstacion,omitempty"`
	NombreEstacion    *string `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico *string `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraSalida   string  `xml:"FechaHoraSalida,attr" bson:"FechaHoraSalida"`
}

type DestinoCartaPorte10 struct {
	IDDestino            *string `xml:"IDDestino,attr" bson:"IDDestino,omitempty"`
	RFCDestinatario      *string `xml:"RFCDestinatario,attr" bson:"RFCDestinatario,omitempty"`       // Cifrado
	NombreDestinatario   *string `xml:"NombreDestinatario,attr" bson:"NombreDestinatario,omitempty"` // Cifrado
	NumRegIdTrib         *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`             // Cifrado
	ResidenciaFiscal     *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`     // Cifrado
	NumEstacion          *string `xml:"NumEstacion,attr" bson:"NumEstacion,omitempty"`
	NombreEstacion       *string `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico    *string `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraProgLlegada string  `xml:"FechaHoraProgLlegada,attr" bson:"FechaHoraProgLlegada"`
}

type DomicilioCartaPorte10 struct {
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

type MercanciasCartaPorte10 struct {
	PesoBrutoTotal        *float64                           `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal,omitempty"`
	UnidadPeso            *string                            `xml:"UnidadPeso,attr" bson:"UnidadPeso,omitempty"`
	PesoNetoTotal         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumTotalMercancias    float64                            `xml:"NumTotalMercancias,attr" bson:"NumTotalMercancias"`
	CargoPorTasacion      *string                            `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty"`
	Mercancia             []MercanciaCartaPorte10            `xml:"Mercancia" bson:"Mercancia"`
	AutotransporteFederal *AutotransporteFederalCartaPorte10 `xml:"AutotransporteFederal" bson:"AutotransporteFederal,omitemtpy"`
	TransporteMaritimo    *TransporteMaritimoCartaPorte10    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty"`
	TransporteAereo       *TransporteAereoCartaPorte10       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty"`
	TransporteFerroviario *TransporteFerroviarioCartaPorte10 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte10 struct {
	BienesTransp         *string                           `xml:"BienesTransp,attr" bson:"BienesTransp,omitempty"`
	ClaveSTCC            *string                           `xml:"ClaveSTCC,attr" bson:"ClaveSTCC,omitempty"`
	Descripcion          *string                           `xml:"Descripcion,attr" bson:"Descripcion,omitempty"`
	Cantidad             *float64                          `xml:"Cantidad,attr" bson:"Cantidad,omitempty"`
	ClaveUnidad          *string                           `xml:"ClaveUnidad,attr" bson:"ClaveUnidad,omitempty"`
	Unidad               *string                           `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones          *string                           `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`
	MaterialPeligroso    *string                           `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	CveMaterialPeligroso *string                           `xml:"CveMaterialPeligroso,attr" bson:"CveMaterialPeligroso,omitempty"`
	Embalaje             *string                           `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripEmbalaje      *string                           `xml:"DescripEmbalaje,attr" bson:"DescripEmbalaje,omitempty"`
	PesoEnKg             float64                           `xml:"PesoEnKg,attr" bson:"PesoEnKg"`
	ValorMercancia       *float64                          `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda               *string                           `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria  *string                           `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UUIDComercioExt      *string                           `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty"`
	CantidadTransporta   *[]CantidadTransportaCartaPorte10 `xml:"CantidadTrasporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia     *DetalleMercanciaCartaPorte10     `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
}

type CantidadTransportaCartaPorte10 struct {
	Cantidad       float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IDOrigen       string  `xml:"IDOrigen,attr" bson:"IDOrigen"`
	IDDestino      string  `xml:"IDDestina,attr" bson:"IDDestino"`
	CvesTransporte *string `xml:"CvesTransporte,attr" bson:"CvesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte10 struct {
	UnidadPeso string   `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoBruto  float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	PesoNeto   float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoTara   float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	NumPiezas  *float64 `xml:"NumPiezas,attr" bson:"NumPiezas,omitempty"`
}

type AutotransporteFederalCartaPorte10 struct {
	PermSCT                 string                              `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT           string                              `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"`     // Cifrado
	NombreAseg              string                              `xml:"NombreAseg,attr" bson:"NombreAseg"`           // Cifrado
	NumPolizaSeguro         string                              `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro"` // Cifrado
	IdentificacionVehicular IdentificacionVehicularCartaPorte10 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Remolques               *[]RemolqueCartaPorte10             `xml:"Renolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte10 struct {
	ConfigVehicular string `xml:"ConfigVehicular,attr" bson:"ConfigVehicular"`
	PlacaVM         string `xml:"PlacaVM,attr" bson:"PlacaVM"` // Cifrado
	AnioModelo      string `xml:"AnioModelo,attr" bson:"AnioModelo"`
}

type RemolqueCartaPorte10 struct {
	SubTipoRem string `xml:"SubTipoRem,attr" bson:"SubTipoRem"`
	Placa      string `xml:"Place,attr" bson:"Placa"` // Cifrado
}

type TransporteMaritimoCartaPorte10 struct {
	PermSCT                *string                  `xml:"PermSCT,attr" bson:"PermSCT,omitempty"`
	NumPermisoSCT          *string                  `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT,omitempty"`     // Cifrado
	NombreAseg             *string                  `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`           // Cifrado
	NumPolizaSeguro        *string                  `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"` // Cifrado
	TipoEmbarcacion        string                   `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion"`
	Matricula              string                   `xml:"Matricula,attr" bson:"Matricula"` // Cifrado
	NumeroOMI              string                   `xml:"NumeroOMI,attr" bson:"NumeroOMI"`
	AnioEmbarcacion        *string                  `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty"`
	NombreEmbarc           *string                  `xml:"NombreEmbarc,attr" bson:"NombreEmbarc,omitempty"` // Cifrado
	NacionalidadEmbarc     string                   `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarc"`
	UnidadesDeArqBruto     float64                  `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesDeArqBruto"`
	TipoCarga              string                   `xml:"TipoCarga,attr" bson:"TipoCarga"`
	NumCertITC             string                   `xml:"NumCertITC,attr" bson:"NumCertITC"`
	Eslora                 *float64                 `xml:"Eslora,attr" bson:"Eslora,omitempty"`
	Manga                  *float64                 `xml:"Manga,attr" bson:"Manga,omitempty"`
	Calado                 *float64                 `xml:"Calado,attr" bson:"Calado,omitempty"`
	LineaNaviera           *string                  `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty"`
	NombreAgenteNaviero    string                   `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero"`       // Cifrado
	NumAutorizacionNaviero string                   `xml:"NumAutorizacionNaviero,attr" bson:"NumAutorizacionNaviero"` // Cifrado
	NumViaje               *string                  `xml:"NumViaje,attr" bson:"NumViaje,omitempty"`                   // Cifrado
	NumConocEmbarc         *string                  `xml:"NumConocEmbarc,attr" bson:"NumConocEmbarc,omitempty"`
	Contenedor             []ContenedorCartaPorte10 `xml:"Contenedor" bson:"Contenedor"`
}

type ContenedorCartaPorte10 struct {
	MatriculaContenedor string  `xml:"MatriculaContenedor,attr" bson:"MatriculaContenedor"` // Cifrado
	TipoContenedor      string  `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	NumPrecinto         *string `xml:"NumPrecinto,attr" bson:"NumPrecinto,omitempty"`
}

type TransporteAereoCartaPorte10 struct {
	PermSCT                  string  `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT            string  `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"`               // Cifrado
	MatriculaAeronave        string  `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave"`       // Cifrado
	NombreAseg               *string `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`           // Cifrado
	NumPolizaSeguro          *string `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"` // Cifrado
	NumeroGuia               string  `xml:"NumeroGuia,attr" bson:"NumeroGuia"`
	LugarContrato            *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty"`       // Cifrado
	RFCTransportista         *string `xml:"RFCTransportista,attr" bson:"RFCTransportista,omitempty"` // Cifrado
	CodigoTransportista      string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista"`
	NumRegIdTribTranspor     *string `xml:"NumRegIdTribTranspor,attr" bson:"NumRegIdTribTranspor,omitempty"`         // Cifrado
	ResidenciaFiscalTranspor *string `xml:"ResidenciaFiscalTranspor,attr" bson:"ResidenciaFiscalTranspor,omitempty"` // Cifrado
	NombreTransportista      *string `xml:"NombreTransportista,attr" bson:"NombreTransportista,omitempty"`           // Cifrado
	RFCEmbarcador            *string `xml:"RFCEmbarcador,attr" bson:"RFCEmbarcador,omitempty"`                       // Cifrado
	NumRegIdTribEmbarc       *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarc,omitempty"`             // Cifrado
	NombreEmbarcador         *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty"`                 // Cifrado
}

type TransporteFerroviarioCartaPorte10 struct {
	TipoServicio    string                      `xml:"TipoServicio,attr" bson:"TipoServicio"`
	NombreAseg      *string                     `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`           // Cifrado
	NumPolizaSeguro *string                     `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"` // Cifrado
	Concesionario   *string                     `xml:"Concesionario,attr" bson:"Concesionario,omitempty"`
	DerechosPaso    *[]DerechosPasoCartaPorte10 `xml:"DerechosDePaso" bson:"DerechosPaso,omitempty"`
	Carro           []CarroCartaPorte10         `xml:"Carro" bson:"Carro"`
}

type DerechosPasoCartaPorte10 struct {
	TipoDerechoDePaso string  `xml:"TipoDerechoDePaso,attr" bson:"TipoDerechoDePaso"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte10 struct {
	TipoCarro           string                         `xml:"TipoCarro,attr" bson:"TipoCarro"`
	MatriculaCarro      string                         `xml:"MatriculaCarro,attr" bson:"MatriculaCarro"` // Cifrado
	GuiaCarro           string                         `xml:"GuiaCarro,attr" bson:"GuiaCarro"`
	ToneladasNetasCarro float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetasCarro"`
	Contenedor          *[]ContenedorCarroCartaPorte10 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte10 struct {
	TipoContenedor      string  `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte10 struct {
	CvesTransporte string                      `xml:"CvesTransporte,attr" bson:"CvesTransporte"`
	Operadores     *[]OperadoresCartaPorte10   `xml:"Operadores" bson:"Operadores,omitempty"`
	Propietario    *[]ProprietarioCartaPorte10 `xml:"Propietario" bson:"Propietario,omitempty"`
	Arrendatario   *[]ArrendatarioCartaPorte10 `xml:"Arrendatario" bson:"Arrendatario,omitempty"`
	Notificado     *[]NotificadoCartaPorte10   `xml:"Notificado" bson:"Notificado,omitempty"`
}

type OperadoresCartaPorte10 struct {
	Operador []OperadorCartaPorte10 `xml:"Operador" bson:"Operador"`
}

type OperadorCartaPorte10 struct {
	RFCOperador              *string                `xml:"RFCOperador,attr" bson:"RFCOperador,omitempty"`                           // Cifrado
	NumLicencia              *string                `xml:"NumLicencia,attr" bson:"NumLicencia,omitempty"`                           // Cifrado
	NombreOperador           *string                `xml:"NombreOperador,attr" bson:"NombreOperador,omitempty"`                     // Cifrado
	NumRegIdTribOperador     *string                `xml:"NumRegIdTribOperador,attr" bson:"NumRegIdTribOperador,omitempty"`         // Cifrado
	ResidenciaFiscalOperador *string                `xml:"ResidenciaFiscalOperador,attr" bson:"ResidenciaFiscalOperador,omitempty"` // Cifrado
	Domicilio                *DomicilioCartaPorte10 `xmml:"Domicilio" bson:"Domicilio,omitempty"`
}

type ProprietarioCartaPorte10 struct {
	RFCPropietario              *string                `xml:"RFCPropietario,attr" bson:"RFCPropietario,omitempty"`                           // Cifrado
	NombrePropietario           *string                `xml:"NombrePropietario,attr" bson:"NombrePropietario,omitempty"`                     // Cifrado
	NumRegIdTribPropietario     *string                `xml:"NumRegIdTribPropietario,attr" bson:"NumRegIdTribPropietario,omitempty"`         // Cifrado
	ResidenciaFiscalPropietario *string                `xml:"ResidenciaFiscalPropietario,attr" bson:"ResidenciaFiscalPropietario,omitempty"` // Cifrado
	Domicilio                   *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type ArrendatarioCartaPorte10 struct {
	RFCArrendatario              *string                `xml:"RFCArrendatario,attr" bson:"RFCArrendatario,omitempty"`                           // Cifrado
	NombreArrendatario           *string                `xml:"NombreArrendatario,attr" bson:"NombreArrendatario,omitempty"`                     // Cifrado
	NumRegIdTribArrendatario     *string                `xml:"NumRegIdTribArrendatario,attr" bson:"NumRegIdTribArrendatario,omitempty"`         // Cifrado
	ResidenciaFiscalArrendatario *string                `xml:"ResidenciaFiscalArrendatario,attr" bson:"ResidenciaFiscalArrendatario,omitempty"` // Cifrado
	Domicilio                    *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type NotificadoCartaPorte10 struct {
	RFCNotificado              *string                `xml:"RFCNotificado,attr" bson:"RFCNotificado,omitempty"`                           // Cifrado
	ResidenciaFiscalNotificado *string                `xml:"ResidenciaFiscalNotificado,attr" bson:"ResidenciaFiscalNotificado,omitempty"` // Cifrado
	NumRegIdTribNotificado     *string                `xml:"NumRegIdTribNotificado,attr" bson:"NumRegIdTribNotificado,omitempty"`         // Cifrado
	NombreNotificado           *string                `xml:"NombreNotificado,attr" bson:"NombreNotificado,omitempty"`                     // Cifrado
	Domicilio                  *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

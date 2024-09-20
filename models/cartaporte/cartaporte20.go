package cartaporte

import "encoding/xml"

type CartaPorte20 struct {
	XMLName           xml.Name                      `xml:"CartaPorte"`
	Version           string                        `xml:"Version,attr" bson:"Version"`
	TranspInternac    string                        `xml:"TranspInternac,attr" bson:"TranspInternac"`
	EntradaSalidaMerc *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMerc,omitempty"`
	PaisOrigenDestino *string                       `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida  *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistRec      *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistRec,omitempty"`
	Ubicaciones       []UbicacionCartaPorte20       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias        MercanciasCartaPorte20        `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte  *FiguraTransporteCartaPorte20 `xml:"FiguraTransporte" bson:"FiguraTransporte"`
}

type UbicacionCartaPorte20 struct {
	TipoUbicacion               string                 `xml:"TipoUbicacion,attr" bson:"TipoUbicacion"`
	IDUbicacion                 *string                `xml:"IDUbicacion,attr" bson:"IDUbicacion,omitempty"`
	RFCRemitenteDestinatario    string                 `xml:"RFCRemitenteDestinatario,attr" bson:"RFCRemitenteDestinatario"`                 // Cifrado
	NombreRemitenteDestinatario *string                `xml:"NombreRemitenteDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty"` // Cifrado
	NumRegIdTrib                *string                `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`                               // Cifrado
	ResidenciaFiscal            *string                `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`                       // Cifrado
	NumEstacion                 *string                `xml:"NumEstacion,attr" bson:"NumEstacion,omitempty"`
	NombreEstacion              *string                `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico           *string                `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegada      string                 `xml:"FechaHoraSalidaLlegada,attr" bson:"FechaHoraSalidaLlegada"`
	TipoEstacion                *string                `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty"`
	DistanciaRecorrida          *float64               `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty"`
	Domicilio                   *DomicilioCartaPorte20 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DomicilioCartaPorte20 struct {
	Calle          *string `xml:"Calle,attr" bson:"Calle,omitempty"`
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

type MercanciasCartaPorte20 struct {
	PesoBrutoTotal        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal"`
	UnidadPeso            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoNetoTotal         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumTotalMercancias    float64                            `xml:"NumTotalMercancias,attr" bson:"NumTotalMercancias"`
	CargoPorTasacion      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty"`
	Mercancia             []MercanciaCartaPorte20            `xml:"Mercancia" bson:"Mercancia"`
	Autotransporte        *AutotransporteCartaPorte20        `xml:"Autotransporte" bson:"Autotrasporte,omitempty"`
	TransporteMaritimo    *TransporteMaritimoCartaPorte20    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty"`
	TransporteAereo       *TransporteAereoCartaPorte20       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty"`
	TransporteFerroviario *TransporteFerroviarioCartaPorte20 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte20 struct {
	BienesTransp         string                             `xml:"BienesTransp,attr" bson:"BienesTransp"`
	ClaveSTCC            *string                            `xml:"ClaveSTCC,attr" bson:"ClaveSTCC,omitempty"`
	Descripcion          string                             `xml:"Descripcion,attr" bson:"Descripcion"`
	Cantidad             float64                            `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad          string                             `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad               *string                            `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones          *string                            `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`
	MaterialPeligroso    *string                            `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	CveMaterialPeligroso *string                            `xml:"CveMaterialPeligroso,attr" bson:"CveMaterialPeligroso,omitempty"`
	Embalaje             *string                            `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripEmbalaje      *string                            `xml:"DescripEmbalaje,attr" bson:"DescripEmbalaje,omitempty"`
	PesoKg               float64                            `xml:"PesoEnKg,attr" bson:"PesoKg"`
	ValorMercancia       *float64                           `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda               *string                            `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria  *string                            `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UUIDComercioExt      *string                            `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty"`
	Pedimentos           *[]PedimentosCartaPorte20          `xml:"Pedimentos" bson:"Pedimentos,omitempty"`
	GuiasIdentificacion  *[]GuiasIdentificacionCartaPorte20 `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty"`
	CantidadTransporta   *[]CantidadTransportaCartaPorte20  `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia     *DetalleMercanciaCartaPorte20      `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
}

type PedimentosCartaPorte20 struct {
	Pedimento string `xml:"Pedimento,attr" bson:"Pedimento"`
}

type GuiasIdentificacionCartaPorte20 struct {
	NumeroGuiaIdentificacion  string  `xml:"NumeroGuiaIdentificacion,attr" bson:"NumeroGuiaIdentificacion"`
	DescripGuiaIdentificacion string  `xml:"DescripGuiaIdentificacion,attr" bson:"DescripGuiaIdentificacion"`
	PesoGuiaIdentificacion    float64 `xml:"PesoGuiaIdentificacion,attr" bson:"PesoGuiaIdentificacion"`
}

type CantidadTransportaCartaPorte20 struct {
	Cantidad       float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IDOrigen       string  `xml:"IDOrigen,attr" bson:"IDOrigen"`
	IDDestino      string  `xml:"IDDestino,attr" bson:"IDDestino"`
	CvesTransporte *string `xml:"CvesTransporte,attr" bson:"CvesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte20 struct {
	UnidadPesoMerc string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPesoMerc"`
	PesoBruto      float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	PesoNeto       float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoTara       float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	NumPiezas      *float64 `xml:"NumPiezas,attr" bson:"NumPiezas,omitempty"`
}

type AutotransporteCartaPorte20 struct {
	PermSCT                 string                              `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT           string                              `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"` // Cifrado
	IdentificacionVehicular IdentificacionVehicularCartaPorte20 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte20                 `xml:"Seguros" bson:"Seguros"`
	Remolques               *[]RemolqueCartaPorte20             `xml:"Remolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte20 struct {
	ConfigVehicular string `xml:"ConfigVehicular,attr" bson:"ConfigVehicular"`
	PlacaVM         string `xml:"PlacaVM,attr" bson:"PlacaVM"` // Cifrado
	AnioModeloVM    string `xml:"AnioModeloVM,attr" bson:"AnioModeloVM"`
}

type SegurosCartaPorte20 struct {
	AseguraRespCivil   string   `xml:"AseguraRespCivil,attr" bson:"AseguraRespCivil"`
	PolizaRespCivil    string   `xml:"PolizaRespCivil,attr" bson:"PolizaRespCivil"` // Cifrado
	AseguraMedAmbiente *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguraMedAmbiente,omitempty"`
	PolizaMedAmbiente  *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedAmbiente,omitempty"` // Cifrado
	AseguraCarga       *string  `xml:"AseguraCarga,attr" bson:"AseguraCarga,omitempty"`
	PolizaCarga        *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty"` // Cifrado
	PrimaSeguro        *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte20 struct {
	SubTipoRem string `xml:"SubTipoRem,attr" bson:"SubTipoRem"`
	Placa      string `xml:"Placa,attr" bson:"Placa"` // Cifrado
}

type TransporteMaritimoCartaPorte20 struct {
	PermSCT                *string                  `xml:"PermSCT,attr" bson:"PermSCT,omitempty"`
	NumPermisoSCT          *string                  `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT,omitempty"`     // Cifrado
	NombreAseg             *string                  `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`           // Cifrado
	NumPolizaSeguro        *string                  `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"` // Cifrado
	TipoEmbarcacion        string                   `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion"`
	Matricula              string                   `xml:"Matricula,attr" bson:"Matricula"`
	NumeroOMI              string                   `xml:"NumeroOMI,attr" bson:"NumeroOMI"`
	AnioEmbarcacion        *string                  `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty"`
	NombreEmbarc           *string                  `xml:"NombreEmbarc,attr" bson:"NombreEmbarc,omitempty"`   // Cifrado
	NacionalidadEmbarc     string                   `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarc"` // Cifrado
	UnidadesArqBruto       float64                  `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesArqBruto"`
	TipoCarga              string                   `xml:"TipoCarga,attr" bson:"TipoCarga"`
	NumCertITC             string                   `xml:"NumCertITC,attr" bson:"NumCertITC"`
	Eslora                 *float64                 `xml:"Eslora,attr" bson:"Eslora,omitempty"`
	Manga                  *float64                 `xml:"Manga,attr" bson:"Manga,omitempty"`
	Calado                 *float64                 `xml:"Calado,attr" bson:"Calado,omitempty"`
	LineaNaviera           *string                  `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty"`
	NombreAgenteNaviero    string                   `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero"`
	NumAutorizacionNaviero string                   `xml:"NumAutorizacionNaviero,attr" bson:"NumAutorizacionNaviero"`
	NumViaje               *string                  `xml:"NumViaje,attr" bson:"NumViaje,omitempty"`
	NumConocEmbarc         *string                  `xml:"NumConocEmbarc,attr" bson:"NumConocEmbarc,omitempty"`
	Contenedor             []ContenedorCartaPorte20 `xml:"Contenedor,attr" bson:"Contenedor"`
}

type ContenedorCartaPorte20 struct {
	MatriculaContenedor string  `xml:"MatriculaContenedor,attr" bson:"MatriculaContenedor"` // Cifrado
	TipoContenedor      string  `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	NumPrecinto         *string `xml:"NumPrecinto,attr" bson:"NumPrecinto,omitempty"`
}

type TransporteAereoCartaPorte20 struct {
	PermSCT                string  `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT          string  `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"`                   // Cifrado
	MatriculaAeronave      *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty"` // Cifrado
	NombreAseg             *string `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`               // Cifrado
	NumPolizaSeguro        *string `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"`     // Cifrado
	NumeroGuia             string  `xml:"NumeroGuia,attr" bson:"NumeroGuia"`
	LugarContrato          *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty"` // Cifrado
	CodigoTransportista    string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista"`
	RFCEmbarcador          *string `xml:"RFCEmbarcador,attr" bson:"RFCEmbarcador,omitempty"`                   // Cifrado
	NumRegIdTribEmbarc     *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarc,omitempty"`         // Cifrado
	ResidenciaFiscalEmbarc *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarc,omitempty"` // Cifrado
	NombreEmbarcador       *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty"`             // Cifrado
}

type TransporteFerroviarioCartaPorte20 struct {
	TipoServicio    string                     `xml:"TipoDeServicio,attr" bson:"TipoServicio"`
	TipoTrafico     string                     `xml:"TipoDeTrafico,attr" bson:"TipoTrafico"`
	NombreAseg      *string                    `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`           // Cifrado
	NumPolizaSeguro *string                    `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"` // Cifrado
	DerechoPaso     *[]DerechoPasoCartaPorte20 `xml:"DerechoDePaso" bson:"DerechoPaso,omitempty"`
	Carro           []CarroCartaPorte20        `xml:"Carro" bson:"Carro"`
}

type DerechoPasoCartaPorte20 struct {
	TipoDerechoPaso   string  `xml:"TipoDerechoDePaso,attr" bson:"TipoDerechoPaso"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte20 struct {
	TipoCarro           string                    `xml:"TipoCarro,attr" bson:"TipoCarro"`
	MatriculaCarro      string                    `xml:"MatriculaCarro,attr" bson:"MatriculaCarro"` // Cifrado
	GuiaCarro           string                    `xml:"GuiaCarro,attr" bson:"GuiaCarro"`
	ToneladasNetasCarro float64                   `xml:"ToneladasNetasCarro" bson:"ToneladasNetasCarro"`
	Contenedor          *[]ContenedorCartaPorte20 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte20 struct {
	TipoContenedor      string  `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte20 struct {
	TiposFigura []TiposFiguraCartaPorte20 `xml:"TiposFigura" bson:"TiposFigura"`
}

type TiposFiguraCartaPorte20 struct {
	TipoFigura             string                          `xml:"TipoFigura,attr" bson:"TipoFigura"`
	RFCFigura              *string                         `xml:"RFCFigura,attr" bson:"RFCFigura,omitempty"`                           // Cifrado
	NumLicencia            *string                         `xml:"NumLicencia,attr" bson:"NumLicencia,omitempty"`                       // Cifrado
	NombreFigura           *string                         `xml:"NombreFigura,attr" bson:"NombreFigura,omitempty"`                     // Cifrado
	NumRegIdTribFigura     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTribFigura,omitempty"`         // Cifrado
	ResidenciaFiscalFigura *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscalFigura,omitempty"` // Cifrado
	PartesTransporte       *[]PartesTransporteCartaPorte20 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty"`
	Domicilio              *DomicilioCartaPorte20          `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte20 struct {
	ParteTransporte string `xml:"ParteTransporte,attr" bson:"ParteTransporte"`
}

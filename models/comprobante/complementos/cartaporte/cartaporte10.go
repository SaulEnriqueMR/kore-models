package cartaporte

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"strings"
	"time"
)

type CartaPorte10 struct {
	Version                   string                        `xml:"Version,attr" bson:"Version" json:"Version"`
	TransporteInternacional   string                        `xml:"TranspInternac,attr" bson:"TransporteInternacional" json:"TransporteInternacional"`
	EsTransporteInternacional bool                          `bson:"EsTransporteInternacional" json:"EsTransporteInternacional"`
	EntradaSalidaMercancia    *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty" json:"EntradaSalidaMercancia,omitempty"`
	ViaEntradaSalida          *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty" json:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida   *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty" json:"TotalDistanciaRecorrida,omitempty"`
	Ubicaciones               []UbicacionCartaPorte10       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones" json:"Ubicaciones"`
	Mercancias                MercanciasCartaPorte10        `xml:"Mercancias" bson:"Mercancias" json:"Mercancias"`
	FiguraTransporte          *FiguraTransporteCartaPorte10 `xml:"FiguraTransporte" bson:"FiguraTransporte" json:"FiguraTransporte"`
}

func (ccp *CartaPorte10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias CartaPorte10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*ccp = CartaPorte10(aux)
	ccp.EsTransporteInternacional = helpers.ResolveSatBoolean(ccp.TransporteInternacional)

	return nil
}

type UbicacionCartaPorte10 struct {
	Origen             *OrigenCartaPorte10    `xml:"Origen" bson:"Origen,omitempty" json:"Origen,omitempty"`
	Destino            *DestinoCartaPorte10   `xml:"Destino" bson:"Destino,omitempty" json:"Destino,omitempty"`
	Domicilio          *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
	TipoEstacion       *string                `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty" json:"TipoEstacion,omitempty"`
	DistanciaRecorrida *float64               `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty" json:"DistanciaRecorrida,omitempty"`
}

type OrigenCartaPorte10 struct {
	IdUbicacion                  *string   `xml:"IDOrigen,attr" bson:"IdUbicacion,omitempty" json:"IdUbicacion,omitempty"`
	RfcRemitenteDestinatario     *string   `xml:"RFCRemitente,attr" bson:"RfcRemitenteDestinatario,omitempty" json:"RfcRemitenteDestinatario,omitempty"`
	NombreRemitenteDestinatario  *string   `xml:"NombreRemitente,attr" bson:"NombreRemitenteDestinatario,omitempty" json:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                 *string   `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal             *string   `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	NoEstacion                   *string   `xml:"NumEstacion,attr" bson:"NoEstacion,omitempty" json:"NoEstacion,omitempty"`
	NombreEstacion               *string   `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty" json:"NombreEstacion,omitempty"`
	NavegacionTrafico            *string   `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty" json:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegadaString string    `xml:"FechaHoraSalida,attr" bson:"FechaHoraSalidaLlegadaString" json:"FechaHoraSalidaLlegadaString"`
	FechaHoraSalidaLlegada       time.Time `bson:"FechaHoraSalidaLlegada" json:"FechaHoraSalidaLlegada"`
}

func (u *OrigenCartaPorte10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias OrigenCartaPorte10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*u = OrigenCartaPorte10(aux)
	fecha, errFecha := helpers.ParseDatetime(u.FechaHoraSalidaLlegadaString)
	if errFecha == nil {
		u.FechaHoraSalidaLlegada = fecha
	}
	return nil
}

type DestinoCartaPorte10 struct {
	IdUbicacion                  *string   `xml:"IDDestino,attr" bson:"IdUbicacion,omitempty" json:"IdUbicacion,omitempty"`
	RfcRemitenteDestinatario     *string   `xml:"RFCDestinatario,attr" bson:"RfcRemitenteDestinatario,omitempty" json:"RfcRemitenteDestinatario,omitempty"`
	NombreRemitenteDestinatario  *string   `xml:"NombreDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty" json:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                 *string   `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal             *string   `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	NoEstacion                   *string   `xml:"NumEstacion,attr" bson:"NoEstacion,omitempty" json:"NoEstacion,omitempty"`
	NombreEstacion               *string   `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty" json:"NombreEstacion,omitempty"`
	NavegacionTrafico            *string   `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty" json:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegadaString string    `xml:"FechaHoraProgLlegada,attr" bson:"FechaHoraSalidaLlegadaString" json:"FechaHoraSalidaLlegadaString"`
	FechaHoraSalidaLlegada       time.Time `bson:"FechaHoraSalidaLlegada" json:"FechaHoraSalidaLlegada"`
}

func (u *DestinoCartaPorte10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DestinoCartaPorte10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*u = DestinoCartaPorte10(aux)
	fecha, errFecha := helpers.ParseDatetime(u.FechaHoraSalidaLlegadaString)
	if errFecha == nil {
		u.FechaHoraSalidaLlegada = fecha
	}
	return nil
}

type DomicilioCartaPorte10 struct {
	Calle        string  `xml:"Calle,attr" bson:"Calle" json:"Calle"`
	NoExterior   *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty" json:"Colonia,omitempty"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`
	Municipio    *string `xml:"Municipio,attr" bson:"Municipio,omitempty" json:"Municipio,omitempty"`
	Estado       string  `xml:"Estado,attr" bson:"Estado" json:"Estado"`
	Pais         string  `xml:"Pais,attr" bson:"Pais" json:"Pais"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal" json:"CodigoPostal"`
}

type MercanciasCartaPorte10 struct {
	PesoBrutoTotal        *float64                           `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal,omitempty" json:"PesoBrutoTotal,omitempty"`
	UnidadPeso            *string                            `xml:"UnidadPeso,attr" bson:"UnidadPeso,omitempty" json:"UnidadPeso,omitempty"`
	PesoNetoTotal         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty" json:"PesoNetoTotal,omitempty"`
	NumeroTotalMercancias float64                            `xml:"NumTotalMercancias,attr" bson:"NumeroTotalMercancias" json:"NumeroTotalMercancias"`
	CargoPorTasacion      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty" json:"CargoPorTasacion,omitempty"`
	Mercancia             []MercanciaCartaPorte10            `xml:"Mercancia" bson:"Mercancia" json:"Mercancia"`
	Autotransporte        *AutotransporteCartaPorte10        `xml:"AutotransporteFederal" bson:"Autotrasporte,omitempty" json:"Autotrasporte,omitempty"`
	TransporteMaritimo    *TransporteMaritimoCartaPorte10    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty" json:"TransporteMaritimo,omitempty"`
	TransporteAereo       *TransporteAereoCartaPorte10       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty" json:"TransporteAereo,omitempty"`
	TransporteFerroviario *TransporteFerroviarioCartaPorte10 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty" json:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte10 struct {
	BienesTransportados    *string                           `xml:"BienesTransp,attr" bson:"BienesTransportados,omitempty" json:"BienesTransportados,omitempty"`
	ClaveStcc              *string                           `xml:"ClaveSTCC,attr" bson:"ClaveStcc,omitempty" json:"ClaveStcc,omitempty"`
	Descripcion            *string                           `xml:"Descripcion,attr" bson:"Descripcion,omitempty" json:"Descripcion,omitempty"`
	Cantidad               *string                           `xml:"Cantidad,attr" bson:"Cantidad,omitempty" json:"Cantidad,omitempty"`
	ClaveUnidad            *string                           `xml:"ClaveUnidad,attr" bson:"ClaveUnidad,omitempty" json:"ClaveUnidad,omitempty"`
	Unidad                 *string                           `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Dimensiones            *string                           `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty" json:"Dimensiones,omitempty"`
	MaterialPeligroso      *string                           `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty" json:"MaterialPeligroso,omitempty"`
	EsMaterialPeligroso    *bool                             `bson:"EsMaterialPeligroso,omitempty" json:"EsMaterialPeligroso,omitempty"`
	ClaveMaterialPeligroso *string                           `xml:"CveMaterialPeligroso,attr" bson:"ClaveMaterialPeligroso,omitempty" json:"ClaveMaterialPeligroso,omitempty"`
	Embalaje               *string                           `xml:"Embalaje,attr" bson:"Embalaje,omitempty" json:"Embalaje,omitempty"`
	DescripcionEmbalaje    *string                           `xml:"DescripEmbalaje,attr" bson:"DescripcionEmbalaje,omitempty" json:"DescripcionEmbalaje,omitempty"`
	PesoKg                 float64                           `xml:"PesoEnKg,attr" bson:"PesoKg" json:"PesoKg"`
	ValorMercancia         *float64                          `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty" json:"ValorMercancia,omitempty"`
	Moneda                 *string                           `xml:"Moneda,attr" bson:"Moneda,omitempty" json:"Moneda,omitempty"`
	FraccionArancelaria    *string                           `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty" json:"FraccionArancelaria,omitempty"`
	UUIDComercioExt        *string                           `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty" json:"UUIDComercioExt,omitempty"`
	UuidComercioExterior   *string                           `bson:"UuidComercioExterior,omitempty" json:"UuidComercioExterior,omitempty"`
	CantidadTransporta     *[]CantidadTransportaCartaPorte10 `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty" json:"CantidadTransporta,omitempty"`
	DetalleMercancia       *DetalleMercanciaCartaPorte10     `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty" json:"DetalleMercancia,omitempty"`
}

func (m *MercanciaCartaPorte10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias MercanciaCartaPorte10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*m = MercanciaCartaPorte10(aux)
	if m.MaterialPeligroso != nil {
		isMaterialPeligroso := helpers.ResolveSatBoolean(*aux.MaterialPeligroso)
		m.EsMaterialPeligroso = &isMaterialPeligroso
	}

	if m.UUIDComercioExt != nil && *m.UUIDComercioExt != "" {
		uuidComercioExterior := strings.ToUpper(*m.UUIDComercioExt)
		m.UuidComercioExterior = &uuidComercioExterior
	}

	return nil
}

type CantidadTransportaCartaPorte10 struct {
	Cantidad         float64 `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	IdOrigen         string  `xml:"IDOrigen,attr" bson:"IdOrigen" json:"IdOrigen"`
	IdDestino        string  `xml:"IDDestino,attr" bson:"IdDestino" json:"IdDestino"`
	ClavesTransporte *string `xml:"CvesTransporte,attr" bson:"ClavesTransporte,omitempty" json:"ClavesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte10 struct {
	UnidadPeso   string   `xml:"UnidadPeso,attr" bson:"UnidadPeso" json:"UnidadPeso"`
	PesoBruto    float64  `xml:"PesoBruto,attr" bson:"PesoBruto" json:"PesoBruto"`
	PesoNeto     float64  `xml:"PesoNeto,attr" bson:"PesoNeto" json:"PesoNeto"`
	PesoTara     float64  `xml:"PesoTara,attr" bson:"PesoTara" json:"PesoTara"`
	NumeroPiezas *float64 `xml:"NumPiezas,attr" bson:"NumeroPiezas,omitempty" json:"NumeroPiezas,omitempty"`
}

type AutotransporteCartaPorte10 struct {
	PermisoSct              string                              `xml:"PermSCT,attr" bson:"PermisoSct" json:"PermisoSct"`
	NoPermisoSct            string                              `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct" json:"NoPermisoSct"`
	NombreAseguradora       string                              `xml:"NombreAseg,attr" bson:"NombreAseguradora" json:"NombreAseguradora"`
	NoPoliza                string                              `xml:"NumPolizaSeguro,attr" bson:"NoPoliza" json:"NoPoliza"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte10 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular" json:"IdentificacionVehicular"`
	Remolques               *[]RemolqueCartaPorte10             `xml:"Remolques>Remolque" bson:"Remolques,omitempty" json:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte10 struct {
	ConfiguracionVehicular string `xml:"ConfigVehicular,attr" bson:"ConfiguracionVehicular" json:"ConfiguracionVehicular"`
	PlacaVm                string `xml:"PlacaVM,attr" bson:"PlacaVm" json:"PlacaVm"`
	AnioModeloVm           string `xml:"AnioModeloVM,attr" bson:"AnioModeloVm" json:"AnioModeloVm"`
}

type RemolqueCartaPorte10 struct {
	Subtipo string `xml:"SubTipoRem,attr" bson:"Subtipo" json:"Subtipo"`
	Placa   string `xml:"Placa,attr" bson:"Placa" json:"Placa"`
}

type TransporteMaritimoCartaPorte10 struct {
	PermisoSct              *string                  `xml:"PermSCT,attr" bson:"PermisoSct,omitempty" json:"PermisoSct,omitempty"`
	NoPermisoSct            *string                  `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct,omitempty" json:"NoPermisoSct,omitempty"`
	NombreAseguradora       *string                  `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty" json:"NombreAseguradora,omitempty"`
	NoPolizaSeguro          *string                  `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty" json:"NoPolizaSeguro,omitempty"`
	TipoEmbarcacion         string                   `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion" json:"TipoEmbarcacion"`
	Matricula               string                   `xml:"Matricula,attr" bson:"Matricula" json:"Matricula"`
	NoOmi                   string                   `xml:"NumeroOMI,attr" bson:"NoOmi" json:"NoOmi"`
	AnioEmbarcacion         *string                  `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty" json:"AnioEmbarcacion,omitempty"`
	NombreEmbarcacion       *string                  `xml:"NombreEmbarc,attr" bson:"NombreEmbarcacion,omitempty" json:"NombreEmbarcacion,omitempty"`
	NacionalidadEmbarcacion string                   `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarcacion" json:"NacionalidadEmbarcacion"`
	UnidadesArqueoBruto     float64                  `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesArqueoBruto" json:"UnidadesArqueoBruto"`
	TipoCarga               string                   `xml:"TipoCarga,attr" bson:"TipoCarga" json:"TipoCarga"`
	NumCertITC              string                   `xml:"NumCertITC,attr" bson:"NumCertITC" json:"NumCertITC"`
	Eslora                  *float64                 `xml:"Eslora,attr" bson:"Eslora,omitempty" json:"Eslora,omitempty"`
	Manga                   *float64                 `xml:"Manga,attr" bson:"Manga,omitempty" json:"Manga,omitempty"`
	Calado                  *float64                 `xml:"Calado,attr" bson:"Calado,omitempty" json:"Calado,omitempty"`
	LineaNaviera            *string                  `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty" json:"LineaNaviera,omitempty"`
	NombreAgenteNaviero     string                   `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero" json:"NombreAgenteNaviero"`
	NoAutorizacionNaviero   string                   `xml:"NumAutorizacionNaviero,attr" bson:"NoAutorizacionNaviero" json:"NoAutorizacionNaviero"`
	NoViaje                 *string                  `xml:"NumViaje,attr" bson:"NoViaje,omitempty" json:"NoViaje,omitempty"`
	NoConocimientoEmbarque  *string                  `xml:"NumConocEmbarc,attr" bson:"NoConocimientoEmbarque,omitempty" json:"NoConocimientoEmbarque,omitempty"`
	Contenedores            []ContenedorCartaPorte10 `xml:"Contenedor" bson:"Contenedores" json:"Contenedores"`
}

type ContenedorCartaPorte10 struct {
	Matricula  string  `xml:"MatriculaContenedor,attr" bson:"Matricula" json:"Matricula"`
	Tipo       string  `xml:"TipoContenedor,attr" bson:"Tipo" json:"Tipo"`
	NoPrecinto *string `xml:"NumPrecinto,attr" bson:"NoPrecinto,omitempty" json:"NoPrecinto,omitempty"`
}

type TransporteAereoCartaPorte10 struct {
	PermisoSct                    string  `xml:"PermSCT,attr" bson:"PermisoSct" json:"PermisoSct"`
	NoPermisoSct                  string  `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct" json:"NoPermisoSct"`
	MatriculaAeronave             *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty" json:"MatriculaAeronave,omitempty"`
	NombreAseguradora             *string `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty" json:"NombreAseguradora,omitempty"`
	NoPolizaSeguro                *string `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty" json:"NoPolizaSeguro,omitempty"`
	NoGuia                        string  `xml:"NumeroGuia,attr" bson:"NoGuia" json:"NoGuia"`
	LugarContrato                 *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty" json:"LugarContrato,omitempty"`
	RfcTransportista              *string `xml:"RFCTransportista,attr" bson:"RfcTransportista,omitempty" json:"RfcTransportista,omitempty"`
	CodigoTransportista           string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista" json:"CodigoTransportista"`
	NumRegIdTribTransportista     *string `xml:"NumRegIdTribTranspor,attr" bson:"NumRegIdTribTransportista,omitempty" json:"NumRegIdTribTransportista,omitempty"`
	ResidenciaFiscalTransportista *string `xml:"ResidenciaFiscalTranspor,attr" bson:"ResidenciaFiscalTransportista,omitempty" json:"ResidenciaFiscalTransportista,omitempty"`
	NombreTransportista           *string `xml:"NombreTransportista,attr" bson:"NombreTransportista,omitempty" json:"NombreTransportista,omitempty"`
	RfcEmbarcador                 *string `xml:"RFCEmbarcador,attr" bson:"RfcEmbarcador,omitempty" json:"RfcEmbarcador,omitempty"`
	NumRegIdTribEmbarcador        *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarcador,omitempty" json:"NumRegIdTribEmbarcador,omitempty"`
	ResidenciaFiscalEmbarcador    *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarcador,omitempty" json:"ResidenciaFiscalEmbarcador,omitempty"`
	NombreEmbarcador              *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty" json:"NombreEmbarcador,omitempty"`
}

type TransporteFerroviarioCartaPorte10 struct {
	TipoServicio      string                     `xml:"TipoDeServicio,attr" bson:"TipoServicio" json:"TipoServicio"`
	NombreAseguradora *string                    `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty" json:"NombreAseguradora,omitempty"`
	NoPolizaSeguro    *string                    `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty" json:"NoPolizaSeguro,omitempty"`
	Concesionario     *string                    `xml:"Concesionario,attr" bson:"Concesionario,omitempty" json:"Concesionario,omitempty"`
	DerechoPaso       *[]DerechoPasoCartaPorte10 `xml:"DerechosDePaso" bson:"DerechoPaso,omitempty" json:"DerechoPaso,omitempty"`
	Carro             []CarroCartaPorte10        `xml:"Carro" bson:"Carro" json:"Carro"`
}

type DerechoPasoCartaPorte10 struct {
	Tipo              string  `xml:"TipoDerechoDePaso,attr" bson:"Tipo" json:"Tipo"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado" json:"KilometrajePagado"`
}

type CarroCartaPorte10 struct {
	Tipo           string                         `xml:"TipoCarro,attr" bson:"Tipo" json:"Tipo"`
	Matricula      string                         `xml:"MatriculaCarro,attr" bson:"Matricula" json:"Matricula"`
	Guia           string                         `xml:"GuiaCarro,attr" bson:"Guia" json:"Guia"`
	ToneladasNetas float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetas" json:"ToneladasNetas"`
	Contenedor     *[]ContenedorCarroCartaPorte10 `xml:"Contenedor" bson:"Contenedor,omitempty" json:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte10 struct {
	Tipo                string  `xml:"TipoContenedor,attr" bson:"Tipo" json:"Tipo"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio" json:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia" json:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte10 struct {
	ClaveTransporte string                      `xml:"CveTransporte,attr" bson:"ClaveTransporte" json:"ClaveTransporte"`
	Operadores      *[]OperadorCartaPorte10     `xml:"Operadores>Operador" bson:"Operadores" json:"Operadores"`
	Propietarios    *[]PropietarioCartaPorte10  `xml:"Propietario" bson:"Propietario" json:"Propietario"`
	Arrendatarios   *[]ArrendatarioCartaPorte10 `xml:"Arrendatario" bson:"Arrendatarios" json:"Arrendatarios"`
	Notificados     *[]NotificadoCartaPorte10   `xml:"Notificado" bson:"Notificado" json:"Notificado"`
}

type OperadorCartaPorte10 struct {
	Rfc              *string                `xml:"RFCOperador,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
	NoLicencia       *string                `xml:"NumLicencia,attr" bson:"NoLicencia,omitempty" json:"NoLicencia,omitempty"`
	Nombre           *string                `xml:"NombreOperador,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribOperador,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalOperador,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type PropietarioCartaPorte10 struct {
	Rfc              *string                `xml:"RFCPropietario,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
	Nombre           *string                `xml:"NombrePropietario,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribPropietario,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalPropietario,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type ArrendatarioCartaPorte10 struct {
	Rfc              *string                `xml:"RFCArrendatario,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
	Nombre           *string                `xml:"NombreArrendatario,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribArrendatario,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalArrendatario,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type NotificadoCartaPorte10 struct {
	Rfc              *string                `xml:"RFCNotificado,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
	Nombre           *string                `xml:"NombreNotificado,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribNotificado,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalNotificado,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

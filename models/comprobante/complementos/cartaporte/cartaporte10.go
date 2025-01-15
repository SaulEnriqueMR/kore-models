package cartaporte

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"strings"
	"time"
)

type CartaPorte10 struct {
	Version                   string                        `xml:"Version,attr" bson:"Version"`
	TransporteInternacional   string                        `xml:"TranspInternac,attr" bson:"TransporteInternacional"`
	EsTransporteInternacional bool                          `bson:"EsTransporteInternacional"`
	EntradaSalidaMercancia    *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty"`
	ViaEntradaSalida          *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida   *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty"`
	Ubicaciones               []UbicacionCartaPorte10       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias                MercanciasCartaPorte10        `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte          *FiguraTransporteCartaPorte10 `xml:"FiguraTransporte" bson:"FiguraTransporte"`
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
	Origen             *OrigenCartaPorte10    `xml:"Origen" bson:"Origen,omitempty"`
	Destino            *DestinoCartaPorte10   `xml:"Destino" bson:"Destino,omitempty"`
	Domicilio          *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
	TipoEstacion       *string                `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty"`
	DistanciaRecorrida *float64               `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty"`
}

type OrigenCartaPorte10 struct {
	IdUbicacion                  *string   `xml:"IDOrigen,attr" bson:"IdUbicacion,omitempty"`
	RfcRemitenteDestinatario     *string   `xml:"RFCRemitente,attr" bson:"RfcRemitenteDestinatario,omitempty"`
	NombreRemitenteDestinatario  *string   `xml:"NombreRemitente,attr" bson:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                 *string   `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal             *string   `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`
	NoEstacion                   *string   `xml:"NumEstacion,attr" bson:"NoEstacion,omitempty"`
	NombreEstacion               *string   `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico            *string   `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegadaString string    `xml:"FechaHoraSalida,attr" bson:"FechaHoraSalidaLlegadaString"`
	FechaHoraSalidaLlegada       time.Time `bson:"FechaHoraSalidaLlegada"`
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
	IdUbicacion                  *string   `xml:"IDDestino,attr" bson:"IdUbicacion,omitempty"`
	RfcRemitenteDestinatario     *string   `xml:"RFCDestinatario,attr" bson:"RfcRemitenteDestinatario,omitempty"`
	NombreRemitenteDestinatario  *string   `xml:"NombreDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                 *string   `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal             *string   `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`
	NoEstacion                   *string   `xml:"NumEstacion,attr" bson:"NoEstacion,omitempty"`
	NombreEstacion               *string   `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico            *string   `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegadaString string    `xml:"FechaHoraProgLlegada,attr" bson:"FechaHoraSalidaLlegadaString"`
	FechaHoraSalidaLlegada       time.Time `bson:"FechaHoraSalidaLlegada"`
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
	Calle        string  `xml:"Calle,attr" bson:"Calle"`
	NoExterior   *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`
	Municipio    *string `xml:"Municipio,attr" bson:"Municipio,omitempty"`
	Estado       string  `xml:"Estado,attr" bson:"Estado"`
	Pais         string  `xml:"Pais,attr" bson:"Pais"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal"`
}

type MercanciasCartaPorte10 struct {
	PesoBrutoTotal        *float64                           `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal,omitempty"`
	UnidadPeso            *string                            `xml:"UnidadPeso,attr" bson:"UnidadPeso,omitempty"`
	PesoNetoTotal         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumeroTotalMercancias float64                            `xml:"NumTotalMercancias,attr" bson:"NumeroTotalMercancias"`
	CargoPorTasacion      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty"`
	Mercancia             []MercanciaCartaPorte10            `xml:"Mercancia" bson:"Mercancia"`
	Autotransporte        *AutotransporteCartaPorte10        `xml:"AutotransporteFederal" bson:"Autotrasporte,omitempty"`
	TransporteMaritimo    *TransporteMaritimoCartaPorte10    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty"`
	TransporteAereo       *TransporteAereoCartaPorte10       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty"`
	TransporteFerroviario *TransporteFerroviarioCartaPorte10 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte10 struct {
	BienesTransportados    *string                           `xml:"BienesTransp,attr" bson:"BienesTransportados,omitempty"`
	ClaveStcc              *string                           `xml:"ClaveSTCC,attr" bson:"ClaveStcc,omitempty"`
	Descripcion            *string                           `xml:"Descripcion,attr" bson:"Descripcion,omitempty"`
	Cantidad               *string                           `xml:"Cantidad,attr" bson:"Cantidad,omitempty"`
	ClaveUnidad            *string                           `xml:"ClaveUnidad,attr" bson:"ClaveUnidad,omitempty"`
	Unidad                 *string                           `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones            *string                           `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`
	MaterialPeligroso      *string                           `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	EsMaterialPeligroso    *bool                             `bson:"EsMaterialPeligroso,omitempty"`
	ClaveMaterialPeligroso *string                           `xml:"CveMaterialPeligroso,attr" bson:"ClaveMaterialPeligroso,omitempty"`
	Embalaje               *string                           `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripcionEmbalaje    *string                           `xml:"DescripEmbalaje,attr" bson:"DescripcionEmbalaje,omitempty"`
	PesoKg                 float64                           `xml:"PesoEnKg,attr" bson:"PesoKg"`
	ValorMercancia         *float64                          `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda                 *string                           `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria    *string                           `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UUIDComercioExt        *string                           `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty"`
	UuidComercioExterior   *string                           `bson:"UuidComercioExterior,omitempty"`
	CantidadTransporta     *[]CantidadTransportaCartaPorte10 `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia       *DetalleMercanciaCartaPorte10     `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
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
	Cantidad         float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IdOrigen         string  `xml:"IDOrigen,attr" bson:"IdOrigen"`
	IdDestino        string  `xml:"IDDestino,attr" bson:"IdDestino"`
	ClavesTransporte *string `xml:"CvesTransporte,attr" bson:"ClavesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte10 struct {
	UnidadPeso   string   `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoBruto    float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	PesoNeto     float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoTara     float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	NumeroPiezas *float64 `xml:"NumPiezas,attr" bson:"NumeroPiezas,omitempty"`
}

type AutotransporteCartaPorte10 struct {
	PermisoSct              string                              `xml:"PermSCT,attr" bson:"PermisoSct"`
	NoPermisoSct            string                              `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct"`
	NombreAseguradora       string                              `xml:"NombreAseg,attr" bson:"NombreAseguradora"`
	NoPoliza                string                              `xml:"NumPolizaSeguro,attr" bson:"NoPoliza"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte10 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Remolques               *[]RemolqueCartaPorte10             `xml:"Remolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte10 struct {
	ConfiguracionVehicular string `xml:"ConfigVehicular,attr" bson:"ConfiguracionVehicular"`
	PlacaVm                string `xml:"PlacaVM,attr" bson:"PlacaVm"`
	AnioModeloVm           string `xml:"AnioModeloVM,attr" bson:"AnioModeloVm"`
}

type RemolqueCartaPorte10 struct {
	Subtipo string `xml:"SubTipoRem,attr" bson:"Subtipo"`
	Placa   string `xml:"Placa,attr" bson:"Placa"`
}

type TransporteMaritimoCartaPorte10 struct {
	PermisoSct              *string                  `xml:"PermSCT,attr" bson:"PermisoSct,omitempty"`
	NoPermisoSct            *string                  `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct,omitempty"`
	NombreAseguradora       *string                  `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro          *string                  `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	TipoEmbarcacion         string                   `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion"`
	Matricula               string                   `xml:"Matricula,attr" bson:"Matricula"`
	NoOmi                   string                   `xml:"NumeroOMI,attr" bson:"NoOmi"`
	AnioEmbarcacion         *string                  `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty"`
	NombreEmbarcacion       *string                  `xml:"NombreEmbarc,attr" bson:"NombreEmbarcacion,omitempty"`
	NacionalidadEmbarcacion string                   `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarcacion"`
	UnidadesArqueoBruto     float64                  `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesArqueoBruto"`
	TipoCarga               string                   `xml:"TipoCarga,attr" bson:"TipoCarga"`
	NumCertITC              string                   `xml:"NumCertITC,attr" bson:"NumCertITC"`
	Eslora                  *float64                 `xml:"Eslora,attr" bson:"Eslora,omitempty"`
	Manga                   *float64                 `xml:"Manga,attr" bson:"Manga,omitempty"`
	Calado                  *float64                 `xml:"Calado,attr" bson:"Calado,omitempty"`
	LineaNaviera            *string                  `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty"`
	NombreAgenteNaviero     string                   `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero"`
	NoAutorizacionNaviero   string                   `xml:"NumAutorizacionNaviero,attr" bson:"NoAutorizacionNaviero"`
	NoViaje                 *string                  `xml:"NumViaje,attr" bson:"NoViaje,omitempty"`
	NoConocimientoEmbarque  *string                  `xml:"NumConocEmbarc,attr" bson:"NoConocimientoEmbarque,omitempty"`
	Contenedores            []ContenedorCartaPorte10 `xml:"Contenedor" bson:"Contenedores"`
}

type ContenedorCartaPorte10 struct {
	Matricula  string  `xml:"MatriculaContenedor,attr" bson:"Matricula"`
	Tipo       string  `xml:"TipoContenedor,attr" bson:"Tipo"`
	NoPrecinto *string `xml:"NumPrecinto,attr" bson:"NoPrecinto,omitempty"`
}

type TransporteAereoCartaPorte10 struct {
	PermisoSct                    string  `xml:"PermSCT,attr" bson:"PermisoSct"`
	NoPermisoSct                  string  `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct"`
	MatriculaAeronave             *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty"`
	NombreAseguradora             *string `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro                *string `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	NoGuia                        string  `xml:"NumeroGuia,attr" bson:"NoGuia"`
	LugarContrato                 *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty"`
	RfcTransportista              *string `xml:"RFCTransportista,attr" bson:"RfcTransportista,omitempty"`
	CodigoTransportista           string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista"`
	NumRegIdTribTransportista     *string `xml:"NumRegIdTribTranspor,attr" bson:"NumRegIdTribTransportista,omitempty"`
	ResidenciaFiscalTransportista *string `xml:"ResidenciaFiscalTranspor,attr" bson:"ResidenciaFiscalTransportista,omitempty"`
	NombreTransportista           *string `xml:"NombreTransportista,attr" bson:"NombreTransportista,omitempty"`
	RfcEmbarcador                 *string `xml:"RFCEmbarcador,attr" bson:"RfcEmbarcador,omitempty"`
	NumRegIdTribEmbarcador        *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarcador,omitempty"`
	ResidenciaFiscalEmbarcador    *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarcador,omitempty"`
	NombreEmbarcador              *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty"`
}

type TransporteFerroviarioCartaPorte10 struct {
	TipoServicio      string                     `xml:"TipoDeServicio,attr" bson:"TipoServicio"`
	NombreAseguradora *string                    `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro    *string                    `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	Concesionario     *string                    `xml:"Concesionario,attr" bson:"Concesionario,omitempty"`
	DerechoPaso       *[]DerechoPasoCartaPorte10 `xml:"DerechosDePaso" bson:"DerechoPaso,omitempty"`
	Carro             []CarroCartaPorte10        `xml:"Carro" bson:"Carro"`
}

type DerechoPasoCartaPorte10 struct {
	Tipo              string  `xml:"TipoDerechoDePaso,attr" bson:"Tipo"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte10 struct {
	Tipo           string                         `xml:"TipoCarro,attr" bson:"Tipo"`
	Matricula      string                         `xml:"MatriculaCarro,attr" bson:"Matricula"`
	Guia           string                         `xml:"GuiaCarro,attr" bson:"Guia"`
	ToneladasNetas float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetas"`
	Contenedor     *[]ContenedorCarroCartaPorte10 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte10 struct {
	Tipo                string  `xml:"TipoContenedor,attr" bson:"Tipo"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte10 struct {
	ClaveTransporte string                      `xml:"CveTransporte,attr" bson:"ClaveTransporte"`
	Operadores      *[]OperadorCartaPorte10     `xml:"Operadores>Operador" bson:"Operadores"`
	Propietarios    *[]PropietarioCartaPorte10  `xml:"Propietario" bson:"Propietario"`
	Arrendatarios   *[]ArrendatarioCartaPorte10 `xml:"Arrendatario" bson:"Arrendatarios"`
	Notificados     *[]NotificadoCartaPorte10   `xml:"Notificado" bson:"Notificado"`
}

type OperadorCartaPorte10 struct {
	Rfc              *string                `xml:"RFCOperador,attr" bson:"Rfc,omitempty"`
	NoLicencia       *string                `xml:"NumLicencia,attr" bson:"NoLicencia,omitempty"`
	Nombre           *string                `xml:"NombreOperador,attr" bson:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribOperador,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalOperador,attr" bson:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type PropietarioCartaPorte10 struct {
	Rfc              *string                `xml:"RFCPropietario,attr" bson:"Rfc,omitempty"`
	Nombre           *string                `xml:"NombrePropietario,attr" bson:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribPropietario,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalPropietario,attr" bson:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type ArrendatarioCartaPorte10 struct {
	Rfc              *string                `xml:"RFCArrendatario,attr" bson:"Rfc,omitempty"`
	Nombre           *string                `xml:"NombreArrendatario,attr" bson:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribArrendatario,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalArrendatario,attr" bson:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type NotificadoCartaPorte10 struct {
	Rfc              *string                `xml:"RFCNotificado,attr" bson:"Rfc,omitempty"`
	Nombre           *string                `xml:"NombreNotificado,attr" bson:"Nombre,omitempty"`
	NumRegIdTrib     *string                `xml:"NumRegIdTribNotificado,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                `xml:"ResidenciaFiscalNotificado,attr" bson:"ResidenciaFiscal,omitempty"`
	Domicilio        *DomicilioCartaPorte10 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

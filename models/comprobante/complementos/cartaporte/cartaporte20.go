package cartaporte

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type CartaPorte20 struct {
	Version string `xml:"Version,attr" bson:"Version"`

	TransporteInternacional   string `xml:"TranspInternac,attr" bson:"TransporteInternacional"`
	EsTransporteInternacional bool   `bson:"EsTransporteInternacional"`

	EntradaSalidaMercancia  *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty"`
	PaisOrigenDestino       *string                       `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida        *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty"`
	Ubicaciones             []UbicacionCartaPorte20       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias              MercanciasCartaPorte20        `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte        *FiguraTransporteCartaPorte20 `xml:"FiguraTransporte" bson:"FiguraTransporte"`
}

func (ccp *CartaPorte20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias CartaPorte20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*ccp = CartaPorte20(aux)
	ccp.EsTransporteInternacional = helpers.ResolveSatBoolean(ccp.TransporteInternacional)

	return nil
}

type UbicacionCartaPorte20 struct {
	TipoUbicacion                string                 `xml:"TipoUbicacion,attr" bson:"TipoUbicacion"`
	IdUbicacion                  *string                `xml:"IDUbicacion,attr" bson:"IdUbicacion,omitempty"`
	RfcRemitenteDestinatario     string                 `xml:"RFCRemitenteDestinatario,attr" bson:"RfcRemitenteDestinatario"`
	NombreRemitenteDestinatario  *string                `xml:"NombreRemitenteDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                 *string                `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal             *string                `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`
	NoEstacion                   *string                `xml:"NumEstacion,attr" bson:"NoEstacion,omitempty"`
	NombreEstacion               *string                `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico            *string                `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegadaString string                 `xml:"FechaHoraSalidaLlegada,attr" bson:"FechaHoraSalidaLlegadaString"`
	FechaHoraSalidaLlegada       time.Time              `bson:"FechaHoraSalidaLlegada"`
	TipoEstacion                 *string                `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty"`
	DistanciaRecorrida           *float64               `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty"`
	Domicilio                    *DomicilioCartaPorte20 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

func (u *UbicacionCartaPorte20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias UbicacionCartaPorte20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*u = UbicacionCartaPorte20(aux)
	fecha, errFecha := helpers.ParseDatetime(u.FechaHoraSalidaLlegadaString)
	if errFecha == nil {
		u.FechaHoraSalidaLlegada = fecha
	}
	return nil
}

type DomicilioCartaPorte20 struct {
	Calle        *string `xml:"Calle,attr" bson:"Calle,omitempty"`
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

type MercanciasCartaPorte20 struct {
	PesoBrutoTotal        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal"`
	UnidadPeso            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoNetoTotal         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumeroTotalMercancias float64                            `xml:"NumTotalMercancias,attr" bson:"NumeroTotalMercancias"`
	CargoPorTasacion      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty"`
	Mercancia             []MercanciaCartaPorte20            `xml:"Mercancia" bson:"Mercancia"`
	Autotransporte        *AutotransporteCartaPorte20        `xml:"Autotransporte" bson:"Autotrasporte,omitempty"`
	TransporteMaritimo    *TransporteMaritimoCartaPorte20    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty"`
	TransporteAereo       *TransporteAereoCartaPorte20       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty"`
	TransporteFerroviario *TransporteFerroviarioCartaPorte20 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte20 struct {
	BienesTransportados string  `xml:"BienesTransp,attr" bson:"BienesTransportados"`
	ClaveSTCC           *string `xml:"ClaveSTCC,attr" bson:"ClaveSTCC,omitempty"`
	Descripcion         string  `xml:"Descripcion,attr" bson:"Descripcion"`
	Cantidad            float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad         string  `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad              *string `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones         *string `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`

	MaterialPeligroso   *string `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	EsMaterialPeligroso *bool   `bson:"EsMaterialPeligroso,omitempty"`

	ClaveMaterialPeligroso *string                            `xml:"CveMaterialPeligroso,attr" bson:"ClaveMaterialPeligroso,omitempty"`
	Embalaje               *string                            `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripcionEmbalaje    *string                            `xml:"DescripEmbalaje,attr" bson:"DescripcionEmbalaje,omitempty"`
	PesoKg                 float64                            `xml:"PesoEnKg,attr" bson:"PesoKg"`
	ValorMercancia         *float64                           `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda                 *string                            `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria    *string                            `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UuidComercioExterior   *string                            `xml:"UUIDComercioExt,attr" bson:"UuidComercioExterior,omitempty"`
	Pedimentos             *[]PedimentosCartaPorte20          `xml:"Pedimentos" bson:"Pedimentos,omitempty"`
	GuiasIdentificacion    *[]GuiasIdentificacionCartaPorte20 `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty"`
	CantidadTransporta     *[]CantidadTransportaCartaPorte20  `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia       *DetalleMercanciaCartaPorte20      `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
}

func (m *MercanciaCartaPorte20) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias MercanciaCartaPorte20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*m = MercanciaCartaPorte20(aux)
	if m.MaterialPeligroso != nil {
		isMaterialPeligroso := helpers.ResolveSatBoolean(*aux.MaterialPeligroso)
		m.EsMaterialPeligroso = &isMaterialPeligroso
	}
	return nil
}

type PedimentosCartaPorte20 struct {
	Pedimento string `xml:"Pedimento,attr" bson:"Pedimento"`
}

type GuiasIdentificacionCartaPorte20 struct {
	Numero      string  `xml:"NumeroGuiaIdentificacion,attr" bson:"Numero"`
	Descripcion string  `xml:"DescripGuiaIdentificacion,attr" bson:"Descripcion"`
	Peso        float64 `xml:"PesoGuiaIdentificacion,attr" bson:"Peso"`
}

type CantidadTransportaCartaPorte20 struct {
	Cantidad         float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IdOrigen         string  `xml:"IDOrigen,attr" bson:"IdOrigen"`
	IdDestino        string  `xml:"IDDestino,attr" bson:"IdDestino"`
	ClavesTransporte *string `xml:"CvesTransporte,attr" bson:"ClavesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte20 struct {
	UnidadPeso   string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPeso"`
	PesoBruto    float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	PesoNeto     float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoTara     float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	NumeroPiezas *float64 `xml:"NumPiezas,attr" bson:"NumeroPiezas,omitempty"`
}

type AutotransporteCartaPorte20 struct {
	PermisoSct              string                              `xml:"PermSCT,attr" bson:"PermisoSct"`
	NoPermisoSct            string                              `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte20 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte20                 `xml:"Seguros" bson:"Seguros"`
	Remolques               *[]RemolqueCartaPorte20             `xml:"Remolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte20 struct {
	ConfiguracionVehicular string `xml:"ConfigVehicular,attr" bson:"Configuracion"`
	PlacaVm                string `xml:"PlacaVM,attr" bson:"PlacaVm"`
	AnioModeloVm           string `xml:"AnioModeloVM,attr" bson:"AnioModeloVm"`
}

type SegurosCartaPorte20 struct {
	AseguradoraResponsabilidadCivil string   `xml:"AseguraRespCivil,attr" bson:"AseguradoraResponsabilidadCivil"`
	PolizaResponsabilidadCivil      string   `xml:"PolizaRespCivil,attr" bson:"PolizaResponsabilidadCivil"`
	AseguradoraMedioAmbiente        *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguradoraMedioAmbiente,omitempty"`
	PolizaMedioAmbiente             *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedioAmbiente,omitempty"`
	AseguradoraCarga                *string  `xml:"AseguraCarga,attr" bson:"AseguradoraCarga,omitempty"`
	PolizaCarga                     *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty"`
	PrimaSeguro                     *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte20 struct {
	Subtipo string `xml:"SubTipoRem,attr" bson:"Subtipo"`
	Placa   string `xml:"Placa,attr" bson:"Placa"`
}

type TransporteMaritimoCartaPorte20 struct {
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
	Contenedor              []ContenedorCartaPorte20 `xml:"Contenedor" bson:"Contenedor"`
}

type ContenedorCartaPorte20 struct {
	Matricula  string  `xml:"MatriculaContenedor,attr" bson:"Matricula"`
	Tipo       string  `xml:"TipoContenedor,attr" bson:"Tipo"`
	NoPrecinto *string `xml:"NumPrecinto,attr" bson:"NoPrecinto,omitempty"`
}

type TransporteAereoCartaPorte20 struct {
	PermSct                    string  `xml:"PermSCT,attr" bson:"PermSct"`
	NoPermisoSct               string  `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct"`
	MatriculaAeronave          *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty"`
	NombreAseguradora          *string `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro             *string `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	NoGuia                     string  `xml:"NumeroGuia,attr" bson:"NoGuia"`
	LugarContrato              *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty"`
	CodigoTransportista        string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista"`
	RfcEmbarcador              *string `xml:"RFCEmbarcador,attr" bson:"RfcEmbarcador,omitempty"`
	NumRegIdTribEmbarcador     *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarcador,omitempty"`
	ResidenciaFiscalEmbarcador *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarcador,omitempty"`
	NombreEmbarcador           *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty"`
}

type TransporteFerroviarioCartaPorte20 struct {
	TipoServicio      string                     `xml:"TipoDeServicio,attr" bson:"TipoServicio"`
	TipoTrafico       string                     `xml:"TipoDeTrafico,attr" bson:"TipoTrafico"`
	NombreAseguradora *string                    `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro    *string                    `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	DerechoPaso       *[]DerechoPasoCartaPorte20 `xml:"DerechosDePaso" bson:"DerechoPaso,omitempty"`
	Carro             []CarroCartaPorte20        `xml:"Carro" bson:"Carro"`
}

type DerechoPasoCartaPorte20 struct {
	Tipo              string  `xml:"TipoDerechoDePaso,attr" bson:"Tipo"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte20 struct {
	Tipo           string                         `xml:"TipoCarro,attr" bson:"Tipo"`
	Matricula      string                         `xml:"MatriculaCarro,attr" bson:"Matricula"`
	Guia           string                         `xml:"GuiaCarro,attr" bson:"Guia"`
	ToneladasNetas float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetas"`
	Contenedor     *[]ContenedorCarroCartaPorte20 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte20 struct {
	Tipo                string  `xml:"TipoContenedor,attr" bson:"Tipo"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte20 struct {
	TiposFigura []TiposFiguraCartaPorte20 `xml:"TiposFigura" bson:"TiposFigura"`
}

type TiposFiguraCartaPorte20 struct {
	Tipo             string                          `xml:"TipoFigura,attr" bson:"Tipo"`
	Rfc              *string                         `xml:"RFCFigura,attr" bson:"Rfc,omitempty"`
	NoLicencia       *string                         `xml:"NumLicencia,attr" bson:"NoLicencia,omitempty"`
	Nombre           *string                         `xml:"NombreFigura,attr" bson:"Nombre,omitempty"`
	NumRegIdTrib     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscal,omitempty"`
	PartesTransporte *[]PartesTransporteCartaPorte20 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty"`
	Domicilio        *DomicilioCartaPorte20          `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte20 struct {
	Parte string `xml:"ParteTransporte,attr" bson:"Parte"`
}

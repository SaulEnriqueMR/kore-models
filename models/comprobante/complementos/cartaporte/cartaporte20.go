package cartaporte

import (
	"encoding/json"
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"strings"
	"time"
)

type CartaPorte20 struct {
	Version                   string                        `xml:"Version,attr" bson:"Version" json:"Version"`
	TransporteInternacional   string                        `xml:"TranspInternac,attr" bson:"TransporteInternacional" json:"TransporteInternacional"`
	EsTransporteInternacional bool                          `bson:"EsTransporteInternacional" json:"EsTransporteInternacional"`
	EntradaSalidaMercancia    *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty" json:"EntradaSalidaMercancia,omitempty"`
	PaisOrigenDestino         *string                       `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty" json:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida          *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty" json:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida   *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty" json:"TotalDistanciaRecorrida,omitempty"`
	Ubicaciones               []UbicacionCartaPorte20       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones" json:"Ubicaciones"`
	Mercancias                MercanciasCartaPorte20        `xml:"Mercancias" bson:"Mercancias" json:"Mercancias"`
	FiguraTransporte          *FiguraTransporteCartaPorte20 `xml:"FiguraTransporte" bson:"FiguraTransporte" json:"FiguraTransporte"`
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
	TipoUbicacion                string                 `xml:"TipoUbicacion,attr" bson:"TipoUbicacion" json:"TipoUbicacion"`
	IdUbicacion                  *string                `xml:"IDUbicacion,attr" bson:"IdUbicacion,omitempty" json:"IdUbicacion,omitempty"`
	RfcRemitenteDestinatario     string                 `xml:"RFCRemitenteDestinatario,attr" bson:"RfcRemitenteDestinatario" json:"RfcRemitenteDestinatario"`
	NombreRemitenteDestinatario  *string                `xml:"NombreRemitenteDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty" json:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                 *string                `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal             *string                `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	NoEstacion                   *string                `xml:"NumEstacion,attr" bson:"NoEstacion,omitempty" json:"NoEstacion,omitempty"`
	NombreEstacion               *string                `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty" json:"NombreEstacion,omitempty"`
	NavegacionTrafico            *string                `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty" json:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegadaString string                 `xml:"FechaHoraSalidaLlegada,attr" bson:"FechaHoraSalidaLlegadaString" json:"FechaHoraSalidaLlegadaString"`
	FechaHoraSalidaLlegada       time.Time              `bson:"FechaHoraSalidaLlegada" json:"FechaHoraSalidaLlegada"`
	TipoEstacion                 *string                `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty" json:"TipoEstacion,omitempty"`
	DistanciaRecorrida           *float64               `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty" json:"DistanciaRecorrida,omitempty"`
	Domicilio                    *DomicilioCartaPorte20 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
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

func (u *UbicacionCartaPorte20) UnmarshalJSON(data []byte) error {
	// Create an alias to avoid recursion
	type Alias UbicacionCartaPorte20
	var aux Alias

	// Unmarshal the XML into the alias
	if err := json.Unmarshal(data, &aux); err != nil {
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
	Calle        *string `xml:"Calle,attr" bson:"Calle,omitempty" json:"Calle,omitempty"`
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

type MercanciasCartaPorte20 struct {
	PesoBrutoTotal        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal" json:"PesoBrutoTotal"`
	UnidadPeso            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso" json:"UnidadPeso"`
	PesoNetoTotal         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty" json:"PesoNetoTotal,omitempty"`
	NumeroTotalMercancias float64                            `xml:"NumTotalMercancias,attr" bson:"NumeroTotalMercancias" json:"NumeroTotalMercancias"`
	CargoPorTasacion      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty" json:"CargoPorTasacion,omitempty"`
	Mercancia             []MercanciaCartaPorte20            `xml:"Mercancia" bson:"Mercancia" json:"Mercancia"`
	Autotransporte        *AutotransporteCartaPorte20        `xml:"Autotransporte" bson:"Autotrasporte,omitempty" json:"Autotrasporte,omitempty"`
	TransporteMaritimo    *TransporteMaritimoCartaPorte20    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty" json:"TransporteMaritimo,omitempty"`
	TransporteAereo       *TransporteAereoCartaPorte20       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty" json:"TransporteAereo,omitempty"`
	TransporteFerroviario *TransporteFerroviarioCartaPorte20 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty" json:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte20 struct {
	BienesTransportados    string                             `xml:"BienesTransp,attr" bson:"BienesTransportados" json:"BienesTransportados"`
	ClaveStcc              *string                            `xml:"ClaveSTCC,attr" bson:"ClaveStcc,omitempty" json:"ClaveStcc,omitempty"`
	Descripcion            string                             `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	Cantidad               float64                            `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	ClaveUnidad            string                             `xml:"ClaveUnidad,attr" bson:"ClaveUnidad" json:"ClaveUnidad"`
	Unidad                 *string                            `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Dimensiones            *string                            `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty" json:"Dimensiones,omitempty"`
	MaterialPeligroso      *string                            `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty" json:"MaterialPeligroso,omitempty"`
	EsMaterialPeligroso    *bool                              `bson:"EsMaterialPeligroso,omitempty" json:"EsMaterialPeligroso,omitempty"`
	ClaveMaterialPeligroso *string                            `xml:"CveMaterialPeligroso,attr" bson:"ClaveMaterialPeligroso,omitempty" json:"ClaveMaterialPeligroso,omitempty"`
	Embalaje               *string                            `xml:"Embalaje,attr" bson:"Embalaje,omitempty" json:"Embalaje,omitempty"`
	DescripcionEmbalaje    *string                            `xml:"DescripEmbalaje,attr" bson:"DescripcionEmbalaje,omitempty" json:"DescripcionEmbalaje,omitempty"`
	PesoKg                 float64                            `xml:"PesoEnKg,attr" bson:"PesoKg" json:"PesoKg"`
	ValorMercancia         *float64                           `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty" json:"ValorMercancia,omitempty"`
	Moneda                 *string                            `xml:"Moneda,attr" bson:"Moneda,omitempty" json:"Moneda,omitempty"`
	FraccionArancelaria    *string                            `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty" json:"FraccionArancelaria,omitempty"`
	UUIDComercioExt        *string                            `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty" json:"UUIDComercioExt,omitempty"`
	UuidComercioExterior   *string                            `bson:"UuidComercioExterior,omitempty" json:"UuidComercioExterior,omitempty"`
	Pedimentos             *[]PedimentosCartaPorte20          `xml:"Pedimentos" bson:"Pedimentos,omitempty" json:"Pedimentos,omitempty"`
	GuiasIdentificacion    *[]GuiasIdentificacionCartaPorte20 `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty" json:"GuiasIdentificacion,omitempty"`
	CantidadTransporta     *[]CantidadTransportaCartaPorte20  `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty" json:"CantidadTransporta,omitempty"`
	DetalleMercancia       *DetalleMercanciaCartaPorte20      `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty" json:"DetalleMercancia,omitempty"`
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

	if m.UUIDComercioExt != nil && *m.UUIDComercioExt != "" {
		uuidComercioExterior := strings.ToUpper(*m.UUIDComercioExt)
		m.UuidComercioExterior = &uuidComercioExterior
	}

	return nil
}

type PedimentosCartaPorte20 struct {
	Pedimento string `xml:"Pedimento,attr" bson:"Pedimento" json:"Pedimento"`
}

type GuiasIdentificacionCartaPorte20 struct {
	Numero      string  `xml:"NumeroGuiaIdentificacion,attr" bson:"Numero" json:"Numero"`
	Descripcion string  `xml:"DescripGuiaIdentificacion,attr" bson:"Descripcion" json:"Descripcion"`
	Peso        float64 `xml:"PesoGuiaIdentificacion,attr" bson:"Peso" json:"Peso"`
}

type CantidadTransportaCartaPorte20 struct {
	Cantidad         float64 `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	IdOrigen         string  `xml:"IDOrigen,attr" bson:"IdOrigen" json:"IdOrigen"`
	IdDestino        string  `xml:"IDDestino,attr" bson:"IdDestino" json:"IdDestino"`
	ClavesTransporte *string `xml:"CvesTransporte,attr" bson:"ClavesTransporte,omitempty" json:"ClavesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte20 struct {
	UnidadPeso   string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPeso" json:"UnidadPeso"`
	PesoBruto    float64  `xml:"PesoBruto,attr" bson:"PesoBruto" json:"PesoBruto"`
	PesoNeto     float64  `xml:"PesoNeto,attr" bson:"PesoNeto" json:"PesoNeto"`
	PesoTara     float64  `xml:"PesoTara,attr" bson:"PesoTara" json:"PesoTara"`
	NumeroPiezas *float64 `xml:"NumPiezas,attr" bson:"NumeroPiezas,omitempty" json:"NumeroPiezas,omitempty"`
}

type AutotransporteCartaPorte20 struct {
	PermisoSct              string                              `xml:"PermSCT,attr" bson:"PermisoSct" json:"PermisoSct"`
	NoPermisoSct            string                              `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct" json:"NoPermisoSct"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte20 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular" json:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte20                 `xml:"Seguros" bson:"Seguros" json:"Seguros"`
	Remolques               *[]RemolqueCartaPorte20             `xml:"Remolques>Remolque" bson:"Remolques,omitempty" json:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte20 struct {
	ConfiguracionVehicular string `xml:"ConfigVehicular,attr" bson:"ConfiguracionVehicular" json:"ConfiguracionVehicular"`
	PlacaVm                string `xml:"PlacaVM,attr" bson:"PlacaVm" json:"PlacaVm"`
	AnioModeloVm           string `xml:"AnioModeloVM,attr" bson:"AnioModeloVm" json:"AnioModeloVm"`
}

type SegurosCartaPorte20 struct {
	AseguradoraResponsabilidadCivil string   `xml:"AseguraRespCivil,attr" bson:"AseguradoraResponsabilidadCivil" json:"AseguradoraResponsabilidadCivil"`
	PolizaResponsabilidadCivil      string   `xml:"PolizaRespCivil,attr" bson:"PolizaResponsabilidadCivil" json:"PolizaResponsabilidadCivil"`
	AseguradoraMedioAmbiente        *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguradoraMedioAmbiente,omitempty" json:"AseguradoraMedioAmbiente,omitempty"`
	PolizaMedioAmbiente             *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedioAmbiente,omitempty" json:"PolizaMedioAmbiente,omitempty"`
	AseguradoraCarga                *string  `xml:"AseguraCarga,attr" bson:"AseguradoraCarga,omitempty" json:"AseguradoraCarga,omitempty"`
	PolizaCarga                     *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty" json:"PolizaCarga,omitempty"`
	PrimaSeguro                     *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty" json:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte20 struct {
	Subtipo string `xml:"SubTipoRem,attr" bson:"Subtipo" json:"Subtipo"`
	Placa   string `xml:"Placa,attr" bson:"Placa" json:"Placa"`
}

type TransporteMaritimoCartaPorte20 struct {
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
	NumCertItc              string                   `xml:"NumCertITC,attr" bson:"NumCertItc" json:"NumCertItc"`
	Eslora                  *float64                 `xml:"Eslora,attr" bson:"Eslora,omitempty" json:"Eslora,omitempty"`
	Manga                   *float64                 `xml:"Manga,attr" bson:"Manga,omitempty" json:"Manga,omitempty"`
	Calado                  *float64                 `xml:"Calado,attr" bson:"Calado,omitempty" json:"Calado,omitempty"`
	LineaNaviera            *string                  `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty" json:"LineaNaviera,omitempty"`
	NombreAgenteNaviero     string                   `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero" json:"NombreAgenteNaviero"`
	NoAutorizacionNaviero   string                   `xml:"NumAutorizacionNaviero,attr" bson:"NoAutorizacionNaviero" json:"NoAutorizacionNaviero"`
	NoViaje                 *string                  `xml:"NumViaje,attr" bson:"NoViaje,omitempty" json:"NoViaje,omitempty"`
	NoConocimientoEmbarque  *string                  `xml:"NumConocEmbarc,attr" bson:"NoConocimientoEmbarque,omitempty" json:"NoConocimientoEmbarque,omitempty"`
	Contenedor              []ContenedorCartaPorte20 `xml:"Contenedor" bson:"Contenedor" json:"Contenedor"`
}

type ContenedorCartaPorte20 struct {
	Matricula  string  `xml:"MatriculaContenedor,attr" bson:"Matricula" json:"Matricula"`
	Tipo       string  `xml:"TipoContenedor,attr" bson:"Tipo" json:"Tipo"`
	NoPrecinto *string `xml:"NumPrecinto,attr" bson:"NoPrecinto,omitempty" json:"NoPrecinto,omitempty"`
}

type TransporteAereoCartaPorte20 struct {
	PermSct                    string  `xml:"PermSCT,attr" bson:"PermSct" json:"PermSct"`
	NoPermisoSct               string  `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct" json:"NoPermisoSct"`
	MatriculaAeronave          *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty" json:"MatriculaAeronave,omitempty"`
	NombreAseguradora          *string `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty" json:"NombreAseguradora,omitempty"`
	NoPolizaSeguro             *string `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty" json:"NoPolizaSeguro,omitempty"`
	NoGuia                     string  `xml:"NumeroGuia,attr" bson:"NoGuia" json:"NoGuia"`
	LugarContrato              *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty" json:"LugarContrato,omitempty"`
	CodigoTransportista        string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista" json:"CodigoTransportista"`
	RfcEmbarcador              *string `xml:"RFCEmbarcador,attr" bson:"RfcEmbarcador,omitempty" json:"RfcEmbarcador,omitempty"`
	NumRegIdTribEmbarcador     *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarcador,omitempty" json:"NumRegIdTribEmbarcador,omitempty"`
	ResidenciaFiscalEmbarcador *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarcador,omitempty" json:"ResidenciaFiscalEmbarcador,omitempty"`
	NombreEmbarcador           *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty" json:"NombreEmbarcador,omitempty"`
}

type TransporteFerroviarioCartaPorte20 struct {
	TipoServicio      string                     `xml:"TipoDeServicio,attr" bson:"TipoServicio" json:"TipoServicio"`
	TipoTrafico       string                     `xml:"TipoDeTrafico,attr" bson:"TipoTrafico" json:"TipoTrafico"`
	NombreAseguradora *string                    `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty" json:"NombreAseguradora,omitempty"`
	NoPolizaSeguro    *string                    `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty" json:"NoPolizaSeguro,omitempty"`
	DerechoPaso       *[]DerechoPasoCartaPorte20 `xml:"DerechosDePaso" bson:"DerechoPaso,omitempty" json:"DerechoPaso,omitempty"`
	Carro             []CarroCartaPorte20        `xml:"Carro" bson:"Carro" json:"Carro"`
}

type DerechoPasoCartaPorte20 struct {
	Tipo              string  `xml:"TipoDerechoDePaso,attr" bson:"Tipo" json:"Tipo"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado" json:"KilometrajePagado"`
}

type CarroCartaPorte20 struct {
	Tipo           string                         `xml:"TipoCarro,attr" bson:"Tipo" json:"Tipo"`
	Matricula      string                         `xml:"MatriculaCarro,attr" bson:"Matricula" json:"Matricula"`
	Guia           string                         `xml:"GuiaCarro,attr" bson:"Guia" json:"Guia"`
	ToneladasNetas float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetas" json:"ToneladasNetas"`
	Contenedor     *[]ContenedorCarroCartaPorte20 `xml:"Contenedor" bson:"Contenedor,omitempty" json:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte20 struct {
	Tipo                string  `xml:"TipoContenedor,attr" bson:"Tipo" json:"Tipo"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio" json:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia" json:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte20 struct {
	TiposFigura []TiposFiguraCartaPorte20 `xml:"TiposFigura" bson:"TiposFigura" json:"TiposFigura"`
}

type TiposFiguraCartaPorte20 struct {
	Tipo             string                          `xml:"TipoFigura,attr" bson:"Tipo" json:"Tipo"`
	Rfc              *string                         `xml:"RFCFigura,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
	NoLicencia       *string                         `xml:"NumLicencia,attr" bson:"NoLicencia,omitempty" json:"NoLicencia,omitempty"`
	Nombre           *string                         `xml:"NombreFigura,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	NumRegIdTrib     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	PartesTransporte *[]PartesTransporteCartaPorte20 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty" json:"PartesTransporte,omitempty"`
	Domicilio        *DomicilioCartaPorte20          `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte20 struct {
	Parte string `xml:"ParteTransporte,attr" bson:"Parte" json:"Parte"`
}

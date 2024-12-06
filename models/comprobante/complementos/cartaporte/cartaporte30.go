package cartaporte

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type CartaPorte30 struct {
	Version                   string                        `xml:"Version,attr" bson:"Version"`
	IdCCP                     string                        `xml:"IdCCP,attr" bson:"IdCCP"`
	IdCcp                     string                        `bson:"IdCcp"`
	TransporteInternacional   string                        `xml:"TranspInternac,attr" bson:"TransporteInternacional"`
	EsTransporteInternacional bool                          `bson:"EsTransporteInternacional"`
	RegimenAduanero           *string                       `xml:"RegimenAduanero,attr" bson:"RegimenAduanero,omitempty"`
	EntradaSalidaMercancia    *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty"`
	PaisOrigenDestino         *string                       `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida          *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida   *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty"`
	RegistroIstmo             *string                       `xml:"RegistroISTMO,attr" bson:"RegistroIstmo,omitempty"`
	UbicacionPoloOrigen       *string                       `xml:"UbicacionPoloOrigen,attr" bson:"UbicacionPoloOrigen,omitempty"`
	UbicacionPoloDestino      *string                       `xml:"UbicacionPoloDestino,attr" bson:"UbicacionPoloDestino,omitempty"`
	Ubicaciones               []UbicacionCartaPorte30       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias                MercanciasCartaPorte30        `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte          *FiguraTransporteCartaPorte30 `xml:"FiguraTransporte" bson:"FiguraTransporte,omitempty"`
}

func (ccp *CartaPorte30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias CartaPorte30
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*ccp = CartaPorte30(aux)
	ccp.EsTransporteInternacional = helpers.ResolveSatBoolean(ccp.TransporteInternacional)

	ccp.IdCcp = strings.ToUpper(aux.IdCCP)

	return nil
}

type UbicacionCartaPorte30 struct {
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
	Domicilio                    *DomicilioCartaPorte30 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

func (u *UbicacionCartaPorte30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias UbicacionCartaPorte30
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*u = UbicacionCartaPorte30(aux)
	fecha, errFecha := helpers.ParseDatetime(u.FechaHoraSalidaLlegadaString)
	if errFecha == nil {
		u.FechaHoraSalidaLlegada = fecha
	}
	return nil
}

type DomicilioCartaPorte30 struct {
	Calle        *string `xml:"Calle,attr" bson:"Calle,omitempty"`
	NoExterior   *string `xml:"NumeroExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"NumeroInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`
	Localidad    *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`
	Municipio    *string `xml:"Municipio,attr" bson:"Municipio,omitempty"`
	Estado       string  `xml:"Estado,attr" bson:"Estado,omitempty"`
	Pais         string  `xml:"Pais,attr" bson:"Pais,omitempty"`
	CodigoPostal string  `xml:"CodigoPostal,attr" bson:"CodigoPostal,omitempty"`
}

type MercanciasCartaPorte30 struct {
	PesoBrutoTotal                        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal"`
	UnidadPeso                            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoNetoTotal                         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumeroTotalMercancias                 float64                            `xml:"NumTotalMercancias,attr" bson:"NumeroTotalMercancias"`
	CargoPorTasacion                      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty"`
	LogisticaInversaRecoleccionDevolucion *string                            `xml:"LogisticaInversaRecoleccionDevolucion,attr" bson:"LogisticaInversaRecoleccionDevolucion,omitempty"`
	Mercancia                             []MercanciaCartaPorte30            `xml:"Mercancia" bson:"Mercancia"`
	Autotransporte                        *AutotransporteCartaPorte30        `xml:"Autotransporte" bson:"Autotransporte,omitempty"`
	TransporteMaritimo                    *TransporteMaritimoCartaPorte30    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty"`
	TransporteAereo                       *TransporteAereoCartaPorte30       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty"`
	TransporteFerroviario                 *TransporteFerroviarioCartaPorte30 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte30 struct {
	BienesTransportados string  `xml:"BienesTransp,attr" bson:"BienesTransportados"`
	ClaveStcc           *string `xml:"ClaveSTCC,attr" bson:"ClaveStcc,omitempty"`
	Descripcion         string  `xml:"Descripcion,attr" bson:"Descripcion"`
	Cantidad            float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad         string  `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad              *string `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones         *string `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`

	MaterialPeligroso   *string `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	EsMaterialPeligroso *bool   `bson:"EsMaterialPeligroso,omitempty"`

	ClaveMaterialPeligroso                 *string                              `xml:"CveMaterialPeligroso,attr" bson:"ClaveMaterialPeligroso,omitempty"`
	Embalaje                               *string                              `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripcionEmbalaje                    *string                              `xml:"DescripEmbalaje,attr" bson:"DescripcionEmbalaje,omitempty"`
	SectorCofepris                         *string                              `xml:"SectorCOFEPRIS,attr" bson:"SectorCofepris,omitempty"`
	NombreIngredienteActivo                *string                              `xml:"NombreIngredienteActivo,attr" bson:"NombreIngredienteActivo,omitempty"`
	NombreQuimico                          *string                              `xml:"NomQuimico,attr" bson:"NombreQuimico,omitempty"`
	DenominacionGenericaProducto           *string                              `xml:"DenominacionGenericaProd,attr" bson:"DenominacionGenericaProducto,omitempty"`
	DenominacionDistintivaProducto         *string                              `xml:"DenominacionDistintivaProd,attr" bson:"DenominacionDistintivaProducto,omitempty"`
	Fabricante                             *string                              `xml:"Fabricante,attr" bson:"Fabricante,omitempty"`
	FechaCaducidadString                   *string                              `xml:"FechaCaducidad,attr" bson:"FechaCaducidadString,omitempty"`
	FechaCaducidad                         *time.Time                           `bson:"FechaCaducidad,omitempty"`
	LoteMedicamento                        *string                              `xml:"LoteMedicamento,attr" bson:"LoteMedicamento,omitempty"`
	FormaFarmaceutica                      *string                              `xml:"FormaFarmaceutica,attr" bson:"FormaFarmaceutica,omitempty"`
	CondicionesEspecialesTransporte        *string                              `xml:"CondicionesEspTransp,attr" bson:"CondicionesEspecialesTransporte,omitempty"`
	RegistroSanitarioFolioAutorizacion     *string                              `xml:"RegistroSanitarioFolioAutorizacion,attr" bson:"RegistroSanitarioFolioAutorizacion,omitempty"`
	PermisoImportacion                     *string                              `xml:"PermisoImportacion,attr" bson:"PermisoImportacion,omitempty"`
	FolioImportacionVucem                  *string                              `xml:"FolioImpoVUCEM,attr" bson:"FolioImportacionVucem,omitempty"`
	NoCas                                  *string                              `xml:"NumCAS,attr" bson:"NoCas,omitempty"`
	RazonSocialEmpresaImportadora          *string                              `xml:"RazonSocialEmpImp,attr" bson:"RazonSocialEmpresaImportadora,omitempty"`
	NoRegistroSanitarioPlaguicidasCofepris *string                              `xml:"NumRegSanPlagCOFEPRIS,attr" bson:"NoRegistroSanitarioPlaguicidasCofepris,omitempty"`
	DatosFabricante                        *string                              `xml:"DatosFabricante,attr" bson:"DatosFabricante,omitempty"`
	DatosFormulador                        *string                              `xml:"DatosFormulador,attr" bson:"DatosFormulador,omitempty"`
	DatosMaquilador                        *string                              `xml:"DatosMaquilador,attr" bson:"DatosMaquilador,omitempty"`
	UsoAutorizado                          *string                              `xml:"UsoAutorizado,attr" bson:"UsoAutorizado,omitempty"`
	PesoKg                                 float64                              `xml:"PesoEnKg,attr" bson:"PesoKg"`
	ValorMercancia                         *float64                             `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda                                 *string                              `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria                    *string                              `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UUIDComercioExt                        *string                              `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty"`
	UuidComercioExterior                   *string                              `bson:"UuidComercioExterior,omitempty"`
	TipoMateria                            *string                              `xml:"TipoMateria,attr" bson:"TipoMateria,omitempty"`
	DescripcionMateria                     *string                              `xml:"DescripcionMateria,attr" bson:"DescripcionMateria,omitempty"`
	DocumentacionAduanera                  *[]DocumentacionAduaneraCartaPorte30 `xml:"DocumentacionAduanera" bson:"DocumentacionAduanera,omitempty"`
	GuiasIdentificacion                    *[]GuiasIdentificacionCartaPorte30   `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty"`
	CantidadTransporta                     *[]CantidadTransportaCartaPorte30    `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia                       *DetalleMercanciaCartaPorte30        `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
}

func (m *MercanciaCartaPorte30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias MercanciaCartaPorte30
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*m = MercanciaCartaPorte30(aux)
	if m.MaterialPeligroso != nil {
		isMaterialPeligroso := helpers.ResolveSatBoolean(*aux.MaterialPeligroso)
		m.EsMaterialPeligroso = &isMaterialPeligroso
	}
	if m.FechaCaducidadString != nil && *m.FechaCaducidadString != "" {
		fecha, errFecha := helpers.ParseDatetime(*m.FechaCaducidadString)
		if errFecha == nil {
			m.FechaCaducidad = &fecha
		}
	}
	if m.UUIDComercioExt != nil && *m.UUIDComercioExt != "" {
		uuidComercioExterior := strings.ToUpper(*m.UUIDComercioExt)
		m.UuidComercioExterior = &uuidComercioExterior
	}
	return nil
}

type DocumentacionAduaneraCartaPorte30 struct {
	TipoDocumento          string  `xml:"TipoDocumento,attr" bson:"TipoDocumento"`
	NumeroPedimento        *string `xml:"NumPedimento,attr" bson:"NumeroPedimento,omitemtpy"`
	IdentificadorDocumento *string `xml:"IdentDocAduanero,attr" bson:"IdentificadorDocumento,omitempty"`
	RfcImportador          *string `xml:"RFCImpo,attr" bson:"RfcImportador,omitempty"`
}

type GuiasIdentificacionCartaPorte30 struct {
	Numero      string  `xml:"NumeroGuiaIdentificacion,attr" bson:"NumeroGuiaIdentificacion"`
	Descripcion string  `xml:"DescripGuiaIdentificacion,attr" bson:"DescripGuiaIdentificacion"`
	Peso        float64 `xml:"PesoGuiaIdentificacion,attr" bson:"PesoGuiaIdentificacion"`
}

type CantidadTransportaCartaPorte30 struct {
	Cantidad         float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IdOrigen         string  `xml:"IDOrigen,attr" bson:"IdOrigen"`
	IdDestino        string  `xml:"IDDestino,attr" bson:"IdDestino"`
	ClavesTransporte *string `xml:"CvesTransporte,attr" bson:"ClavesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte30 struct {
	UnidadPeso   string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPeso"`
	PesoTara     float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	PesoNeto     float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoBruto    float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	NumeroPiezas *float64 `xml:"NumPiezas,attr" bson:"NumeroPiezas,omitempty"`
}

type AutotransporteCartaPorte30 struct {
	PermisoSct              string                              `xml:"PermSCT,attr" bson:"PermisoSct"`
	NoPermisoSct            string                              `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte30 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte30                 `xml:"Seguros" bson:"Seguros"`
	Remolques               *[]RemolqueCartaPorte30             `xml:"Remolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte30 struct {
	Configuracion string  `xml:"ConfigVehicular,attr" bson:"Configuracion"`
	PesoBruto     float64 `xml:"PesoBrutoVehicular,attr" bson:"PesoBruto"`
	PlacaVm       string  `xml:"PlacaVM,attr" bson:"PlacaVm"`
	AnioModeloVm  float64 `xml:"AnioModeloVM,attr" bson:"AnioModeloVm"`
}

type SegurosCartaPorte30 struct {
	AseguradoraResponsabilidadCivil string   `xml:"AseguraRespCivil,attr" bson:"AseguradoraResponsabilidadCivil"`
	PolizaResponsabilidadCivil      string   `xml:"PolizaRespCivil,attr" bson:"PolizaResponsabilidadCivil"`
	AseguradoraMedioAmbiente        *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguradoraMedioAmbiente,omitempty"`
	PolizaMedioAmbiente             *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedioAmbiente,omitempty"`
	AseguradoraCarga                *string  `xml:"AseguraCarga,attr" bson:"AseguradoraCarga,omitempty"`
	PolizaCarga                     *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty"`
	PrimaSeguro                     *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte30 struct {
	Subtipo string `xml:"SubTipoRem,attr" bson:"Subtipo"`
	Placa   string `xml:"Placa,attr" bson:"Placa"`
}

type TransporteMaritimoCartaPorte30 struct {
	PermisoSct                *string                    `xml:"PermSCT,attr" bson:"PermisoSct,omitempty"`
	NoPermisoSct              *string                    `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct,omitempty"`
	NombreAseguradora         *string                    `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro            *string                    `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	TipoEmbarcacion           string                     `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion"`
	Matricula                 string                     `xml:"Matricula,attr" bson:"Matricula"`
	NoOmi                     string                     `xml:"NumeroOMI,attr" bson:"NoOmi"`
	AnioEmbarcacion           *float64                   `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty"`
	NombreEmbarcacion         *string                    `xml:"NombreEmbarc,attr" bson:"NombreEmbarcacion,omitempty"`
	NacionalidadEmbarcacion   string                     `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarcacion"`
	UnidadesArqueoBruto       float64                    `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesArqueoBruto"`
	TipoCarga                 string                     `xml:"TipoCarga,attr" bson:"TipoCarga"`
	Eslora                    *float64                   `xml:"Eslora,attr" bson:"Eslora,omitempty"`
	Manga                     *float64                   `xml:"Manga,attr" bson:"Manga,omitempty"`
	Calado                    *float64                   `xml:"Calado,attr" bson:"Calado,omitempty"`
	Puntal                    *float64                   `xml:"Puntal,attr" bson:"Puntal,omitempty"`
	LineaNaviera              *string                    `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty"`
	NombreAgenteNaviero       string                     `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero"`
	NoAutorizacionNaviero     string                     `xml:"NumAutorizacionNaviero,attr" bson:"NoAutorizacionNaviero"`
	NoViaje                   *string                    `xml:"NumViaje,attr" bson:"NoViaje,omitempty"`
	NoConocimientoEmbarque    *string                    `xml:"NumConocEmbarc,attr" bson:"NoConocimientoEmbarque,omitempty"`
	PermisoTemporalNavegacion *string                    `xml:"PermisoTempNavegacion,attr" bson:"PermisoTemporalNavegacion,omitempty"`
	Contenedor                *[]ContenedorCartaPorte30  `xml:"Contenedor" bson:"Contenedor,omitempty"`
	RemolquesCcp              *[]RemolqueCCPCartaPorte30 `xml:"RemolquesCCP>RemolqueCCP" bson:"RemolquesCcp,omitempty"`
}

type ContenedorCartaPorte30 struct {
	Tipo                  string  `xml:"TipoContenedor,attr" bson:"Tipo"`
	Matricula             *string `xml:"MatriculaContenedor,attr" bson:"Matricula,omitempty"` // Cifrado
	NoPrecinto            *string `xml:"NumPrecinto,attr" bson:"NoPrecinto,omitempty"`
	IdCcpRelacionado      *string `xml:"IdCCPRelacionado,attr" bson:"IdCcpRelacionado,omitempty"`
	PlacaVmCcp            *string `xml:"PlacaVMCCP,attr" bson:"PlacaVmCcp,omitempty"` // Cifrado
	FechaCertificacionCcp *string `xml:"FechaCertificacionCCP,attr" bson:"FechaCertificacionCcp,omitempty"`
}

type RemolqueCCPCartaPorte30 struct {
	Subtipo string `xml:"SubTipoRemCCP,attr" bson:"Subtipo"`
	Placa   string `xml:"PlacaCCP,attr" bson:"Placa"` // Cifrado
}

type TransporteAereoCartaPorte30 struct {
	PermisoSct                 string  `xml:"PermSCT,attr" bson:"PermisoSct"`
	NoPermisoSct               string  `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct"`
	MatriculaAeronave          *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty"`
	NombreAseguradora          *string `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro             *string `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	NoGuia                     string  `xml:"NumeroGuia,attr" bson:"NoGuia"`
	RfcEmbarcador              *string `xml:"RFCEmbarcador,attr" bson:"RfcEmbarcador,omitempty"`
	ResidenciaFiscalEmbarcador *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarcador,omitempty"`
	NumRegIdTribEmbarcador     *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarcador,omitempty"`
	NombreEmbarcador           *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty"`
	LugarContrato              *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty"`
	CodigoTransportista        string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista"`
}

type TransporteFerroviarioCartaPorte30 struct {
	TipoTrafico       string                        `xml:"TipoDeTrafico,attr" bson:"TipoTrafico"`
	TipoServicio      string                        `xml:"TipoDeServicio,attr" bson:"TipoServicio"`
	NoPolizaSeguro    *string                       `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	NombreAseguradora *string                       `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	DerechosPaso      *[]DerechosDePasoCartaPorte30 `xml:"DerechosDePaso" bson:"DerechosPaso,omitempty"`
	Carro             []CarroCartaPorte30           `xml:"Carro" bson:"Carro"`
}

type DerechosDePasoCartaPorte30 struct {
	Tipo              string  `xml:"TipoDerechoDePaso,attr" bson:"Tipo"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte30 struct {
	ToneladasNetas float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetas"`
	Tipo           string                         `xml:"TipoCarro,attr" bson:"Tipo"`
	Matricula      string                         `xml:"MatriculaCarro,attr" bson:"Matricula"`
	Guia           string                         `xml:"GuiaCarro,attr" bson:"Guia"`
	Contenedor     *[]ContenedorCarroCartaPorte30 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte30 struct {
	Tipo                string  `xml:"TipoContenedor,attr" bson:"Tipo"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
}

type FiguraTransporteCartaPorte30 struct {
	TiposFigura []TiposFiguraCartaPorte30 `xml:"TiposFigura" bson:"TiposFigura"`
}

type TiposFiguraCartaPorte30 struct {
	Tipo             string                          `xml:"TipoFigura,attr" bson:"Tipo"`
	Rfc              *string                         `xml:"RFCFigura,attr" bson:"Rfc,omitempty"`                           // Cifrado
	ResidenciaFiscal *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscal,omitempty"` // Cifrado
	NumRegIdTrib     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTrib,omitempty"`         // Cifrado
	NoLicencia       *string                         `xml:"NumLicencia,attr" bson:"NoLicencia,omitempty"`                  // Cifrado
	Nombre           string                          `xml:"NombreFigura,attr" bson:"Nombre"`
	PartesTransporte *[]PartesTransporteCartaPorte30 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty"`
	Domicilio        *DomicilioCartaPorte30          `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte30 struct {
	Parte string `xml:"ParteTransporte,attr" bson:"Parte"`
}

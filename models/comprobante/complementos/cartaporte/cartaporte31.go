package cartaporte

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type CartaPorte31 struct {
	// Attr
	Version string `xml:"Version,attr" bson:"Version"`
	IdCcp   string `xml:"IdCCP,attr" bson:"IdCcp"`

	TransporteInternacional   string `xml:"TranspInternac,attr" bson:"TransporteInternacional"`
	EsTransporteInternacional bool   `bson:"EsTransporteInternacional"`

	EntradaSalidaMercancia  *string  `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty"`
	PaisOrigenDestino       *string  `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida        *string  `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida *float64 `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty"`
	RegistroIstmo           *string  `xml:"RegistroISTMO,attr" bson:"RegistroIstmo,omitempty"`
	UbicacionPoloOrigen     *string  `xml:"UbicacionPoloOrigen,attr" bson:"UbicacionPoloOrigen,omitempty"`
	UbicacionPoloDestino    *string  `xml:"UbicacionPoloDestino,attr" bson:"UbicacionPoloDestino,omitempty"`

	// Nodos
	RegimenesAduaneros *[]RegimenAduaneroCartaPorte31 `xml:"RegimenesAduaneros>RegimenAduaneroCCP" bson:"RegimenesAduaneros,omitempty"`
	Ubicaciones        []UbicacionCartaPorte31        `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias         MercanciasCartaPorte31         `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte   *FiguraTransporteCartaPorte31  `xml:"FiguraTransporte" bson:"FiguraTransporte,omitempty"`
}

func (ccp *CartaPorte31) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias CartaPorte31
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*ccp = CartaPorte31(aux)
	ccp.EsTransporteInternacional = helpers.ResolveSatBoolean(ccp.TransporteInternacional)

	return nil
}

type RegimenAduaneroCartaPorte31 struct {
	RegimenAduanero string `xml:"RegimenAduanero,attr" bson:"RegimenAduanero"`
}

type UbicacionCartaPorte31 struct {
	// Attr
	TipoUbicacion               string  `xml:"TipoUbicacion,attr" bson:"TipoUbicacion"`
	IdUbicacion                 *string `xml:"IDUbicacion,attr" bson:"IdUbicacion,omitempty"`
	RfcRemitenteDestinatario    string  `xml:"RFCRemitenteDestinatario,attr" bson:"RfcRemitenteDestinatario"`
	NombreRemitenteDestinatario *string `xml:"NombreRemitenteDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                *string `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal            *string `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`
	NoEstacion                  *string `xml:"NumEstacion,attr" bson:"NoEstacion,omitempty"`
	NombreEstacion              *string `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico           *string `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`

	FechaHoraSalidaLlegadaString string    `xml:"FechaHoraSalidaLlegada,attr"`
	FechaHoraSalidaLlegada       time.Time `bson:"FechaHoraSalidaLlegada"`

	TipoEstacion       *string  `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty"`
	DistanciaRecorrida *float64 `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty"`
	// Nodos
	Domicilio *DomicilioCartaPorte31 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

func (u *UbicacionCartaPorte31) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias UbicacionCartaPorte31
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*u = UbicacionCartaPorte31(aux)
	fecha, errFecha := helpers.ParseDatetime(u.FechaHoraSalidaLlegadaString)
	if errFecha == nil {
		u.FechaHoraSalidaLlegada = fecha
	}
	return nil
}

type DomicilioCartaPorte31 struct {
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

type MercanciasCartaPorte31 struct {
	// Attr
	PesoBrutoTotal                        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal"`
	UnidadPeso                            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoNetoTotal                         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumeroTotalMercancias                 float64                            `xml:"NumTotalMercancias,attr" bson:"NumeroTotalMercancias"`
	CargoPorTasacion                      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty"`
	LogisticaInversaRecoleccionDevolucion *string                            `xml:"LogisticaInversaRecoleccionDevolucion,attr" bson:"LogisticaInversaRecoleccionDevolucion,omitempty"`
	Mercancia                             []MercanciaCartaPorte31            `xml:"Mercancia" bson:"Mercancia"`
	Autotransporte                        *AutotransporteCartaPorte31        `xml:"Autotransporte" bson:"Autotransporte,omitempty"`
	TransporteMaritimo                    *TransporteMaritimoCartaPorte31    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty"`
	TransporteAereo                       *TransporteAereoCartaPorte31       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty"`
	TransporteFerroviario                 *TransporteFerroviarioCartaPorte31 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte31 struct {
	// Attrs
	BienesTransportados string  `xml:"BienesTransp,attr" bson:"BienesTransportados"`
	ClaveStcc           *string `xml:"ClaveSTCC,attr" bson:"ClaveStcc,omitempty"`
	Descripcion         string  `xml:"Descripcion,attr" bson:"Descripcion"`
	Cantidad            float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad         string  `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad              *string `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones         *string `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`

	MaterialPeligroso   *string `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	EsMaterialPeligroso *bool   `bson:"EsMaterialPeligroso,omitempty"`

	ClaveMaterialPeligroso         *string `xml:"CveMaterialPeligroso,attr" bson:"ClaveMaterialPeligroso,omitempty"`
	Embalaje                       *string `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripcionEmbalaje            *string `xml:"DescripEmbalaje,attr" bson:"DescripcionEmbalaje,omitempty"`
	SectorCofepris                 *string `xml:"SectorCOFEPRIS,attr" bson:"SectorCofepris,omitempty"`
	NombreIngredienteActivo        *string `xml:"NombreIngredienteActivo,attr" bson:"NombreIngredienteActivo,omitempty"`
	NombreQuimico                  *string `xml:"NomQuimico,attr" bson:"NombreQuimico,omitempty"`
	DenominacionGenericaProducto   *string `xml:"DenominacionGenericaProd,attr" bson:"DenominacionGenericaProducto,omitempty"`
	DenominacionDistintivaProducto *string `xml:"DenominacionDistintivaProd,attr" bson:"DenominacionDistintivaProducto,omitempty"`
	Fabricante                     *string `xml:"Fabricante,attr" bson:"Fabricante,omitempty"`

	FechaCaducidadString *string    `xml:"FechaCaducidad,attr"`
	FechaCaducidad       *time.Time `bson:"FechaCaducidad,omitempty"`

	LoteMedicamento                        *string  `xml:"LoteMedicamento,attr" bson:"LoteMedicamento,omitempty"`
	FormaFarmaceutica                      *string  `xml:"FormaFarmaceutica,attr" bson:"FormaFarmaceutica,omitempty"`
	CondicionesEspecialesTransporte        *string  `xml:"CondicionesEspTransp,attr" bson:"CondicionesEspecialesTransporte,omitempty"`
	RegistroSanitarioFolioAutorizacion     *string  `xml:"RegistroSanitarioFolioAutorizacion,attr" bson:"RegistroSanitarioFolioAutorizacion,omitempty"`
	PermisoImportacion                     *string  `xml:"PermisoImportacion,attr" bson:"PermisoImportacion,omitempty"`
	FolioImportacionVucem                  *string  `xml:"FolioImpoVUCEM,attr" bson:"FolioImportacionVucem,omitempty"`
	NoCas                                  *string  `xml:"NumCAS,attr" bson:"NoCas,omitempty"`
	RazonSocialEmpresaImportadora          *string  `xml:"RazonSocialEmpImp,attr" bson:"RazonSocialEmpresaImportadora,omitempty"`
	NoRegistroSanitarioPlaguicidasCofepris *string  `xml:"NumRegSanPlagCOFEPRIS,attr" bson:"NoRegistroSanitarioPlaguicidasCofepris,omitempty"`
	DatosFabricante                        *string  `xml:"DatosFabricante,attr" bson:"DatosFabricante,omitempty"`
	DatosFormulador                        *string  `xml:"DatosFormulador,attr" bson:"DatosFormulador,omitempty"`
	DatosMaquilador                        *string  `xml:"DatosMaquilador,attr" bson:"DatosMaquilador,omitempty"`
	UsoAutorizado                          *string  `xml:"UsoAutorizado,attr" bson:"UsoAutorizado,omitempty"`
	PesoKg                                 float64  `xml:"PesoEnKg,attr" bson:"PesoKg"`
	ValorMercancia                         *float64 `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda                                 *string  `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria                    *string  `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UuidComercioExterior                   *string  `xml:"UUIDComercioExt,attr" bson:"UuidComercioExterior,omitempty"`
	TipoMateria                            *string  `xml:"TipoMateria,attr" bson:"TipoMateria,omitempty"`
	DescripcionMateria                     *string  `xml:"DescripcionMateria,attr" bson:"DescripcionMateria,omitempty"`
	// Nodos
	DocumentacionAduanera *[]DocumentacionAduaneraCartaPorte31 `xml:"DocumentacionAduanera" bson:"DocumentacionAduanera,omitempty"`
	GuiasIdentificacion   *[]GuiasIdentificacionCartaPorte31   `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty"`
	CantidadTransporta    *[]CantidadTransporteCartaPorte31    `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia      *DetalleMercanciaCartaPorte31        `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
}

func (m *MercanciaCartaPorte31) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias MercanciaCartaPorte31
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*m = MercanciaCartaPorte31(aux)
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
	return nil
}

type DocumentacionAduaneraCartaPorte31 struct {
	Tipo                   string  `xml:"TipoDocumento,attr" bson:"TipoDocumento"`
	NumeroPedimento        *string `xml:"NumPedimento,attr" bson:"NumeroPedimento,omitempty"`
	IdentificadorDocumento *string `xml:"IdentDocAduanero,attr" bson:"IdentificadorDocumento,omitempty"`
	RfcImportador          *string `xml:"RFCImpo,attr" bson:"RfcImportador,omitempty"`
}

type GuiasIdentificacionCartaPorte31 struct {
	Numero      string  `xml:"NumeroGuiaIdentificacion,attr" bson:"Numero"`
	Descripcion string  `xml:"DescripGuiaIdentificacion,attr" bson:"Descripcion"`
	Peso        float64 `xml:"PesoGuiaIdentificacion,attr" bson:"Peso"`
}

type CantidadTransporteCartaPorte31 struct {
	Cantidad         float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IdOrigen         string  `xml:"IDOrigen,attr" bson:"IdOrigen"`
	IdDestino        string  `xml:"IDDestino,attr" bson:"IdDestino"`
	ClavesTransporte *string `xml:"CvesTransporte,attr" bson:"ClavesTransporte"`
}

type DetalleMercanciaCartaPorte31 struct {
	UnidadPeso   string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPeso"`
	PesoBruto    float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	PesoNeto     float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoTara     float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	NumeroPiezas *float64 `xml:"NumPiezas,attr" bson:"NumeroPiezas,omitempty"`
}

type AutotransporteCartaPorte31 struct {
	PermisoSct              string                              `xml:"PermSCT,attr" bson:"PermisoSct"`
	NoPermisoSct            string                              `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte31 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte31                 `xml:"Seguros" bson:"Seguros"`
	Remolques               *[]RemolqueCartaPorte31             `xml:"Remolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte31 struct {
	Configuracion string  `xml:"ConfigVehicular,attr" bson:"Configuracion"`
	PesoBruto     float64 `xml:"PesoBrutoVehicular,attr" bson:"PesoBruto"`
	PlacaVm       string  `xml:"PlacaVM,attr" bson:"PlacaVm"`
	AnioModeloVm  float64 `xml:"AnioModeloVM,attr" bson:"AnioModeloVm"`
}

type SegurosCartaPorte31 struct {
	AseguradoraResponsabilidadCivil string   `xml:"AseguraRespCivil,attr" bson:"AseguradoraResponsabilidadCivil"`
	PolizaResponsabilidadCivil      string   `xml:"PolizaRespCivil,attr" bson:"PolizaResponsabilidadCivil"`
	AseguradoraMedioAmbiente        *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguradoraMedioAmbiente,omitempty"`
	PolizaMedioAmbiente             *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedioAmbiente,omitempty"`
	AseguradoraCarga                *string  `xml:"AseguraCarga,attr" bson:"AseguradoraCarga,omitempty"`
	PolizaCarga                     *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty"`
	PrimaSeguro                     *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte31 struct {
	Subtipo string `xml:"SubTipoRem,attr" bson:"Subtipo"`
	Placa   string `xml:"Placa,attr" bson:"Placa"`
}

type TransporteMaritimoCartaPorte31 struct {
	PermSCT                   *string                   `xml:"PermSCT,attr" bson:"PermisoSct,omitempty"`
	NoPermisoSct              *string                   `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct,omitempty"`
	NombreAseguradora         *string                   `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro            *string                   `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	TipoEmbarcacion           string                    `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion"`
	Matricula                 string                    `xml:"Matricula,attr" bson:"Matricula"`
	NoOmi                     string                    `xml:"NumeroOMI,attr" bson:"NoOmi"`
	AnioEmbarcacion           *float64                  `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty"`
	NombreEmbarcacion         *string                   `xml:"NombreEmbarc,attr" bson:"NombreEmbarcacion,omitempty"`
	NacionalidadEmbarcion     string                    `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarcion"`
	UnidadesArqueoBruto       float64                   `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesArqueoBruto"`
	TipoCarga                 string                    `xml:"TipoCarga,attr" bson:"TipoCarga"`
	Eslora                    *float64                  `xml:"Eslora,attr" bson:"Eslora,omitempty"`
	Manga                     *float64                  `xml:"Manga,attr" bson:"Manga,omitempty"`
	Calado                    *float64                  `xml:"Calado,attr" bson:"Calado,omitempty"`
	Puntal                    *float64                  `xml:"Puntal,attr" bson:"Puntal,omitempty"`
	LineaNaviera              *string                   `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty"`
	NombreAgenteNaviero       string                    `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero"`
	NoAutorizacionNaviero     string                    `xml:"NumAutorizacionNaviero,attr" bson:"NoAutorizacionNaviero"`
	NoViaje                   *string                   `xml:"NumViaje,attr" bson:"NoViaje,omitempty"`
	NoConocimientoEmbarque    *string                   `xml:"NumConocEmbarc,attr" bson:"NoConocimientoEmbarque,omitempty"`
	PermisoTemporalNavegacion *string                   `xml:"PermisoTempNavegacion,attr" bson:"PermisoTemporalNavegacion,omitempty"`
	Contenedor                *[]ContenedorCartaPorte31 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCartaPorte31 struct {
	Tipo                        string                     `xml:"TipoContenedor,attr" bson:"Tipo"`
	Matricula                   *string                    `xml:"MatriculaContenedor,attr" bson:"Matricula,omitempty"`
	NoPrecinto                  *string                    `xml:"NumPrecinto,attr" bson:"NoPrecinto,omitempty"`
	IdCcpRelacionado            *string                    `xml:"IdCCPRelacionado,attr" bson:"IdCcpRelacionado,omitempty"`
	PlacaVmCcp                  *string                    `xml:"PlacaVMCCP,attr" bson:"PlacaVmCcp,omitempty"`
	FechaCertificacionCcpString *string                    `xml:"FechaCertificacionCCP,attr"`
	FechaCertificacionCcp       *time.Time                 `bson:"FechaCertificacionCcp,omitempty"`
	RemolquesCcp                *[]RemolqueCCPCartaPorte31 `xml:"RemolquesCCP>RemolqueCCP" bson:"RemolquesCcp,omitempty"`
}

func (c *ContenedorCartaPorte31) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ContenedorCartaPorte31
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*c = ContenedorCartaPorte31(aux)

	if c.FechaCertificacionCcpString != nil && *c.FechaCertificacionCcpString != "" {
		fecha, errFecha := helpers.ParseDatetime(*c.FechaCertificacionCcpString)
		if errFecha == nil {
			c.FechaCertificacionCcp = &fecha
		}
	}
	return nil
}

type RemolqueCCPCartaPorte31 struct {
	Subtipo string `xml:"SubTipoRemCCP,attr" bson:"Subtipo"`
	Placa   string `xml:"PlacaCCP,attr" bson:"Placa"`
}

type TransporteAereoCartaPorte31 struct {
	PermisoSct                 string  `xml:"PermSCT,attr" bson:"PermisoSct"`
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

type TransporteFerroviarioCartaPorte31 struct {
	TipoServicio      string                        `xml:"TipoDeServicio,attr" bson:"TipoServicio"`
	TipoTrafico       string                        `xml:"TipoDeTrafico,attr" bson:"TipoTrafico"`
	NombreAseguradora *string                       `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty"`
	NoPolizaSeguro    *string                       `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty"`
	DerechosDePaso    *[]DerechosDePasoCartaPorte31 `xml:"DerechosDePaso" bson:"DerechosDePaso,omitempty"`
	Carro             []CarroCartaPorte31           `xml:"Carro" bson:"Carro"`
}

type DerechosDePasoCartaPorte31 struct {
	Tipo              string  `xml:"TipoDerechoDePaso,attr" bson:"Tipo"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte31 struct {
	Tipo           string                         `xml:"TipoCarro,attr" bson:"Tipo"`
	Matricula      string                         `xml:"MatriculaCarro,attr" bson:"Matricula"`
	Guia           string                         `xml:"GuiaCarro,attr" bson:"Guia"`
	ToneladasNetas float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetas"`
	Contenedor     *[]ContenedorCarroCartaPorte31 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte31 struct {
	Tipo                string  `xml:"TipoContenedor,attr" bson:"Tipo"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte31 struct {
	TiposFigura []*TiposFiguraCartaPorte31 `xml:"TiposFigura" bson:"TiposFigura"`
}

type TiposFiguraCartaPorte31 struct {
	Tipo             string                          `xml:"TipoFigura,attr" bson:"Tipo"`
	Rfc              *string                         `xml:"RFCFigura,attr" bson:"Rfc,omitempty"`
	NoLicencia       *string                         `xml:"NumLicencia,attr" bson:"NoLicencia,omitempty"`
	Nombre           string                          `xml:"NombreFigura,attr" bson:"Nombre"`
	NumRegIdTrib     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscal,omitempty"`
	PartesTransporte *[]PartesTransporteCartaPorte31 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty"`
	Domicilio        *DomicilioCartaPorte31          `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte31 struct {
	Parte string `xml:"ParteTransporte,attr" bson:"Parte"`
}

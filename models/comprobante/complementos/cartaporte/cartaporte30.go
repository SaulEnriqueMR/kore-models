package cartaporte

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type CartaPorte30 struct {
	XMLName xml.Name `xml:"CartaPorte"`
	Version string   `xml:"Version,attr" bson:"Version"`
	IdCcp   string   `xml:"IdCCP,attr" bson:"IdCcp"`

	TransporteInternacional   string `xml:"TranspInternac,attr" bson:"TransporteInternacional"`
	EsTransporteInternacional bool   `bson:"EsTransporteInternacional"`

	RegimenAduanero         *string                       `xml:"RegimenAduanero,attr" bson:"RegimenAduanero,omitempty"`
	EntradaSalidaMercancia  *string                       `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty"`
	PaisOrigenDestino       *string                       `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida        *string                       `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida *float64                      `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty"`
	RegistroIstmo           *string                       `xml:"RegistroISTMO,attr" bson:"RegistroIstmo,omitempty"`
	UbicacionPoloOrigen     *string                       `xml:"UbicacionPoloOrigen,attr" bson:"UbicacionPoloOrigen,omitempty"`
	UbicacionPoloDestino    *string                       `xml:"UbicacionPoloDestino,attr" bson:"UbicacionPoloDestino,omitempty"`
	Ubicaciones             []UbicacionCartaPorte30       `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias              MercanciasCartaPorte30        `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte        *FiguraTransporteCartaPorte30 `xml:"FiguraTransporte" bson:"FiguraTransporte,omitempty"`
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

	return nil
}

type UbicacionCartaPorte30 struct {
	TipoUbicacion               string                 `xml:"TipoUbicacion,attr" bson:"TipoUbicacion"`
	IDUbicacion                 *string                `xml:"IDUbicacion,attr" bson:"IDUbicacion,omitempty"`
	RFCRemitenteDestinatario    string                 `xml:"RFCRemitenteDestinatario,attr" bson:"RFCRemitenteDestinatario"`                 // Cifrado
	NombreRemitenteDestinatario *string                `xml:"NombreRemitenteDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty"` // Cifrado
	NumRegIdTrib                *string                `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`                               // Cifrado
	ResidenciaFiscal            *string                `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`                       // Cifrado
	NumEstacion                 *string                `xml:"NumEstacion,attr" bson:"NumEstacion,omitempty"`
	NombreEstacion              *string                `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico           *string                `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegada      string                 `xml:"FechaHoraSalidaLlegada,attr"`
	FechaHoraSalidaLlegadaDate  time.Time              `bson:"FechaHoraSalidaLlegada"`
	TipoEstacion                *string                `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty"`
	DistanciaRecorrida          *float64               `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty"`
	Domicilio                   *DomicilioCartaPorte30 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DomicilioCartaPorte30 struct {
	Calle          *string `xml:"Calle,attr" bson:"Calle,omitempty"`                   // Cifrado
	NumeroExterior *string `xml:"NumeroExterior,attr" bson:"NumeroExterior,omitempty"` // Cifrado
	NumeroInterior *string `xml:"NumeroInterior,attr" bson:"NumeroInterior,omitempty"` // Cifrado
	Colonia        *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`               // Cifrado
	Localidad      *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`           // Cifrado
	Referencia     *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`         // Cifrado
	Municipio      *string `xml:"Municipio,attr" bson:"Municipio,omitempty"`           // Cifrado
	Estado         string  `xml:"Estado,attr" bson:"Estado,omitempty"`                 // Cifrado
	Pais           string  `xml:"Pais,attr" bson:"Pais,omitempty"`                     // Cifrado
	CodigoPostal   string  `xml:"CodigoPostal,attr" bson:"CodigoPostal,omitempty"`     // Cifrado
}

type MercanciasCartaPorte30 struct {
	PesoBrutoTotal                        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal"`
	UnidadPeso                            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoNetoTotal                         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumTotalMercancias                    float64                            `xml:"NumTotalMercancias,attr" bson:"NumTotalMercancias"`
	CargoPorTasacion                      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty"`
	LogisticaInversaRecoleccionDevolucion *string                            `xml:"LogisticaInversaRecoleccionDevolucion,attr" bson:"LogisticaInversaRecoleccionDevolucion,omitempty"`
	Mercancia                             []MercanciaCartaPorte30            `xml:"Mercancia" bson:"Mercancia"`
	Autotransporte                        *AutotransporteCartaPorte30        `xml:"Autotransporte" bson:"Autotransporte,omitempty"`
	TransporteMaritimo                    *TransporteMaritimoCartaPorte30    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty"`
	TransporteAereo                       *TransporteAereoCartaPorte30       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty"`
	TransporteFerroviario                 *TransporteFerroviarioCartaPorte30 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte30 struct {
	BienesTransp                       string                               `xml:"BienesTransp,attr" bson:"BienesTransp"`
	ClaveSTCC                          *string                              `xml:"ClaveSTCC,attr" bson:"ClaveSTCC,omitempty"`
	Descripcion                        string                               `xml:"Descripcion,attr" bson:"Descripcion"`
	Cantidad                           float64                              `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad                        string                               `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad                             *string                              `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones                        *string                              `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`
	MaterialPeligroso                  *string                              `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	CveMaterialPeligroso               *string                              `xml:"CveMaterialPeligroso,attr" bson:"CveMaterialPeligroso,omitempty"`
	Embalaje                           *string                              `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripEmbalaje                    *string                              `xml:"DescripEmbalaje,attr" bson:"DescripEmbalaje,omitempty"`
	SectorCOFEPRIS                     *string                              `xml:"SectorCOFEPRIS,attr" bson:"SectorCOFEPRIS,omitempty"`
	NombreIngredienteActivo            *string                              `xml:"NombreIngredienteActivo,attr" bson:"NombreIngredienteActivo,omitempty"`
	NomQuimico                         *string                              `xml:"NomQuimico,attr" bson:"NomQuimico,omitempty"`
	DenominacionGenericaProd           *string                              `xml:"DenominacionGenericaProd,attr" bson:"DenominacionGenericaProd,omitempty"`
	DenominacionDistintivaProd         *string                              `xml:"DenominacionDistintivaProd,attr" bson:"DenominacionDistintivaProd,omitempty"`
	Fabricante                         *string                              `xml:"Fabricante,attr" bson:"Fabricante,omitempty"`
	FechaCaducidad                     *string                              `xml:"FechaCaducidad,attr" bson:"FechaCaducidad,omitempty"`
	LoteMedicamento                    *string                              `xml:"LoteMedicamento,attr" bson:"LoteMedicamento,omitempty"`
	FormaFarmaceutica                  *string                              `xml:"FormaFarmaceutica,attr" bson:"FormaFarmaceutica,omitempty"`
	CondicionesEspTransp               *string                              `xml:"CondicionesEspTransp,attr" bson:"CondicionesEspTransp,omitempty"`
	RegistroSanitarioFolioAutorizacion *string                              `xml:"RegistroSanitarioFolioAutorizacion,attr" bson:"RegistroSanitarioFolioAutorizacion,omitempty"`
	PermisoImportacion                 *string                              `xml:"PermisoImportacion,attr" bson:"PermisoImportacion,omitempty"`
	FolioImpoVUCEM                     *string                              `xml:"FolioImpoVUCEM,attr" bson:"FolioImpoVUCEM,omitempty"`
	NumCAS                             *string                              `xml:"NumCAS,attr" bson:"NumCAS,omitempty"`
	RazonSocialEmpImp                  *string                              `xml:"RazonSocialEmpImp,attr" bson:"RazonSocialEmpImp,omitempty"`
	NumRegSanPlagCOFEPRIS              *string                              `xml:"NumRegSanPlagCOFEPRIS,attr" bson:"NumRegSanPlagCOFEPRIS,omitempty"`
	DatosFabricante                    *string                              `xml:"DatosFabricante,attr" bson:"DatosFabricante,omitempty"`
	DatosFormulador                    *string                              `xml:"DatosFormulador,attr" bson:"DatosFormulador,omitempty"`
	DatosMaquilador                    *string                              `xml:"DatosMaquilador,attr" bson:"DatosMaquilador,omitempty"`
	UsoAutorizado                      *string                              `xml:"UsoAutorizado,attr" bson:"UsoAutorizado,omitempty"`
	PesoEnKg                           float64                              `xml:"PesoEnKg,attr" bson:"PesoEnKg"`
	ValorMercancia                     *float64                             `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda                             *string                              `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria                *string                              `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UUIDComercioExt                    *string                              `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty"`
	TipoMateria                        *string                              `xml:"TipoMateria,attr" bson:"TipoMateria,omitempty"`
	DescripcionMateria                 *string                              `xml:"DescripcionMateria,attr" bson:"DescripcionMateria,omitempty"`
	DocumentacionAduanera              *[]DocumentacionAduaneraCartaPorte30 `xml:"DocumentacionAduanera" bson:"DocumentacionAduanera,omitempty"`
	GuiasIdentificacion                *[]GuiasIdentificacionCartaPorte30   `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty"`
	CantidadTransporta                 *[]CantidadTransportaCartaPorte30    `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia                   *DetalleMercanciaCartaPorte30        `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
}

type DocumentacionAduaneraCartaPorte30 struct {
	TipoDocumento    string  `xml:"TipoDocumento,attr" bson:"TipoDocumento"`
	NumPedimento     *string `xml:"NumPedimento,attr" bson:"NumPedimento,omitemtpy"`
	IdentDocAduanero *string `xml:"IdentDocAduanero,attr" bson:"IdentDocAduanero,omitempty"`
	RFCImpo          *string `xml:"RFCImpo,attr" bson:"RFCImpo,omitempty"` // Cifrado
}

type GuiasIdentificacionCartaPorte30 struct {
	NumeroGuiaIdentificacion  string  `xml:"NumeroGuiaIdentificacion,attr" bson:"NumeroGuiaIdentificacion"`
	DescripGuiaIdentificacion string  `xml:"DescripGuiaIdentificacion,attr" bson:"DescripGuiaIdentificacion"`
	PesoGuiaIdentificacion    float64 `xml:"PesoGuiaIdentificacion,attr" bson:"PesoGuiaIdentificacion"`
}

type CantidadTransportaCartaPorte30 struct {
	Cantidad       float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IDOrigen       string  `xml:"IDOrigen,attr" bson:"IDOrigen"`
	IDDestino      string  `xml:"IDDestino,attr" bson:"IDDestino"`
	CvesTransporte *string `xml:"CvesTransporte,attr" bson:"CvesTransporte,omitempty"`
}

type DetalleMercanciaCartaPorte30 struct {
	UnidadPesoMerc string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPesoMerc"`
	PesoTara       float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	PesoNeto       float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoBruto      float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	NumPiezas      *float64 `xml:"NumPiezas,attr" bson:"NumPiezas,omitempty"`
}

type AutotransporteCartaPorte30 struct {
	PermSCT                 string                              `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT           string                              `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte30 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte30                 `xml:"Seguros" bson:"Seguros"`
	Remolques               *[]RemolqueCartaPorte30             `xml:"Remolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte30 struct {
	ConfigVehicular    string  `xml:"ConfigVehicular,attr" bson:"ConfigVehicular"`
	PesoBrutoVehicular float64 `xml:"PesoBrutoVehicular,attr" bson:"PesoBrutoVehicular"`
	PlacaVM            string  `xml:"PlacaVM,attr" bson:"PlacaVM"`
	AnioModeloVM       float64 `xml:"AnioModeloVM,attr" bson:"AnioModeloVM"`
}

type SegurosCartaPorte30 struct {
	AseguraRespCivil   string   `xml:"AseguraRespCivil,attr" bson:"AseguraRespCivil"`
	PolizaRespCivil    string   `xml:"PolizaRespCivil,attr" bson:"PolizaRespCivil"` // Cifrado
	AseguraMedAmbiente *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguraMedAmbiente,omitempty"`
	PolizaMedAmbiente  *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedAmbiente,omitempty"` // Cifrado
	AseguraCarga       *string  `xml:"AseguraCarga,attr" bson:"AseguraCarga,omitempty"`
	PolizaCarga        *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty"` // Cifrado
	PrimaSeguro        *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte30 struct {
	SubTipoRem string `xml:"SubTipoRem,attr" bson:"SubTipoRem"`
	Placa      string `xml:"Placa,attr" bson:"Placa"`
}

type TransporteMaritimoCartaPorte30 struct {
	PermSCT                *string                    `xml:"PermSCT,attr" bson:"PermSCT,omitempty"`
	NumPermisoSCT          *string                    `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT,omitempty"`     // Cifrado
	NombreAseg             *string                    `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`           // Cifrado
	NumPolizaSeguro        *string                    `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"` // Cifrado
	TipoEmbarcacion        string                     `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion"`
	Matricula              string                     `xml:"Matricula,attr" bson:"Matricula"` // Cifrado
	NumeroOMI              string                     `xml:"NumeroOMI,attr" bson:"NumeroOMI"`
	AnioEmbarcacion        *float64                   `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty"`
	NombreEmbarc           *string                    `xml:"NombreEmbarc,attr" bson:"NombreEmbarc,omitempty"`
	NacionalidadEmbarc     string                     `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarc"`
	UnidadesDeArqBruto     float64                    `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesDeArqBruto"`
	TipoCarga              string                     `xml:"TipoCarga,attr" bson:"TipoCarga"`
	Eslora                 *float64                   `xml:"Eslora,attr" bson:"Eslora,omitempty"`
	Manga                  *float64                   `xml:"Manga,attr" bson:"Manga,omitempty"`
	Calado                 *float64                   `xml:"Calado,attr" bson:"Calado,omitempty"`
	Puntal                 *float64                   `xml:"Puntal,attr" bson:"Puntal,omitempty"`
	LineaNaviera           *string                    `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty"`
	NombreAgenteNaviero    string                     `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero"`       // Cifrado
	NumAutorizacionNaviero string                     `xml:"NumAutorizacionNaviero,attr" bson:"NumAutorizacionNaviero"` // Cifrado
	NumViaje               *string                    `xml:"NumViaje,attr" bson:"NumViaje,omitempty"`
	NumConocEmbarc         *string                    `xml:"NumConocEmbarc,attr" bson:"NumConocEmbarc,omitempty"`
	PermisoTempNavegacion  *string                    `xml:"PermisoTempNavegacion,attr" bson:"PermisoTempNavegacion,omitempty"`
	Contenedor             *[]ContenedorCartaPorte30  `xml:"Contenedor" bson:"Contenedor,omitempty"`
	RemolquesCCP           *[]RemolqueCCPCartaPorte30 `xml:"RemolquesCCP>RemolqueCCP" bson:"RemolquesCCP,omitempty"`
}

type ContenedorCartaPorte30 struct {
	TipoContenedor        string  `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	MatriculaContenedor   *string `xml:"MatriculaContenedor,attr" bson:"MatriculaContenedor,omitempty"` // Cifrado
	NumPrecinto           *string `xml:"NumPrecinto,attr" bson:"NumPrecinto,omitempty"`
	IdCCPRelacionado      *string `xml:"IdCCPRelacionado,attr" bson:"IdCCPRelacionado,omitempty"`
	PlacaVMCCP            *string `xml:"PlacaVMCCP,attr" bson:"PlacaVMCCP,omitempty"` // Cifrado
	FechaCertificacionCCP *string `xml:"FechaCertificacionCCP,attr" bson:"FechaCertificacionCCP,omitempty"`
}

type RemolqueCCPCartaPorte30 struct {
	SubTipoRemCCP string `xml:"SubTipoRemCCP,attr" bson:"SubTipoRemCCP"`
	PlacaCCP      string `xml:"PlacaCCP,attr" bson:"PlacaCCP"` // Cifrado
}

type TransporteAereoCartaPorte30 struct {
	PermSCT                string  `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT          string  `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"`                   // Cifrado
	MatriculaAeronave      *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty"` // Cifrado
	NombreAseg             *string `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`               // Cifrado
	NumPolizaSeguro        *string `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"`     // Cifrado
	NumeroGuia             string  `xml:"NumeroGuia,attr" bson:"NumeroGuia"`
	RFCEmbarcador          *string `xml:"RFCEmbarcador,attr" bson:"RFCEmbarcador,omitempty"`                   // Cifrado
	ResidenciaFiscalEmbarc *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarc,omitempty"` // Cifrado
	NumRegIdTribEmbarc     *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarc,omitempty"`         // Cifrado
	NombreEmbarcador       *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty"`             // Cifrado
	LugarContrato          *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty"`
	CodigoTransportista    string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista"`
}

type TransporteFerroviarioCartaPorte30 struct {
	TipoDeTrafico   string                        `xml:"TipoDeTrafico,attr" bson:"TipoDeTrafico"`
	TipoDeServicio  string                        `xml:"TipoDeServicio,attr" bson:"TipoDeServicio"`
	NumPolizaSeguro *string                       `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"` // Cifrado
	NombreAseg      *string                       `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`           // Cifrado
	DerechosPaso    *[]DerechosDePasoCartaPorte30 `xml:"DerechosDePaso" bson:"DerechosPaso,omitempty"`
	Carro           []CarroCartaPorte30           `xml:"Carro" bson:"Carro"`
}

type DerechosDePasoCartaPorte30 struct {
	TipoDerechoDePaso string  `xml:"TipoDerechoDePaso,attr" bson:"TipoDerechoDePaso"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte30 struct {
	ToneladasNetasCarro float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetasCarro"`
	TipoCarro           string                         `xml:"TipoCarro,attr" bson:"TipoCarro"`
	MatriculaCarro      string                         `xml:"MatriculaCarro,attr" bson:"MatriculaCarro"` // Cifrado
	GuiaCarro           string                         `xml:"GuiaCarro,attr" bson:"GuiaCarro"`
	Contenedor          *[]ContenedorCarroCartaPorte30 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte30 struct {
	TipoContenedor      string  `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
}

type FiguraTransporteCartaPorte30 struct {
	TiposFigura []TiposFiguraCartaPorte30 `xml:"TiposFigura" bson:"TiposFigura"`
}

type TiposFiguraCartaPorte30 struct {
	TipoFigura             string                          `xml:"TipoFigura,attr" bson:"TipoFigura"`
	RFCFigura              *string                         `xml:"RFCFigura,attr" bson:"RFCFigura,omitempty"`                           // Cifrado
	ResidenciaFiscalFigura *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscalFigura,omitempty"` // Cifrado
	NumRegIdTribFigura     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTribFigura,omitempty"`         // Cifrado
	NumLicencia            *string                         `xml:"NumLicencia,attr" bson:"NumLicencia,omitempty"`                       // Cifrado
	NombreFigura           string                          `xml:"NombreFigura,attr" bson:"NombreFigura"`
	PartesTransporte       *[]PartesTransporteCartaPorte30 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty"`
	Domicilio              *DomicilioCartaPorte30          `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte30 struct {
	ParteTransporte string `xml:"ParteTransporte,attr" bson:"ParteTransporte"`
}

func (ucp30 *UbicacionCartaPorte30) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias UbicacionCartaPorte30
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*ucp30 = UbicacionCartaPorte30(aux)

	if ucp30 != nil {
		fecha, err := helpers.ParseDatetime(ucp30.FechaHoraSalidaLlegada)
		if err != nil {
			return err
		}
		ucp30.FechaHoraSalidaLlegadaDate = fecha
	}

	return nil
}

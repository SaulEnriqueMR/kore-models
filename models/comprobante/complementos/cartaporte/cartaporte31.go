package cartaporte

import "encoding/xml"

type CartaPorte31 struct {
	// Attr
	XMLName              xml.Name `xml:"CartaPorte"`
	Version              string   `xml:"Version,attr" bson:"Version"`
	IdCCP                string   `xml:"IdCCP,attr" bson:"IdCCP"`
	TranspInternac       string   `xml:"TranspInternac,attr" bson:"TranspInternac"`
	EntradaSalidaMerc    *string  `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMerc,omitempty"`
	PaisOrigenDestino    *string  `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida     *string  `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty"`
	TotalDistRec         *float64 `xml:"TotalDistRec,attr" bson:"TotalDistRec,omitempty"`
	RegistroISTMO        *string  `xml:"RegistroISTMO,attr" bson:"RegistroISTMO,omitempty"`
	UbicacionPoloOrigen  *string  `xml:"UbicacionPoloOrigen,attr" bson:"UbicacionPoloOrigen,omitempty"`
	UbicacionPoloDestino *string  `xml:"UbicacionPoloDestino,attr" bson:"UbicacionPoloDestino,omitempty"`
	// Nodos
	RegimenesAduaneros *[]RegimenAduaneroCartaPorte31 `xml:"RegimenesAduaneros>RegimenAduaneroCCP" bson:"RegimenesAduaneros,omitempty"`
	Ubicaciones        []UbicacionCartaPorte31        `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones"`
	Mercancias         MercanciasCartaPorte31         `xml:"Mercancias" bson:"Mercancias"`
	FiguraTransporte   *FiguraTransporteCartaPorte31  `xml:"FiguraTransporte" bson:"FiguraTransporte,omitempty"`
}

type RegimenAduaneroCartaPorte31 struct {
	RegimenAduanero string `xml:"RegimenAduanero,attr" bson:"RegimenAduanero"`
}

type UbicacionCartaPorte31 struct {
	// Attr
	TipoUbicacion               string   `xml:"TipoUbicacion,attr" bson:"TipoUbicacion"`
	IDUbicacion                 *string  `xml:"IDUbicacion,attr" bson:"IDUbicacion,omitempty"`
	RFCRemitenteDestinatario    string   `xml:"RFCRemitenteDestinatario,attr" bson:"RFCRemitenteDestinatario"`
	NombreRemitenteDestinatario *string  `xml:"NombreRemitenteDestinatario,attr" bson:"NombreRemitenteDestinatario,omitempty"`
	NumRegIdTrib                *string  `xml:"NumRegIdTrib,attr" bson:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal            *string  `xml:"ResidenciaFiscal,attr" bson:"ResidenciaFiscal,omitempty"`
	NumEstacion                 *string  `xml:"NumEstacion,attr" bson:"NumEstacion,omitempty"`
	NombreEstacion              *string  `xml:"NombreEstacion,attr" bson:"NombreEstacion,omitempty"`
	NavegacionTrafico           *string  `xml:"NavegacionTrafico,attr" bson:"NavegacionTrafico,omitempty"`
	FechaHoraSalidaLlegada      string   `xml:"FechaHoraSalidaLlegada,attr" bson:"FechaHoraSalidaLlegada"` //Agregar parser
	TipoEstacion                *string  `xml:"TipoEstacion,attr" bson:"TipoEstacion,omitempty"`
	DistanciaRecorrida          *float64 `xml:"DistanciaRecorrida,attr" bson:"DistanciaRecorrida,omitempty"`
	// Nodos
	Domicilio *DomicilioCartaPorte31 `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type DomicilioCartaPorte31 struct {
	Calle          *string `xml:"Calle,attr" bson:"Calle,omitempty"`
	NumeroExterior *string `xml:"NumeroExterior,attr" bson:"NumeroExterior,omitempty"`
	NumeroInterior *string `xml:"NumeroInterior,attr" bson:"NumeroInterior,omitempty"`
	Colonia        *string `xml:"Colonia,attr" bson:"Colonia,omitempty"`
	Localidad      *string `xml:"Localidad,attr" bson:"Localidad,omitempty"`
	Referencia     *string `xml:"Referencia,attr" bson:"Referencia,omitempty"`
	Municipio      *string `xml:"Municipio,attr" bson:"Municipio,omitempty"`
	Estado         string  `xml:"Estado,attr" bson:"Estado"`
	Pais           string  `xml:"Pais,attr" bson:"Pais"`
	CodigoPostal   string  `xml:"CodigoPostal,attr" bson:"CodigoPostal"`
}

type MercanciasCartaPorte31 struct {
	// Attr
	PesoBrutoTotal                        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal"`
	UnidadPeso                            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso"`
	PesoNetoTotal                         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty"`
	NumTotalMercancias                    float64                            `xml:"NumTotalMercancias,attr" bson:"NumTotalMercancias"`
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
	BienesTransp                       string   `xml:"BienesTransp,attr" bson:"BienesTransp"`
	ClaveSTCC                          *string  `xml:"ClaveSTCC,attr" bson:"ClaveSTCC,omitempty"`
	Descripcion                        string   `xml:"Descripcion,attr" bson:"Descripcion"`
	Cantidad                           float64  `xml:"Cantidad,attr" bson:"Cantidad"`
	ClaveUnidad                        string   `xml:"ClaveUnidad,attr" bson:"ClaveUnidad"`
	Unidad                             *string  `xml:"Unidad,attr" bson:"Unidad,omitempty"`
	Dimensiones                        *string  `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty"`
	MaterialPeligroso                  *string  `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty"`
	CveMaterialPeligroso               *string  `xml:"CveMaterialPeligroso,attr" bson:"CveMaterialPeligroso,omitempty"`
	Embalaje                           *string  `xml:"Embalaje,attr" bson:"Embalaje,omitempty"`
	DescripEmbalaje                    *string  `xml:"DescripEmbalaje,attr" bson:"DescripEmbalaje,omitempty"`
	SectorCOFEPRIS                     *string  `xml:"SectorCOFEPRIS,attr" bson:"SectorCOFEPRIS,omitempty"`
	NombreIngredienteActivo            *string  `xml:"NombreIngredienteActivo,attr" bson:"NombreIngredienteActivo,omitempty"`
	NomQuimico                         *string  `xml:"NomQuimico,attr" bson:"NomQuimico,omitempty"`
	DenominacionGenericaProd           *string  `xml:"DenominacionGenericaProd,attr" bson:"DenominacionGenericaProd,omitempty"`
	DenominacionDistintivaProd         *string  `xml:"DenominacionDistintivaProd,attr" bson:"DenominacionDistintivaProd,omitempty"`
	Fabricante                         *string  `xml:"Fabricante,attr" bson:"Fabricante,omitempty"`
	FechaCaducidad                     *string  `xml:"FechaCaducidad,attr" bson:"FechaCaducidad,omitempty"` //Agregar parser
	LoteMedicamento                    *string  `xml:"LoteMedicamento,attr" bson:"LoteMedicamento,omitempty"`
	FormaFarmaceutica                  *string  `xml:"FormaFarmaceutica,attr" bson:"FormaFarmaceutica,omitempty"`
	CondicionesEspTransp               *string  `xml:"CondicionesEspTransp,attr" bson:"CondicionesEspTransp,omitempty"`
	RegistroSanitarioFolioAutorizacion *string  `xml:"RegistroSanitarioFolioAutorizacion,attr" bson:"RegistroSanitarioFolioAutorizacion,omitempty"`
	PermisoImportacion                 *string  `xml:"PermisoImportacion,attr" bson:"PermisoImportacion,omitempty"`
	FolioImpoVUCEM                     *string  `xml:"FolioImpoVUCEM,attr" bson:"FolioImpoVUCEM,omitempty"`
	NumCAS                             *string  `xml:"NumCAS,attr" bson:"NumCAS,omitempty"`
	RazonSocialEmpImp                  *string  `xml:"RazonSocialEmpImp,attr" bson:"RazonSocialEmpImp,omitempty"`
	NumRegSanPlagCOFEPRIS              *string  `xml:"NumRegSanPlagCOFEPRIS,attr" bson:"NumRegSanPlagCOFEPRIS,omitempty"`
	DatosFabricante                    *string  `xml:"DatosFabricante,attr" bson:"DatosFabricante,omitempty"`
	DatosFormulador                    *string  `xml:"DatosFormulador,attr" bson:"DatosFormulador,omitempty"`
	DatosMaquilador                    *string  `xml:"DatosMaquilador,attr" bson:"DatosMaquilador,omitempty"`
	UsoAutorizado                      *string  `xml:"UsoAutorizado,attr" bson:"UsoAutorizado,omitempty"`
	PesoEnKg                           float64  `xml:"PesoEnKg,attr" bson:"PesoEnKg"`
	ValorMercancia                     *float64 `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty"`
	Moneda                             *string  `xml:"Moneda,attr" bson:"Moneda,omitempty"`
	FraccionArancelaria                *string  `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty"`
	UUIDComercioExt                    *string  `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty"`
	TipoMateria                        *string  `xml:"TipoMateria,attr" bson:"TipoMateria,omitempty"`
	DescripcionMateria                 *string  `xml:"DescripcionMateria,attr" bson:"DescripcionMateria,omitempty"`
	// Nodos
	DocumentacionAduanera *[]DocumentacionAduaneraCartaPorte31 `xml:"DocumentacionAduanera" bson:"DocumentacionAduanera,omitempty"`
	GuiasIdentificacion   *[]GuiasIdentificacionCartaPorte31   `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty"`
	CantidadTransporta    *[]CantidadTransporteCartaPorte31    `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty"`
	DetalleMercancia      *DetalleMercanciaCartaPorte31        `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty"`
}

type DocumentacionAduaneraCartaPorte31 struct {
	TipoDocumento    string  `xml:"TipoDocumento,attr" bson:"TipoDocumento"`
	NumPedimento     *string `xml:"NumPedimento,attr" bson:"NumPedimento,omitempty"`
	IdentDocAduanero *string `xml:"IdentDocAduanero,attr" bson:"IdentDocAduanero,omitempty"`
	RFCImpo          *string `xml:"RFCImpo,attr" bson:"RFCImpo,omitempty"`
}

type GuiasIdentificacionCartaPorte31 struct {
	NumeroGuiaIdentificacion  string  `xml:"NumeroGuiaIdentificacion,attr" bson:"NumeroGuiaIdentificacion"`
	DescripGuiaIdentificacion string  `xml:"DescripGuiaIdentificacion,attr" bson:"DescripGuiaIdentificacion"`
	PesoGuiaIdentificacion    float64 `xml:"PesoGuiaIdentificacion,attr" bson:"PesoGuiaIdentificacion"`
}

type CantidadTransporteCartaPorte31 struct {
	Cantidad       float64 `xml:"Cantidad,attr" bson:"Cantidad"`
	IDOrigen       string  `xml:"IDOrigen,attr" bson:"IDOrigen"`
	IDDestino      string  `xml:"IDDestino,attr" bson:"IDDestino"`
	CvesTransporte *string `xml:"CvesTransporte,attr" bson:"CvesTransporte"`
}

type DetalleMercanciaCartaPorte31 struct {
	UnidadPesoMerc string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPesoMerc"`
	PesoBruto      float64  `xml:"PesoBruto,attr" bson:"PesoBruto"`
	PesoNeto       float64  `xml:"PesoNeto,attr" bson:"PesoNeto"`
	PesoTara       float64  `xml:"PesoTara,attr" bson:"PesoTara"`
	NumPiezas      *float64 `xml:"NumPiezas,attr" bson:"NumPiezas,omitempty"`
}

type AutotransporteCartaPorte31 struct {
	PermSCT                 string                              `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT           string                              `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte31 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte31                 `xml:"Seguros" bson:"Seguros"`
	Remolques               *[]RemolqueCartaPorte31             `xml:"Remolques>Remolque" bson:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte31 struct {
	ConfigVehicular    string  `xml:"ConfigVehicular,attr" bson:"ConfigVehicular"`
	PesoBrutoVehicular float64 `xml:"PesoBrutoVehicular,attr" bson:"PesoBrutoVehicular"`
	PlacaVM            string  `xml:"PlacaVM,attr" bson:"PlacaVM"`
	AnioModeloVM       float64 `xml:"AnioModeloVM,attr" bson:"AnioModeloVM"`
}

type SegurosCartaPorte31 struct {
	AseguraRespCivil   string   `xml:"AseguraRespCivil,attr" bson:"AseguraRespCivil"`
	PolizaRespCivil    string   `xml:"PolizaRespCivil,attr" bson:"PolizaRespCivil"`
	AseguraMedAmbiente *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguraMedAmbiente,omitempty"`
	PolizaMedAmbiente  *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedAmbiente,omitempty"`
	AseguraCarga       *string  `xml:"AseguraCarga,attr" bson:"AseguraCarga,omitempty"`
	PolizaCarga        *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty"`
	PrimaSeguro        *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte31 struct {
	SubTipoRem string `xml:"SubTipoRem,attr" bson:"SubTipoRem"`
	Placa      string `xml:"Placa,attr" bson:"Placa"`
}

type TransporteMaritimoCartaPorte31 struct {
	PermSCT                *string                   `xml:"PermSCT,attr" bson:"PermSCT,omitempty"`
	NumPermisoSCT          *string                   `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT,omitempty"`
	NombreAseg             *string                   `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`
	NumPolizaSeguro        *string                   `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"`
	TipoEmbarcacion        string                    `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion"`
	Matricula              string                    `xml:"Matricula,attr" bson:"Matricula"`
	NumeroOMI              string                    `xml:"NumeroOMI,attr" bson:"NumeroOMI"`
	AnioEmbarcacion        *float64                  `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty"`
	NombreEmbarc           *string                   `xml:"NombreEmbarc,attr" bson:"NombreEmbarc,omitempty"`
	NacionalidadEmbarc     string                    `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarc"`
	UnidadesDeArqBruto     float64                   `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesDeArqBruto"`
	TipoCarga              string                    `xml:"TipoCarga,attr" bson:"TipoCarga"`
	Eslora                 *float64                  `xml:"Eslora,attr" bson:"Eslora,omitempty"`
	Manga                  *float64                  `xml:"Manga,attr" bson:"Manga,omitempty"`
	Calado                 *float64                  `xml:"Calado,attr" bson:"Calado,omitempty"`
	Puntal                 *float64                  `xml:"Puntal,attr" bson:"Puntal,omitempty"`
	LineaNaviera           *string                   `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty"`
	NombreAgenteNaviero    string                    `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero"`
	NumAutorizacionNaviero string                    `xml:"NumAutorizacionNaviero,attr" bson:"NumAutorizacionNaviero"`
	NumViaje               *string                   `xml:"NumViaje,attr" bson:"NumViaje,omitempty"`
	NumConocEmbarc         *string                   `xml:"NumConocEmbarc,attr" bson:"NumConocEmbarc,omitempty"`
	PermisoTempNavegacion  *string                   `xml:"PermisoTempNavegacion,attr" bson:"PermisoTempNavegacion,omitempty"`
	Contenedor             *[]ContenedorCartaPorte31 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCartaPorte31 struct {
	TipoContenedor        string                     `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	MatriculaContenedor   *string                    `xml:"MatriculaContenedor,attr" bson:"MatriculaContenedor,omitempty"`
	NumPrecinto           *string                    `xml:"NumPrecinto,attr" bson:"NumPrecinto,omitempty"`
	IdCCPRelacionado      *string                    `xml:"IdCCPRelacionado,attr" bson:"IdCCPRelacionado,omitempty"`
	PlacaVMCCP            *string                    `xml:"PlacaVMCCP,attr" bson:"PlacaVMCCP,omitempty"`
	FechaCertificacionCCP *string                    `xml:"FechaCertificacionCCP,attr" bson:"FechaCertificacionCCP,omitempty"`
	RemolquesCCP          *[]RemolqueCCPCartaPorte31 `xml:"RemolquesCCP>RemolqueCCP" bson:"RemolquesCCP,omitempty"`
}

type RemolqueCCPCartaPorte31 struct {
	SubTipoRemCCP string `xml:"SubTipoRemCCP,attr" bson:"SubTipoRemCCP"`
	PlacaCCP      string `xml:"PlacaCCP,attr" bson:"PlacaCCP"`
}

type TransporteAereoCartaPorte31 struct {
	PermSCT                string  `xml:"PermSCT,attr" bson:"PermSCT"`
	NumPermisoSCT          string  `xml:"NumPermisoSCT,attr" bson:"NumPermisoSCT"`
	MatriculaAeronave      *string `xml:"MatriculaAeronave,attr" bson:"MatriculaAeronave,omitempty"`
	NombreAseg             *string `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`
	NumPolizaSeguro        *string `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"`
	NumeroGuia             string  `xml:"NumeroGuia,attr" bson:"NumeroGuia"`
	LugarContrato          *string `xml:"LugarContrato,attr" bson:"LugarContrato,omitempty"`
	CodigoTransportista    string  `xml:"CodigoTransportista,attr" bson:"CodigoTransportista"`
	RFCEmbarcador          *string `xml:"RFCEmbarcador,attr" bson:"RFCEmbarcador,omitempty"`
	NumRegIdTribEmbarc     *string `xml:"NumRegIdTribEmbarc,attr" bson:"NumRegIdTribEmbarc,omitempty"`
	ResidenciaFiscalEmbarc *string `xml:"ResidenciaFiscalEmbarc,attr" bson:"ResidenciaFiscalEmbarc,omitempty"`
	NombreEmbarcador       *string `xml:"NombreEmbarcador,attr" bson:"NombreEmbarcador,omitempty"`
}

type TransporteFerroviarioCartaPorte31 struct {
	TipoDeServicio  string                        `xml:"TipoDeServicio,attr" bson:"TipoDeServicio"`
	TipoDeTrafico   string                        `xml:"TipoDeTrafico,attr" bson:"TipoDeTrafico"`
	NombreAseg      *string                       `xml:"NombreAseg,attr" bson:"NombreAseg,omitempty"`
	NumPolizaSeguro *string                       `xml:"NumPolizaSeguro,attr" bson:"NumPolizaSeguro,omitempty"`
	DerechosDePaso  *[]DerechosDePasoCartaPorte31 `xml:"DerechosDePaso" bson:"DerechosDePaso,omitempty"`
	Carro           []CarroCartaPorte31           `xml:"Carro" bson:"Carro"`
}

type DerechosDePasoCartaPorte31 struct {
	TipoDerechoDePaso string  `xml:"TipoDerechoDePaso,attr" bson:"TipoDerechoDePaso"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado"`
}

type CarroCartaPorte31 struct {
	TipoCarro           string                         `xml:"TipoCarro,attr" bson:"TipoCarro"`
	MatriculaCarro      string                         `xml:"MatriculaCarro,attr" bson:"MatriculaCarro"`
	GuiaCarro           string                         `xml:"GuiaCarro,attr" bson:"GuiaCarro"`
	ToneladasNetasCarro float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetasCarro"`
	Contenedor          *[]ContenedorCarroCartaPorte31 `xml:"Contenedor" bson:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte31 struct {
	TipoContenedor      string  `xml:"TipoContenedor,attr" bson:"TipoContenedor"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte31 struct {
	TiposFigura []*TiposFiguraCartaPorte31 `xml:"TiposFigura" bson:"TiposFigura"`
}

type TiposFiguraCartaPorte31 struct {
	TipoFigura             string                          `xml:"TipoFigura,attr" bson:"TipoFigura"`
	RFCFigura              *string                         `xml:"RFCFigura,attr" bson:"RFCFigura,omitempty"`
	NumLicencia            *string                         `xml:"NumLicencia,attr" bson:"NumLicencia,omitempty"`
	NombreFigura           string                          `xml:"NombreFigura,attr" bson:"NombreFigura"`
	NumRegIdTribFigura     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTribFigura,omitempty"`
	ResidenciaFiscalFigura *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscalFigura,omitempty"`
	PartesTransporte       *[]PartesTransporteCartaPorte31 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty"`
	Domicilio              *DomicilioCartaPorte31          `xml:"Domicilio" bson:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte31 struct {
	ParteTransporte string `xml:"ParteTransporte,attr" bson:"ParteTransporte"`
}

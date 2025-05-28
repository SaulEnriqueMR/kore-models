package cartaporte

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type CartaPorte31 struct {
	Version                   string                         `xml:"Version,attr" bson:"Version" json:"Version"`
	IdCCP                     string                         `xml:"IdCCP,attr" bson:"IdCCP" json:"IdCCP"`
	IdCcp                     string                         `bson:"IdCcp" json:"IdCcp"`
	TransporteInternacional   string                         `xml:"TranspInternac,attr" bson:"TransporteInternacional" json:"TransporteInternacional"`
	EsTransporteInternacional bool                           `bson:"EsTransporteInternacional" json:"EsTransporteInternacional"`
	EntradaSalidaMercancia    *string                        `xml:"EntradaSalidaMerc,attr" bson:"EntradaSalidaMercancia,omitempty" json:"EntradaSalidaMercancia,omitempty"`
	PaisOrigenDestino         *string                        `xml:"PaisOrigenDestino,attr" bson:"PaisOrigenDestino,omitempty" json:"PaisOrigenDestino,omitempty"`
	ViaEntradaSalida          *string                        `xml:"ViaEntradaSalida,attr" bson:"ViaEntradaSalida,omitempty" json:"ViaEntradaSalida,omitempty"`
	TotalDistanciaRecorrida   *float64                       `xml:"TotalDistRec,attr" bson:"TotalDistanciaRecorrida,omitempty" json:"TotalDistanciaRecorrida,omitempty"`
	RegistroIstmo             *string                        `xml:"RegistroISTMO,attr" bson:"RegistroIstmo,omitempty" json:"RegistroIstmo,omitempty"`
	UbicacionPoloOrigen       *string                        `xml:"UbicacionPoloOrigen,attr" bson:"UbicacionPoloOrigen,omitempty" json:"UbicacionPoloOrigen,omitempty"`
	UbicacionPoloDestino      *string                        `xml:"UbicacionPoloDestino,attr" bson:"UbicacionPoloDestino,omitempty" json:"UbicacionPoloDestino,omitempty"`
	RegimenesAduaneros        *[]RegimenAduaneroCartaPorte31 `xml:"RegimenesAduaneros>RegimenAduaneroCCP" bson:"RegimenesAduaneros,omitempty" json:"RegimenesAduaneros,omitempty"`
	Ubicaciones               []UbicacionCartaPorte31        `xml:"Ubicaciones>Ubicacion" bson:"Ubicaciones" json:"Ubicaciones"`
	Mercancias                MercanciasCartaPorte31         `xml:"Mercancias" bson:"Mercancias" json:"Mercancias"`
	FiguraTransporte          *FiguraTransporteCartaPorte31  `xml:"FiguraTransporte" bson:"FiguraTransporte,omitempty" json:"FiguraTransporte,omitempty"`
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

	ccp.IdCcp = strings.ToUpper(aux.IdCCP)

	return nil
}

type RegimenAduaneroCartaPorte31 struct {
	RegimenAduanero string `xml:"RegimenAduanero,attr" bson:"RegimenAduanero" json:"RegimenAduanero"`
}

type UbicacionCartaPorte31 struct {
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
	Domicilio                    *DomicilioCartaPorte31 `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
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

func (u *UbicacionCartaPorte31) UnmarshalJSON(data []byte) error {
	type Alias UbicacionCartaPorte31
	var aux Alias

	// Unmarshal the XML into the alias
	if err := json.Unmarshal(data, &aux); err != nil {
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

type MercanciasCartaPorte31 struct {
	PesoBrutoTotal                        float64                            `xml:"PesoBrutoTotal,attr" bson:"PesoBrutoTotal" json:"PesoBrutoTotal"`
	UnidadPeso                            string                             `xml:"UnidadPeso,attr" bson:"UnidadPeso" json:"UnidadPeso"`
	PesoNetoTotal                         *float64                           `xml:"PesoNetoTotal,attr" bson:"PesoNetoTotal,omitempty" json:"PesoNetoTotal,omitempty"`
	NumeroTotalMercancias                 float64                            `xml:"NumTotalMercancias,attr" bson:"NumeroTotalMercancias" json:"NumeroTotalMercancias"`
	CargoPorTasacion                      *float64                           `xml:"CargoPorTasacion,attr" bson:"CargoPorTasacion,omitempty" json:"CargoPorTasacion,omitempty"`
	LogisticaInversaRecoleccionDevolucion *string                            `xml:"LogisticaInversaRecoleccionDevolucion,attr" bson:"LogisticaInversaRecoleccionDevolucion,omitempty" json:"LogisticaInversaRecoleccionDevolucion,omitempty"`
	Mercancia                             []MercanciaCartaPorte31            `xml:"Mercancia" bson:"Mercancia" json:"Mercancia"`
	Autotransporte                        *AutotransporteCartaPorte31        `xml:"Autotransporte" bson:"Autotransporte,omitempty" json:"Autotransporte,omitempty"`
	TransporteMaritimo                    *TransporteMaritimoCartaPorte31    `xml:"TransporteMaritimo" bson:"TransporteMaritimo,omitempty" json:"TransporteMaritimo,omitempty"`
	TransporteAereo                       *TransporteAereoCartaPorte31       `xml:"TransporteAereo" bson:"TransporteAereo,omitempty" json:"TransporteAereo,omitempty"`
	TransporteFerroviario                 *TransporteFerroviarioCartaPorte31 `xml:"TransporteFerroviario" bson:"TransporteFerroviario,omitempty" json:"TransporteFerroviario,omitempty"`
}

type MercanciaCartaPorte31 struct {
	// Attrs
	BienesTransportados                    string                               `xml:"BienesTransp,attr" bson:"BienesTransportados" json:"BienesTransportados"`
	ClaveStcc                              *string                              `xml:"ClaveSTCC,attr" bson:"ClaveStcc,omitempty" json:"ClaveStcc,omitempty"`
	Descripcion                            string                               `xml:"Descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	Cantidad                               float64                              `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	ClaveUnidad                            string                               `xml:"ClaveUnidad,attr" bson:"ClaveUnidad" json:"ClaveUnidad"`
	Unidad                                 *string                              `xml:"Unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	Dimensiones                            *string                              `xml:"Dimensiones,attr" bson:"Dimensiones,omitempty" json:"Dimensiones,omitempty"`
	MaterialPeligroso                      *string                              `xml:"MaterialPeligroso,attr" bson:"MaterialPeligroso,omitempty" json:"MaterialPeligroso,omitempty"`
	EsMaterialPeligroso                    *bool                                `bson:"EsMaterialPeligroso,omitempty" json:"EsMaterialPeligroso,omitempty"`
	ClaveMaterialPeligroso                 *string                              `xml:"CveMaterialPeligroso,attr" bson:"ClaveMaterialPeligroso,omitempty" json:"ClaveMaterialPeligroso,omitempty"`
	Embalaje                               *string                              `xml:"Embalaje,attr" bson:"Embalaje,omitempty" json:"Embalaje,omitempty"`
	DescripcionEmbalaje                    *string                              `xml:"DescripEmbalaje,attr" bson:"DescripcionEmbalaje,omitempty" json:"DescripcionEmbalaje,omitempty"`
	SectorCofepris                         *string                              `xml:"SectorCOFEPRIS,attr" bson:"SectorCofepris,omitempty" json:"SectorCofepris,omitempty"`
	NombreIngredienteActivo                *string                              `xml:"NombreIngredienteActivo,attr" bson:"NombreIngredienteActivo,omitempty" json:"NombreIngredienteActivo,omitempty"`
	NombreQuimico                          *string                              `xml:"NomQuimico,attr" bson:"NombreQuimico,omitempty" json:"NombreQuimico,omitempty"`
	DenominacionGenericaProducto           *string                              `xml:"DenominacionGenericaProd,attr" bson:"DenominacionGenericaProducto,omitempty" json:"DenominacionGenericaProducto,omitempty"`
	DenominacionDistintivaProducto         *string                              `xml:"DenominacionDistintivaProd,attr" bson:"DenominacionDistintivaProducto,omitempty" json:"DenominacionDistintivaProducto,omitempty"`
	Fabricante                             *string                              `xml:"Fabricante,attr" bson:"Fabricante,omitempty" json:"Fabricante,omitempty"`
	FechaCaducidadString                   *string                              `xml:"FechaCaducidad,attr" bson:"FechaCaducidadString" json:"FechaCaducidadString"`
	FechaCaducidad                         *time.Time                           `bson:"FechaCaducidad,omitempty" json:"FechaCaducidad,omitempty"`
	LoteMedicamento                        *string                              `xml:"LoteMedicamento,attr" bson:"LoteMedicamento,omitempty" json:"LoteMedicamento,omitempty"`
	FormaFarmaceutica                      *string                              `xml:"FormaFarmaceutica,attr" bson:"FormaFarmaceutica,omitempty" json:"FormaFarmaceutica,omitempty"`
	CondicionesEspecialesTransporte        *string                              `xml:"CondicionesEspTransp,attr" bson:"CondicionesEspecialesTransporte,omitempty" json:"CondicionesEspecialesTransporte,omitempty"`
	RegistroSanitarioFolioAutorizacion     *string                              `xml:"RegistroSanitarioFolioAutorizacion,attr" bson:"RegistroSanitarioFolioAutorizacion,omitempty" json:"RegistroSanitarioFolioAutorizacion,omitempty"`
	PermisoImportacion                     *string                              `xml:"PermisoImportacion,attr" bson:"PermisoImportacion,omitempty" json:"PermisoImportacion,omitempty"`
	FolioImportacionVucem                  *string                              `xml:"FolioImpoVUCEM,attr" bson:"FolioImportacionVucem,omitempty" json:"FolioImportacionVucem,omitempty"`
	NoCas                                  *string                              `xml:"NumCAS,attr" bson:"NoCas,omitempty" json:"NoCas,omitempty"`
	RazonSocialEmpresaImportadora          *string                              `xml:"RazonSocialEmpImp,attr" bson:"RazonSocialEmpresaImportadora,omitempty" json:"RazonSocialEmpresaImportadora,omitempty"`
	NoRegistroSanitarioPlaguicidasCofepris *string                              `xml:"NumRegSanPlagCOFEPRIS,attr" bson:"NoRegistroSanitarioPlaguicidasCofepris,omitempty" json:"NoRegistroSanitarioPlaguicidasCofepris,omitempty"`
	DatosFabricante                        *string                              `xml:"DatosFabricante,attr" bson:"DatosFabricante,omitempty" json:"DatosFabricante,omitempty"`
	DatosFormulador                        *string                              `xml:"DatosFormulador,attr" bson:"DatosFormulador,omitempty" json:"DatosFormulador,omitempty"`
	DatosMaquilador                        *string                              `xml:"DatosMaquilador,attr" bson:"DatosMaquilador,omitempty" json:"DatosMaquilador,omitempty"`
	UsoAutorizado                          *string                              `xml:"UsoAutorizado,attr" bson:"UsoAutorizado,omitempty" json:"UsoAutorizado,omitempty"`
	PesoKg                                 float64                              `xml:"PesoEnKg,attr" bson:"PesoKg" json:"PesoKg"`
	ValorMercancia                         *float64                             `xml:"ValorMercancia,attr" bson:"ValorMercancia,omitempty" json:"ValorMercancia,omitempty"`
	Moneda                                 *string                              `xml:"Moneda,attr" bson:"Moneda,omitempty" json:"Moneda,omitempty"`
	FraccionArancelaria                    *string                              `xml:"FraccionArancelaria,attr" bson:"FraccionArancelaria,omitempty" json:"FraccionArancelaria,omitempty"`
	UUIDComercioExt                        *string                              `xml:"UUIDComercioExt,attr" bson:"UUIDComercioExt,omitempty" json:"UUIDComercioExt,omitempty"`
	UuidComercioExterior                   *string                              `bson:"UuidComercioExterior,omitempty" json:"UuidComercioExterior,omitempty"`
	TipoMateria                            *string                              `xml:"TipoMateria,attr" bson:"TipoMateria,omitempty" json:"TipoMateria,omitempty"`
	DescripcionMateria                     *string                              `xml:"DescripcionMateria,attr" bson:"DescripcionMateria,omitempty" json:"DescripcionMateria,omitempty"`
	DocumentacionAduanera                  *[]DocumentacionAduaneraCartaPorte31 `xml:"DocumentacionAduanera" bson:"DocumentacionAduanera,omitempty" json:"DocumentacionAduanera,omitempty"`
	GuiasIdentificacion                    *[]GuiasIdentificacionCartaPorte31   `xml:"GuiasIdentificacion" bson:"GuiasIdentificacion,omitempty" json:"GuiasIdentificacion,omitempty"`
	CantidadTransporta                     *[]CantidadTransporteCartaPorte31    `xml:"CantidadTransporta" bson:"CantidadTransporta,omitempty" json:"CantidadTransporta,omitempty"`
	DetalleMercancia                       *DetalleMercanciaCartaPorte31        `xml:"DetalleMercancia" bson:"DetalleMercancia,omitempty" json:"DetalleMercancia,omitempty"`
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
	if m.UUIDComercioExt != nil && *m.UUIDComercioExt != "" {
		uuidComercioExterior := strings.ToUpper(*m.UUIDComercioExt)
		m.UuidComercioExterior = &uuidComercioExterior
	}
	return nil
}

type DocumentacionAduaneraCartaPorte31 struct {
	Tipo                   string  `xml:"TipoDocumento,attr" bson:"TipoDocumento" json:"TipoDocumento"`
	NumeroPedimento        *string `xml:"NumPedimento,attr" bson:"NumeroPedimento,omitempty" json:"NumeroPedimento,omitempty"`
	IdentificadorDocumento *string `xml:"IdentDocAduanero,attr" bson:"IdentificadorDocumento,omitempty" json:"IdentificadorDocumento,omitempty"`
	RfcImportador          *string `xml:"RFCImpo,attr" bson:"RfcImportador,omitempty" json:"RfcImportador,omitempty"`
}

type GuiasIdentificacionCartaPorte31 struct {
	Numero      string  `xml:"NumeroGuiaIdentificacion,attr" bson:"Numero" json:"Numero"`
	Descripcion string  `xml:"DescripGuiaIdentificacion,attr" bson:"Descripcion" json:"Descripcion"`
	Peso        float64 `xml:"PesoGuiaIdentificacion,attr" bson:"Peso" json:"Peso"`
}

type CantidadTransporteCartaPorte31 struct {
	Cantidad         float64 `xml:"Cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	IdOrigen         string  `xml:"IDOrigen,attr" bson:"IdOrigen" json:"IdOrigen"`
	IdDestino        string  `xml:"IDDestino,attr" bson:"IdDestino" json:"IdDestino"`
	ClavesTransporte *string `xml:"CvesTransporte,attr" bson:"ClavesTransporte" json:"ClavesTransporte"`
}

type DetalleMercanciaCartaPorte31 struct {
	UnidadPeso   string   `xml:"UnidadPesoMerc,attr" bson:"UnidadPeso" json:"UnidadPeso"`
	PesoBruto    float64  `xml:"PesoBruto,attr" bson:"PesoBruto" json:"PesoBruto"`
	PesoNeto     float64  `xml:"PesoNeto,attr" bson:"PesoNeto" json:"PesoNeto"`
	PesoTara     float64  `xml:"PesoTara,attr" bson:"PesoTara" json:"PesoTara"`
	NumeroPiezas *float64 `xml:"NumPiezas,attr" bson:"NumeroPiezas,omitempty" json:"NumeroPiezas,omitempty"`
}

type AutotransporteCartaPorte31 struct {
	PermisoSct              string                              `xml:"PermSCT,attr" bson:"PermisoSct" json:"PermisoSct"`
	NoPermisoSct            string                              `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct" json:"NoPermisoSct"`
	IdentificacionVehicular IdentificacionVehicularCartaPorte31 `xml:"IdentificacionVehicular" bson:"IdentificacionVehicular" json:"IdentificacionVehicular"`
	Seguros                 SegurosCartaPorte31                 `xml:"Seguros" bson:"Seguros" json:"Seguros"`
	Remolques               *[]RemolqueCartaPorte31             `xml:"Remolques>Remolque" bson:"Remolques,omitempty" json:"Remolques,omitempty"`
}

type IdentificacionVehicularCartaPorte31 struct {
	Configuracion string  `xml:"ConfigVehicular,attr" bson:"Configuracion" json:"Configuracion"`
	PesoBruto     float64 `xml:"PesoBrutoVehicular,attr" bson:"PesoBruto" json:"PesoBruto"`
	PlacaVm       string  `xml:"PlacaVM,attr" bson:"PlacaVm" json:"PlacaVm"`
	AnioModeloVm  float64 `xml:"AnioModeloVM,attr" bson:"AnioModeloVm" json:"AnioModeloVm"`
}

type SegurosCartaPorte31 struct {
	AseguradoraResponsabilidadCivil string   `xml:"AseguraRespCivil,attr" bson:"AseguradoraResponsabilidadCivil" json:"AseguradoraResponsabilidadCivil"`
	PolizaResponsabilidadCivil      string   `xml:"PolizaRespCivil,attr" bson:"PolizaResponsabilidadCivil" json:"PolizaResponsabilidadCivil"`
	AseguradoraMedioAmbiente        *string  `xml:"AseguraMedAmbiente,attr" bson:"AseguradoraMedioAmbiente,omitempty" json:"AseguradoraMedioAmbiente,omitempty"`
	PolizaMedioAmbiente             *string  `xml:"PolizaMedAmbiente,attr" bson:"PolizaMedioAmbiente,omitempty" json:"PolizaMedioAmbiente,omitempty"`
	AseguradoraCarga                *string  `xml:"AseguraCarga,attr" bson:"AseguradoraCarga,omitempty" json:"AseguradoraCarga,omitempty"`
	PolizaCarga                     *string  `xml:"PolizaCarga,attr" bson:"PolizaCarga,omitempty" json:"PolizaCarga,omitempty"`
	PrimaSeguro                     *float64 `xml:"PrimaSeguro,attr" bson:"PrimaSeguro,omitempty" json:"PrimaSeguro,omitempty"`
}

type RemolqueCartaPorte31 struct {
	Subtipo string `xml:"SubTipoRem,attr" bson:"Subtipo" json:"Subtipo"`
	Placa   string `xml:"Placa,attr" bson:"Placa" json:"Placa"`
}

type TransporteMaritimoCartaPorte31 struct {
	PermisoSct                *string                   `xml:"PermSCT,attr" bson:"PermisoSct,omitempty" json:"PermisoSct,omitempty"`
	NoPermisoSct              *string                   `xml:"NumPermisoSCT,attr" bson:"NoPermisoSct,omitempty" json:"NoPermisoSct,omitempty"`
	NombreAseguradora         *string                   `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty" json:"NombreAseguradora,omitempty"`
	NoPolizaSeguro            *string                   `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty" json:"NoPolizaSeguro,omitempty"`
	TipoEmbarcacion           string                    `xml:"TipoEmbarcacion,attr" bson:"TipoEmbarcacion" json:"TipoEmbarcacion"`
	Matricula                 string                    `xml:"Matricula,attr" bson:"Matricula" json:"Matricula"`
	NoOmi                     string                    `xml:"NumeroOMI,attr" bson:"NoOmi" json:"NoOmi"`
	AnioEmbarcacion           *float64                  `xml:"AnioEmbarcacion,attr" bson:"AnioEmbarcacion,omitempty" json:"AnioEmbarcacion,omitempty"`
	NombreEmbarcacion         *string                   `xml:"NombreEmbarc,attr" bson:"NombreEmbarcacion,omitempty" json:"NombreEmbarcacion,omitempty"`
	NacionalidadEmbarcacion   string                    `xml:"NacionalidadEmbarc,attr" bson:"NacionalidadEmbarcacion" json:"NacionalidadEmbarcacion"`
	UnidadesArqueoBruto       float64                   `xml:"UnidadesDeArqBruto,attr" bson:"UnidadesArqueoBruto" json:"UnidadesArqueoBruto"`
	TipoCarga                 string                    `xml:"TipoCarga,attr" bson:"TipoCarga" json:"TipoCarga"`
	Eslora                    *float64                  `xml:"Eslora,attr" bson:"Eslora,omitempty" json:"Eslora,omitempty"`
	Manga                     *float64                  `xml:"Manga,attr" bson:"Manga,omitempty" json:"Manga,omitempty"`
	Calado                    *float64                  `xml:"Calado,attr" bson:"Calado,omitempty" json:"Calado,omitempty"`
	Puntal                    *float64                  `xml:"Puntal,attr" bson:"Puntal,omitempty" json:"Puntal,omitempty"`
	LineaNaviera              *string                   `xml:"LineaNaviera,attr" bson:"LineaNaviera,omitempty" json:"LineaNaviera,omitempty"`
	NombreAgenteNaviero       string                    `xml:"NombreAgenteNaviero,attr" bson:"NombreAgenteNaviero" json:"NombreAgenteNaviero"`
	NoAutorizacionNaviero     string                    `xml:"NumAutorizacionNaviero,attr" bson:"NoAutorizacionNaviero" json:"NoAutorizacionNaviero"`
	NoViaje                   *string                   `xml:"NumViaje,attr" bson:"NoViaje,omitempty" json:"NoViaje,omitempty"`
	NoConocimientoEmbarque    *string                   `xml:"NumConocEmbarc,attr" bson:"NoConocimientoEmbarque,omitempty" json:"NoConocimientoEmbarque,omitempty"`
	PermisoTemporalNavegacion *string                   `xml:"PermisoTempNavegacion,attr" bson:"PermisoTemporalNavegacion,omitempty" json:"PermisoTemporalNavegacion,omitempty"`
	Contenedor                *[]ContenedorCartaPorte31 `xml:"Contenedor" bson:"Contenedor,omitempty" json:"Contenedor,omitempty"`
}

type ContenedorCartaPorte31 struct {
	Tipo                  string                     `xml:"TipoContenedor,attr" bson:"Tipo" json:"Tipo"`
	Matricula             *string                    `xml:"MatriculaContenedor,attr" bson:"Matricula,omitempty" json:"Matricula,omitempty"`
	NoPrecinto            *string                    `xml:"NumPrecinto,attr" bson:"NoPrecinto,omitempty" json:"NoPrecinto,omitempty"`
	IdCcpRelacionado      *string                    `xml:"IdCCPRelacionado,attr" bson:"IdCcpRelacionado,omitempty" json:"IdCcpRelacionado,omitempty"`
	PlacaVmCcp            *string                    `xml:"PlacaVMCCP,attr" bson:"PlacaVmCcp,omitempty" json:"PlacaVmCcp,omitempty"`
	FechaCertificacionCCP *string                    `xml:"FechaCertificacionCCP,attr" bson:"FechaCertificacionCCP" json:"FechaCertificacionCCP"`
	FechaCertificacionCcp *time.Time                 `bson:"FechaCertificacionCcp,omitempty" json:"FechaCertificacionCcp,omitempty"`
	RemolquesCcp          *[]RemolqueCCPCartaPorte31 `xml:"RemolquesCCP>RemolqueCCP" bson:"RemolquesCcp,omitempty" json:"RemolquesCcp,omitempty"`
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

	if c.FechaCertificacionCCP != nil && *c.FechaCertificacionCCP != "" {
		fecha, errFecha := helpers.ParseDatetime(*c.FechaCertificacionCCP)
		if errFecha == nil {
			c.FechaCertificacionCcp = &fecha
		}
	}
	return nil
}

func (c *ContenedorCartaPorte31) UnmarshalJSON(data []byte) error {
	// Create an alias to avoid recursion
	type Alias ContenedorCartaPorte31
	var aux Alias

	// Unmarshal the XML into the alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*c = ContenedorCartaPorte31(aux)

	if c.FechaCertificacionCCP != nil && *c.FechaCertificacionCCP != "" {
		fecha, errFecha := helpers.ParseDatetime(*c.FechaCertificacionCCP)
		if errFecha == nil {
			c.FechaCertificacionCcp = &fecha
		}
	}
	return nil
}

type RemolqueCCPCartaPorte31 struct {
	Subtipo string `xml:"SubTipoRemCCP,attr" bson:"Subtipo" json:"Subtipo"`
	Placa   string `xml:"PlacaCCP,attr" bson:"Placa" json:"Placa"`
}

type TransporteAereoCartaPorte31 struct {
	PermisoSct                 string  `xml:"PermSCT,attr" bson:"PermisoSct" json:"PermisoSct"`
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

type TransporteFerroviarioCartaPorte31 struct {
	TipoServicio      string                        `xml:"TipoDeServicio,attr" bson:"TipoServicio" json:"TipoServicio"`
	TipoTrafico       string                        `xml:"TipoDeTrafico,attr" bson:"TipoTrafico" json:"TipoTrafico"`
	NombreAseguradora *string                       `xml:"NombreAseg,attr" bson:"NombreAseguradora,omitempty" json:"NombreAseguradora,omitempty"`
	NoPolizaSeguro    *string                       `xml:"NumPolizaSeguro,attr" bson:"NoPolizaSeguro,omitempty" json:"NoPolizaSeguro,omitempty"`
	DerechosDePaso    *[]DerechosDePasoCartaPorte31 `xml:"DerechosDePaso" bson:"DerechosDePaso,omitempty" json:"DerechosDePaso,omitempty"`
	Carro             []CarroCartaPorte31           `xml:"Carro" bson:"Carro" json:"Carro"`
}

type DerechosDePasoCartaPorte31 struct {
	Tipo              string  `xml:"TipoDerechoDePaso,attr" bson:"Tipo" json:"Tipo"`
	KilometrajePagado float64 `xml:"KilometrajePagado,attr" bson:"KilometrajePagado" json:"KilometrajePagado"`
}

type CarroCartaPorte31 struct {
	Tipo           string                         `xml:"TipoCarro,attr" bson:"Tipo" json:"Tipo"`
	Matricula      string                         `xml:"MatriculaCarro,attr" bson:"Matricula" json:"Matricula"`
	Guia           string                         `xml:"GuiaCarro,attr" bson:"Guia" json:"Guia"`
	ToneladasNetas float64                        `xml:"ToneladasNetasCarro,attr" bson:"ToneladasNetas" json:"ToneladasNetas"`
	Contenedor     *[]ContenedorCarroCartaPorte31 `xml:"Contenedor" bson:"Contenedor,omitempty" json:"Contenedor,omitempty"`
}

type ContenedorCarroCartaPorte31 struct {
	Tipo                string  `xml:"TipoContenedor,attr" bson:"Tipo" json:"Tipo"`
	PesoContenedorVacio float64 `xml:"PesoContenedorVacio,attr" bson:"PesoContenedorVacio" json:"PesoContenedorVacio"`
	PesoNetoMercancia   float64 `xml:"PesoNetoMercancia,attr" bson:"PesoNetoMercancia" json:"PesoNetoMercancia"`
}

type FiguraTransporteCartaPorte31 struct {
	TiposFigura []*TiposFiguraCartaPorte31 `xml:"TiposFigura" bson:"TiposFigura" json:"TiposFigura"`
}

type TiposFiguraCartaPorte31 struct {
	Tipo             string                          `xml:"TipoFigura,attr" bson:"Tipo" json:"Tipo"`
	Rfc              *string                         `xml:"RFCFigura,attr" bson:"Rfc,omitempty" json:"Rfc,omitempty"`
	NoLicencia       *string                         `xml:"NumLicencia,attr" bson:"NoLicencia,omitempty" json:"NoLicencia,omitempty"`
	Nombre           string                          `xml:"NombreFigura,attr" bson:"Nombre" json:"Nombre"`
	NumRegIdTrib     *string                         `xml:"NumRegIdTribFigura,attr" bson:"NumRegIdTrib,omitempty" json:"NumRegIdTrib,omitempty"`
	ResidenciaFiscal *string                         `xml:"ResidenciaFiscalFigura,attr" bson:"ResidenciaFiscal,omitempty" json:"ResidenciaFiscal,omitempty"`
	PartesTransporte *[]PartesTransporteCartaPorte31 `xml:"PartesTransporte" bson:"PartesTransporte,omitempty" json:"PartesTransporte,omitempty"`
	Domicilio        *DomicilioCartaPorte31          `xml:"Domicilio" bson:"Domicilio,omitempty" json:"Domicilio,omitempty"`
}

type PartesTransporteCartaPorte31 struct {
	Parte string `xml:"ParteTransporte,attr" bson:"Parte" json:"Parte"`
}

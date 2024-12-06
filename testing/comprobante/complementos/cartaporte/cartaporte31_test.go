package cartaporte

import (
	"encoding/xml"
	"strings"
	"testing"

	cartaporte1 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cartaporte"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetCartaPorteForTest(filename string, t *testing.T) (cartaporte1.CartaPorte31, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed cartaporte1.CartaPorte31
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	testing2.GenerateJSONFromStructure("cartaporte31.json", parsed)
	return parsed, errUnmashal
}

func TestFullCartaPorte31(t *testing.T) {
	cartaporte, _ := GetCartaPorteForTest("./cartaporte31.xml", t)
	assert.NotNil(t, cartaporte)
	InternalTestCartaPorte(t, cartaporte)
}

func InternalTestCartaPorte(t *testing.T, carta cartaporte1.CartaPorte31) {
	assert.Equal(t, "3.1", carta.Version)
	assert.Equal(t, strings.ToUpper("CCCfa34e2-5bcd-4a5f-b8de-66bdeb4a2e12"), carta.IdCcp)
	assert.Equal(t, "CCCfa34e2-5bcd-4a5f-b8de-66bdeb4a2e12", carta.IdCCP)
	assert.Equal(t, "Sí", carta.TransporteInternacional)
	assert.Equal(t, true, carta.EsTransporteInternacional)
	assert.NotNil(t, carta.EntradaSalidaMercancia)
	assert.Equal(t, "Entrada", *carta.EntradaSalidaMercancia)
	assert.Equal(t, "MEX", *carta.PaisOrigenDestino)
	assert.Equal(t, "01", *carta.ViaEntradaSalida)
	assert.Equal(t, 1500.00, *carta.TotalDistanciaRecorrida)
	assert.Equal(t, "Sí", *carta.RegistroIstmo)
	assert.Equal(t, "001", *carta.UbicacionPoloOrigen)
	assert.Equal(t, "002", *carta.UbicacionPoloDestino)
	InternalTestRegimenAduanero(t, carta.RegimenesAduaneros)
	InternalTestUbicacion(t, carta.Ubicaciones)
	InternalTestMercancias(t, carta.Mercancias)
	InternalTestTiposFigura(t, carta.FiguraTransporte)
}

func InternalTestRegimenAduanero(t *testing.T, regimenes *[]cartaporte1.RegimenAduaneroCartaPorte31) {
	assert.NotNil(t, regimenes)
	assert.Equal(t, len(*regimenes), 3)
	regimen := (*regimenes)[0]
	assert.Equal(t, "A1", regimen.RegimenAduanero)
}

func InternalTestUbicacion(t *testing.T, ubicaciones []cartaporte1.UbicacionCartaPorte31) {
	assert.NotNil(t, ubicaciones)
	assert.Equal(t, len(ubicaciones), 3)
	ubicacion := ubicaciones[0]
	assert.Equal(t, "Origen", ubicacion.TipoUbicacion)
	assert.Equal(t, "OR123456", *ubicacion.IdUbicacion)
	assert.Equal(t, "XAXX010101000", ubicacion.RfcRemitenteDestinatario)
	assert.Equal(t, "Empresa Origen", *ubicacion.NombreRemitenteDestinatario)
	assert.Equal(t, "123456", *ubicacion.NumRegIdTrib)
	assert.Equal(t, "MEX", *ubicacion.ResidenciaFiscal)
	assert.Equal(t, "E123456", *ubicacion.NoEstacion)
	assert.Equal(t, "Estación 1", *ubicacion.NombreEstacion)
	assert.Equal(t, "Altura", *ubicacion.NavegacionTrafico)
	assert.Equal(t, "2024-09-20T08:30:00", ubicacion.FechaHoraSalidaLlegadaString)
	assert.Equal(t, "Terrestre", *ubicacion.TipoEstacion)
	assert.Equal(t, 150.50, *ubicacion.DistanciaRecorrida)
	InternalTestDomicilio(t, ubicacion.Domicilio)
}

func InternalTestDomicilio(t *testing.T, domicilio *cartaporte1.DomicilioCartaPorte31) {
	assert.Equal(t, "Avenida Principal", *domicilio.Calle)
	assert.Equal(t, "123", *domicilio.NoExterior)
	assert.Equal(t, "A", *domicilio.NoInterior)
	assert.Equal(t, "Colonia Centro", *domicilio.Colonia)
	assert.Equal(t, "Ciudad de México", *domicilio.Localidad)
	assert.Equal(t, "Cerca del parque", *domicilio.Referencia)
	assert.Equal(t, "Cuauhtémoc", *domicilio.Municipio)
	assert.Equal(t, "CDMX", domicilio.Estado)
	assert.Equal(t, "MEX", domicilio.Pais)
	assert.Equal(t, "01000", domicilio.CodigoPostal)
}

func InternalTestMercancias(t *testing.T, mercancias cartaporte1.MercanciasCartaPorte31) {
	assert.Equal(t, 1500.250, mercancias.PesoBrutoTotal)
	assert.Equal(t, "KGM", mercancias.UnidadPeso)
	assert.Equal(t, 1400.200, *mercancias.PesoNetoTotal)
	assert.Equal(t, 5.0, mercancias.NumeroTotalMercancias)
	assert.Equal(t, 200.00, *mercancias.CargoPorTasacion)
	assert.Equal(t, "Sí", *mercancias.LogisticaInversaRecoleccionDevolucion)
	InternalTestMercancia(t, mercancias.Mercancia)
	InternalTestAutotransporte(t, mercancias.Autotransporte)
	InternalTestTransporteMaritimo(t, mercancias.TransporteMaritimo)
	InternalTestTransporteAereo(t, mercancias.TransporteAereo)
	InternalTestTransporteFerroviario(t, mercancias.TransporteFerroviario)
}

func InternalTestMercancia(t *testing.T, mercancias []cartaporte1.MercanciaCartaPorte31) {
	assert.NotNil(t, mercancias)
	assert.Equal(t, len(mercancias), 1)
	mercancia := mercancias[0]
	assert.Equal(t, "123456", mercancia.BienesTransportados)
	assert.Equal(t, "1234567", *mercancia.ClaveStcc)
	assert.Equal(t, "Medicamento XYZ", mercancia.Descripcion)
	assert.Equal(t, 10.0, mercancia.Cantidad)
	assert.Equal(t, "KGM", mercancia.ClaveUnidad)
	assert.Equal(t, "Kilogramos", *mercancia.Unidad)
	assert.Equal(t, "30/40/30cm", *mercancia.Dimensiones)
	assert.Equal(t, "No", *mercancia.MaterialPeligroso)
	assert.Equal(t, "C1", *mercancia.ClaveMaterialPeligroso)
	assert.Equal(t, "Caja", *mercancia.Embalaje)
	assert.Equal(t, "Caja de cartón", *mercancia.DescripcionEmbalaje)
	assert.Equal(t, "No", *mercancia.SectorCofepris)
	assert.Equal(t, "N/A", *mercancia.NombreIngredienteActivo)
	assert.Equal(t, "N/A", *mercancia.NombreQuimico)
	assert.Equal(t, "N/A", *mercancia.DenominacionGenericaProducto)
	assert.Equal(t, "Marca ABC", *mercancia.DenominacionDistintivaProducto)
	assert.Equal(t, "Laboratorio XYZ", *mercancia.Fabricante)
	assert.Equal(t, "2025-12-31", *mercancia.FechaCaducidadString)
	assert.Equal(t, "L123456", *mercancia.LoteMedicamento)
	assert.Equal(t, "Tableta", *mercancia.FormaFarmaceutica)
	assert.Equal(t, "N/A", *mercancia.CondicionesEspecialesTransporte)
	assert.Equal(t, "RS123456", *mercancia.RegistroSanitarioFolioAutorizacion)
	assert.Equal(t, "P1234", *mercancia.PermisoImportacion)
	assert.Equal(t, "VUCEM123456", *mercancia.FolioImportacionVucem)
	assert.Equal(t, "123-45-6", *mercancia.NoCas)
	assert.Equal(t, "Importadora ABC S.A.", *mercancia.RazonSocialEmpresaImportadora)
	assert.Equal(t, "RS-ABC-123", *mercancia.NoRegistroSanitarioPlaguicidasCofepris)
	assert.Equal(t, "México, Laboratorio XYZ", *mercancia.DatosFabricante)
	assert.Equal(t, "México, Formulador ABC", *mercancia.DatosFormulador)
	assert.Equal(t, "México, Maquilador XYZ", *mercancia.DatosMaquilador)
	assert.Equal(t, "Agricultura, plagas específicas", *mercancia.UsoAutorizado)
	assert.Equal(t, 1.5, mercancia.PesoKg)
	assert.Equal(t, 1000.00, *mercancia.ValorMercancia)
	assert.Equal(t, "MXN", *mercancia.Moneda)
	assert.Equal(t, "123456", *mercancia.FraccionArancelaria)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", *mercancia.UUIDComercioExt)
	assert.Equal(t, strings.ToUpper("123e4567-e89b-12d3-a456-426614174000"), *mercancia.UuidComercioExterior)
	assert.Equal(t, "Producto", *mercancia.TipoMateria)
	assert.Equal(t, "Tabletas de medicamento", *mercancia.DescripcionMateria)
	InternalTestDocumentacionAduanera(t, mercancia.DocumentacionAduanera)
	InternalTestGuiasIdentificacion(t, mercancia.GuiasIdentificacion)
	InternalTestCantidadTransporta(t, mercancia.CantidadTransporta)
	InternalTestDetalleMercancia(t, mercancia.DetalleMercancia)
}

func InternalTestDocumentacionAduanera(t *testing.T, docs *[]cartaporte1.DocumentacionAduaneraCartaPorte31) {
	assert.NotNil(t, docs)
	assert.Equal(t, len(*docs), 4)
	doc := (*docs)[0]
	assert.Equal(t, "tipo_doc_aduanero", doc.Tipo)
	assert.Equal(t, "22  99  1234  0000000", *doc.NumeroPedimento)
	assert.Equal(t, "Folio123456", *doc.IdentificadorDocumento)
	assert.Equal(t, "XAXX010101000", *doc.RfcImportador)
}

func InternalTestGuiasIdentificacion(t *testing.T, guias *[]cartaporte1.GuiasIdentificacionCartaPorte31) {
	assert.NotNil(t, guias)
	assert.Equal(t, len(*guias), 3)
	guia := (*guias)[0]
	assert.Equal(t, "GU1234567890", guia.Numero)
	assert.Equal(t, "Paquete de medicamentos", guia.Descripcion)
	assert.Equal(t, 10.500, guia.Peso)
}

func InternalTestCantidadTransporta(t *testing.T, cantidades *[]cartaporte1.CantidadTransporteCartaPorte31) {
	assert.NotNil(t, cantidades)
	assert.Equal(t, len(*cantidades), 3)
	cantidad := (*cantidades)[0]
	assert.Equal(t, 100.500, cantidad.Cantidad)
	assert.Equal(t, "OR123456", cantidad.IdOrigen)
	assert.Equal(t, "DE654321", cantidad.IdDestino)
	assert.Equal(t, "CVE01", *cantidad.ClavesTransporte)
}

func InternalTestDetalleMercancia(t *testing.T, detalle *cartaporte1.DetalleMercanciaCartaPorte31) {
	assert.NotNil(t, detalle)
	assert.Equal(t, "KG", detalle.UnidadPeso)
	assert.Equal(t, 1500.500, detalle.PesoBruto)
	assert.Equal(t, 1450.250, detalle.PesoNeto)
	assert.Equal(t, 50.250, detalle.PesoTara)
	assert.Equal(t, 10.0, *detalle.NumeroPiezas)
}

func InternalTestAutotransporte(t *testing.T, autotransporte *cartaporte1.AutotransporteCartaPorte31) {
	assert.Equal(t, "TIPO_PERMISO", autotransporte.PermisoSct)
	assert.Equal(t, "NUM_PERMISO", autotransporte.NoPermisoSct)
	InternalTestIdentificacionVehicular(t, autotransporte.IdentificacionVehicular)
	InternalTestSeguros(t, autotransporte.Seguros)
	InternalTestRemolque(t, autotransporte.Remolques)
}

func InternalTestIdentificacionVehicular(t *testing.T, vehiculo cartaporte1.IdentificacionVehicularCartaPorte31) {
	assert.Equal(t, "CLAVE_CONFIG", vehiculo.Configuracion)
	assert.Equal(t, 20.00, vehiculo.PesoBruto)
	assert.Equal(t, "ABC1234", vehiculo.PlacaVm)
	assert.Equal(t, 2020.0, vehiculo.AnioModeloVm)
}

func InternalTestSeguros(t *testing.T, seguros cartaporte1.SegurosCartaPorte31) {
	assert.Equal(t, "Nombre Aseguradora", seguros.AseguradoraResponsabilidadCivil)
	assert.Equal(t, "POLIZA_123456", seguros.PolizaResponsabilidadCivil)
	assert.Equal(t, "Nombre Aseguradora Ambiental", *seguros.AseguradoraMedioAmbiente)
	assert.Equal(t, "POLIZA_AMB_123", *seguros.PolizaMedioAmbiente)
	assert.Equal(t, "Nombre Aseguradora Carga", *seguros.AseguradoraCarga)
	assert.Equal(t, "POLIZA_CARGA_123", *seguros.PolizaCarga)
	assert.Equal(t, 1500.00, *seguros.PrimaSeguro)
}

func InternalTestRemolque(t *testing.T, remolques *[]cartaporte1.RemolqueCartaPorte31) {
	assert.NotNil(t, remolques)
	assert.Equal(t, len(*remolques), 2)
	remolque := (*remolques)[0]
	assert.Equal(t, "SUBTIPO_REM", remolque.Subtipo)
	assert.Equal(t, "REM1234", remolque.Placa)
}

func InternalTestTransporteMaritimo(t *testing.T, transporte *cartaporte1.TransporteMaritimoCartaPorte31) {
	assert.Equal(t, "ABC123", *transporte.PermSCT)
	assert.Equal(t, "1234567890", *transporte.NoPermisoSct)
	assert.Equal(t, "Aseguradora XYZ", *transporte.NombreAseguradora)
	assert.Equal(t, "POL123456", *transporte.NoPolizaSeguro)
	assert.Equal(t, "Tipo1", transporte.TipoEmbarcacion)
	assert.Equal(t, "MAT12345", transporte.Matricula)
	assert.Equal(t, "IMO1234567", transporte.NoOmi)
	assert.Equal(t, 2020.0, *transporte.AnioEmbarcacion)
	assert.Equal(t, "Embarcación del Mar", *transporte.NombreEmbarcacion)
	assert.Equal(t, "MX", transporte.NacionalidadEmbarcacion)
	assert.Equal(t, 100.500, transporte.UnidadesArqueoBruto)
	assert.Equal(t, "CargaGeneral", transporte.TipoCarga)
	assert.Equal(t, 150.25, *transporte.Eslora)
	assert.Equal(t, 30.50, *transporte.Manga)
	assert.Equal(t, 10.75, *transporte.Calado)
	assert.Equal(t, 15.00, *transporte.Puntal)
	assert.Equal(t, "Naviera ABC", *transporte.LineaNaviera)
	assert.Equal(t, "Agente XYZ", transporte.NombreAgenteNaviero)
	assert.Equal(t, "AUT123456", transporte.NoAutorizacionNaviero)
	assert.Equal(t, "VIAJE001", *transporte.NoViaje)
	assert.Equal(t, "CONOC123456", *transporte.NoConocimientoEmbarque)
	assert.Equal(t, "TEMP123456", *transporte.PermisoTemporalNavegacion)
	InternalTestContenedorMatirimo(t, transporte.Contenedor)
}

func InternalTestContenedorMatirimo(t *testing.T, contenedores *[]cartaporte1.ContenedorCartaPorte31) {
	assert.NotNil(t, contenedores)
	assert.Equal(t, len(*contenedores), 2)
	contenedor := (*contenedores)[0]
	assert.Equal(t, "ContenedorTipo1", contenedor.Tipo)
	assert.Equal(t, "ABCD1234567", *contenedor.Matricula)
	assert.Equal(t, "PREC12345", *contenedor.NoPrecinto)
	assert.Equal(t, "CCC1234-abcd-5678-efgh-1234567890ab", *contenedor.IdCcpRelacionado)
	assert.Equal(t, "XYZ123", *contenedor.PlacaVmCcp)
	assert.Equal(t, "2023-09-20T12:00:00", *contenedor.FechaCertificacionCCP)
	InternalTestRemolqueCCP(t, contenedor.RemolquesCcp)
}

func InternalTestRemolqueCCP(t *testing.T, remolques *[]cartaporte1.RemolqueCCPCartaPorte31) {
	assert.NotNil(t, remolques)
	assert.Equal(t, len(*remolques), 2)
	remolque := (*remolques)[0]
	assert.Equal(t, "SubTipo1", remolque.Subtipo)
	assert.Equal(t, "REM123", remolque.Placa)
}

func InternalTestTransporteAereo(t *testing.T, transporteAereo *cartaporte1.TransporteAereoCartaPorte31) {
	assert.Equal(t, "PermisoTipo1", transporteAereo.PermisoSct)
	assert.Equal(t, "1234567890", transporteAereo.NoPermisoSct)
	assert.Equal(t, "ABC-123", *transporteAereo.MatriculaAeronave)
	assert.Equal(t, "Aseguradora XYZ", *transporteAereo.NombreAseguradora)
	assert.Equal(t, "POL123456", *transporteAereo.NoPolizaSeguro)
	assert.Equal(t, "GUIDE123456789", transporteAereo.NoGuia)
	assert.Equal(t, "Ciudad de México", *transporteAereo.LugarContrato)
	assert.Equal(t, "T123", transporteAereo.CodigoTransportista)
	assert.Equal(t, "RFC123456789", *transporteAereo.RfcEmbarcador)
	assert.Equal(t, "ID123456", *transporteAereo.NumRegIdTribEmbarcador)
	assert.Equal(t, "MEX", *transporteAereo.ResidenciaFiscalEmbarcador)
	assert.Equal(t, "Embarcador S.A.", *transporteAereo.NombreEmbarcador)
}

func InternalTestTransporteFerroviario(t *testing.T, transporteFerroviario *cartaporte1.TransporteFerroviarioCartaPorte31) {
	assert.Equal(t, "01", transporteFerroviario.TipoServicio)
	assert.Equal(t, "A", transporteFerroviario.TipoTrafico)
	assert.Equal(t, "Aseguradora S.A.", *transporteFerroviario.NombreAseguradora)
	assert.Equal(t, "PZ123456789", *transporteFerroviario.NoPolizaSeguro)
	InternalTestDerechosDePaso(t, transporteFerroviario.DerechosDePaso)
	InternalTestCarroFerroviario(t, transporteFerroviario.Carro)
}

func InternalTestDerechosDePaso(t *testing.T, derechosDerechoPaso *[]cartaporte1.DerechosDePasoCartaPorte31) {
	assert.NotNil(t, derechosDerechoPaso)
	assert.Equal(t, len(*derechosDerechoPaso), 5)
	derechoDerechoPaso := (*derechosDerechoPaso)[4]
	assert.Equal(t, "D01", derechoDerechoPaso.Tipo)
	assert.Equal(t, 150.50, derechoDerechoPaso.KilometrajePagado)
}

func InternalTestCarroFerroviario(t *testing.T, carros []cartaporte1.CarroCartaPorte31) {
	assert.NotNil(t, carros)
	assert.Equal(t, len(carros), 3)
	carro := (carros)[1]
	assert.Equal(t, "C01", carro.Tipo)
	assert.Equal(t, "CAR123", carro.Matricula)
	assert.Equal(t, "G12345", carro.Guia)
	assert.Equal(t, 10.500, carro.ToneladasNetas)
	InternalTestContenedor(t, carro.Contenedor)
}

func InternalTestContenedor(t *testing.T, contenedores *[]cartaporte1.ContenedorCarroCartaPorte31) {
	assert.NotNil(t, contenedores)
	assert.Equal(t, len(*contenedores), 6)
	contenedor := (*contenedores)[4]
	assert.Equal(t, "CT06", contenedor.Tipo)
	assert.Equal(t, 2.500, contenedor.PesoContenedorVacio)
	assert.Equal(t, 8.000, contenedor.PesoNetoMercancia)
}
func InternalTestTiposFigura(t *testing.T, figuratransporte *cartaporte1.FiguraTransporteCartaPorte31) {
	var tiposfiguras = figuratransporte.TiposFigura
	assert.NotNil(t, tiposfiguras)
	assert.Equal(t, len(tiposfiguras), 2)
	tipoFigura := tiposfiguras[0]
	assert.Equal(t, "FT01", tipoFigura.Tipo)
	assert.Equal(t, "RFC123456", *tipoFigura.Rfc)
	assert.Equal(t, "LIC123456", *tipoFigura.NoLicencia)
	assert.Equal(t, "Transportista Ejemplo", tipoFigura.Nombre)
	InternalTestPartesTransporte(t, tipoFigura.PartesTransporte)
	InternalTestDomiciliFigura(t, tipoFigura.Domicilio)
}

func InternalTestPartesTransporte(t *testing.T, partesTransporte *[]cartaporte1.PartesTransporteCartaPorte31) {
	assert.NotNil(t, partesTransporte)
	assert.Equal(t, len(*partesTransporte), 5)
	parteTransporte := (*partesTransporte)[2]
	assert.Equal(t, "PT02", parteTransporte.Parte)
}

func InternalTestDomiciliFigura(t *testing.T, domicilio *cartaporte1.DomicilioCartaPorte31) {
	assert.Equal(t, "Av. Principal", *domicilio.Calle)
	assert.Equal(t, "123", *domicilio.NoExterior)
	assert.Equal(t, "Municipio Ejemplo", *domicilio.Municipio)
	assert.Equal(t, "Estado Ejemplo", domicilio.Estado)
	assert.Equal(t, "MX", domicilio.Pais)
	assert.Equal(t, "12345", domicilio.CodigoPostal)
}

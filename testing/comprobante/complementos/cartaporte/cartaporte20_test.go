package cartaporte

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	"testing"

	cartaporte2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cartaporte"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetCartaPorte20ForTest(filename string, t *testing.T) (cartaporte2.CartaPorte20, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed cartaporte2.CartaPorte20
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("cartaporte20.json", parsed)
	return parsed, errUnmarshal
}

func GetCartaPorte20ForTestJSON(filename string, t *testing.T) (cartaporte2.CartaPorte20, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed cartaporte2.CartaPorte20
	errUnmarshal := json.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestFullCartaPorte20JSON(t *testing.T) {
	cartaPorte20, _ := GetCartaPorte20ForTestJSON("./cartaporte20.json", t)
	InternalTestFullAtributesCartaPorte20(t, cartaPorte20)
	InternalTestFullAtributesUbicacionCartaPorte20(t, cartaPorte20.Ubicaciones)
	InternalTestFullAtributesMercancias(t, cartaPorte20.Mercancias)
	InternalTestFullAtributesFiguraTransporte20(t, *cartaPorte20.FiguraTransporte)
}

func TestFullCartaPorte20(t *testing.T) {
	cartaPorte20, _ := GetCartaPorte20ForTest("./cartaporte20.xml", t)
	InternalTestFullAtributesCartaPorte20(t, cartaPorte20)
	InternalTestFullAtributesUbicacionCartaPorte20(t, cartaPorte20.Ubicaciones)
	InternalTestFullAtributesMercancias(t, cartaPorte20.Mercancias)
	InternalTestFullAtributesFiguraTransporte20(t, *cartaPorte20.FiguraTransporte)
}

func InternalTestFullAtributesCartaPorte20(t *testing.T, porte20 cartaporte2.CartaPorte20) {
	assert.Equal(t, "2.0", porte20.Version)
	assert.Equal(t, "Si", porte20.TransporteInternacional)
	assert.Equal(t, true, porte20.EsTransporteInternacional)
	assert.Equal(t, "Salida", *porte20.EntradaSalidaMercancia)
	assert.Equal(t, "MEX", *porte20.PaisOrigenDestino)
	assert.Equal(t, "T1", *porte20.ViaEntradaSalida)
	assert.Equal(t, 120.50, *porte20.TotalDistanciaRecorrida)
}

func InternalTestFullAtributesUbicacionCartaPorte20(t *testing.T, ubicaciones []cartaporte2.UbicacionCartaPorte20) {
	assert.Equal(t, 2, len(ubicaciones))
	assert.Equal(t, "Origen", ubicaciones[0].TipoUbicacion)
	assert.Equal(t, "OR123456", *ubicaciones[0].IdUbicacion)
	assert.Equal(t, "XAXX010101000", ubicaciones[0].RfcRemitenteDestinatario)
	assert.Equal(t, "Empresa Origen", *ubicaciones[0].NombreRemitenteDestinatario)
	assert.Equal(t, "IdTrib1", *ubicaciones[0].NumRegIdTrib)
	assert.Equal(t, "ResidenciaFisc1", *ubicaciones[0].ResidenciaFiscal)
	assert.Equal(t, "1", *ubicaciones[0].NoEstacion)
	assert.Equal(t, "Estacion1", *ubicaciones[0].NombreEstacion)
	assert.Equal(t, "Trafico1", *ubicaciones[0].NavegacionTrafico)
	assert.Equal(t, "2024-09-23T12:30:00", ubicaciones[0].FechaHoraSalidaLlegadaString)
	assert.Equal(t, "Tipo1", *ubicaciones[0].TipoEstacion)
	assert.Equal(t, 100.50, *ubicaciones[0].DistanciaRecorrida)

	assert.Equal(t, "Origen", ubicaciones[1].TipoUbicacion)
	assert.Equal(t, "OR123456", *ubicaciones[1].IdUbicacion)
	assert.Equal(t, "XAXX010101000", ubicaciones[1].RfcRemitenteDestinatario)
	assert.Equal(t, "Empresa Origen", *ubicaciones[1].NombreRemitenteDestinatario)
	assert.Equal(t, "IdTrib2", *ubicaciones[1].NumRegIdTrib)
	assert.Equal(t, "ResidenciaFisc2", *ubicaciones[1].ResidenciaFiscal)
	assert.Equal(t, "2", *ubicaciones[1].NoEstacion)
	assert.Equal(t, "Estacion2", *ubicaciones[1].NombreEstacion)
	assert.Equal(t, "Trafico2", *ubicaciones[1].NavegacionTrafico)
	assert.Equal(t, "2024-09-23T12:30:00", ubicaciones[1].FechaHoraSalidaLlegadaString)
	assert.Equal(t, "Tipo2", *ubicaciones[1].TipoEstacion)
	assert.Equal(t, 100.50, *ubicaciones[1].DistanciaRecorrida)

	domicilios := []cartaporte2.DomicilioCartaPorte20{*ubicaciones[0].Domicilio, *ubicaciones[1].Domicilio}
	InternalTestFullAtributesDomicilioCartaPorte20(t, domicilios)
}

func InternalTestFullAtributesDomicilioCartaPorte20(t *testing.T, porte20 []cartaporte2.DomicilioCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 2, len(porte20))
	assert.Equal(t, "Avenida Principal", *porte20[0].Calle)
	assert.Equal(t, "123", *porte20[0].NoExterior)
	assert.Equal(t, "A", *porte20[0].NoInterior)
	assert.Equal(t, "Centro", *porte20[0].Colonia)
	assert.Equal(t, "Ciudad Ejemplo", *porte20[0].Localidad)
	assert.Equal(t, "Cerca de la Plaza Central", *porte20[0].Referencia)
	assert.Equal(t, "Municipio Ejemplo", *porte20[0].Municipio)
	assert.Equal(t, "Estado Ejemplo", porte20[0].Estado)
	assert.Equal(t, "MEX", porte20[0].Pais)
	assert.Equal(t, "12345", porte20[0].CodigoPostal)

	assert.Equal(t, "Avenida Principal", *porte20[1].Calle)
	assert.Equal(t, "123", *porte20[1].NoExterior)
	assert.Equal(t, "A", *porte20[1].NoInterior)
	assert.Equal(t, "Centro", *porte20[1].Colonia)
	assert.Equal(t, "Ciudad Ejemplo", *porte20[1].Localidad)
	assert.Equal(t, "Cerca de la Plaza Central", *porte20[1].Referencia)
	assert.Equal(t, "Municipio Ejemplo", *porte20[1].Municipio)
	assert.Equal(t, "Estado Ejemplo", porte20[1].Estado)
	assert.Equal(t, "MEX", porte20[1].Pais)
	assert.Equal(t, "12345", porte20[1].CodigoPostal)
}

func InternalTestFullAtributesMercancias(t *testing.T, porte20 cartaporte2.MercanciasCartaPorte20) {
	assert.Equal(t, 1500.250, porte20.PesoBrutoTotal)
	assert.Equal(t, "KGM", porte20.UnidadPeso)
	assert.Equal(t, 1400.000, *porte20.PesoNetoTotal)
	assert.Equal(t, float64(10), porte20.NumeroTotalMercancias)
	assert.Equal(t, 100.00, *porte20.CargoPorTasacion)
	InternalFullTestAtributesMercancia20(t, porte20.Mercancia)
	InternalTestFullAtributesAutotransporte20(t, *porte20.Autotransporte)
	InternalTestFullAtributesTransporteMaritimo20(t, *porte20.TransporteMaritimo)
	InternalTestFullAtributesTransporteAereo20(t, *porte20.TransporteAereo)
	InternalTestFullAtributesTransporteFerroviario20(t, *porte20.TransporteFerroviario)
}

func InternalFullTestAtributesMercancia20(t *testing.T, porte20 []cartaporte2.MercanciaCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "12345678", porte20[0].BienesTransportados)
	assert.Equal(t, "123456", *porte20[0].ClaveStcc)
	assert.Equal(t, "Electrodomésticos de cocina", porte20[0].Descripcion)
	assert.Equal(t, 10.000000, porte20[0].Cantidad)
	assert.Equal(t, "E48", porte20[0].ClaveUnidad)
	assert.Equal(t, "Cajas", *porte20[0].Unidad)
	assert.Equal(t, "30/40/30cm", *porte20[0].Dimensiones)
	assert.Equal(t, "No", *porte20[0].MaterialPeligroso)
	assert.Equal(t, "No", *porte20[0].ClaveMaterialPeligroso)
	assert.Equal(t, "Si", *porte20[0].Embalaje)
	assert.Equal(t, "Tipo1", *porte20[0].DescripcionEmbalaje)
	assert.Equal(t, 25.500, porte20[0].PesoKg)
	assert.Equal(t, 1500.00, *porte20[0].ValorMercancia)
	assert.Equal(t, "MXN", *porte20[0].Moneda)
	assert.Equal(t, "1234.56", *porte20[0].FraccionArancelaria)
	assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", *porte20[0].UUIDComercioExt)
	assert.Equal(t, strings.ToUpper("550e8400-e29b-41d4-a716-446655440000"), *porte20[0].UuidComercioExterior)
	InternalTestFullAtributesPedimentos20(t, *porte20[0].Pedimentos)
	InternalTestFullAtributesGuiasIdentificacion20(t, *porte20[0].GuiasIdentificacion)
	InternalTestFullAtributesCantidadTransporta20(t, *porte20[0].CantidadTransporta)
	InternalTestFullAtributesDetalleMerancia20(t, *porte20[0].DetalleMercancia)
}

func InternalTestFullAtributesPedimentos20(t *testing.T, porte20 []cartaporte2.PedimentosCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "230112345678901", porte20[0].Pedimento)
}

func InternalTestFullAtributesGuiasIdentificacion20(t *testing.T, porte20 []cartaporte2.GuiasIdentificacionCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "1234567890", porte20[0].Numero)
	assert.Equal(t, "Paquete de libros de texto", porte20[0].Descripcion)
	assert.Equal(t, 2.500, porte20[0].Peso)
}

func InternalTestFullAtributesCantidadTransporta20(t *testing.T, porte20 []cartaporte2.CantidadTransportaCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, 100.500000, porte20[0].Cantidad)
	assert.Equal(t, "OR123456", porte20[0].IdOrigen)
	assert.Equal(t, "DE654321", porte20[0].IdDestino)
	assert.Equal(t, "TRK", *porte20[0].ClavesTransporte)
}

func InternalTestFullAtributesDetalleMerancia20(t *testing.T, porte20 cartaporte2.DetalleMercanciaCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "MTR", porte20.UnidadPeso)
	assert.Equal(t, 1500.500, porte20.PesoBruto)
	assert.Equal(t, 1400.200, porte20.PesoNeto)
	assert.Equal(t, 100.300, porte20.PesoTara)
	assert.Equal(t, float64(10), *porte20.NumeroPiezas)
}

func InternalTestFullAtributesAutotransporte20(t *testing.T, porte20 cartaporte2.AutotransporteCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "PERM001", porte20.PermisoSct)
	assert.Equal(t, "NUMPERM12345", porte20.NoPermisoSct)
	InternalTestFullAtributesIdentificacionVehicular20(t, porte20.IdentificacionVehicular)
	InternalTestFullAtributesSeguros20(t, porte20.Seguros)
	InternalTestFullAtributesRemolques20(t, *porte20.Remolques)
}

func InternalTestFullAtributesIdentificacionVehicular20(t *testing.T, porte20 cartaporte2.IdentificacionVehicularCartaPorte20) {
	assert.Equal(t, "ClaveConfig123", porte20.ConfiguracionVehicular)
	assert.Equal(t, "ABC1234", porte20.PlacaVm)
	assert.Equal(t, "2020", porte20.AnioModeloVm)
}

func InternalTestFullAtributesSeguros20(t *testing.T, porte20 cartaporte2.SegurosCartaPorte20) {
	assert.Equal(t, "Aseguradora XYZ", porte20.AseguradoraResponsabilidadCivil)
	assert.Equal(t, "POL123456", porte20.PolizaResponsabilidadCivil)
	assert.Equal(t, "Aseguradora ABC", *porte20.AseguradoraMedioAmbiente)
	assert.Equal(t, "POLMED789", *porte20.PolizaMedioAmbiente)
	assert.Equal(t, "Aseguradora DEF", *porte20.AseguradoraCarga)
	assert.Equal(t, "POLCARG456", *porte20.PolizaCarga)
	assert.Equal(t, 1500.00, *porte20.PrimaSeguro)
}

func InternalTestFullAtributesRemolques20(t *testing.T, porte20 []cartaporte2.RemolqueCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "REM01", porte20[0].Subtipo)
	assert.Equal(t, "ABC1234", porte20[0].Placa)
}

func InternalTestFullAtributesTransporteMaritimo20(t *testing.T, porte20 cartaporte2.TransporteMaritimoCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "PERM002", *porte20.PermisoSct)
	assert.Equal(t, "987654", *porte20.NoPermisoSct)
	assert.Equal(t, "Aseguradora ABC", *porte20.NombreAseguradora)
	assert.Equal(t, "POL987654", *porte20.NoPolizaSeguro)
	assert.Equal(t, "TIPO001", porte20.TipoEmbarcacion)
	assert.Equal(t, "MATRICULA123", porte20.Matricula)
	assert.Equal(t, "IMO1234567", porte20.NoOmi)
	assert.Equal(t, "MX", porte20.NacionalidadEmbarcacion)
	assert.Equal(t, 250.500, porte20.UnidadesArqueoBruto)
	assert.Equal(t, "CARGA001", porte20.TipoCarga)
	assert.Equal(t, "CERT123", porte20.NumCertItc)
	assert.Equal(t, "2010", *porte20.AnioEmbarcacion)
	assert.Equal(t, "Embarcación XYZ", *porte20.NombreEmbarcacion)
	assert.Equal(t, 30.00, *porte20.Eslora)
	assert.Equal(t, 8.50, *porte20.Manga)
	assert.Equal(t, 5.00, *porte20.Calado)
	assert.Equal(t, "Naviera 123", *porte20.LineaNaviera)
	assert.Equal(t, "Agente Naviero ABC", porte20.NombreAgenteNaviero)
	assert.Equal(t, "AUTH123", porte20.NoAutorizacionNaviero)
	assert.Equal(t, "TRIP001", *porte20.NoViaje)
	assert.Equal(t, "CONOC123", *porte20.NoConocimientoEmbarque)
	InternalTestFullAtributesContenedor20(t, porte20.Contenedor)
}

func InternalTestFullAtributesContenedor20(t *testing.T, porte20 []cartaporte2.ContenedorCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "ABC1234567D", porte20[0].Matricula)
	assert.Equal(t, "CONT001", porte20[0].Tipo)
	assert.Equal(t, "PRECINT001", *porte20[0].NoPrecinto)
}

func InternalTestFullAtributesTransporteAereo20(t *testing.T, porte20 cartaporte2.TransporteAereoCartaPorte20) {
	assert.Equal(t, "PERM001", porte20.PermSct)
	assert.Equal(t, "123456", porte20.NoPermisoSct)
	assert.Equal(t, "ABC-123", *porte20.MatriculaAeronave)
	assert.Equal(t, "Aseguradora XYZ", *porte20.NombreAseguradora)
	assert.Equal(t, "POL123456", *porte20.NoPolizaSeguro)
	assert.Equal(t, "GUIDE123456789", porte20.NoGuia)
	assert.Equal(t, "CODE001", porte20.CodigoTransportista)
	assert.Equal(t, "Ciudad de México", *porte20.LugarContrato)
	assert.Equal(t, "RFC123456789", *porte20.RfcEmbarcador)
	assert.Equal(t, "REG123456", *porte20.NumRegIdTribEmbarcador)
	assert.Equal(t, "MX", *porte20.ResidenciaFiscalEmbarcador)
	assert.Equal(t, "Embarcador S.A. de C.V.", *porte20.NombreEmbarcador)
}

func InternalTestFullAtributesTransporteFerroviario20(t *testing.T, porte20 cartaporte2.TransporteFerroviarioCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "01", porte20.TipoServicio)
	assert.Equal(t, "A", porte20.TipoTrafico)
	assert.Equal(t, "Aseguradora Ejemplo", *porte20.NombreAseguradora)
	assert.Equal(t, "POL123456", *porte20.NoPolizaSeguro)
	InternalTestFullAtributesDerechosDePaso20(t, *porte20.DerechoPaso)
	InternalTestFullAtributesCarro20(t, porte20.Carro)
}

func InternalTestFullAtributesDerechosDePaso20(t *testing.T, porte20 []cartaporte2.DerechoPasoCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "DERECHO001", porte20[0].Tipo)
	assert.Equal(t, 15.50, porte20[0].KilometrajePagado)
}

func InternalTestFullAtributesCarro20(t *testing.T, porte20 []cartaporte2.CarroCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "01", porte20[0].Tipo)
	assert.Equal(t, "CAR12345", porte20[0].Matricula)
	assert.Equal(t, "GUIA123", porte20[0].Guia)
	assert.Equal(t, 15.500, porte20[0].ToneladasNetas)
	InternalTestFullAtributesContenedorCarro20(t, *porte20[0].Contenedor)
}

func InternalTestFullAtributesContenedorCarro20(t *testing.T, porte20 []cartaporte2.ContenedorCarroCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "CTN001", porte20[0].Tipo)
	assert.Equal(t, 2.500, porte20[0].PesoContenedorVacio)
	assert.Equal(t, 13.000, porte20[0].PesoNetoMercancia)
}

func InternalTestFullAtributesFiguraTransporte20(t *testing.T, porte20 cartaporte2.FiguraTransporteCartaPorte20) {
	InternalTestFullAtributesTiposFigura20(t, porte20.TiposFigura)
}

func InternalTestFullAtributesTiposFigura20(t *testing.T, porte20 []cartaporte2.TiposFiguraCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "Figura1", porte20[0].Tipo)
	assert.Equal(t, "RFCFigura1", *porte20[0].Rfc)
	assert.Equal(t, "12456", *porte20[0].NoLicencia)
	assert.Equal(t, "Nombre1", *porte20[0].Nombre)
	assert.Equal(t, "IdTrib1", *porte20[0].NumRegIdTrib)
	assert.Equal(t, "Residencia1", *porte20[0].ResidenciaFiscal)
	InternalTestFullAtributesPartesTransporte(t, *porte20[0].PartesTransporte)
	InternalTestFullAtributesDomicilioFiguraTransporte20(t, *porte20[0].Domicilio)
}

func InternalTestFullAtributesPartesTransporte(t *testing.T, porte20 []cartaporte2.PartesTransporteCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "Vehículo", porte20[0].Parte)
}

func InternalTestFullAtributesDomicilioFiguraTransporte20(t *testing.T, porte20 cartaporte2.DomicilioCartaPorte20) {
	assert.Equal(t, "Avenida Principal", *porte20.Calle)
	assert.Equal(t, "123", *porte20.NoExterior)
	assert.Equal(t, "A", *porte20.NoInterior)
	assert.Equal(t, "Centro", *porte20.Colonia)
	assert.Equal(t, "Ciudad Ejemplo", *porte20.Localidad)
	assert.Equal(t, "Cerca de la Plaza Central", *porte20.Referencia)
	assert.Equal(t, "Municipio Ejemplo", *porte20.Municipio)
	assert.Equal(t, "Estado Ejemplo", porte20.Estado)
	assert.Equal(t, "MEX", porte20.Pais)
	assert.Equal(t, "12345", porte20.CodigoPostal)
}

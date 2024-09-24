package cartaporte

import (
	"encoding/xml"
	cartaporte2 "github.com/SaulEnriqueMR/kore-models/models/cartaporte"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetCartaPorte20ForTest(filename string, t *testing.T) (cartaporte2.CartaPorte20, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed cartaporte2.CartaPorte20
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	return parsed, errUnmarshal
}

func TestCartaPorte20(t *testing.T) {
	cartaPorte20, _ := GetCartaPorte20ForTest("./cartaporte20.xml", t)
	InternalTestFullAtributesCartaPorte20(t, cartaPorte20)
	InternalTestFullAtributesUbicacionCartaPorte20(t, cartaPorte20.Ubicaciones)
	InternalTestFullAtributesMercancias(t, cartaPorte20.Mercancias)
	InternalTestFullAtributesFiguraTransporte20(t, *cartaPorte20.FiguraTransporte)
}

func InternalTestFullAtributesCartaPorte20(t *testing.T, porte20 cartaporte2.CartaPorte20) {
	assert.Equal(t, "2.0", porte20.Version)
	assert.Equal(t, "Si", porte20.TranspInternac)
	assert.Equal(t, "Salida", *porte20.EntradaSalidaMerc)
	assert.Equal(t, "MEX", *porte20.PaisOrigenDestino)
	assert.Equal(t, "T1", *porte20.ViaEntradaSalida)
	assert.Equal(t, 120.50, *porte20.TotalDistRec)
}

func InternalTestFullAtributesUbicacionCartaPorte20(t *testing.T, ubicaciones []cartaporte2.UbicacionCartaPorte20) {
	assert.Equal(t, 2, len(ubicaciones))
	assert.Equal(t, "Origen", ubicaciones[0].TipoUbicacion)
	assert.Equal(t, "OR123456", *ubicaciones[0].IDUbicacion)
	assert.Equal(t, "XAXX010101000", ubicaciones[0].RFCRemitenteDestinatario)
	assert.Equal(t, "Empresa Origen", *ubicaciones[0].NombreRemitenteDestinatario)
	assert.Equal(t, "IdTrib1", *ubicaciones[0].NumRegIdTrib)
	assert.Equal(t, "ResidenciaFisc1", *ubicaciones[0].ResidenciaFiscal)
	assert.Equal(t, "1", *ubicaciones[0].NumEstacion)
	assert.Equal(t, "Estacion1", *ubicaciones[0].NombreEstacion)
	assert.Equal(t, "Trafico1", *ubicaciones[0].NavegacionTrafico)
	assert.Equal(t, "2024-09-23T12:30:00", ubicaciones[0].FechaHoraSalidaLlegada)
	assert.Equal(t, "Tipo1", *ubicaciones[0].TipoEstacion)
	assert.Equal(t, 100.50, *ubicaciones[0].DistanciaRecorrida)

	assert.Equal(t, "Origen", ubicaciones[1].TipoUbicacion)
	assert.Equal(t, "OR123456", *ubicaciones[1].IDUbicacion)
	assert.Equal(t, "XAXX010101000", ubicaciones[1].RFCRemitenteDestinatario)
	assert.Equal(t, "Empresa Origen", *ubicaciones[1].NombreRemitenteDestinatario)
	assert.Equal(t, "IdTrib2", *ubicaciones[1].NumRegIdTrib)
	assert.Equal(t, "ResidenciaFisc2", *ubicaciones[1].ResidenciaFiscal)
	assert.Equal(t, "2", *ubicaciones[1].NumEstacion)
	assert.Equal(t, "Estacion2", *ubicaciones[1].NombreEstacion)
	assert.Equal(t, "Trafico2", *ubicaciones[1].NavegacionTrafico)
	assert.Equal(t, "2024-09-23T12:30:00", ubicaciones[1].FechaHoraSalidaLlegada)
	assert.Equal(t, "Tipo2", *ubicaciones[1].TipoEstacion)
	assert.Equal(t, 100.50, *ubicaciones[1].DistanciaRecorrida)

	domicilios := []cartaporte2.DomicilioCartaPorte20{*ubicaciones[0].Domicilio, *ubicaciones[1].Domicilio}
	InternalTestFullAtributesDomicilioCartaPorte20(t, domicilios)
}

func InternalTestFullAtributesDomicilioCartaPorte20(t *testing.T, porte20 []cartaporte2.DomicilioCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 2, len(porte20))
	assert.Equal(t, "Avenida Principal", *porte20[0].Calle)
	assert.Equal(t, "123", *porte20[0].NumeroExterior)
	assert.Equal(t, "A", *porte20[0].NumeroInterior)
	assert.Equal(t, "Centro", *porte20[0].Colonia)
	assert.Equal(t, "Ciudad Ejemplo", *porte20[0].Localidad)
	assert.Equal(t, "Cerca de la Plaza Central", *porte20[0].Referencia)
	assert.Equal(t, "Municipio Ejemplo", *porte20[0].Municipio)
	assert.Equal(t, "Estado Ejemplo", porte20[0].Estado)
	assert.Equal(t, "MEX", porte20[0].Pais)
	assert.Equal(t, "12345", porte20[0].CodigoPostal)

	assert.Equal(t, "Avenida Principal", *porte20[1].Calle)
	assert.Equal(t, "123", *porte20[1].NumeroExterior)
	assert.Equal(t, "A", *porte20[1].NumeroInterior)
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
	assert.Equal(t, float64(10), porte20.NumTotalMercancias)
	assert.Equal(t, 100.00, *porte20.CargoPorTasacion)
	InternalFullTestAtributesMercancia20(t, porte20.Mercancia)
	InternalTestFullAtributesAutotransporte20(t, *porte20.Autotransporte)
	InternalTestFullAtributesTransporteMaritimo20(t, *porte20.TransporteMaritimo)
	InternalTestFullAtributesTransporteAereo20(t, *porte20.TransporteAereo)
	InternalTestFullAtributesTransporteFerroviario20(t, *porte20.TransporteFerroviario)
}

func InternalFullTestAtributesMercancia20(t *testing.T, porte20 []cartaporte2.MercanciaCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "12345678", porte20[0].BienesTransp)
	assert.Equal(t, "123456", *porte20[0].ClaveSTCC)
	assert.Equal(t, "Electrodomésticos de cocina", porte20[0].Descripcion)
	assert.Equal(t, 10.000000, porte20[0].Cantidad)
	assert.Equal(t, "E48", porte20[0].ClaveUnidad)
	assert.Equal(t, "Cajas", *porte20[0].Unidad)
	assert.Equal(t, "30/40/30cm", *porte20[0].Dimensiones)
	assert.Equal(t, "No", *porte20[0].MaterialPeligroso)
	assert.Equal(t, "No", *porte20[0].CveMaterialPeligroso)
	assert.Equal(t, "Si", *porte20[0].Embalaje)
	assert.Equal(t, "Tipo1", *porte20[0].DescripEmbalaje)
	assert.Equal(t, 25.500, porte20[0].PesoKg)
	assert.Equal(t, 1500.00, *porte20[0].ValorMercancia)
	assert.Equal(t, "MXN", *porte20[0].Moneda)
	assert.Equal(t, "1234.56", *porte20[0].FraccionArancelaria)
	assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", *porte20[0].UUIDComercioExt)
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
	assert.Equal(t, "1234567890", porte20[0].NumeroGuiaIdentificacion)
	assert.Equal(t, "Paquete de libros de texto", porte20[0].DescripGuiaIdentificacion)
	assert.Equal(t, 2.500, porte20[0].PesoGuiaIdentificacion)
}

func InternalTestFullAtributesCantidadTransporta20(t *testing.T, porte20 []cartaporte2.CantidadTransportaCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, 100.500000, porte20[0].Cantidad)
	assert.Equal(t, "OR123456", porte20[0].IDOrigen)
	assert.Equal(t, "DE654321", porte20[0].IDDestino)
	assert.Equal(t, "TRK", *porte20[0].CvesTransporte)
}

func InternalTestFullAtributesDetalleMerancia20(t *testing.T, porte20 cartaporte2.DetalleMercanciaCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "MTR", porte20.UnidadPesoMerc)
	assert.Equal(t, 1500.500, porte20.PesoBruto)
	assert.Equal(t, 1400.200, porte20.PesoNeto)
	assert.Equal(t, 100.300, porte20.PesoTara)
	assert.Equal(t, float64(10), *porte20.NumPiezas)
}

func InternalTestFullAtributesAutotransporte20(t *testing.T, porte20 cartaporte2.AutotransporteCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "PERM001", porte20.PermSCT)
	assert.Equal(t, "NUMPERM12345", porte20.NumPermisoSCT)
	InternalTestFullAtributesIdentificacionVehicular20(t, porte20.IdentificacionVehicular)
	InternalTestFullAtributesSeguros20(t, porte20.Seguros)
	InternalTestFullAtributesRemolques20(t, *porte20.Remolques)
}

func InternalTestFullAtributesIdentificacionVehicular20(t *testing.T, porte20 cartaporte2.IdentificacionVehicularCartaPorte20) {
	assert.Equal(t, "ClaveConfig123", porte20.ConfigVehicular)
	assert.Equal(t, "ABC1234", porte20.PlacaVM)
	assert.Equal(t, "2020", porte20.AnioModeloVM)
}

func InternalTestFullAtributesSeguros20(t *testing.T, porte20 cartaporte2.SegurosCartaPorte20) {
	assert.Equal(t, "Aseguradora XYZ", porte20.AseguraRespCivil)
	assert.Equal(t, "POL123456", porte20.PolizaRespCivil)
	assert.Equal(t, "Aseguradora ABC", *porte20.AseguraMedAmbiente)
	assert.Equal(t, "POLMED789", *porte20.PolizaMedAmbiente)
	assert.Equal(t, "Aseguradora DEF", *porte20.AseguraCarga)
	assert.Equal(t, "POLCARG456", *porte20.PolizaCarga)
	assert.Equal(t, 1500.00, *porte20.PrimaSeguro)
}

func InternalTestFullAtributesRemolques20(t *testing.T, porte20 []cartaporte2.RemolqueCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "REM01", porte20[0].SubTipoRem)
	assert.Equal(t, "ABC1234", porte20[0].Placa)
}

func InternalTestFullAtributesTransporteMaritimo20(t *testing.T, porte20 cartaporte2.TransporteMaritimoCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "PERM002", *porte20.PermSCT)
	assert.Equal(t, "987654", *porte20.NumPermisoSCT)
	assert.Equal(t, "Aseguradora ABC", *porte20.NombreAseg)
	assert.Equal(t, "POL987654", *porte20.NumPolizaSeguro)
	assert.Equal(t, "TIPO001", porte20.TipoEmbarcacion)
	assert.Equal(t, "MATRICULA123", porte20.Matricula)
	assert.Equal(t, "IMO1234567", porte20.NumeroOMI)
	assert.Equal(t, "MX", porte20.NacionalidadEmbarc)
	assert.Equal(t, 250.500, porte20.UnidadesArqBruto)
	assert.Equal(t, "CARGA001", porte20.TipoCarga)
	assert.Equal(t, "CERT123", porte20.NumCertITC)
	assert.Equal(t, "2010", *porte20.AnioEmbarcacion)
	assert.Equal(t, "Embarcación XYZ", *porte20.NombreEmbarc)
	assert.Equal(t, 30.00, *porte20.Eslora)
	assert.Equal(t, 8.50, *porte20.Manga)
	assert.Equal(t, 5.00, *porte20.Calado)
	assert.Equal(t, "Naviera 123", *porte20.LineaNaviera)
	assert.Equal(t, "Agente Naviero ABC", porte20.NombreAgenteNaviero)
	assert.Equal(t, "AUTH123", porte20.NumAutorizacionNaviero)
	assert.Equal(t, "TRIP001", *porte20.NumViaje)
	assert.Equal(t, "CONOC123", *porte20.NumConocEmbarc)
	InternalTestFullAtributesContenedor20(t, porte20.Contenedor)
}

func InternalTestFullAtributesContenedor20(t *testing.T, porte20 []cartaporte2.ContenedorCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "ABC1234567D", porte20[0].MatriculaContenedor)
	assert.Equal(t, "CONT001", porte20[0].TipoContenedor)
	assert.Equal(t, "PRECINT001", *porte20[0].NumPrecinto)
}

func InternalTestFullAtributesTransporteAereo20(t *testing.T, porte20 cartaporte2.TransporteAereoCartaPorte20) {
	assert.Equal(t, "PERM001", porte20.PermSCT)
	assert.Equal(t, "123456", porte20.NumPermisoSCT)
	assert.Equal(t, "ABC-123", *porte20.MatriculaAeronave)
	assert.Equal(t, "Aseguradora XYZ", *porte20.NombreAseg)
	assert.Equal(t, "POL123456", *porte20.NumPolizaSeguro)
	assert.Equal(t, "GUIDE123456789", porte20.NumeroGuia)
	assert.Equal(t, "CODE001", porte20.CodigoTransportista)
	assert.Equal(t, "Ciudad de México", *porte20.LugarContrato)
	assert.Equal(t, "RFC123456789", *porte20.RFCEmbarcador)
	assert.Equal(t, "REG123456", *porte20.NumRegIdTribEmbarc)
	assert.Equal(t, "MX", *porte20.ResidenciaFiscalEmbarc)
	assert.Equal(t, "Embarcador S.A. de C.V.", *porte20.NombreEmbarcador)
}

func InternalTestFullAtributesTransporteFerroviario20(t *testing.T, porte20 cartaporte2.TransporteFerroviarioCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, "01", porte20.TipoServicio)
	assert.Equal(t, "A", porte20.TipoTrafico)
	assert.Equal(t, "Aseguradora Ejemplo", *porte20.NombreAseg)
	assert.Equal(t, "POL123456", *porte20.NumPolizaSeguro)
	InternalTestFullAtributesDerechosDePaso20(t, *porte20.DerechoPaso)
	InternalTestFullAtributesCarro20(t, porte20.Carro)
}

func InternalTestFullAtributesDerechosDePaso20(t *testing.T, porte20 []cartaporte2.DerechoPasoCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "DERECHO001", porte20[0].TipoDerechoPaso)
	assert.Equal(t, 15.50, porte20[0].KilometrajePagado)
}

func InternalTestFullAtributesCarro20(t *testing.T, porte20 []cartaporte2.CarroCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "01", porte20[0].TipoCarro)
	assert.Equal(t, "CAR12345", porte20[0].MatriculaCarro)
	assert.Equal(t, "GUIA123", porte20[0].GuiaCarro)
	assert.Equal(t, 15.500, porte20[0].ToneladasNetasCarro)
	InternalTestFullAtributesContenedorCarro20(t, *porte20[0].Contenedor)
}

func InternalTestFullAtributesContenedorCarro20(t *testing.T, porte20 []cartaporte2.ContenedorCarroCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "CTN001", porte20[0].TipoContenedor)
	assert.Equal(t, 2.500, porte20[0].PesoContenedorVacio)
	assert.Equal(t, 13.000, porte20[0].PesoNetoMercancia)
}

func InternalTestFullAtributesFiguraTransporte20(t *testing.T, porte20 cartaporte2.FiguraTransporteCartaPorte20) {
	InternalTestFullAtributesTiposFigura20(t, porte20.TiposFigura)
}

func InternalTestFullAtributesTiposFigura20(t *testing.T, porte20 []cartaporte2.TiposFiguraCartaPorte20) {
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "Figura1", porte20[0].TipoFigura)
	assert.Equal(t, "RFCFigura1", *porte20[0].RFCFigura)
	assert.Equal(t, "12456", *porte20[0].NumLicencia)
	assert.Equal(t, "Nombre1", *porte20[0].NombreFigura)
	assert.Equal(t, "IdTrib1", *porte20[0].NumRegIdTribFigura)
	assert.Equal(t, "Residencia1", *porte20[0].ResidenciaFiscalFigura)
	InternalTestFullAtributesPartesTransporte(t, *porte20[0].PartesTransporte)
	InternalTestFullAtributesDomicilioFiguraTransporte20(t, *porte20[0].Domicilio)
}

func InternalTestFullAtributesPartesTransporte(t *testing.T, porte20 []cartaporte2.PartesTransporteCartaPorte20) {
	assert.NotNil(t, porte20)
	assert.Equal(t, 1, len(porte20))
	assert.Equal(t, "Vehículo", porte20[0].ParteTransporte)
}

func InternalTestFullAtributesDomicilioFiguraTransporte20(t *testing.T, porte20 cartaporte2.DomicilioCartaPorte20) {
	assert.Equal(t, "Avenida Principal", *porte20.Calle)
	assert.Equal(t, "123", *porte20.NumeroExterior)
	assert.Equal(t, "A", *porte20.NumeroInterior)
	assert.Equal(t, "Centro", *porte20.Colonia)
	assert.Equal(t, "Ciudad Ejemplo", *porte20.Localidad)
	assert.Equal(t, "Cerca de la Plaza Central", *porte20.Referencia)
	assert.Equal(t, "Municipio Ejemplo", *porte20.Municipio)
	assert.Equal(t, "Estado Ejemplo", porte20.Estado)
	assert.Equal(t, "MEX", porte20.Pais)
	assert.Equal(t, "12345", porte20.CodigoPostal)
}

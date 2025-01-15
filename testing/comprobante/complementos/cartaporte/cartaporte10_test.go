package cartaporte

import (
	"encoding/xml"
	cartaporte2 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cartaporte"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetCartaPorte10ForTest(filename string, t *testing.T) (cartaporte2.CartaPorte10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed cartaporte2.CartaPorte10
	errUnmarshal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmarshal)
	testing2.GenerateJSONFromStructure("cartaporte10.json", parsed)
	return parsed, errUnmarshal
}

func TestFullCartaPorte10(t *testing.T) {
	cp10, _ := GetCartaPorte10ForTest("./cartaporte10.xml", t)
	InternalTestFullAttributesCartaPorte10(t, cp10)
	InternalTestFullUbicaciones(t, cp10.Ubicaciones)
	InternalTestFullMercancias(t, cp10.Mercancias)
	InternalTestFullFiguraTransporte(t, cp10)
}

func InternalTestFullAttributesCartaPorte10(t *testing.T, cp cartaporte2.CartaPorte10) {
	assert.Equal(t, "1.0", cp.Version)
	assert.Equal(t, "No", cp.TransporteInternacional)
	assert.Equal(t, "Entrada", *cp.EntradaSalidaMercancia)
	assert.Equal(t, "01", *cp.ViaEntradaSalida)
	assert.Equal(t, 2500.00, *cp.TotalDistanciaRecorrida)
}

func InternalTestFullUbicaciones(t *testing.T, cp []cartaporte2.UbicacionCartaPorte10) {
	assert.Equal(t, 2, len(cp))
	InternalTestFullOrigen(t, cp[0])
	InternalTestFullDestino(t, cp[1])
}

func InternalTestFullOrigen(t *testing.T, cp cartaporte2.UbicacionCartaPorte10) {
	og := cp.Origen
	assert.Equal(t, "OR00001", *og.IdUbicacion)
	assert.Equal(t, "MME861201NB3", *og.RfcRemitenteDestinatario)
	assert.Equal(t, "MANE MEXICO S.A. DE C.V.", *og.NombreRemitenteDestinatario)
	assert.Equal(t, "1232131211", *og.NumRegIdTrib)
	assert.Equal(t, "MEX", *og.ResidenciaFiscal)
	assert.Equal(t, "91", *og.NoEstacion)
	assert.Equal(t, "ESTACION 1", *og.NombreEstacion)
	assert.Equal(t, "Altura", *og.NavegacionTrafico)
	assert.Equal(t, "2021-11-24T16:16:39", og.FechaHoraSalidaLlegadaString)
	InternalTestFullDomicilioOrigen(t, *cp.Domicilio)
}

func InternalTestFullDomicilioOrigen(t *testing.T, cp cartaporte2.DomicilioCartaPorte10) {
	assert.Equal(t, "NEMESIO DIEZ RIEGA 9", cp.Calle)
	assert.Equal(t, "12", *cp.NoExterior)
	assert.Equal(t, "13", *cp.NoInterior)
	assert.Equal(t, "BENITO", *cp.Colonia)
	assert.Equal(t, "12", *cp.Localidad)
	assert.Equal(t, "Entre", *cp.Referencia)
	assert.Equal(t, "051", *cp.Municipio)
	assert.Equal(t, "MEX", cp.Estado)
	assert.Equal(t, "MEX", cp.Pais)
	assert.Equal(t, "52004", cp.CodigoPostal)
}

func InternalTestFullDestino(t *testing.T, cp cartaporte2.UbicacionCartaPorte10) {
	assert.Equal(t, 2500.00, *cp.DistanciaRecorrida)
	assert.Equal(t, "01", *cp.TipoEstacion)
	de := cp.Destino
	assert.Equal(t, "DE000005", *de.IdUbicacion)
	assert.Equal(t, "ERFDSFGR312", *de.RfcRemitenteDestinatario)
	assert.Equal(t, "DROGUERIA SINERGIA INTERNATIONAL, SOCIEDAD ANONIMA", *de.NombreRemitenteDestinatario)
	assert.Equal(t, "12770698", *de.NumRegIdTrib)
	assert.Equal(t, "GTM", *de.ResidenciaFiscal)
	assert.Equal(t, "02", *de.NoEstacion)
	assert.Equal(t, "ESTACION", *de.NombreEstacion)
	assert.Equal(t, "Cabotaje", *de.NavegacionTrafico)
	assert.Equal(t, "2021-11-24T16:16:40", de.FechaHoraSalidaLlegadaString)
	InternalTestFullDomicilioDestino(t, *cp.Domicilio)
}

func InternalTestFullDomicilioDestino(t *testing.T, cp cartaporte2.DomicilioCartaPorte10) {
	assert.Equal(t, "9A. CALLE 2-71", cp.Calle)
	assert.Equal(t, "12", *cp.NoExterior)
	assert.Equal(t, "13", *cp.NoInterior)
	assert.Equal(t, "ZONA 1", *cp.Colonia)
	assert.Equal(t, "No Aplica", *cp.Localidad)
	assert.Equal(t, "Entre", *cp.Referencia)
	assert.Equal(t, "No Aplica", *cp.Municipio)
	assert.Equal(t, "No Aplica", cp.Estado)
	assert.Equal(t, "GTM", cp.Pais)
	assert.Equal(t, "1000", cp.CodigoPostal)
}

func InternalTestFullMercancias(t *testing.T, cp cartaporte2.MercanciasCartaPorte10) {
	assert.Equal(t, 10.00, *cp.PesoBrutoTotal)
	assert.Equal(t, "E48", *cp.UnidadPeso)
	assert.Equal(t, 5660.00, *cp.PesoNetoTotal)
	assert.Equal(t, 1.00, cp.NumeroTotalMercancias)
	assert.Equal(t, 100.00, *cp.CargoPorTasacion)

	InternalTestFullMercancia(t, cp.Mercancia[0])
	InternalTestFullAutotransporte10(t, cp)
	InternalTestFullTransporteMaritimo(t, cp)
	InternalTestFullTransporteAereo(t, cp)
	InternalTestFullTransporteFerroviario(t, cp)
}

func InternalTestFullMercancia(t *testing.T, cp cartaporte2.MercanciaCartaPorte10) {
	assert.Equal(t, "12164503", *cp.BienesTransportados)
	assert.Equal(t, "123456", *cp.ClaveStcc)
	assert.Equal(t, "ADITIVOS DE FRAGANCIA", *cp.Descripcion)
	assert.Equal(t, "13", *cp.Cantidad)
	assert.Equal(t, "H87", *cp.ClaveUnidad)
	assert.Equal(t, "PIEZA", *cp.Unidad)
	assert.Equal(t, "CUBOS", *cp.Dimensiones)
	assert.Equal(t, "Sí", *cp.MaterialPeligroso)
	assert.Equal(t, "M0508", *cp.ClaveMaterialPeligroso)
	assert.Equal(t, "3H1", *cp.Embalaje)
	assert.Equal(t, "JERRICANES (PORRONES) DE PLÁSTICO DE TAPA NO DESMONTABLE", *cp.DescripcionEmbalaje)
	assert.Equal(t, 5660.00, cp.PesoKg)
	assert.Equal(t, 100.00, *cp.ValorMercancia)
	assert.Equal(t, "MXN", *cp.Moneda)
	assert.Equal(t, "01", *cp.FraccionArancelaria)
	assert.Equal(t, "94072f17-4f13-40b4-bd07-e2e027650898", *cp.UUIDComercioExt)

	InternalTestFullCantidadTransporta(t, cp)
	InternalTestFullDetalleMercancia(t, cp)
}

func InternalTestFullCantidadTransporta(t *testing.T, cp cartaporte2.MercanciaCartaPorte10) {
	assert.NotNil(t, cp.CantidadTransporta)
	cantidad := *cp.CantidadTransporta
	assert.Equal(t, 2, len(cantidad))
	assert.Equal(t, 10.00, cantidad[0].Cantidad)
	assert.Equal(t, "OR00001", cantidad[0].IdOrigen)
	assert.Equal(t, "DE00002", cantidad[0].IdDestino)
	assert.Equal(t, "02", *cantidad[0].ClavesTransporte)
}

func InternalTestFullDetalleMercancia(t *testing.T, cp cartaporte2.MercanciaCartaPorte10) {
	assert.NotNil(t, cp.DetalleMercancia)
	detalle := *cp.DetalleMercancia

	assert.Equal(t, "KG", detalle.UnidadPeso)
	assert.Equal(t, 100.00, detalle.PesoBruto)
	assert.Equal(t, 101.00, detalle.PesoNeto)
	assert.Equal(t, 102.00, detalle.PesoTara)
	assert.Equal(t, 1.00, *detalle.NumeroPiezas)
}

func InternalTestFullAutotransporte10(t *testing.T, cp cartaporte2.MercanciasCartaPorte10) {
	assert.NotNil(t, cp.Autotransporte)
	at := *cp.Autotransporte

	assert.Equal(t, "TPAF01", at.PermisoSct)
	assert.Equal(t, "1329TERS24112011230301006", at.NoPermisoSct)
	assert.Equal(t, "QUALITAS", at.NombreAseguradora)
	assert.Equal(t, "2440038161", at.NoPoliza)

	InternalTestFullIdentificacionVehicular(t, at)
	InternalTestFullRemolquesAutotransporte(t, at)
}

func InternalTestFullIdentificacionVehicular(t *testing.T, cp cartaporte2.AutotransporteCartaPorte10) {
	assert.NotNil(t, cp.IdentificacionVehicular)

	iv := &cp.IdentificacionVehicular
	assert.Equal(t, "2009", iv.AnioModeloVm)
	assert.Equal(t, "11AL4L", iv.PlacaVm)
	assert.Equal(t, "C3R2", iv.ConfiguracionVehicular)
}

func InternalTestFullRemolquesAutotransporte(t *testing.T, cp cartaporte2.AutotransporteCartaPorte10) {
	assert.NotNil(t, cp.Remolques)

	r := cp.Remolques
	assert.Equal(t, 2, len(*r))

	r1 := &(*r)[0]
	assert.Equal(t, "02UJ6C", r1.Placa)
	assert.Equal(t, "CTR004", r1.Subtipo)
}

func InternalTestFullTransporteMaritimo(t *testing.T, cp cartaporte2.MercanciasCartaPorte10) {
	assert.NotNil(t, cp.TransporteMaritimo)
	tm := *cp.TransporteMaritimo

	assert.Equal(t, "PermSCT", *tm.PermisoSct)
	assert.Equal(t, "NumPermisoSCT", *tm.NoPermisoSct)
	assert.Equal(t, "NombreAseg", *tm.NombreAseguradora)
	assert.Equal(t, "NumPolizaSeguro", *tm.NoPolizaSeguro)
	assert.Equal(t, "TipoEmbarcacion", tm.TipoEmbarcacion)
	assert.Equal(t, "Matricula", tm.Matricula)
	assert.Equal(t, "NumeroOMI", tm.NoOmi)
	assert.Equal(t, "AnioEmbarcacion", *tm.AnioEmbarcacion)
	assert.Equal(t, "NombreEmbarc", *tm.NombreEmbarcacion)
	assert.Equal(t, "NacionalidadEmbarc", tm.NacionalidadEmbarcacion)
	assert.Equal(t, 3.0, tm.UnidadesArqueoBruto)
	assert.Equal(t, "TipoCarga", tm.TipoCarga)
	assert.Equal(t, "NumCertITC", tm.NumCertITC)
	assert.Equal(t, 2.0, *tm.Eslora)
	assert.Equal(t, 3.0, *tm.Manga)
	assert.Equal(t, 4.0, *tm.Calado)
	assert.Equal(t, "LineaNaviera", *tm.LineaNaviera)
	assert.Equal(t, "NombreAgenteNaviero", tm.NombreAgenteNaviero)
	assert.Equal(t, "NumAutorizacionNaviero", tm.NoAutorizacionNaviero)
	assert.Equal(t, "NumViaje", *tm.NoViaje)
	assert.Equal(t, "NumConocEmbarc", *tm.NoConocimientoEmbarque)
	InternalTestFullContenedorMaritimo(t, tm)
}

func InternalTestFullContenedorMaritimo(t *testing.T, cp cartaporte2.TransporteMaritimoCartaPorte10) {
	assert.Equal(t, 2, len(cp.Contenedores))

	c := cp.Contenedores[0]
	assert.Equal(t, "MatriculaContenedor", c.Matricula)
	assert.Equal(t, "TipoContenedor", c.Tipo)
	assert.Equal(t, "NumPrecinto", *c.NoPrecinto)
}

func InternalTestFullTransporteAereo(t *testing.T, cp cartaporte2.MercanciasCartaPorte10) {
	assert.NotNil(t, cp.TransporteAereo)

	ta := cp.TransporteAereo
	assert.Equal(t, "PermSCT", ta.PermisoSct)
	assert.Equal(t, "NumPermisoSCT", ta.NoPermisoSct)
	assert.Equal(t, "MatriculaAeronave", *ta.MatriculaAeronave)
	assert.Equal(t, "NombreAseg", *ta.NombreAseguradora)
	assert.Equal(t, "NumPolizaSeguro", *ta.NoPolizaSeguro)
	assert.Equal(t, "NumeroGuia", ta.NoGuia)
	assert.Equal(t, "LugarContrato", *ta.LugarContrato)
	assert.Equal(t, "RFCTransportista", *ta.RfcTransportista)
	assert.Equal(t, "CodigoTransportista", ta.CodigoTransportista)
	assert.Equal(t, "NumRegIdTribTranspor", *ta.NumRegIdTribTransportista)
	assert.Equal(t, "ResidenciaFiscalTranspor", *ta.ResidenciaFiscalTransportista)
	assert.Equal(t, "NombreTransportista", *ta.NombreTransportista)
	assert.Equal(t, "RFCEmbarcador", *ta.RfcEmbarcador)
	assert.Equal(t, "NumRegIdTribEmbarc", *ta.NumRegIdTribEmbarcador)
	assert.Equal(t, "ResidenciaFiscalEmbarc", *ta.ResidenciaFiscalEmbarcador)
	assert.Equal(t, "NombreEmbarcador", *ta.NombreEmbarcador)
}

func InternalTestFullTransporteFerroviario(t *testing.T, cp cartaporte2.MercanciasCartaPorte10) {
	assert.NotNil(t, cp.TransporteFerroviario)

	tf := cp.TransporteFerroviario
	assert.Equal(t, "TipoDeServicio", tf.TipoServicio)
	assert.Equal(t, "NombreAseg", *tf.NombreAseguradora)
	assert.Equal(t, "NumPolizaSeguro", *tf.NoPolizaSeguro)
	assert.Equal(t, "Concesionario", *tf.Concesionario)

	InternalTestFullDerechoPaso(t, *tf)
	InternalTestFullCarro(t, *tf)
}

func InternalTestFullDerechoPaso(t *testing.T, cp cartaporte2.TransporteFerroviarioCartaPorte10) {
	assert.NotNil(t, cp.DerechoPaso)
	assert.Equal(t, 2, len(*cp.DerechoPaso))

	dps := *cp.DerechoPaso

	assert.Equal(t, "TipoDerechoDePaso", dps[0].Tipo)
	assert.Equal(t, 10.0, dps[0].KilometrajePagado)
}

func InternalTestFullCarro(t *testing.T, cp cartaporte2.TransporteFerroviarioCartaPorte10) {
	assert.Equal(t, 2, len(cp.Carro))
	c := cp.Carro[0]

	assert.Equal(t, "TipoCarro", c.Tipo)
	assert.Equal(t, "MatriculaCarro", c.Matricula)
	assert.Equal(t, "GuiaCarro", c.Guia)
	assert.Equal(t, 10.0, c.ToneladasNetas)
	InternalTestFullContenedorFerroviario(t, c)
}

func InternalTestFullContenedorFerroviario(t *testing.T, cp cartaporte2.CarroCartaPorte10) {
	assert.NotNil(t, cp.Contenedor)
	assert.Equal(t, 2, len(*cp.Contenedor))

	c := *cp.Contenedor

	assert.Equal(t, "TipoContenedor", c[0].Tipo)
	assert.Equal(t, 10.0, c[0].PesoContenedorVacio)
	assert.Equal(t, 10.0, c[0].PesoNetoMercancia)
}

func InternalTestFullFiguraTransporte(t *testing.T, cp cartaporte2.CartaPorte10) {
	assert.NotNil(t, cp.FiguraTransporte)

	ft := *cp.FiguraTransporte
	assert.Equal(t, "01", ft.ClaveTransporte)

	InternalTestFullOperadores(t, ft)
	InternalTestFullPropietario(t, ft)
	InternalTestFullArrendatario(t, ft)
	InternalTestFullNotificado(t, ft)
}

func InternalTestFullOperadores(t *testing.T, cp cartaporte2.FiguraTransporteCartaPorte10) {
	assert.NotNil(t, cp.Operadores)
	assert.Equal(t, 2, len(*cp.Operadores))

	operadores := *cp.Operadores
	op1 := operadores[0]
	assert.Equal(t, "RFCOperador", *op1.Rfc)
	assert.Equal(t, "NumLicencia", *op1.NoLicencia)
	assert.Equal(t, "NombreOperador", *op1.Nombre)
	assert.Equal(t, "NumRegIdTribOperador", *op1.NumRegIdTrib)
	assert.Equal(t, "ResidenciaFiscalOperador", *op1.ResidenciaFiscal)

	assert.NotNil(t, op1.Domicilio)
	InternalTestDomicilioTipoFigura(t, *op1.Domicilio)
}

func InternalTestDomicilioTipoFigura(t *testing.T, cp cartaporte2.DomicilioCartaPorte10) {
	assert.Equal(t, "CONOCIDO", cp.Calle)
	assert.Equal(t, "10", *cp.NoExterior)
	assert.Equal(t, "10", *cp.NoInterior)
	assert.Equal(t, "1510", *cp.Colonia)
	assert.Equal(t, "01", *cp.Localidad)
	assert.Equal(t, "Referencia", *cp.Referencia)
	assert.Equal(t, "012", *cp.Municipio)
	assert.Equal(t, "HID", cp.Estado)
	assert.Equal(t, "MEX", cp.Pais)
	assert.Equal(t, "43300", cp.CodigoPostal)
}

func InternalTestFullPropietario(t *testing.T, cp cartaporte2.FiguraTransporteCartaPorte10) {
	assert.NotNil(t, cp.Propietarios)
	assert.Equal(t, 2, len(*cp.Propietarios))

	p := *cp.Propietarios
	p1 := p[0]

	assert.Equal(t, "RFCPropietario", *p1.Rfc)
	assert.Equal(t, "NombrePropietario", *p1.Nombre)
	assert.Equal(t, "NumRegIdTribPropietario", *p1.NumRegIdTrib)
	assert.Equal(t, "ResidenciaFiscalPropietario", *p1.ResidenciaFiscal)
	InternalTestDomicilioTipoFigura(t, *p1.Domicilio)
}

func InternalTestFullArrendatario(t *testing.T, cp cartaporte2.FiguraTransporteCartaPorte10) {
	assert.NotNil(t, cp.Arrendatarios)
	assert.Equal(t, 2, len(*cp.Arrendatarios))

	a := *cp.Arrendatarios
	a1 := a[0]

	assert.Equal(t, "RFCArrendatario", *a1.Rfc)
	assert.Equal(t, "NombreArrendatario", *a1.Nombre)
	assert.Equal(t, "NumRegIdTribArrendatario", *a1.NumRegIdTrib)
	assert.Equal(t, "ResidenciaFiscalArrendatario", *a1.ResidenciaFiscal)
	InternalTestDomicilioTipoFigura(t, *a1.Domicilio)
}

func InternalTestFullNotificado(t *testing.T, cp cartaporte2.FiguraTransporteCartaPorte10) {
	assert.NotNil(t, cp.Notificados)
	assert.Equal(t, 2, len(*cp.Notificados))

	n := *cp.Notificados
	n1 := n[0]

	assert.Equal(t, "RFCNotificado", *n1.Rfc)
	assert.Equal(t, "NombreNotificado", *n1.Nombre)
	assert.Equal(t, "NumRegIdTribNotificado", *n1.NumRegIdTrib)
	assert.Equal(t, "ResidenciaFiscalNotificado", *n1.ResidenciaFiscal)
	InternalTestDomicilioTipoFigura(t, *n1.Domicilio)
}

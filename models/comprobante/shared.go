package comprobante

import (
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/aerolineas"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cartaporte"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/certificadodestruccion"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cfdiregistrofiscal"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/comercioexterior"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/complementospei"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/consumocombustible"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/detallista"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/divisas"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/donatorias"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/estadodecuentacombustible"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/gastohidrocarburos"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/impuestoslocales"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/ine"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/leyendasfiscales"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/nomina"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/notariospublicos"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/obrasantiguedades"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pagoespecie"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pagos"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/parcialesconstruccion"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/pfintegrantecoordinado"
	renovacionsustitucion "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/renovacionsustitucionvehi"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/turistapasajeroextranjero"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/valesdespensa"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/vehiculousado"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/acuentaterceros"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/institeducativas"
	"github.com/SaulEnriqueMR/kore-models/models/timbrefiscaldigital"
)

type ComplementoConcepto struct {
	ACuentaTerceros         *[]acuentaterceros.ACuentaTerceros         `xml:"PorCuentadeTerceros" bson:"ACuentaTerceros,omitempty"`
	InstitucionesEducativas *[]institeducativas.InstitucioneEducativas `xml:"instEducativas" bson:"InstitucionesEducativas,omitempty"`
}

type Complemento struct {
	CartaPorte                     *cartaporte.CartaPorte                               `xml:"CartaPorte" bson:"CartaPorte,omitempty"`
	ComercioExterior               *comercioexterior.ComercioExterior                   `xml:"ComercioExterior" bson:"ComercioExterior,omitempty"`
	ImpuestoLocales                *impuestoslocales.ImpuestosLocales                   `xml:"ImpuestosLocales" bson:"ImpuestosLocales,omitempty"`
	Aerolineas                     *aerolineas.Aerolineas                               `xml:"Aerolineas" bson:"Aerolineas,omitempty"`
	CertificadoDestruccion         *certificadodestruccion.CertificadoDeDestruccion     `xml:"certificadodedestruccion" bson:"CertificadoDeDestruccion,omitempty"`
	CFDIRegistroFiscal             *cfdiregistrofiscal.CfdiRegistroFiscal               `xml:"CFDIRegistroFiscal" bson:"CFDIRegistroFiscal,omitempty"`
	ComplementoSPEI                *complementos.ComplementoSpei                        `xml:"Complemento_SPEI" bson:"ComplementoSPEI,omitempty"`
	ConsumoCombustibles            *consumocombustible.ConsumoDeCombustible             `xml:"ConsumoDeCombustibles" bson:"ConsumoDeCombustibles,omitempty"`
	Detallista                     *detallista.Detallista                               `xml:"detallista" bson:"detallista,omitempty"`
	Divisas                        *divisas.Divisas                                     `xml:"Divisas" bson:"Divisas,omitempty"`
	Donatarias                     *donatorias.Donatarias                               `xml:"Donatarias" bson:"Donatarias,omitempty"`
	EstadoCuentaCombustible        *estadodecuentacombustible.EstadoDeCuentaCombustible `xml:"EstadoDeCuentaCombustible" bson:"EstadoDeCuentaCombustible,omitempty"`
	GastoHidrocarburos             *gastohidrocarburos.GastoHidrocarburos               `xml:"GastosHidrocarburos" bson:"GastosHidrocarburos,omitempty"`
	Ine                            *ine.INE                                             `xml:"INE" bson:"Ine,omitempty"`
	LeyendasFiscales               *leyendasfiscales.LeyendasFiscales                   `xml:"LeyendasFiscales" bson:"LeyendasFiscales,omitempty"`
	Nomina                         *nomina.Nomina                                       `xml:"Nomina" bson:"Nomina,omitempty"`
	NotariosPublicos               *notariospublicos.NotariosPublicos                   `xml:"NotariosPublicos" bson:"NotariosPublicos,omitempty"`
	Pagos                          *pagos.Pagos                                         `xml:"Pagos" bson:"Pagos,omitempty"`
	TimbreFiscalDigital            *timbrefiscaldigital.TimbreFiscalDigital             `xml:"TimbreFiscalDigital" bson:"TimbreFiscalDigital,omitempty"`
	ObrasAntiguedades              *obrasantiguedades.ObrasAntiguedades                 `xml:"obrasarteantiguedades" bson:"ObrasArteAntiguedaes,omitempty"`
	PagoEnEspecie                  *pagoespecie.PagoEnEspecie                           `xml:"PagoEnEspecie" bson:"PagoEnEspecie,omitempty"`
	ParcialesConstruccion          *parcialesconstruccion.ParcialesConstruccion         `xml:"parcialesconstruccion" bson:"parcialesconstruccion,omitempty"`
	PFIntegrantesCoordinado        *pfintegrantecoordinado.PFIntegranteCoordinado       `xml:"PFintegranteCoordinado" bson:"PFintegranteCoordinado,omitempty"`
	RenovacionSustitucionVehiculos *renovacionsustitucion.RenovacionSustitucion         `xml:"renovacionysustitucionvehiculos" bson:"RenovacionSustitucionVehiculos,omitempty"`
	TuristaPasajeroExtranjero      *turistapasajeroextranjero.TuristaPasajeroExtranjero `xml:"TuristaPasajeroExtranjero" bson:"TuristaPasajeroExtranjero,omitempty"`
	ValesDeDespensa                *valesdespensa.ValesDespensa                         `xml:"ValesDeDespensa" bson:"ValesDeDespensa,omitempty"`
	VehiculoUsado                  *vehiculousado.VehiculoUsado                         `xml:"VehiculoUsado" bson:"VehiculoUsado,omitempty"`
}

package comprobante

import (
	"github.com/SaulEnriqueMR/kore-models/models/addenda"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/aerolineas"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cartaporte"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/certificadodestruccion"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/cfdiregistrofiscal"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/comercioexterior"
	complementos "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/complementospei"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/consumocombustible"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/detallista"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/divisas"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/donatarias"
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
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/acreditamientoieps"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/acuentaterceros"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/institeducativas"
	"github.com/SaulEnriqueMR/kore-models/models/comprobante/complementosconcepto/ventavehiculos"
	"github.com/SaulEnriqueMR/kore-models/models/timbrefiscaldigital"
)

type ComplementoConcepto struct {
	ACuentaTerceros         *acuentaterceros.ACuentaTerceros         `xml:"PorCuentadeTerceros" bson:"ACuentaTerceros,omitempty" json:"ACuentaTerceros,omitempty"`
	InstitucionesEducativas *institeducativas.InstitucioneEducativas `xml:"instEducativas" bson:"InstitucionesEducativas,omitempty" json:"InstitucionesEducativas,omitempty"`
	VentaVehiculos          *ventavehiculos.VentaVehiculos           `xml:"VentaVehiculos" bson:"VentaVehiculos,omitempty" json:"VentaVehiculos,omitempty"`
	AcreditamientoIeps      *acreditamientoieps.AcreditamientoIeps   `xml:"acreditamientoIEPS" bson:"AcreditamientoIeps,omitempty" json:"AcreditamientoIeps,omitempty"`
}

type Complemento struct {
	CartaPorte                     *cartaporte.CartaPorte                               `xml:"CartaPorte" bson:"CartaPorte,omitempty" json:"CartaPorte,omitempty"`
	ComercioExterior               *comercioexterior.ComercioExterior                   `xml:"ComercioExterior" bson:"ComercioExterior,omitempty" json:"ComercioExterior,omitempty"`
	ImpuestoLocales                *impuestoslocales.ImpuestosLocales                   `xml:"ImpuestosLocales" bson:"ImpuestosLocales,omitempty" json:"ImpuestosLocales,omitempty"`
	Aerolineas                     *aerolineas.Aerolineas                               `xml:"Aerolineas" bson:"Aerolineas,omitempty" json:"Aerolineas,omitempty"`
	CertificadoDestruccion         *certificadodestruccion.CertificadoDeDestruccion     `xml:"certificadodedestruccion" bson:"CertificadoDeDestruccion,omitempty" json:"CertificadoDeDestruccion,omitempty"`
	CFDIRegistroFiscal             *cfdiregistrofiscal.CfdiRegistroFiscal               `xml:"CFDIRegistroFiscal" bson:"CFDIRegistroFiscal,omitempty" json:"CFDIRegistroFiscal,omitempty"`
	ComplementoSPEI                *complementos.ComplementoSpei                        `xml:"Complemento_SPEI" bson:"ComplementoSPEI,omitempty" json:"ComplementoSPEI,omitempty"`
	ConsumoCombustibles            *consumocombustible.ConsumoDeCombustible             `xml:"ConsumoDeCombustibles" bson:"ConsumoDeCombustibles,omitempty" json:"ConsumoDeCombustibles,omitempty"`
	Detallista                     *detallista.Detallista                               `xml:"detallista" bson:"detallista,omitempty" json:"detallista,omitempty"`
	Divisas                        *divisas.Divisas                                     `xml:"Divisas" bson:"Divisas,omitempty" json:"Divisas,omitempty"`
	Donatarias                     *donatarias.Donatarias                               `xml:"Donatarias" bson:"Donatarias,omitempty" json:"Donatarias,omitempty"`
	EstadoCuentaCombustible        *estadodecuentacombustible.EstadoDeCuentaCombustible `xml:"EstadoDeCuentaCombustible" bson:"EstadoDeCuentaCombustible,omitempty" json:"EstadoDeCuentaCombustible,omitempty"`
	GastoHidrocarburos             *gastohidrocarburos.GastoHidrocarburos               `xml:"GastosHidrocarburos" bson:"GastosHidrocarburos,omitempty" json:"GastosHidrocarburos,omitempty"`
	Ine                            *ine.INE                                             `xml:"INE" bson:"Ine,omitempty" json:"Ine,omitempty"`
	LeyendasFiscales               *leyendasfiscales.LeyendasFiscales                   `xml:"LeyendasFiscales" bson:"LeyendasFiscales,omitempty" json:"LeyendasFiscales,omitempty"`
	Nomina                         *nomina.Nomina                                       `xml:"Nomina" bson:"Nomina,omitempty" json:"Nomina,omitempty"`
	NotariosPublicos               *notariospublicos.NotariosPublicos                   `xml:"NotariosPublicos" bson:"NotariosPublicos,omitempty" json:"NotariosPublicos,omitempty"`
	Pagos                          *pagos.Pagos                                         `xml:"Pagos" bson:"Pagos,omitempty" json:"Pagos,omitempty"`
	TimbreFiscalDigital            *timbrefiscaldigital.TimbreFiscalDigital             `xml:"TimbreFiscalDigital" bson:"TimbreFiscalDigital,omitempty" json:"TimbreFiscalDigital,omitempty"`
	ObrasAntiguedades              *obrasantiguedades.ObrasAntiguedades                 `xml:"obrasarteantiguedades" bson:"ObrasArteAntiguedaes,omitempty" json:"ObrasArteAntiguedaes,omitempty"`
	PagoEnEspecie                  *pagoespecie.PagoEnEspecie                           `xml:"PagoEnEspecie" bson:"PagoEnEspecie,omitempty" json:"PagoEnEspecie,omitempty"`
	ParcialesConstruccion          *parcialesconstruccion.ParcialesConstruccion         `xml:"parcialesconstruccion" bson:"parcialesconstruccion,omitempty" json:"parcialesconstruccion,omitempty"`
	PFIntegrantesCoordinado        *pfintegrantecoordinado.PFIntegranteCoordinado       `xml:"PFintegranteCoordinado" bson:"PFintegranteCoordinado,omitempty" json:"PFintegranteCoordinado,omitempty"`
	RenovacionSustitucionVehiculos *renovacionsustitucion.RenovacionSustitucion         `xml:"renovacionysustitucionvehiculos" bson:"RenovacionSustitucionVehiculos,omitempty" json:"RenovacionSustitucionVehiculos,omitempty"`
	TuristaPasajeroExtranjero      *turistapasajeroextranjero.TuristaPasajeroExtranjero `xml:"TuristaPasajeroExtranjero" bson:"TuristaPasajeroExtranjero,omitempty" json:"TuristaPasajeroExtranjero,omitempty"`
	ValesDeDespensa                *valesdespensa.ValesDespensa                         `xml:"ValesDeDespensa" bson:"ValesDeDespensa,omitempty" json:"ValesDeDespensa,omitempty"`
	VehiculoUsado                  *vehiculousado.VehiculoUsado                         `xml:"VehiculoUsado" bson:"VehiculoUsado,omitempty" json:"VehiculoUsado,omitempty"`
}

type Addenda struct {
	Value             string                       `xml:",innerxml" bson:"Value" json:"Value"`
	BuzonFiscal       *[]addenda.BuzonFiscal       `xml:"AddendaBuzonFiscal" bson:"BuzonFiscal,omitempty" json:"BuzonFiscal,omitempty"`
	AddendaFacto      *[]addenda.AddendaFacto      `xml:"addendaFacto" bson:"AddendaFacto,omitempty" json:"AddendaFacto,omitempty"`
	EdicomGenerica    *[]addenda.EdicomGenerica    `xml:"customized>EDICOM_GENERICA" bson:"EdicomGenerica,omitempty" json:"EdicomGenerica,omitempty"`
	RBoschSap         *[]addenda.RBoschSap         `xml:"customized>RBOSCH_SAP" bson:"RBoschSap,omitempty" json:"RBoschSap,omitempty"`
	RequestForPayment *[]addenda.RequestForPayment `xml:"requestForPayment" bson:"RequestForPayment,omitempty" json:"RequestForPayment,omitempty"`
	Innsist           *[]addenda.Inssist           `xml:"Inssist" bson:"Inssist,omitempty" json:"Inssist,omitempty"`
}

type KuantikMetadata struct {
	SerieFolio string `bson:"SerieFolio" json:"SerieFolio"`
}

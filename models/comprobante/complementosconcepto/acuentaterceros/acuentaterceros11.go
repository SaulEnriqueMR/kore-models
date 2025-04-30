package acuentaterceros

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type ACuentaTerceros11 struct {
	Version                  string                              `xml:"version,attr" bson:"Version" json:"Version"`
	Rfc                      string                              `xml:"rfc,attr" bson:"Rfc" json:"Rfc"`
	Nombre                   *string                             `xml:"nombre,attr" bson:"Nombre,omitempty" json:"Nombre,omitempty"`
	InformacionFiscalTercero *InformacionFiscalTerceroTerceros11 `xml:"InformacionFiscalTercero" bson:"InformacionFiscalTercero,omitempty" json:"InformacionFiscalTercero,omitempty"`
	InformacionAduanera      *InformacionAduaneraTerceros11      `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
	Parte                    *[]ParteTerceros11                  `xml:"Parte" bson:"Parte,omitempty" json:"Parte,omitempty"`
	CuentaPredial            *CuentaPredialTerceros11            `xml:"CuentaPredial" bson:"CuentaPredial,omitempty" json:"CuentaPredial,omitempty"`
	Impuestos                ImpuestosTerceros11                 `xml:"Impuestos" bson:"Impuestos" json:"Impuestos"`
}

type InformacionFiscalTerceroTerceros11 struct {
	Calle        string  `xml:"calle,attr" bson:"Calle" json:"Calle"`
	NoExterior   *string `xml:"noExterior,attr" bson:"NoExterior,omitempty" json:"NoExterior,omitempty"`
	NoInterior   *string `xml:"noInterior,attr" bson:"NoInterior,omitempty" json:"NoInterior,omitempty"`
	Colonia      *string `xml:"colonia,attr" bson:"Colonia,omitempty" json:"Colonia,omitempty"`
	Localidad    *string `xml:"localidad,attr" bson:"Localidad,omitempty" json:"Localidad,omitempty"`
	Referencia   *string `xml:"referencia,attr" bson:"Referencia,omitempty" json:"Referencia,omitempty"`
	Municipio    string  `xml:"municipio,attr" bson:"Municipio" json:"Municipio"`
	Estado       string  `xml:"estado,attr" bson:"Estado" json:"Estado"`
	Pais         string  `xml:"pais,attr" bson:"Pais" json:"Pais"`
	CodigoPostal string  `xml:"codigoPostal,attr" bson:"CodigoPostal" json:"CodigoPostal"`
}

type InformacionAduaneraTerceros11 struct {
	Numero      string    `xml:"numero,attr" bson:"Numero" json:"Numero"`
	FechaString string    `xml:"fecha,attr" bson:"FechaString" json:"FechaString"`
	Fecha       time.Time `bson:"Fecha" json:"Fecha"`
	Aduana      *string   `xml:"aduana,attr" bson:"Aduana,omitempty" json:"Aduana,omitempty"`
}

func (iat *InformacionAduaneraTerceros11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias InformacionAduaneraTerceros11
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*iat = InformacionAduaneraTerceros11(aux)

	if iat != nil {
		fecha, err := helpers.ParseDatetime(iat.FechaString)
		if err != nil {
			return err
		}
		iat.Fecha = fecha
	}

	return nil
}

type ParteTerceros11 struct {
	Cantidad            float64                          `xml:"cantidad,attr" bson:"Cantidad" json:"Cantidad"`
	Unidad              *string                          `xml:"unidad,attr" bson:"Unidad,omitempty" json:"Unidad,omitempty"`
	NoIdentificacion    *string                          `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty" json:"NoIdentificacion,omitempty"`
	Descripcion         string                           `xml:"descripcion,attr" bson:"Descripcion" json:"Descripcion"`
	ValorUnitario       *float64                         `xml:"valorUnitario,attr" bson:"ValorUnitario,omitempty" json:"ValorUnitario,omitempty"`
	Importe             *float64                         `xml:"importe,attr" bson:"Importe,omitempty" json:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduaneraTerceros11 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty" json:"InformacionAduanera,omitempty"`
}

type CuentaPredialTerceros11 struct {
	Numero string `xml:"numero,attr" bson:"Numero" json:"Numero"`
}

type ImpuestosTerceros11 struct {
	Retenciones *[]RetencionTerceros11 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty" json:"Retenciones,omitempty"`
	Traslados   *[]TrasladoTerceros11  `xml:"Traslados>Traslado" bson:"Traslados,omitempty" json:"Traslados,omitempty"`
}

type RetencionTerceros11 struct {
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto" json:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto" json:"Impuesto"`
	Importe      float64 `xml:"importe,attr" bson:"Importe" json:"Importe"`
}

func (i *RetencionTerceros11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias RetencionTerceros11
	var aux Alias

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*i = RetencionTerceros11(aux)

	switch strings.ToLower(i.TipoImpuesto) {
	case "isr":
		i.Impuesto = "001"
		break
	case "iva":
		i.Impuesto = "002"
		break
	case "ieps":
		i.Impuesto = "003"
		break
	}

	return nil
}

type TrasladoTerceros11 struct {
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto" json:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto" json:"Impuesto"`
	Tasa         float64 `xml:"tasa,attr" bson:"Tasa" json:"Tasa"`
	Importe      float64 `xml:"importe,attr" bson:"Importe" json:"Importe"`
}

func (i *TrasladoTerceros11) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias TrasladoTerceros11
	var aux Alias

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	*i = TrasladoTerceros11(aux)

	switch strings.ToLower(i.TipoImpuesto) {
	case "isr":
		i.Impuesto = "001"
		break
	case "iva":
		i.Impuesto = "002"
		break
	case "ieps":
		i.Impuesto = "003"
		break
	}

	return nil
}

package acuentaterceros

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type ACuentaTerceros11 struct {
	Version                  string                              `xml:"version,attr" bson:"Version"`
	Rfc                      string                              `xml:"rfc,attr" bson:"Rfc"`
	Nombre                   *string                             `xml:"nombre,attr" bson:"Nombre,omitempty"`
	InformacionFiscalTercero *InformacionFiscalTerceroTerceros11 `xml:"InformacionFiscalTercero" bson:"InformacionFiscalTercero,omitempty"`
	InformacionAduanera      *InformacionAduaneraTerceros11      `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
	Parte                    *[]ParteTerceros11                  `xml:"Parte" bson:"Parte,omitempty"`
	CuentaPredial            *CuentaPredialTerceros11            `xml:"CuentaPredial" bson:"CuentaPredial,omitempty"`
	Impuestos                ImpuestosTerceros11                 `xml:"Impuestos" bson:"Impuestos"`
}

type InformacionFiscalTerceroTerceros11 struct {
	Calle        string  `xml:"calle,attr" bson:"Calle"`
	NoExterior   *string `xml:"noExterior,attr" bson:"NoExterior,omitempty"`
	NoInterior   *string `xml:"noInterior,attr" bson:"NoInterior,omitempty"`
	Colonia      *string `xml:"colonia,attr" bson:"Colonia,omitempty"`
	Localidad    *string `xml:"localidad,attr" bson:"Localidad,omitempty"`
	Referencia   *string `xml:"referencia,attr" bson:"Referencia,omitempty"`
	Municipio    string  `xml:"municipio,attr" bson:"Municipio"`
	Estado       string  `xml:"estado,attr" bson:"Estado"`
	Pais         string  `xml:"pais,attr" bson:"Pais"`
	CodigoPostal string  `xml:"codigoPostal,attr" bson:"CodigoPostal"`
}

type InformacionAduaneraTerceros11 struct {
	Numero      string    `xml:"numero,attr" bson:"Numero"`
	FechaString string    `xml:"fecha,attr" bson:"FechaString"`
	Fecha       time.Time `bson:"Fecha"`
	Aduana      *string   `xml:"aduana,attr" bson:"Aduana,omitempty"`
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
	Cantidad            float64                          `xml:"cantidad,attr" bson:"Cantidad"`
	Unidad              *string                          `xml:"unidad,attr" bson:"Unidad,omitempty"`
	NoIdentificacion    *string                          `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Descripcion         string                           `xml:"descripcion,attr" bson:"Descripcion"`
	ValorUnitario       *float64                         `xml:"valorUnitario,attr" bson:"ValorUnitario,omitempty"`
	Importe             *float64                         `xml:"importe,attr" bson:"Importe,omitempty"`
	InformacionAduanera *[]InformacionAduaneraTerceros11 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
}

type CuentaPredialTerceros11 struct {
	Numero string `xml:"numero,attr" bson:"Numero"`
}

type ImpuestosTerceros11 struct {
	Retenciones *[]RetencionTerceros11 `xml:"Retenciones>Retencion" bson:"Retenciones,omitempty"`
	Traslados   *[]TrasladoTerceros11  `xml:"Traslados>Traslado" bson:"Traslados,omitempty"`
}

type RetencionTerceros11 struct {
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto"`
	Importe      float64 `xml:"importe,attr" bson:"Importe"`
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
	TipoImpuesto string  `xml:"impuesto,attr" bson:"TipoImpuesto"`
	Impuesto     string  `bson:"Impuesto"`
	Tasa         float64 `xml:"tasa,attr" bson:"Tasa"`
	Importe      float64 `xml:"importe,attr" bson:"Importe"`
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

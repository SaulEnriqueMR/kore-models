package ventavehiculos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type VentaVehiculos10 struct {
	Version        string             `xml:"version,attr" bson:"Version"`
	ClaveVehicular string             `xml:"ClaveVehicular,attr" bson:"ClaveVehicular"`
	Parte          *[]ParteVentVehi10 `xml:"Parte" bson:"Parte,omitempty"`
}

type ParteVentVehi10 struct {
	InformacionAduanera *[]InformacionAduaneraVentVehi10 `xml:"InformacionAduanera" bson:"InformacionAduanera,omitempty"`
	Cantidad            string                           `xml:"cantidad,attr" bson:"Cantidad"`
	Unidad              *string                          `xml:"unidad,attr" bson:"Unidad,omitempty"`
	NoIdentificacion    *string                          `xml:"noIdentificacion,attr" bson:"NoIdentificacion,omitempty"`
	Descripcion         string                           `xml:"descripcion,attr" bson:"Descripcion"`
	ValorUnitario       *string                          `xml:"valorUnitario,attr" bson:"ValorUnitario,omitempty"`
	Importe             *string                          `xml:"importe,attr" bson:"Importe,omitempty"`
}

type InformacionAduaneraVentVehi10 struct {
	Numero      string    `xml:"numero,attr" bson:"Numero"`
	FechaString string    `xml:"fecha,attr" bson:"FechaString"`
	Fecha       time.Time `bson:"Fecha"`
	Aduana      *string   `xml:"aduana,attr" bson:"Aduana,omitempty"`
}

func (iavv10 *InformacionAduaneraVentVehi10) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias InformacionAduaneraVentVehi10
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*iavv10 = InformacionAduaneraVentVehi10(aux)

	if iavv10 != nil {
		fecha, err := helpers.ParseDatetime(iavv10.FechaString)
		if err != nil {
			return err
		}
		iavv10.Fecha = fecha
	}

	return nil
}

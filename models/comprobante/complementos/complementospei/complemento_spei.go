package complementos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ComplementoSpei struct {
	SpeiTercero []SpeiTerceroCompSpei `xml:"SPEI_Tercero" bson:"SpeiTercero"`
}

type SpeiTerceroCompSpei struct {
	FechaOperacionString string                 `xml:"FechaOperacion,attr"`
	FechaOperacion       time.Time              `bson:"FechaOperacion"`
	Hora                 string                 `xml:"Hora,attr" bson:"Hora"`
	ClaveSpei            string                 `xml:"ClaveSPEI,attr" bson:"ClaveSpei"`
	Sello                string                 `xml:"sello,attr" bson:"Sello"`
	NoCertificado        string                 `xml:"numeroCertificado,attr" bson:"NoCertificado"`
	CadenaCda            string                 `xml:"cadenaCDA,attr" bson:"cadenaCda,omitempty"`
	Ordenante            []OrdenanteCompSpei    `xml:"Ordenante" bson:"Ordenante"`
	Beneficiario         []BeneficiarioCompSpei `xml:"Beneficiario" bson:"Beneficiario"`
}

type OrdenanteCompSpei struct {
	BancoEmisor string `xml:"BancoEmisor,attr" bson:"BancoEmisor"`
	Nombre      string `xml:"Nombre,attr" bson:"Nombre"`
	TipoCuenta  string `xml:"TipoCuenta,attr" bson:"TipoCuenta"`
	Cuenta      string `xml:"Cuenta,attr" bson:"Cuenta"`
	Rfc         string `xml:"RFC,attr" bson:"Rfc"`
}

type BeneficiarioCompSpei struct {
	BancoReceptor string  `xml:"BancoReceptor,attr" bson:"BancoReceptor"`
	Nombre        string  `xml:"Nombre,attr" bson:"Nombre"`
	TipoCuenta    string  `xml:"TipoCuenta,attr" bson:"TipoCuenta"`
	Cuenta        string  `xml:"Cuenta,attr" bson:"Cuenta"`
	Rfc           string  `xml:"RFC,attr" bson:"Rfc"`
	Concepto      string  `xml:"Concepto,attr" bson:"Concepto"`
	Iva           *string `xml:"IVA,attr" bson:"Iva,omitempty"`
	MontoPago     string  `xml:"MontoPago,attr" bson:"MontoPago"`
}

func (stcs *SpeiTerceroCompSpei) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias SpeiTerceroCompSpei
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*stcs = SpeiTerceroCompSpei(aux)

	if stcs != nil {
		fecha, err := helpers.ParseDatetime(stcs.FechaOperacionString)
		if err != nil {
			return err
		}
		stcs.FechaOperacion = fecha
	}

	return nil
}

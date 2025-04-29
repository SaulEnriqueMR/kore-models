package complementos

import (
	"encoding/xml"
	"github.com/SaulEnriqueMR/kore-models/models/helpers"
	"time"
)

type ComplementoSpei struct {
	SpeiTercero []SpeiTerceroCompSpei `xml:"SPEI_Tercero" bson:"SpeiTercero" json:"SpeiTercero"`
}

type SpeiTerceroCompSpei struct {
	FechaOperacionString string                 `xml:"FechaOperacion,attr" bson:"FechaOperacionString" json:"FechaOperacionString"`
	FechaOperacion       time.Time              `bson:"FechaOperacion" json:"FechaOperacion"`
	Hora                 string                 `xml:"Hora,attr" bson:"Hora" json:"Hora"`
	ClaveSpei            string                 `xml:"ClaveSPEI,attr" bson:"ClaveSpei" json:"ClaveSpei"`
	Sello                string                 `xml:"sello,attr" bson:"Sello" json:"Sello"`
	NoCertificado        string                 `xml:"numeroCertificado,attr" bson:"NoCertificado" json:"NoCertificado"`
	CadenaCda            string                 `xml:"cadenaCDA,attr" bson:"cadenaCda,omitempty" json:"cadenaCda,omitempty"`
	Ordenante            []OrdenanteCompSpei    `xml:"Ordenante" bson:"Ordenante" json:"Ordenante"`
	Beneficiario         []BeneficiarioCompSpei `xml:"Beneficiario" bson:"Beneficiario" json:"Beneficiario"`
}

type OrdenanteCompSpei struct {
	BancoEmisor string `xml:"BancoEmisor,attr" bson:"BancoEmisor" json:"BancoEmisor"`
	Nombre      string `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	TipoCuenta  string `xml:"TipoCuenta,attr" bson:"TipoCuenta" json:"TipoCuenta"`
	Cuenta      string `xml:"Cuenta,attr" bson:"Cuenta" json:"Cuenta"`
	Rfc         string `xml:"RFC,attr" bson:"Rfc" json:"Rfc"`
}

type BeneficiarioCompSpei struct {
	BancoReceptor string  `xml:"BancoReceptor,attr" bson:"BancoReceptor" json:"BancoReceptor"`
	Nombre        string  `xml:"Nombre,attr" bson:"Nombre" json:"Nombre"`
	TipoCuenta    string  `xml:"TipoCuenta,attr" bson:"TipoCuenta" json:"TipoCuenta"`
	Cuenta        string  `xml:"Cuenta,attr" bson:"Cuenta" json:"Cuenta"`
	Rfc           string  `xml:"RFC,attr" bson:"Rfc" json:"Rfc"`
	Concepto      string  `xml:"Concepto,attr" bson:"Concepto" json:"Concepto"`
	Iva           *string `xml:"IVA,attr" bson:"Iva,omitempty" json:"Iva,omitempty"`
	MontoPago     string  `xml:"MontoPago,attr" bson:"MontoPago" json:"MontoPago"`
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

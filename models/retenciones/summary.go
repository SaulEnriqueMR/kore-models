package retenciones

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type SummarizedRetenciones struct {
	EmisorRfc string    `json:"emisorRfc" bson:"emisorRfc"`
	Receptor  string    `json:"receptor" bson:"receptor"`
	Uuid      string    `json:"uuid" bson:"uuid"`
	Fecha     time.Time `json:"fecha" bson:"fecha"`
}

func (f *SummarizedRetenciones) GetFileName() string {
	year := f.Fecha.Year()
	month := f.Fecha.Month()
	// log.Println(f.EmisorRfc + "/" + f.ReceptorRfc + "/" + fmt.Sprintf("%d", year) + "/" + fmt.Sprintf("%d", month) + "/" + f.Uuid + ".xml")
	return f.EmisorRfc + "/" + f.Receptor + "/" + fmt.Sprintf("%d", year) + "/" + fmt.Sprintf("%d", month) + "/" + f.Uuid + ".xml"
}

func MyFilename(d *xml.Decoder) (string, error) {
	var EmisorRfc string
	var ReceptorRfc string
	var Uuid string
	var Fecha time.Time

	summarized := SummarizedRetenciones{}
	for {
		t, _ := d.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "TimbreFiscalDigital" {
				start := xml.StartElement(se)
				for _, attributes := range start.Attr {
					if strings.ToLower(attributes.Name.Local) == "uuid" {
						// filename.Uuid = attributes.Value
						Uuid = attributes.Value
					}
				}
			}
			if se.Name.Local == "Emisor" {
				start := xml.StartElement(se)
				for _, attributes := range start.Attr {
					if strings.ToLower(attributes.Name.Local) == "rfc" {
						EmisorRfc = attributes.Value
					}
				}
			}
			if se.Name.Local == "Receptor" {
				start := xml.StartElement(se)
				for _, attributes := range start.Attr {
					if strings.ToLower(attributes.Name.Local) == "nacional" {
						// ReceptorRfc = attributes.Value
						startNacional := xml.StartElement(se)
						for _, attributesNacional := range startNacional.Attr {
							if strings.ToLower(attributesNacional.Name.Local) == "rfc" {
								ReceptorRfc = attributesNacional.Value
							}
						}
					} else if strings.ToLower(attributes.Name.Local) == "extranjero" {
						// ReceptorRfc = attributes.Value
						startExtranjero := xml.StartElement(se)
						for _, attributesExtranjero := range startExtranjero.Attr {
							if strings.ToLower(attributesExtranjero.Name.Local) == "numregidtrib" && attributesExtranjero.Value != "" {
								ReceptorRfc = attributesExtranjero.Value
							} else if strings.ToLower(attributesExtranjero.Name.Local) == "nombre" && attributesExtranjero.Value != "" {
								ReceptorRfc = attributesExtranjero.Value
							} else {
								ReceptorRfc = "EXTRANJERO"
							}
						}
					}
				}
			}
			if se.Name.Local == "Retenciones" {
				for _, attributes := range se.Attr {
					if strings.ToLower(attributes.Name.Local) == "fechaexp" {
						fecha, err := helpers.ParseDatetime(attributes.Value)
						if err != nil {
							spliData := strings.Split(attributes.Value, "-")
							if len(spliData) >= 3 {
								year := spliData[0]
								month := spliData[1]
								day := spliData[2]
								fecha, err = time.Parse("2006-01-02", year+"-"+month+"-"+day)
								if err != nil {
									continue
								}
							}
						}
						Fecha = fecha
					}
				}
			}
			// case xml.EndElement:
			// log.Println("end", se.Name.Local)
			// default:
			// 	log.Println("default", se)
		}
	}
	if Fecha.IsZero() {
		return "", errors.New("error parsing datetime. Fecha is empty")
	}
	summarized.EmisorRfc = strings.ToUpper(EmisorRfc)
	summarized.Receptor = strings.ToUpper(ReceptorRfc)
	summarized.Uuid = strings.ToUpper(Uuid)
	summarized.Fecha = Fecha

	return summarized.GetFileName(), nil
}

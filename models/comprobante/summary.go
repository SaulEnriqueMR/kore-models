package comprobante

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type SummarizedComprobante struct {
	EmisorRfc   string    `json:"emisorRfc" bson:"emisorRfc"`
	ReceptorRfc string    `json:"receptorRfc" bson:"receptorRfc"`
	Uuid        string    `json:"uuid" bson:"uuid"`
	Fecha       time.Time `json:"fecha" bson:"fecha"`
}

func (f *SummarizedComprobante) GetFileName() string {
	year := f.Fecha.Year()
	month := f.Fecha.Month()
	log.Println(f.EmisorRfc + "/" + f.ReceptorRfc + "/" + fmt.Sprintf("%d", year) + "/" + fmt.Sprintf("%d", month) + "/" + f.Uuid + ".xml")
	return f.EmisorRfc + "/" + f.ReceptorRfc + "/" + fmt.Sprintf("%d", year) + "/" + fmt.Sprintf("%d", month) + "/" + f.Uuid + ".xml"
}

func MyFilename(d *xml.Decoder) (string, error) {
	var EmisorRfc string
	var ReceptorRfc string
	var Uuid string
	var Fecha time.Time

	summarized := SummarizedComprobante{}
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
					if attributes.Name.Local == "UUID" {
						// filename.Uuid = attributes.Value
						Uuid = attributes.Value
					}
				}
			}
			if se.Name.Local == "Emisor" {
				start := xml.StartElement(se)
				for _, attributes := range start.Attr {
					if strings.ToLower(attributes.Name.Local) == "rfc" {
						// filename.EmisorRfc = attributes.Value
						EmisorRfc = attributes.Value
					}
				}
			}
			if se.Name.Local == "Receptor" {
				start := xml.StartElement(se)
				for _, attributes := range start.Attr {
					if strings.ToLower(attributes.Name.Local) == "rfc" {
						// filename.EmisorRfc = attributes.Value
						ReceptorRfc = attributes.Value
					}
				}
			}
			if se.Name.Local == "Comprobante" {
				for _, attributes := range se.Attr {
					if strings.ToLower(attributes.Name.Local) == "fecha" {
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
	summarized.EmisorRfc = EmisorRfc
	summarized.ReceptorRfc = ReceptorRfc
	summarized.Uuid = Uuid

	return summarized.GetFileName(), nil
}

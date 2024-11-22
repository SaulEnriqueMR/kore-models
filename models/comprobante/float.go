package comprobante

import (
	"encoding/xml"
	"strconv"
)

type Float float64

func (f *Float) UnmarshalXMLAttr(attr xml.Attr) error {
	value, err := strconv.ParseFloat(attr.Value, 64)
	if err != nil {
		return err
	}
	*f = Float(value)
	return nil
}

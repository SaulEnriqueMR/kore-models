package serializate

import "github.com/SaulEnriqueMR/kore-models/models/helpers"

func SerializeComprobanteFromXml[comprobanteType any](inputXml []byte) (comprobanteType, error) {
	var parsed comprobanteType
	trimmedXml, errOnTrim := helpers.TrimXml(inputXml)
	if errOnTrim != nil {
		return parsed, errOnTrim
	}

	errUnmarshal := helpers.UnmarshalXMLWithEncoding(trimmedXml, &parsed)
	if errUnmarshal != nil {
		return parsed, errUnmarshal
	}
	return parsed, nil
}

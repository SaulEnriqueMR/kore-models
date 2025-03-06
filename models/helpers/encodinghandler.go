package helpers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"io"
)

// CharsetReader is a reusable function to handle multiple encodings
func CharsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch charset {
	case "ISO-8859-1":
		return charmap.ISO8859_1.NewDecoder().Reader(input), nil
	case "Windows-1252":
		return charmap.Windows1252.NewDecoder().Reader(input), nil
	case "UTF-8", "UTF-8-SIG":
		// UTF-8 doesn't need decoding
		return input, nil
	case "UTF-16":
		// Handle UTF-16 encoding
		return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder().Reader(input), nil
	default:
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}
}

// UnmarshalXMLWithEncoding handles []byte input and supports multiple encodings
func UnmarshalXMLWithEncoding(data []byte, v interface{}) error {
	// Wrap the []byte in a bytes.Reader to make it an io.Reader
	reader := bytes.NewReader(data)

	// Create an xml.Decoder
	decoder := xml.NewDecoder(reader)

	// Set the CharsetReader to handle multiple encodings
	decoder.CharsetReader = CharsetReader

	// Decode the XML into the provided struct
	return decoder.Decode(v)
}

package cfdi

type Contribuyente struct {
	Rfc           string
	Nombre        string
	RegimenFiscal string
}

type Emisor struct {
	Contribuyente
	FactAtrAdquirentEmisor string
}

type Receptor struct {
	Contribuyente
	DomicilioFiscal  string
	ResidenciaFiscal string
	NumRegIdTrib     string
	UsoCFDI          string
}

package cfdi

import (
	"time"

	cfdicomplementos "github.com/SaulEnriqueMR/kore-models/models/comprobante/cfdi/cfdi_complementos"
)

// CFDI Modelo unificado para el manejo de CFDI en la plataforma Kore independientemente de la version.
type CFDI struct {
	Uuid             string
	Serie            string
	Folio            string
	SerieFolio       string
	Version          string
	Vigencia         bool
	Tipo             string
	MetodoPago       string
	FechaEmision     *time.Time
	FechaCancelacion *time.Time
	FechaTimbrado    *time.Time
	Emisor           Emisor
	Receptor         Receptor
	Moneda           string
	FormaPago        string
	CondicionesPago  string
	Descuento        *float64
	TipoCambio       float64
	Subtotal         float64
	Total            float64
	Complemento      cfdicomplementos.ComplementoCfdi
	CfdiRelacionados []CfdiRelacionado
}

type CfdiRelacionado struct {
	Uuid         string
	TipoRelacion string
}

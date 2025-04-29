package addenda

type RBoschSap struct {
	Cabecera RBoschSapCabecera `xml:"CABECERA" bson:"Cabecera" json:"Cabecera"`
	Detalle  RBoschSapDetalle  `xml:"DETALLE" bson:"Detalle" json:"Detalle"`
}

type RBoschSapCabecera struct {
	ClaveCli          string          `xml:"CLAVECLI,attr" bson:"ClaveCli" json:"ClaveCli"`
	ClaveNum          string          `xml:"CLAVENUM,attr" bson:"ClaveNum" json:"ClaveNum"`
	CondicionesPago   string          `xml:"CONDICIONESPAGO,attr" bson:"CondicionesPago" json:"CondicionesPago"`
	Delivery          string          `xml:"DELIVERY,attr" bson:"Delivery" json:"Delivery"`
	Fecha             string          `xml:"FECHA,attr" bson:"Fecha" json:"Fecha"`
	Folio             string          `xml:"FOLIO,attr" bson:"Folio" json:"Folio"`
	FolioErp          string          `xml:"FOLIOERP,attr" bson:"FolioErp" json:"FolioErp"`
	ImporteTotalLetra string          `xml:"IMPTOTALLETRA,attr" bson:"ImporteTotalLetra" json:"ImporteTotalLetra"`
	Nodo              string          `xml:"NODO,attr" bson:"Nodo" json:"Nodo"`
	Peso              string          `xml:"PESO,attr" bson:"Peso" json:"Peso"`
	TelefX            string          `xml:"TELF_X,attr" bson:"TelefX" json:"TelefX"`
	TotalFactura      string          `xml:"TOTALFAC,attr" bson:"TotalFactura" json:"TotalFactura"`
	TotalNeto         string          `xml:"TOTNETO,attr" bson:"TotalNeto" json:"TotalNeto"`
	Moneda            RBoschSapMoneda `xml:"MONEDA" bson:"Moneda" json:"Moneda"`
	BillTo            RBoschSapBillTo `xml:"BILLTO" bson:"BillTo" json:"BillTo"`
	ShipTo            RBoschSapShipTo `xml:"SHIPTO" bson:"ShipTo" json:"ShipTo"`
}

type RBoschSapDetalle struct {
	Lineas RBoschSapLineas `xml:"LINEAS" bson:"Lineas" json:"Lineas"`
}

type RBoschSapMoneda struct {
	Moneda     string `xml:"MONEDA,attr" bson:"Moneda" json:"Moneda"`
	TipoCambio string `xml:"TIPO_CAMBIO,attr" bson:"TipoCambio" json:"TipoCambio"`
}

type RBoschSapBillTo struct {
	Calle        string `xml:"CALLE,attr" bson:"Calle" json:"Calle"`
	Ciudad       string `xml:"CIUDAD,attr" bson:"Ciudad" json:"Ciudad"`
	CodigoPostal string `xml:"CP,attr" bson:"CodigoPostal" json:"CodigoPostal"`
	Email        string `xml:"EMAIL,attr" bson:"Email" json:"Email"`
	Estado       string `xml:"ESTADO,attr" bson:"Estado" json:"Estado"`
	Nombre       string `xml:"NOMBRE,attr" bson:"Nombre" json:"Nombre"`
	Pais         string `xml:"PAIS,attr" bson:"Pais" json:"Pais"`
}

type RBoschSapShipTo struct {
	Calle        string `xml:"CALLE_RECEPTOR,attr" bson:"Calle" json:"Calle"`
	Ciudad       string `xml:"CIUDAD_RECEPTOR,attr" bson:"Ciudad" json:"Ciudad"`
	CodigoPostal string `xml:"CP_RECEPTOR,attr" bson:"CodigoPostal" json:"CodigoPostal"`
	Estado       string `xml:"ESTADO_RECEPTOR,attr" bson:"Estado" json:"Estado"`
	Nombre       string `xml:"NOMBRE_RECEPTOR,attr" bson:"Nombre" json:"Nombre"`
}

type RBoschSapLineas struct {
	Cantidad          string `xml:"CANTIDAD,attr" bson:"Cantidad" json:"Cantidad"`
	MontoLinea        string `xml:"MONTOLIN,attr" bson:"MontoLinea" json:"MontoLinea"`
	NoParte           string `xml:"NPARTE,attr" bson:"NoParte" json:"NoParte"`
	NoLinea           string `xml:"NUMLIN,attr" bson:"NoLinea" json:"NoLinea"`
	Pedido            string `xml:"PEDIDO,attr" bson:"Pedido" json:"Pedido"`
	PrecioNeto        string `xml:"PRECIONETO,attr" bson:"PrecioNeto" json:"PrecioNeto"`
	ReferenciaCliente string `xml:"REFCLIENTE,attr" bson:"ReferenciaCliente" json:"ReferenciaCliente"`
	UnidadMedida      string `xml:"UMEDIDA,attr" bson:"UnidadMedida" json:"UnidadMedida"`
}

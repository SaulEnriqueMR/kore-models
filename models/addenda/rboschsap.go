package addenda

type RBoschSap struct {
	Cabecera RBoschSapCabecera `xml:"CABECERA" bson:"Cabecera"`
	Detalle  RBoschSapDetalle  `xml:"DETALLE" bson:"Detalle"`
}

type RBoschSapCabecera struct {
	ClaveCli          string          `xml:"CLAVECLI,attr" bson:"ClaveCli"`
	ClaveNum          string          `xml:"CLAVENUM,attr" bson:"ClaveNum"`
	CondicionesPago   string          `xml:"CONDICIONESPAGO,attr" bson:"CondicionesPago"`
	Delivery          string          `xml:"DELIVERY,attr" bson:"Delivery"`
	Fecha             string          `xml:"FECHA,attr" bson:"Fecha"`
	Folio             string          `xml:"FOLIO,attr" bson:"Folio"`
	FolioErp          string          `xml:"FOLIOERP,attr" bson:"FolioErp"`
	ImporteTotalLetra string          `xml:"IMPTOTALLETRA,attr" bson:"ImporteTotalLetra"`
	Nodo              string          `xml:"NODO,attr" bson:"Nodo"`
	Peso              string          `xml:"PESO,attr" bson:"Peso"`
	TelefX            string          `xml:"TELF_X,attr" bson:"TelefX"`
	TotalFactura      string          `xml:"TOTALFAC,attr" bson:"TotalFactura"`
	TotalNeto         string          `xml:"TOTNETO,attr" bson:"TotalNeto"`
	Moneda            RBoschSapMoneda `xml:"MONEDA" bson:"Moneda"`
	BillTo            RBoschSapBillTo `xml:"BILLTO" bson:"BillTo"`
	ShipTo            RBoschSapShipTo `xml:"SHIPTO" bson:"ShipTo"`
}

type RBoschSapDetalle struct {
	Lineas RBoschSapLineas `xml:"LINEAS" bson:"Lineas"`
}

type RBoschSapMoneda struct {
	Moneda     string `xml:"MONEDA,attr" bson:"Moneda"`
	TipoCambio string `xml:"TIPO_CAMBIO,attr" bson:"TipoCambio"`
}

type RBoschSapBillTo struct {
	Calle        string `xml:"CALLE,attr" bson:"Calle"`
	Ciudad       string `xml:"CIUDAD,attr" bson:"Ciudad"`
	CodigoPostal string `xml:"CP,attr" bson:"CodigoPostal"`
	Email        string `xml:"EMAIL,attr" bson:"Email"`
	Estado       string `xml:"ESTADO,attr" bson:"Estado"`
	Nombre       string `xml:"NOMBRE,attr" bson:"Nombre"`
	Pais         string `xml:"PAIS,attr" bson:"Pais"`
}

type RBoschSapShipTo struct {
	Calle        string `xml:"CALLE_RECEPTOR,attr" bson:"Calle"`
	Ciudad       string `xml:"CIUDAD_RECEPTOR,attr" bson:"Ciudad"`
	CodigoPostal string `xml:"CP_RECEPTOR,attr" bson:"CodigoPostal"`
	Estado       string `xml:"ESTADO_RECEPTOR,attr" bson:"Estado"`
	Nombre       string `xml:"NOMBRE_RECEPTOR,attr" bson:"Nombre"`
}

type RBoschSapLineas struct {
	Cantidad          string `xml:"CANTIDAD,attr" bson:"Cantidad"`
	MontoLinea        string `xml:"MONTOLIN,attr" bson:"MontoLinea"`
	NoParte           string `xml:"NPARTE,attr" bson:"NoParte"`
	NoLinea           string `xml:"NUMLIN,attr" bson:"NoLinea"`
	Pedido            string `xml:"PEDIDO,attr" bson:"Pedido"`
	PrecioNeto        string `xml:"PRECIONETO,attr" bson:"PrecioNeto"`
	ReferenciaCliente string `xml:"REFCLIENTE,attr" bson:"ReferenciaCliente"`
	UnidadMedida      string `xml:"UMEDIDA,attr" bson:"UnidadMedida"`
}

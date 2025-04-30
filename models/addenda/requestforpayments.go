package addenda

// RequestForPayment represents the root element <requestForPayment>
type RequestForPayment struct {
	DeliveryDate                    string                                 `xml:"DeliveryDate,attr"`
	ContentVersion                  string                                 `xml:"contentVersion,attr"`
	DocumentStatus                  string                                 `xml:"documentStatus,attr"`
	DocumentStructureVersion        string                                 `xml:"documentStructureVersion,attr"`
	Type                            string                                 `xml:"type,attr"`
	RequestForPaymentIdentification AddendaRequestForPaymentIdentification `xml:"requestForPaymentIdentification"`
	SpecialInstructions             []AddendaSpecialInstruction            `xml:"specialInstruction"`
	OrderIdentification             AddendaOrderIdentification             `xml:"orderIdentification"`
	AdditionalInformation           AddendaAdditionalInformation           `xml:"AdditionalInformation"`
	Buyer                           AddendaBuyer                           `xml:"buyer"`
	Seller                          AddendaSeller                          `xml:"seller"`
	Currency                        AddendaCurrency                        `xml:"currency"`
	LineItem                        AddendaLineItem                        `xml:"lineItem"`
	TotalAmount                     AddendaTotalAmount                     `xml:"totalAmount"`
	TotalAllowanceCharge            AddendaTotalAllowanceCharge            `xml:"TotalAllowanceCharge"`
	BaseAmount                      AddendaBaseAmount                      `xml:"baseAmount"`
	Tax                             AddendaTax                             `xml:"tax"`
	PayableAmount                   AddendaPayableAmount                   `xml:"payableAmount"`
}

// AddendaRequestForPaymentIdentification represents the <requestForPaymentIdentification> element
type AddendaRequestForPaymentIdentification struct {
	EntityType                  string `xml:"entityType"`
	UniqueCreatorIdentification string `xml:"uniqueCreatorIdentification"`
}

// AddendaSpecialInstruction represents the <specialInstruction> element
type AddendaSpecialInstruction struct {
	Code string `xml:"code,attr"`
	Text string `xml:"text"`
}

// AddendaOrderIdentification represents the <orderIdentification> element
type AddendaOrderIdentification struct {
	ReferenceIdentification AddendaReferenceIdentification `xml:"referenceIdentification"`
	ReferenceDate           string                         `xml:"ReferenceDate"`
}

// AddendaReferenceIdentification represents the <referenceIdentification> element
type AddendaReferenceIdentification struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// AddendaAdditionalInformation represents the <AddendaAdditionalInformation> element
type AddendaAdditionalInformation struct {
	ReferenceIdentification AddendaReferenceIdentification `xml:"referenceIdentification"`
}

// AddendaBuyer represents the <buyer> element
type AddendaBuyer struct {
	Gln string `xml:"gln"`
}

// AddendaSeller represents the <seller> element
type AddendaSeller struct {
	Gln                          string                              `xml:"gln"`
	AlternatePartyIdentification AddendaAlternatePartyIdentification `xml:"alternatePartyIdentification"`
}

// AddendaAlternatePartyIdentification represents the <alternatePartyIdentification> element
type AddendaAlternatePartyIdentification struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// AddendaCurrency represents the <currency> element
type AddendaCurrency struct {
	CurrencyISOCode  string `xml:"currencyISOCode,attr"`
	CurrencyFunction string `xml:"currencyFunction"`
	RateOfChange     string `xml:"rateOfChange"`
}

// AddendaLineItem represents the <lineItem> element
type AddendaLineItem struct {
	Number                           string                                  `xml:"number,attr"`
	Type                             string                                  `xml:"type,attr"`
	TradeItemIdentification          AddendaTradeItemIdentification          `xml:"tradeItemIdentification"`
	AlternateTradeItemIdentification AddendaAlternateTradeItemIdentification `xml:"alternateTradeItemIdentification"`
	TradeItemDescriptionInformation  AddendaTradeItemDescriptionInformation  `xml:"tradeItemDescriptionInformation"`
	InvoicedQuantity                 AddendaInvoicedQuantity                 `xml:"invoicedQuantity"`
	GrossPrice                       AddendaGrossPrice                       `xml:"grossPrice"`
	PalletInformation                AddendaPalletInformation                `xml:"palletInformation"`
	TradeItemTaxInformation          AddendaTradeItemTaxInformation          `xml:"tradeItemTaxInformation"`
	TotalLineAmount                  AddendaTotalLineAmount                  `xml:"totalLineAmount"`
}

// AddendaTradeItemIdentification represents the <tradeItemIdentification> element
type AddendaTradeItemIdentification struct {
	Gtin string `xml:"gtin"`
}

// AddendaAlternateTradeItemIdentification represents the <alternateTradeItemIdentification> element
type AddendaAlternateTradeItemIdentification struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// AddendaTradeItemDescriptionInformation represents the <tradeItemDescriptionInformation> element
type AddendaTradeItemDescriptionInformation struct {
	Language string `xml:"language,attr"`
	LongText string `xml:"longText"`
}

// AddendaInvoicedQuantity represents the <invoicedQuantity> element
type AddendaInvoicedQuantity struct {
	UnitOfMeasure string `xml:"unitOfMeasure,attr"`
	Value         string `xml:",chardata"`
}

// AddendaGrossPrice represents the <grossPrice> element
type AddendaGrossPrice struct {
	Amount AddendaAmount `xml:"Amount"`
}

// AddendaAmount represents the <AddendaAmount> element
type AddendaAmount struct {
	Value string `xml:",chardata"`
}

// AddendaPalletInformation represents the <palletInformation> element
type AddendaPalletInformation struct {
	PalletQuantity string             `xml:"palletQuantity"`
	Description    AddendaDescription `xml:"description"`
	Transport      AddendaTransport   `xml:"transport"`
}

// AddendaDescription represents the <description> element
type AddendaDescription struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// AddendaTransport represents the <transport> element
type AddendaTransport struct {
	MethodOfPayment string `xml:"methodOfPayment"`
}

// AddendaTradeItemTaxInformation represents the <tradeItemTaxInformation> element
type AddendaTradeItemTaxInformation struct {
	TaxTypeDescription string                    `xml:"taxTypeDescription"`
	TradeItemTaxAmount AddendaTradeItemTaxAmount `xml:"tradeItemTaxAmount"`
}

// AddendaTradeItemTaxAmount represents the <tradeItemTaxAmount> element
type AddendaTradeItemTaxAmount struct {
	TaxPercentage string `xml:"taxPercentage"`
	TaxAmount     string `xml:"taxAmount"`
}

// AddendaTotalLineAmount represents the <totalLineAmount> element
type AddendaTotalLineAmount struct {
	NetAmount AddendaNetAmount `xml:"netAmount"`
}

// AddendaNetAmount represents the <netAmount> element
type AddendaNetAmount struct {
	Amount AddendaAmount `xml:"Amount"`
}

// AddendaTotalAmount represents the <totalAmount> element
type AddendaTotalAmount struct {
	Amount AddendaAmount `xml:"Amount"`
}

// AddendaTotalAllowanceCharge represents the <AddendaTotalAllowanceCharge> element
type AddendaTotalAllowanceCharge struct {
	AllowanceOrChargeType string        `xml:"allowanceOrChargeType,attr"`
	Amount                AddendaAmount `xml:"Amount"`
}

// AddendaBaseAmount represents the <baseAmount> element
type AddendaBaseAmount struct {
	Amount AddendaAmount `xml:"Amount"`
}

// AddendaTax represents the <tax> element
type AddendaTax struct {
	Type          string `xml:"type,attr"`
	TaxPercentage string `xml:"taxPercentage"`
	TaxAmount     string `xml:"taxAmount"`
}

// AddendaPayableAmount represents the <payableAmount> element
type AddendaPayableAmount struct {
	Amount AddendaAmount `xml:"Amount"`
}

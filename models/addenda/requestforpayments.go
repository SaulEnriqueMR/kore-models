package addenda

// RequestForPayment represents the root element <requestForPayment>
type RequestForPayment struct {
	DeliveryDate                    string                          `xml:"DeliveryDate,attr"`
	ContentVersion                  string                          `xml:"contentVersion,attr"`
	DocumentStatus                  string                          `xml:"documentStatus,attr"`
	DocumentStructureVersion        string                          `xml:"documentStructureVersion,attr"`
	Type                            string                          `xml:"type,attr"`
	RequestForPaymentIdentification RequestForPaymentIdentification `xml:"requestForPaymentIdentification"`
	SpecialInstructions             []SpecialInstruction            `xml:"specialInstruction"`
	OrderIdentification             OrderIdentification             `xml:"orderIdentification"`
	AdditionalInformation           AdditionalInformation           `xml:"AdditionalInformation"`
	Buyer                           Buyer                           `xml:"buyer"`
	Seller                          Seller                          `xml:"seller"`
	Currency                        Currency                        `xml:"currency"`
	LineItem                        LineItem                        `xml:"lineItem"`
	TotalAmount                     TotalAmount                     `xml:"totalAmount"`
	TotalAllowanceCharge            TotalAllowanceCharge            `xml:"TotalAllowanceCharge"`
	BaseAmount                      BaseAmount                      `xml:"baseAmount"`
	Tax                             Tax                             `xml:"tax"`
	PayableAmount                   PayableAmount                   `xml:"payableAmount"`
}

// RequestForPaymentIdentification represents the <requestForPaymentIdentification> element
type RequestForPaymentIdentification struct {
	EntityType                  string `xml:"entityType"`
	UniqueCreatorIdentification string `xml:"uniqueCreatorIdentification"`
}

// SpecialInstruction represents the <specialInstruction> element
type SpecialInstruction struct {
	Code string `xml:"code,attr"`
	Text string `xml:"text"`
}

// OrderIdentification represents the <orderIdentification> element
type OrderIdentification struct {
	ReferenceIdentification ReferenceIdentification `xml:"referenceIdentification"`
	ReferenceDate           string                  `xml:"ReferenceDate"`
}

// ReferenceIdentification represents the <referenceIdentification> element
type ReferenceIdentification struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// AdditionalInformation represents the <AdditionalInformation> element
type AdditionalInformation struct {
	ReferenceIdentification ReferenceIdentification `xml:"referenceIdentification"`
}

// Buyer represents the <buyer> element
type Buyer struct {
	Gln string `xml:"gln"`
}

// Seller represents the <seller> element
type Seller struct {
	Gln                          string                       `xml:"gln"`
	AlternatePartyIdentification AlternatePartyIdentification `xml:"alternatePartyIdentification"`
}

// AlternatePartyIdentification represents the <alternatePartyIdentification> element
type AlternatePartyIdentification struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// Currency represents the <currency> element
type Currency struct {
	CurrencyISOCode  string `xml:"currencyISOCode,attr"`
	CurrencyFunction string `xml:"currencyFunction"`
	RateOfChange     string `xml:"rateOfChange"`
}

// LineItem represents the <lineItem> element
type LineItem struct {
	Number                           string                           `xml:"number,attr"`
	Type                             string                           `xml:"type,attr"`
	TradeItemIdentification          TradeItemIdentification          `xml:"tradeItemIdentification"`
	AlternateTradeItemIdentification AlternateTradeItemIdentification `xml:"alternateTradeItemIdentification"`
	TradeItemDescriptionInformation  TradeItemDescriptionInformation  `xml:"tradeItemDescriptionInformation"`
	InvoicedQuantity                 InvoicedQuantity                 `xml:"invoicedQuantity"`
	GrossPrice                       GrossPrice                       `xml:"grossPrice"`
	PalletInformation                PalletInformation                `xml:"palletInformation"`
	TradeItemTaxInformation          TradeItemTaxInformation          `xml:"tradeItemTaxInformation"`
	TotalLineAmount                  TotalLineAmount                  `xml:"totalLineAmount"`
}

// TradeItemIdentification represents the <tradeItemIdentification> element
type TradeItemIdentification struct {
	Gtin string `xml:"gtin"`
}

// AlternateTradeItemIdentification represents the <alternateTradeItemIdentification> element
type AlternateTradeItemIdentification struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// TradeItemDescriptionInformation represents the <tradeItemDescriptionInformation> element
type TradeItemDescriptionInformation struct {
	Language string `xml:"language,attr"`
	LongText string `xml:"longText"`
}

// InvoicedQuantity represents the <invoicedQuantity> element
type InvoicedQuantity struct {
	UnitOfMeasure string `xml:"unitOfMeasure,attr"`
	Value         string `xml:",chardata"`
}

// GrossPrice represents the <grossPrice> element
type GrossPrice struct {
	Amount Amount `xml:"Amount"`
}

// Amount represents the <Amount> element
type Amount struct {
	Value string `xml:",chardata"`
}

// PalletInformation represents the <palletInformation> element
type PalletInformation struct {
	PalletQuantity string      `xml:"palletQuantity"`
	Description    Description `xml:"description"`
	Transport      Transport   `xml:"transport"`
}

// Description represents the <description> element
type Description struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

// Transport represents the <transport> element
type Transport struct {
	MethodOfPayment string `xml:"methodOfPayment"`
}

// TradeItemTaxInformation represents the <tradeItemTaxInformation> element
type TradeItemTaxInformation struct {
	TaxTypeDescription string             `xml:"taxTypeDescription"`
	TradeItemTaxAmount TradeItemTaxAmount `xml:"tradeItemTaxAmount"`
}

// TradeItemTaxAmount represents the <tradeItemTaxAmount> element
type TradeItemTaxAmount struct {
	TaxPercentage string `xml:"taxPercentage"`
	TaxAmount     string `xml:"taxAmount"`
}

// TotalLineAmount represents the <totalLineAmount> element
type TotalLineAmount struct {
	NetAmount NetAmount `xml:"netAmount"`
}

// NetAmount represents the <netAmount> element
type NetAmount struct {
	Amount Amount `xml:"Amount"`
}

// TotalAmount represents the <totalAmount> element
type TotalAmount struct {
	Amount Amount `xml:"Amount"`
}

// TotalAllowanceCharge represents the <TotalAllowanceCharge> element
type TotalAllowanceCharge struct {
	AllowanceOrChargeType string `xml:"allowanceOrChargeType,attr"`
	Amount                Amount `xml:"Amount"`
}

// BaseAmount represents the <baseAmount> element
type BaseAmount struct {
	Amount Amount `xml:"Amount"`
}

// Tax represents the <tax> element
type Tax struct {
	Type          string `xml:"type,attr"`
	TaxPercentage string `xml:"taxPercentage"`
	TaxAmount     string `xml:"taxAmount"`
}

// PayableAmount represents the <payableAmount> element
type PayableAmount struct {
	Amount Amount `xml:"Amount"`
}

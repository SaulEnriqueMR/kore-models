package detallista

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Detallista10 struct {
	Type                            *string                         `xml:"type,attr" bson:"Type,omitempty"`
	ContentVersion                  *string                         `xml:"contentVersion,attr" bson:"ContentVersion,omitempty"`
	DocumentStructureVersion        string                          `xml:"documentStructureVersion,attr" bson:"DocumentStructureVersion"`
	DocumentStatus                  string                          `xml:"documentStatus,attr" bson:"DocumentStatus"`
	RequestForPaymentIdentification RequestForPaymentIdentification `xml:"requestForPaymentIdentification" bson:"RequestForPaymentIdentification"`
	SpecialInstruction              *[]SpecialInstruction           `xml:"specialInstruction" bson:"SpecialInstruction,omitempty"`
	OrderIdentification             OrderIdentification             `xml:"orderIdentification" bson:"OrderIdentification"`
	AdditionalInformation           *AdditionalInformation          `xml:"AdditionalInformation" bson:"AdditionalInformation,omitempty"`
	DeliveryNote                    *DeliveryNote                   `xml:"DeliveryNote" bson:"DeliveryNote,omitempty"`
	Buyer                           Buyer                           `xml:"buyer" bson:"Buyer"`
	Seller                          *Seller                         `xml:"seller" bson:"Seller,omitempty"`
	ShipTo                          *ShipTo                         `xml:"shipTo" bson:"ShipTo,omitempty"`
	InvoiceCreator                  *InvoiceCreator                 `xml:"InvoiceCreator" bson:"InvoiceCreator,omitempty"`
	Customs                         *[]Customs                      `xml:"Customs" bson:"Customs,omitempty"`
	Currency                        *[]Currency                     `xml:"currency" bson:"Currency,omitempty"`
	PaymentTerms                    *PaymentTerms                   `xml:"paymentTerms" bson:"PaymentTerms,omitempty"`
	ShipmentDetail                  *ShipmentDetail                 `xml:"shipmentDetail" bson:"ShipmentDetail,omitempty"`
	AllowanceCharge                 *[]AllowanceCharge              `xml:"allowanceCharge" bson:"AllowanceCharge,omitempty"`
	LineItem                        *[]LineItem                     `xml:"lineItem" bson:"LineItem,omitempty"`
	TotalAmount                     *TotalAmount                    `xml:"totalAmount" bson:"TotalAmount,omitempty"`
	TotalAllowanceCharge            *[]TotalAllowanceCharge         `xml:"TotalAllowanceCharge" bson:"TotalAllowanceCharge,omitempty"`
}

type RequestForPaymentIdentification struct {
	EntityType string `xml:"entityType" bson:"EntityType"`
}

type SpecialInstruction struct {
	Code string   `xml:"code,attr" bson:"Code"`
	Text []string `xml:"text" bson:"Text"`
}
type ReferenceIdentification struct {
	Value string `xml:",chardata" bson:"Value"`
	Type  string `xml:"type,attr" bson:"Type"`
}

type OrderIdentification struct {
	ReferenceIdentification []ReferenceIdentification `xml:"referenceIdentification" bson:"ReferenceIdentification"`
	ReferenceDate           *ReferenceDate            `xml:"ReferenceDate"`
}

type ReferenceDate struct {
	ReferenceDateString string    `xml:",chardata"`
	Date                time.Time `bson:"ReferenceDate"`
}

type AdditionalInformation struct {
	ReferenceIdentification []ReferenceIdentification `xml:"referenceIdentification" bson:"ReferenceIdentification"`
}

type DeliveryNote struct {
	ReferenceIdentification []string       `xml:"referenceIdentification" bson:"ReferenceIdentification"`
	ReferenceDate           *ReferenceDate `xml:"ReferenceDate"`
}

type Buyer struct {
	Gln  string `xml:"gln" bson:"Gln"`
	Text string `xml:"contactInformation>personOrDepartmentName>text" bson:"Text"`
}

type Seller struct {
	Gln                          string                       `xml:"gln" bson:"Gln"`
	AlternatePartyIdentification AlternatePartyIdentification `xml:"alternatePartyIdentification" bson:"AlternatePartyIdentification"`
}

type AlternatePartyIdentification struct {
	Value string `xml:",chardata" bson:"Value"`
	Type  string `xml:"type,attr" bson:"Type"`
}

type ShipTo struct {
	Gln            *string              `xml:"gln" bson:"Gln,omitempty"`
	NameAndAddress *NameAndAddressArray `xml:"nameAndAddress" bson:"NameAndAddress,omitempty"`
}

type NameAndAddressArray struct {
	Name             *[]string `xml:"name" bson:"Name,omitempty"`
	StreetAddressOne *[]string `xml:"streetAddressOne" bson:"StreetAddressOne,omitempty"`
	City             *[]string `xml:"city" bson:"City,omitempty"`
	PostalCode       *[]string `xml:"postalCode" bson:"PostalCode,omitempty"`
}

type InvoiceCreator struct {
	Gln                          *string                       `xml:"gln" bson:"Gln,omitempty"`
	AlternatePartyIdentification *AlternatePartyIdentification `xml:"alternatePartyIdentification" bson:"AlternatePartyIdentification,omitempty"`
	NameAndAddress               *NameAndAddress               `xml:"nameAndAddress" bson:"NameAndAddress,omitempty"`
}

type NameAndAddress struct {
	Name             *string `xml:"name" bson:"Name,omitempty"`
	StreetAddressOne *string `xml:"streetAddressOne" bson:"StreetAddressOne,omitempty"`
	City             *string `xml:"city" bson:"City,omitempty"`
	PostalCode       *string `xml:"postalCode" bson:"PostalCode,omitempty"`
}

type Customs struct {
	Gln string `xml:"gln" bson:"Gln"`
}

type Currency struct {
	CurrencyISOCode  string    `xml:"currencyISOCode,attr" bson:"CurrencyISOCode"`
	CurrencyFunction *[]string `xml:"currencyFunction" bson:"CurrencyFunction,omitempty"`
	RateOfChange     *float64  `xml:"rateOfChange" bson:"RateOfChange,omitempty"`
}

type PaymentTerms struct {
	PaymentTermsEvent        *string          `xml:"paymentTermsEvent,attr" bson:"PaymentTermsEvent,omitempty"`
	PaymentTermsRelationTime *string          `xml:"PaymentTermsRelationTime,attr" bson:"PaymentTermsRelationTime,omitempty"`
	NetPayment               *NetPayment      `xml:"netPayment" bson:"NetPayment,omitempty"`
	DiscountPayment          *DiscountPayment `xml:"discountPayment" bson:"DiscountPayment,omitempty"`
}

type NetPayment struct {
	NetPaymentTermsType *string        `xml:"netPaymentTermsType,attr" bson:"NetPaymentTermsType,omitempty"`
	TimePeriodDue       *TimePeriodDue `xml:"paymentTimePeriod>timePeriodDue" bson:"PaymentTimePeriod,omitempty"`
}

type TimePeriodDue struct {
	TimePeriod string `xml:"timePeriod,attr" bson:"TimePeriod"`
	Value      string `xml:"value" bson:"Value"`
}

type DiscountPayment struct {
	DiscountType string `xml:"discountType,attr" bson:"DiscountType"`
	Percentage   string `xml:"percentage" bson:"Percentage"`
}

type ShipmentDetail struct {
	XMLName xml.Name `xml:"shipmentDetail"`
}

type AllowanceCharge struct {
	AllowanceChargeType        string                      `xml:"allowanceChargeType,attr" bson:"AllowanceChargeType"`
	SettlementType             string                      `xml:"settlementType,attr" bson:"SettlementType"`
	SequenceNumber             *string                     `xml:"sequenceNumber,attr" bson:"SequenceNumber,omitempty"`
	SpecialServicesType        *string                     `xml:"specialServicesType" bson:"SpecialServicesType"`
	MonetaryAmountOrPercentage *MonetaryAmountOrPercentage `xml:"monetaryAmountOrPercentage>rate" bson:"MonetaryAmountOrPercentage"`
}

type MonetaryAmountOrPercentage struct {
	Base       string   `xml:"base,attr" bson:"Base"`
	Percentage *float64 `xml:"percentage" bson:"Percentage,omitempty"`
}

type LineItem struct {
	Type                             string                              `xml:"" bson:""`
	Number                           int32                               `xml:"" bson:""`
	TradeItemIdentification          *TradeItemIdentification            `xml:"tradeItemIdentification" bson:"TradeItemIdentification,omitempty"`
	AlternateTradeItemIdentification *[]AlternateTradeItemIdentification `xml:"alternateTradeItemIdentification" bson:"AlternateTradeItemIdentification,omitempty"`
	TradeItemDescriptionInformation  *TradeItemDescriptionInformation    `xml:"tradeItemDescriptionInformation" bson:"TradeItemDescriptionInformation,omitempty"`
	InvoicedQuantity                 InvoicedQuantity                    `xml:"invoicedQuantity" bson:"InvoicedQuantity,omitempty"`
	AditionalQuantity                *[]AditionalQuantity                `xml:"aditionalQuantity" bson:"AditionalQuantity,omitempty"`
	GrossPrice                       *GrossPrice                         `xml:"grossPrice" bson:"GrossPrice,omitempty"`
	NetPrice                         *NetPrice                           `xml:"netPrice" bson:"NetPrice,omitempty"`
	AdditionalInformation            *AdditionalInformationLine          `xml:"AdditionalInformation" bson:"AdditionalInformation,omitempty"`
	Customs                          *[]CustomsLineItem                  `xml:"Customs" bson:"Customs,omitempty"`
	LogisticUnits                    *LogisticUnits                      `xml:"LogisticUnits" bson:"LogisticUnits,omitempty"`
	PalletInformation                *PalletInformation                  `xml:"palletInformation" bson:"PalletInformation,omitempty"`
	ExtendedAttributes               *ExtendedAttributes                 `xml:"extendedAttributes" bson:"ExtendedAttributes,omitempty"`
	AllowanceCharge                  *[]AllowanceChargeLineItem          `xml:"allowanceCharge" bson:"AllowanceCharge,omitempty"`
	TradeItemTaxInformation          *[]TradeItemTaxInformation          `xml:"tradeItemTaxInformation" bson:"TradeItemTaxInformation,omitempty"`
	TotalLineAmount                  TotalLineAmount                     `xml:"totalLineAmount" bson:"TotalLineAmount"`
}

type TradeItemIdentification struct {
	Gtin string `xml:"gtin" bson:"Gtin"`
}

type AlternateTradeItemIdentification struct {
	Value *string `xml:",chardata" bson:"Value,omitempty"`
	Type  *string `xml:"type,attr" bson:"type,omitempty"`
}

type TradeItemDescriptionInformation struct {
	Language *string `xml:"language,attr" bson:"Language,omitempty"`
	LongText string  `xml:"longText" bson:"LongText"`
}

type InvoicedQuantity struct {
	UnitOfMeasure string `xml:"unitOfMeasure,attr" bson:"UnitOfMeasure"`
	Value         int    `xml:",chardata" bson:"Value"`
}

type AditionalQuantity struct {
	QuantityType string  `xml:"QuantityType,attr" bson:"QuantityType"`
	Value        float64 `xml:",chardata" bson:"Value"`
}

type GrossPrice struct {
	Amount float64 `xml:"Amount" bson:"Amount"`
}

type NetPrice struct {
	Amount float64 `xml:"Amount" bson:"Amount"`
}

type AdditionalInformationLine struct {
	ReferenceIdentification *ReferenceIdentification `xml:"referenceIdentification" bson:"ReferenceIdentification,omitempty"`
}

type CustomsLineItem struct {
	Gln                          *string                      `xml:"gln" bson:"Gln,omitempty"`
	AlternatePartyIdentification AlternatePartyIdentification `xml:"alternatePartyIdentification" bson:"AlternatePartyIdentification"`
	ReferenceDate                ReferenceDate                `xml:"ReferenceDate"`
	NameAndAddress               NameAndAddressCustoms        `xml:"nameAndAddress" bson:"NameAndAddress"`
}

type NameAndAddressCustoms struct {
	Name string `xml:"name" bson:"Name"`
}

type LogisticUnits struct {
	SerialShippingContainerCode SerialShippingContainerCode `xml:"serialShippingContainerCode" bson:"SerialShippingContainerCode"`
}

type SerialShippingContainerCode struct {
	Value string `xml:",chardata" bson:"Value"`
	Type  string `xml:"type,attr" bson:"Type"`
}

type PalletInformation struct {
	PalletQuantity string      `xml:"palletQuantity" bson:"PalletQuantity"`
	Description    Description `xml:"description" bson:"Description"`
	Transport      string      `xml:"transport>methodOfPayment" bson:"Transport"`
}

type Description struct {
	Value string `xml:",chardata" bson:"Value"`
	Type  string `xml:"type,attr" bson:"type"`
}

type Transport struct {
	MethodPayment MethodPayment `xml:"" bson:""`
}

type MethodPayment struct {
	Value string `xml:",chardata" bson:"Value"`
}

type ExtendedAttributes struct {
	LotNumber []LotNumber `xml:"lotNumber" bson:"LotNumber"`
}

type LotNumber struct {
	ProductionDateString string    `xml:"productionDate,attr"`
	ProductionDate       time.Time `bson:"ProductionDate"`
	Value                string    `xml:",chardata" bson:"Value"`
}

type AllowanceChargeLineItem struct {
	AllowanceChargeType        string                              `xml:"allowanceChargeType,attr" bson:"AllowanceChargeType"`
	SettlementType             string                              `xml:"settlementType,attr" bson:"SettlementType"`
	SequenceNumber             *string                             `xml:"sequenceNumber,attr" bson:"SequenceNumber,omitempty"`
	SpecialServicesType        *string                             `xml:"specialServicesType" bson:"SpecialServicesType"`
	MonetaryAmountOrPercentage *MonetaryAmountOrPercentageLineItem `xml:"monetaryAmountOrPercentage" bson:"MonetaryAmountOrPercentage"`
}
type MonetaryAmountOrPercentageLineItem struct {
	PercentagePerUnit string `xml:"percentagePerUnit" bson:"PercentagePerUnit"`
	RatePerUnit       string `xml:"ratePerUnit>amountPerUnit" bson:"RatePerUnit"`
}

type TradeItemTaxInformation struct {
	TaxTypeDescription string              `xml:"taxTypeDescription" bson:"TaxTypeDescription"`
	ReferenceNumber    *string             `xml:"referenceNumber" bson:"ReferenceNumber,omitempty"`
	TaxCategory        *string             `xml:"taxCategory" bson:"TaxCategory,omitempty"`
	TradeItemTaxAmount *TradeItemTaxAmount `xml:"tradeItemTaxAmount" bson:"TradeItemTaxAmount,omitempty"`
}

type TradeItemTaxAmount struct {
	TaxPercentage float64 `xml:"taxPercentage" bson:"TaxPercentage"`
	TaxAmount     float64 `xml:"taxAmount" bson:"TaxAmount"`
}

type TotalLineAmount struct {
	GrossAmount float64  `xml:"grossAmount>Amount" bson:"GrossAmount"`
	NetAmount   *float64 `xml:"netAmount>Amount" bson:"NetAmount,omitempty"`
}

type TotalAmount struct {
	Amount float64 `xml:"Amount" bson:"Amount"`
}

type TotalAllowanceCharge struct {
	AllowanceOrChargeType string   `xml:"allowanceOrChargeType,attr" bson:"AllowanceOrChargeType"`
	SpecialServicesType   *string  `xml:"specialServicesType" bson:"SpecialServicesType"`
	Amount                *float64 `xml:"Amount" bson:"Amount,omitempty"`
}

func (rd *ReferenceDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias ReferenceDate
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*rd = ReferenceDate(aux)

	fecha, err := helpers.ParseDatetime(aux.ReferenceDateString)
	if err != nil {
		return err
	}
	rd.Date = fecha

	return nil
}

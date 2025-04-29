package detallista

import (
	"encoding/xml"
	"time"

	"github.com/SaulEnriqueMR/kore-models/models/helpers"
)

type Detallista10 struct {
	Type                            *string                         `xml:"type,attr" bson:"Type,omitempty" json:"Type,omitempty"`
	ContentVersion                  *string                         `xml:"contentVersion,attr" bson:"ContentVersion,omitempty" json:"ContentVersion,omitempty"`
	DocumentStructureVersion        string                          `xml:"documentStructureVersion,attr" bson:"DocumentStructureVersion" json:"DocumentStructureVersion"`
	DocumentStatus                  string                          `xml:"documentStatus,attr" bson:"DocumentStatus" json:"DocumentStatus"`
	RequestForPaymentIdentification RequestForPaymentIdentification `xml:"requestForPaymentIdentification" bson:"RequestForPaymentIdentification" json:"RequestForPaymentIdentification"`
	SpecialInstruction              *[]SpecialInstruction           `xml:"specialInstruction" bson:"SpecialInstruction,omitempty" json:"SpecialInstruction,omitempty"`
	OrderIdentification             OrderIdentification             `xml:"orderIdentification" bson:"OrderIdentification" json:"OrderIdentification"`
	AdditionalInformation           *AdditionalInformation          `xml:"AdditionalInformation" bson:"AdditionalInformation,omitempty" json:"AdditionalInformation,omitempty"`
	DeliveryNote                    *DeliveryNote                   `xml:"DeliveryNote" bson:"DeliveryNote,omitempty" json:"DeliveryNote,omitempty"`
	Buyer                           Buyer                           `xml:"buyer" bson:"Buyer" json:"Buyer"`
	Seller                          *Seller                         `xml:"seller" bson:"Seller,omitempty" json:"Seller,omitempty"`
	ShipTo                          *ShipTo                         `xml:"shipTo" bson:"ShipTo,omitempty" json:"ShipTo,omitempty"`
	InvoiceCreator                  *InvoiceCreator                 `xml:"InvoiceCreator" bson:"InvoiceCreator,omitempty" json:"InvoiceCreator,omitempty"`
	Customs                         *[]Customs                      `xml:"Customs" bson:"Customs,omitempty" json:"Customs,omitempty"`
	Currency                        *[]Currency                     `xml:"currency" bson:"Currency,omitempty" json:"Currency,omitempty"`
	PaymentTerms                    *PaymentTerms                   `xml:"paymentTerms" bson:"PaymentTerms,omitempty" json:"PaymentTerms,omitempty"`
	ShipmentDetail                  *ShipmentDetail                 `xml:"shipmentDetail" bson:"ShipmentDetail,omitempty" json:"ShipmentDetail,omitempty"`
	AllowanceCharge                 *[]AllowanceCharge              `xml:"allowanceCharge" bson:"AllowanceCharge,omitempty" json:"AllowanceCharge,omitempty"`
	LineItem                        *[]LineItem                     `xml:"lineItem" bson:"LineItem,omitempty" json:"LineItem,omitempty"`
	TotalAmount                     *TotalAmount                    `xml:"totalAmount" bson:"TotalAmount,omitempty" json:"TotalAmount,omitempty"`
	TotalAllowanceCharge            *[]TotalAllowanceCharge         `xml:"TotalAllowanceCharge" bson:"TotalAllowanceCharge,omitempty" json:"TotalAllowanceCharge,omitempty"`
}

type RequestForPaymentIdentification struct {
	EntityType string `xml:"entityType" bson:"EntityType" json:"EntityType"`
}

type SpecialInstruction struct {
	Code string   `xml:"code,attr" bson:"Code" json:"Code"`
	Text []string `xml:"text" bson:"Text" json:"Text"`
}
type ReferenceIdentification struct {
	Value string `xml:",chardata" bson:"Value" json:"Value"`
	Type  string `xml:"type,attr" bson:"Type" json:"Type"`
}

type OrderIdentification struct {
	ReferenceIdentification []ReferenceIdentification `xml:"referenceIdentification" bson:"ReferenceIdentification" json:"ReferenceIdentification"`
	ReferenceDateString     *string                   `xml:"ReferenceDate"`
	Date                    *time.Time                `bson:"Date" json:"Date"`
}

func (oi *OrderIdentification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias OrderIdentification
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*oi = OrderIdentification(aux)
	if aux.ReferenceDateString != nil && *aux.ReferenceDateString != "" {
		fecha, err := helpers.ParseDatetime(*aux.ReferenceDateString)
		if err != nil {
			return err
		}
		oi.Date = &fecha
	}

	return nil
}

type AdditionalInformation struct {
	ReferenceIdentification []ReferenceIdentification `xml:"referenceIdentification" bson:"ReferenceIdentification" json:"ReferenceIdentification"`
}

type DeliveryNote struct {
	ReferenceIdentification []string   `xml:"referenceIdentification" bson:"ReferenceIdentification" json:"ReferenceIdentification"`
	ReferenceDateString     *string    `xml:"ReferenceDate"`
	Date                    *time.Time `bson:"Date" json:"Date"`
}

func (dn *DeliveryNote) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias DeliveryNote
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*dn = DeliveryNote(aux)
	if aux.ReferenceDateString != nil && *aux.ReferenceDateString != "" {
		fecha, err := helpers.ParseDatetime(*aux.ReferenceDateString)
		if err != nil {
			return err
		}
		dn.Date = &fecha
	}

	return nil
}

type Buyer struct {
	Gln  string `xml:"gln" bson:"Gln" json:"Gln"`
	Text string `xml:"contactInformation>personOrDepartmentName>text" bson:"Text" json:"Text"`
}

type Seller struct {
	Gln                          string                       `xml:"gln" bson:"Gln" json:"Gln"`
	AlternatePartyIdentification AlternatePartyIdentification `xml:"alternatePartyIdentification" bson:"AlternatePartyIdentification" json:"AlternatePartyIdentification"`
}

type AlternatePartyIdentification struct {
	Value string `xml:",chardata" bson:"Value" json:"Value"`
	Type  string `xml:"type,attr" bson:"Type" json:"Type"`
}

type ShipTo struct {
	Gln            *string              `xml:"gln" bson:"Gln,omitempty" json:"Gln,omitempty"`
	NameAndAddress *NameAndAddressArray `xml:"nameAndAddress" bson:"NameAndAddress,omitempty" json:"NameAndAddress,omitempty"`
}

type NameAndAddressArray struct {
	Name             *[]string `xml:"name" bson:"Name,omitempty" json:"Name,omitempty"`
	StreetAddressOne *[]string `xml:"streetAddressOne" bson:"StreetAddressOne,omitempty" json:"StreetAddressOne,omitempty"`
	City             *[]string `xml:"city" bson:"City,omitempty" json:"City,omitempty"`
	PostalCode       *[]string `xml:"postalCode" bson:"PostalCode,omitempty" json:"PostalCode,omitempty"`
}

type InvoiceCreator struct {
	Gln                          *string                       `xml:"gln" bson:"Gln,omitempty" json:"Gln,omitempty"`
	AlternatePartyIdentification *AlternatePartyIdentification `xml:"alternatePartyIdentification" bson:"AlternatePartyIdentification,omitempty" json:"AlternatePartyIdentification,omitempty"`
	NameAndAddress               *NameAndAddress               `xml:"nameAndAddress" bson:"NameAndAddress,omitempty" json:"NameAndAddress,omitempty"`
}

type NameAndAddress struct {
	Name             *string `xml:"name" bson:"Name,omitempty" json:"Name,omitempty"`
	StreetAddressOne *string `xml:"streetAddressOne" bson:"StreetAddressOne,omitempty" json:"StreetAddressOne,omitempty"`
	City             *string `xml:"city" bson:"City,omitempty" json:"City,omitempty"`
	PostalCode       *string `xml:"postalCode" bson:"PostalCode,omitempty" json:"PostalCode,omitempty"`
}

type Customs struct {
	Gln string `xml:"gln" bson:"Gln" json:"Gln"`
}

type Currency struct {
	CurrencyISOCode  string    `xml:"currencyISOCode,attr" bson:"CurrencyISOCode" json:"CurrencyISOCode"`
	CurrencyFunction *[]string `xml:"currencyFunction" bson:"CurrencyFunction,omitempty" json:"CurrencyFunction,omitempty"`
	RateOfChange     *float64  `xml:"rateOfChange" bson:"RateOfChange,omitempty" json:"RateOfChange,omitempty"`
}

type PaymentTerms struct {
	PaymentTermsEvent        *string          `xml:"paymentTermsEvent,attr" bson:"PaymentTermsEvent,omitempty" json:"PaymentTermsEvent,omitempty"`
	PaymentTermsRelationTime *string          `xml:"PaymentTermsRelationTime,attr" bson:"PaymentTermsRelationTime,omitempty" json:"PaymentTermsRelationTime,omitempty"`
	NetPayment               *NetPayment      `xml:"netPayment" bson:"NetPayment,omitempty" json:"NetPayment,omitempty"`
	DiscountPayment          *DiscountPayment `xml:"discountPayment" bson:"DiscountPayment,omitempty" json:"DiscountPayment,omitempty"`
}

type NetPayment struct {
	NetPaymentTermsType *string        `xml:"netPaymentTermsType,attr" bson:"NetPaymentTermsType,omitempty" json:"NetPaymentTermsType,omitempty"`
	TimePeriodDue       *TimePeriodDue `xml:"paymentTimePeriod>timePeriodDue" bson:"PaymentTimePeriod,omitempty" json:"PaymentTimePeriod,omitempty"`
}

type TimePeriodDue struct {
	TimePeriod string `xml:"timePeriod,attr" bson:"TimePeriod" json:"TimePeriod"`
	Value      string `xml:"value" bson:"Value" json:"Value"`
}

type DiscountPayment struct {
	DiscountType string `xml:"discountType,attr" bson:"DiscountType" json:"DiscountType"`
	Percentage   string `xml:"percentage" bson:"Percentage" json:"Percentage"`
}

type ShipmentDetail struct {
}

type AllowanceCharge struct {
	AllowanceChargeType        string                      `xml:"allowanceChargeType,attr" bson:"AllowanceChargeType" json:"AllowanceChargeType"`
	SettlementType             string                      `xml:"settlementType,attr" bson:"SettlementType" json:"SettlementType"`
	SequenceNumber             *string                     `xml:"sequenceNumber,attr" bson:"SequenceNumber,omitempty" json:"SequenceNumber,omitempty"`
	SpecialServicesType        *string                     `xml:"specialServicesType" bson:"SpecialServicesType" json:"SpecialServicesType"`
	MonetaryAmountOrPercentage *MonetaryAmountOrPercentage `xml:"monetaryAmountOrPercentage>rate" bson:"MonetaryAmountOrPercentage" json:"MonetaryAmountOrPercentage"`
}

type MonetaryAmountOrPercentage struct {
	Base       string   `xml:"base,attr" bson:"Base" json:"Base"`
	Percentage *float64 `xml:"percentage" bson:"Percentage,omitempty" json:"Percentage,omitempty"`
}

type LineItem struct {
	Type                             *string                             `xml:"type,attr" bson:"Type" json:"Type"`
	Number                           *int32                              `xml:"number,attr" bson:"Number" json:"Number"`
	TradeItemIdentification          *TradeItemIdentification            `xml:"tradeItemIdentification" bson:"TradeItemIdentification,omitempty" json:"TradeItemIdentification,omitempty"`
	AlternateTradeItemIdentification *[]AlternateTradeItemIdentification `xml:"alternateTradeItemIdentification" bson:"AlternateTradeItemIdentification,omitempty" json:"AlternateTradeItemIdentification,omitempty"`
	TradeItemDescriptionInformation  *TradeItemDescriptionInformation    `xml:"tradeItemDescriptionInformation" bson:"TradeItemDescriptionInformation,omitempty" json:"TradeItemDescriptionInformation,omitempty"`
	InvoicedQuantity                 InvoicedQuantity                    `xml:"invoicedQuantity" bson:"InvoicedQuantity,omitempty" json:"InvoicedQuantity,omitempty"`
	AditionalQuantity                *[]AditionalQuantity                `xml:"aditionalQuantity" bson:"AditionalQuantity,omitempty" json:"AditionalQuantity,omitempty"`
	GrossPrice                       *GrossPrice                         `xml:"grossPrice" bson:"GrossPrice,omitempty" json:"GrossPrice,omitempty"`
	NetPrice                         *NetPrice                           `xml:"netPrice" bson:"NetPrice,omitempty" json:"NetPrice,omitempty"`
	AdditionalInformation            *AdditionalInformationLine          `xml:"AdditionalInformation" bson:"AdditionalInformation,omitempty" json:"AdditionalInformation,omitempty"`
	Customs                          *[]CustomsLineItem                  `xml:"Customs" bson:"Customs,omitempty" json:"Customs,omitempty"`
	LogisticUnits                    *LogisticUnits                      `xml:"LogisticUnits" bson:"LogisticUnits,omitempty" json:"LogisticUnits,omitempty"`
	PalletInformation                *PalletInformation                  `xml:"palletInformation" bson:"PalletInformation,omitempty" json:"PalletInformation,omitempty"`
	ExtendedAttributes               *ExtendedAttributes                 `xml:"extendedAttributes" bson:"ExtendedAttributes,omitempty" json:"ExtendedAttributes,omitempty"`
	AllowanceCharge                  *[]AllowanceChargeLineItem          `xml:"allowanceCharge" bson:"AllowanceCharge,omitempty" json:"AllowanceCharge,omitempty"`
	TradeItemTaxInformation          *[]TradeItemTaxInformation          `xml:"tradeItemTaxInformation" bson:"TradeItemTaxInformation,omitempty" json:"TradeItemTaxInformation,omitempty"`
	TotalLineAmount                  TotalLineAmount                     `xml:"totalLineAmount" bson:"TotalLineAmount" json:"TotalLineAmount"`
}

type TradeItemIdentification struct {
	Gtin string `xml:"gtin" bson:"Gtin" json:"Gtin"`
}

type AlternateTradeItemIdentification struct {
	Value *string `xml:",chardata" bson:"Value,omitempty" json:"Value,omitempty"`
	Type  *string `xml:"type,attr" bson:"type,omitempty" json:"type,omitempty"`
}

type TradeItemDescriptionInformation struct {
	Language *string `xml:"language,attr" bson:"Language,omitempty" json:"Language,omitempty"`
	LongText string  `xml:"longText" bson:"LongText" json:"LongText"`
}

type InvoicedQuantity struct {
	UnitOfMeasure string  `xml:"unitOfMeasure,attr" bson:"UnitOfMeasure" json:"UnitOfMeasure"`
	Value         float64 `xml:",chardata" bson:"Value" json:"Value"`
}

type AditionalQuantity struct {
	QuantityType string  `xml:"QuantityType,attr" bson:"QuantityType" json:"QuantityType"`
	Value        float64 `xml:",chardata" bson:"Value" json:"Value"`
}

type GrossPrice struct {
	Amount float64 `xml:"Amount" bson:"Amount" json:"Amount"`
}

type NetPrice struct {
	Amount float64 `xml:"Amount" bson:"Amount" json:"Amount"`
}

type AdditionalInformationLine struct {
	ReferenceIdentification *ReferenceIdentification `xml:"referenceIdentification" bson:"ReferenceIdentification,omitempty" json:"ReferenceIdentification,omitempty"`
}

type CustomsLineItem struct {
	Gln                          *string                      `xml:"gln" bson:"Gln,omitempty" json:"Gln,omitempty"`
	AlternatePartyIdentification AlternatePartyIdentification `xml:"alternatePartyIdentification" bson:"AlternatePartyIdentification" json:"AlternatePartyIdentification"`
	ReferenceDateString          *string                      `xml:"ReferenceDate"`
	Date                         *time.Time                   `bson:"Date" json:"Date"`
	NameAndAddress               NameAndAddressCustoms        `xml:"nameAndAddress" bson:"NameAndAddress" json:"NameAndAddress"`
}

func (cli *CustomsLineItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias CustomsLineItem
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*cli = CustomsLineItem(aux)
	if aux.ReferenceDateString != nil && *aux.ReferenceDateString != "" {
		fecha, err := helpers.ParseDatetime(*aux.ReferenceDateString)
		if err != nil {
			return err
		}
		cli.Date = &fecha
	}

	return nil
}

type NameAndAddressCustoms struct {
	Name string `xml:"name" bson:"Name" json:"Name"`
}

type LogisticUnits struct {
	SerialShippingContainerCode SerialShippingContainerCode `xml:"serialShippingContainerCode" bson:"SerialShippingContainerCode" json:"SerialShippingContainerCode"`
}

type SerialShippingContainerCode struct {
	Value string `xml:",chardata" bson:"Value" json:"Value"`
	Type  string `xml:"type,attr" bson:"Type" json:"Type"`
}

type PalletInformation struct {
	PalletQuantity string      `xml:"palletQuantity" bson:"PalletQuantity" json:"PalletQuantity"`
	Description    Description `xml:"description" bson:"Description" json:"Description"`
	MethodPayment  string      `xml:"transport>methodOfPayment" bson:"MethodPayment" json:"MethodPayment"`
}

type Description struct {
	Value string `xml:",chardata" bson:"Value" json:"Value"`
	Type  string `xml:"type,attr" bson:"type" json:"type"`
}

type ExtendedAttributes struct {
	LotNumber []LotNumber `xml:"lotNumber" bson:"LotNumber" json:"LotNumber"`
}

type LotNumber struct {
	ProductionDateString string    `xml:"productionDate,attr" bson:"ProductionDateString" json:"ProductionDateString"`
	ProductionDate       time.Time `bson:"ProductionDate" json:"ProductionDate"`
	Value                string    `xml:",chardata" bson:"Value" json:"Value"`
}

func (ln *LotNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Create an alias to avoid recursion
	type Alias LotNumber
	var aux Alias

	// Unmarshal the XML into the alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	*ln = LotNumber(aux)

	fecha, err := helpers.ParseDatetime(aux.ProductionDateString)
	if err != nil {
		return err
	}
	ln.ProductionDate = fecha

	return nil
}

type AllowanceChargeLineItem struct {
	AllowanceChargeType        string                              `xml:"allowanceChargeType,attr" bson:"AllowanceChargeType" json:"AllowanceChargeType"`
	SettlementType             string                              `xml:"settlementType,attr" bson:"SettlementType" json:"SettlementType"`
	SequenceNumber             *string                             `xml:"sequenceNumber,attr" bson:"SequenceNumber,omitempty" json:"SequenceNumber,omitempty"`
	SpecialServicesType        *string                             `xml:"specialServicesType" bson:"SpecialServicesType" json:"SpecialServicesType"`
	MonetaryAmountOrPercentage *MonetaryAmountOrPercentageLineItem `xml:"monetaryAmountOrPercentage" bson:"MonetaryAmountOrPercentage" json:"MonetaryAmountOrPercentage"`
}
type MonetaryAmountOrPercentageLineItem struct {
	PercentagePerUnit string `xml:"percentagePerUnit" bson:"PercentagePerUnit" json:"PercentagePerUnit"`
	RatePerUnit       string `xml:"ratePerUnit>amountPerUnit" bson:"RatePerUnit" json:"RatePerUnit"`
}

type TradeItemTaxInformation struct {
	TaxTypeDescription string              `xml:"taxTypeDescription" bson:"TaxTypeDescription" json:"TaxTypeDescription"`
	ReferenceNumber    *string             `xml:"referenceNumber" bson:"ReferenceNumber,omitempty" json:"ReferenceNumber,omitempty"`
	TaxCategory        *string             `xml:"taxCategory" bson:"TaxCategory,omitempty" json:"TaxCategory,omitempty"`
	TradeItemTaxAmount *TradeItemTaxAmount `xml:"tradeItemTaxAmount" bson:"TradeItemTaxAmount,omitempty" json:"TradeItemTaxAmount,omitempty"`
}

type TradeItemTaxAmount struct {
	TaxPercentage float64 `xml:"taxPercentage" bson:"TaxPercentage" json:"TaxPercentage"`
	TaxAmount     float64 `xml:"taxAmount" bson:"TaxAmount" json:"TaxAmount"`
}

type TotalLineAmount struct {
	GrossAmount float64  `xml:"grossAmount>Amount" bson:"GrossAmount" json:"GrossAmount"`
	NetAmount   *float64 `xml:"netAmount>Amount" bson:"NetAmount,omitempty" json:"NetAmount,omitempty"`
}

type TotalAmount struct {
	Amount float64 `xml:"Amount" bson:"Amount" json:"Amount"`
}

type TotalAllowanceCharge struct {
	AllowanceOrChargeType string   `xml:"allowanceOrChargeType,attr" bson:"AllowanceOrChargeType" json:"AllowanceOrChargeType"`
	SpecialServicesType   *string  `xml:"specialServicesType" bson:"SpecialServicesType" json:"SpecialServicesType"`
	Amount                *float64 `xml:"Amount" bson:"Amount,omitempty" json:"Amount,omitempty"`
}

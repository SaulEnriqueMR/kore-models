package detallista

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
	"testing"

	detallista10 "github.com/SaulEnriqueMR/kore-models/models/comprobante/complementos/detallista"
	testing2 "github.com/SaulEnriqueMR/kore-models/testing"
	"github.com/stretchr/testify/assert"
)

func GetDetallista10ForTest(filename string, t *testing.T) (detallista10.Detallista10, error) {
	data := testing2.GetFileContentForTest(filename, t)
	var parsed detallista10.Detallista10
	errUnmashal := xml.Unmarshal(data, &parsed)
	assert.NoError(t, errUnmashal)
	GenerateJSONFromXML("detallista10.json", parsed)
	return parsed, errUnmashal
}

func GenerateJSONFromXML(namefile string, data detallista10.Detallista10) {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println(err)
	}
	err1 := os.WriteFile(namefile, jsonData, 0644)
	if err1 != nil {
		log.Println(err1)
	}
}

func TestFullDetallistas10(t *testing.T) {
	detallista, _ := GetDetallista10ForTest("./detallista.xml", t)
	InternalTestBaseAttributes(t, detallista)
}
func InternalTestBaseAttributes(t *testing.T, detallista detallista10.Detallista10) {
	assert.Equal(t, "SimpleInvoiceType", *detallista.Type)
	assert.Equal(t, "1.3.1", *detallista.ContentVersion)
	assert.Equal(t, "AMC8.1", detallista.DocumentStructureVersion)
	assert.Equal(t, "ORIGINAL", detallista.DocumentStatus)
	InternalRequestForPaymentIdentification(t, detallista.RequestForPaymentIdentification)
	InternalSpecialInstruction(t, detallista.SpecialInstruction)
	InternalOrderIdentification(t, detallista.OrderIdentification)
	InternalAdditionalInformation(t, detallista.AdditionalInformation)
	InternalDeliveryNote(t, *detallista.DeliveryNote)
	InternalBuyer(t, detallista.Buyer)
	InternalSeller(t, *detallista.Seller)
	InternalShipTo(t, detallista.ShipTo)
	InternalInvoiceCreator(t, *detallista.InvoiceCreator)
	InternalCustoms(t, detallista.Customs)
	InternalCurrency(t, detallista.Currency)
	InternalPaymentTerms(t, detallista.PaymentTerms)
	InternalShipmentDetail(t, detallista.ShipmentDetail)
	InternalAllowanceCharge(t, detallista.AllowanceCharge)
	InternalLineItem(t, detallista.LineItem)
	InternalTotalAmount(t, detallista.TotalAmount)
	InternalTotalAllowanceCharge(t, detallista.TotalAllowanceCharge)
}

func InternalRequestForPaymentIdentification(t *testing.T, request detallista10.RequestForPaymentIdentification) {
	assert.Equal(t, "DEBIT_NOTE", request.EntityType)
}

func InternalSpecialInstruction(t *testing.T, specials *[]detallista10.SpecialInstruction) {
	assert.NotNil(t, specials)
	assert.Equal(t, len(*specials), 1)
	special := (*specials)[0]
	assert.Equal(t, "AAB", special.Code)
	text1 := special.Text[0]
	text2 := special.Text[3]
	assert.Equal(t, "TEXT", text1)
	assert.Equal(t, "HOLA", text2)
}

func InternalOrderIdentification(t *testing.T, order detallista10.OrderIdentification) {
	assert.NotNil(t, order)
	assert.Equal(t, "2016-02-07", *order.ReferenceDateString)
	InternalReferenceIdentification1(t, order.ReferenceIdentification)
}

func InternalReferenceIdentification1(t *testing.T, references []detallista10.ReferenceIdentification) {
	assert.NotNil(t, references)
	assert.Equal(t, len(references), 2)
	reference := (references)[0]
	assert.Equal(t, "ON", reference.Type)
	assert.Equal(t, "81283939", reference.Value)
}

func InternalAdditionalInformation(t *testing.T, additional *detallista10.AdditionalInformation) {
	assert.NotNil(t, additional)
	InternalReferenceIdentification2(t, additional.ReferenceIdentification)
}

func InternalReferenceIdentification2(t *testing.T, references []detallista10.ReferenceIdentification) {
	assert.NotNil(t, references)
	assert.Equal(t, len(references), 2)
	reference := (references)[1]
	assert.Equal(t, "AAE", reference.Type)
	assert.Equal(t, "000001", reference.Value)
}
func InternalDeliveryNote(t *testing.T, delivery detallista10.DeliveryNote) {
	assert.NotNil(t, delivery)
	assert.Equal(t, "2016-04-07", *delivery.ReferenceDateString)
	references := delivery.ReferenceIdentification
	assert.NotNil(t, references)
	assert.Equal(t, len(references), 2)
	reference := (references)[1]
	assert.Equal(t, "0000002", reference)
}

func InternalBuyer(t *testing.T, buyer detallista10.Buyer) {
	assert.NotNil(t, buyer)
	assert.Equal(t, "7504000107903", buyer.Gln)
	assert.Equal(t, "0001", buyer.Text)
}

func InternalSeller(t *testing.T, seller detallista10.Seller) {
	assert.NotNil(t, seller)
	assert.Equal(t, "7244020982223", seller.Gln)
	assert.Equal(t, "IEPS_REFERENCE", seller.AlternatePartyIdentification.Type)
	assert.Equal(t, "12382317", seller.AlternatePartyIdentification.Value)
}

func InternalShipTo(t *testing.T, shipto *detallista10.ShipTo) {
	assert.NotNil(t, shipto)
	assert.Equal(t, "8724098266218", *shipto.Gln)
	nameandaddress := shipto.NameAndAddress
	names := nameandaddress.Name
	name1 := (*names)[2]
	assert.Equal(t, len(*nameandaddress.City), 2)
	assert.Equal(t, len(*nameandaddress.StreetAddressOne), 2)
	assert.Equal(t, len(*nameandaddress.PostalCode), 3)
	assert.Equal(t, "EMPRESA DEMO", name1)
}

func InternalInvoiceCreator(t *testing.T, invoice detallista10.InvoiceCreator) {
	assert.NotNil(t, invoice)
	assert.Equal(t, "5529910871321", *invoice.Gln)
	assert.Equal(t, "IA", invoice.AlternatePartyIdentification.Type)
	assert.Equal(t, "1281282", invoice.AlternatePartyIdentification.Value)
	nameandaddress := invoice.NameAndAddress
	assert.Equal(t, "EMPRESA DEMO 1", *nameandaddress.Name)
	assert.Equal(t, "Calle Niños Heres num 1293", *nameandaddress.StreetAddressOne)
	assert.Equal(t, "Nuevo León", *nameandaddress.City)
	assert.Equal(t, "64590", *nameandaddress.PostalCode)

}

func InternalCustoms(t *testing.T, customs *[]detallista10.Customs) {
	assert.NotNil(t, customs)
	custom1 := (*customs)[0]
	assert.Equal(t, "0129829934772", custom1.Gln)
	custom2 := (*customs)[1]
	assert.Equal(t, "0129829934773", custom2.Gln)
}

func InternalCurrency(t *testing.T, currencys *[]detallista10.Currency) {
	assert.NotNil(t, currencys)
	assert.Equal(t, len(*currencys), 3)
	currency := (*currencys)[1]
	assert.Equal(t, "MXN", currency.CurrencyISOCode)
	assert.Equal(t, len(*currency.CurrencyFunction), 3)
	assert.Equal(t, "PRICE_CURRENCY", (*currency.CurrencyFunction)[1])
	assert.Equal(t, 17.10, *currency.RateOfChange)

}

func InternalPaymentTerms(t *testing.T, payment *detallista10.PaymentTerms) {
	assert.NotNil(t, payment)
	assert.Equal(t, "DATE_OF_INVOICE", *payment.PaymentTermsEvent)
	assert.Equal(t, "REFERENCE_AFTER", *payment.PaymentTermsRelationTime)
	netPayment := payment.NetPayment
	assert.Equal(t, "BASIC_NET", *netPayment.NetPaymentTermsType)
	timePeriodDue := netPayment.TimePeriodDue
	assert.Equal(t, "DAYS", timePeriodDue.TimePeriod)
	assert.Equal(t, "10", timePeriodDue.Value)
	discountPayment := payment.DiscountPayment
	assert.Equal(t, "ALLOWANCE_BY_PAYMENT_ON_TIME", discountPayment.DiscountType)
	assert.Equal(t, "0.00", discountPayment.Percentage)

}

func InternalShipmentDetail(t *testing.T, shipment *detallista10.ShipmentDetail) {
	assert.NotNil(t, shipment)
}

func InternalAllowanceCharge(t *testing.T, allowances *[]detallista10.AllowanceCharge) {
	assert.NotNil(t, allowances)
	assert.Equal(t, len(*allowances), 1)
	allowance := (*allowances)[0]
	assert.Equal(t, "ALLOWANCE_GLOBAL", allowance.AllowanceChargeType)
	assert.Equal(t, "BILL_BACK", allowance.SettlementType)
	assert.Equal(t, "12717262", *allowance.SequenceNumber)
	monetaryAmount := allowance.MonetaryAmountOrPercentage
	assert.Equal(t, "INVOICE_VALUE", monetaryAmount.Base)
	assert.Equal(t, 1.00, *monetaryAmount.Percentage)
}

func InternalLineItem(t *testing.T, lineitems *[]detallista10.LineItem) {
	assert.NotNil(t, lineitems)
	assert.Equal(t, len(*lineitems), 1)
	lineitem := (*lineitems)[0]
	InternalTradeItemIdentification(t, lineitem.TradeItemIdentification)
	InternalAlternateTradeItemIdentification(t, lineitem.AlternateTradeItemIdentification)
	InternalTradeItemDescriptionInformation(t, lineitem.TradeItemDescriptionInformation)
	InternalInvoicedQuantity(t, lineitem.InvoicedQuantity)
	InternalAditionalQuantity(t, lineitem.AditionalQuantity)
	InternalGrossPrice(t, lineitem.GrossPrice)
	InternalNetPrice(t, lineitem.NetPrice)
	InternalAdditionalInformationLine(t, lineitem.AdditionalInformation)
	InternalCustomsLine(t, lineitem.Customs)
	InternalLogisticUnits(t, lineitem.LogisticUnits)
	InternalPalletInformation(t, lineitem.PalletInformation)
	InternalExtendedAttributes(t, lineitem.ExtendedAttributes)
	InternalAllowanceChargeLine(t, lineitem.AllowanceCharge)
	InternalTradeItemTaxInformation(t, lineitem.TradeItemTaxInformation)
	InternalTotalLineAmount(t, lineitem.TotalLineAmount)

}

func InternalTradeItemIdentification(t *testing.T, trade *detallista10.TradeItemIdentification) {
	assert.NotNil(t, trade)
	assert.Equal(t, "FGTSDF", trade.Gtin)
}

func InternalAlternateTradeItemIdentification(t *testing.T, alternates *[]detallista10.AlternateTradeItemIdentification) {
	assert.NotNil(t, alternates)
	assert.Equal(t, len(*alternates), 1)
	alternate := (*alternates)[0]
	assert.Equal(t, "SUPPLIER_ASSIGNED", *alternate.Type)
	assert.Equal(t, "1111", *alternate.Value)
}
func InternalTradeItemDescriptionInformation(t *testing.T, trade *detallista10.TradeItemDescriptionInformation) {
	assert.NotNil(t, trade)
	assert.Equal(t, "ES", *trade.Language)
	assert.Equal(t, "ARTICULO", trade.LongText)
}

func InternalInvoicedQuantity(t *testing.T, invoived detallista10.InvoicedQuantity) {
	assert.NotNil(t, invoived)
	assert.Equal(t, "PZA", invoived.UnitOfMeasure)
	assert.Equal(t, 1, invoived.Value)
}

func InternalAditionalQuantity(t *testing.T, aditionals *[]detallista10.AditionalQuantity) {
	assert.NotNil(t, aditionals)
	assert.Equal(t, len(*aditionals), 1)
	aditional := (*aditionals)[0]
	assert.Equal(t, "NUM_CONSUMER_UNITS", aditional.QuantityType)
	assert.Equal(t, 1.00, aditional.Value)
}

func InternalGrossPrice(t *testing.T, gross *detallista10.GrossPrice) {
	assert.NotNil(t, gross)
	assert.Equal(t, 1.0, gross.Amount)
}

func InternalNetPrice(t *testing.T, net *detallista10.NetPrice) {
	assert.NotNil(t, net)
	assert.Equal(t, 1.0, net.Amount)
}

func InternalAdditionalInformationLine(t *testing.T, additional *detallista10.AdditionalInformationLine) {
	assert.NotNil(t, additional)
	InternalReferenceIdentificationLine(t, additional.ReferenceIdentification)
}

func InternalReferenceIdentificationLine(t *testing.T, reference *detallista10.ReferenceIdentification) {
	assert.Equal(t, "ON", reference.Type)
	assert.Equal(t, "TIPOREFERENCIA", reference.Value)
}

func InternalCustomsLine(t *testing.T, customs *[]detallista10.CustomsLineItem) {
	assert.NotNil(t, customs)
	assert.Equal(t, len(*customs), 1)
	custom := (*customs)[0]
	assert.Equal(t, "1111111111111", *custom.Gln)
	assert.Equal(t, "1239K23092", custom.AlternatePartyIdentification.Value)
	assert.Equal(t, "TN", custom.AlternatePartyIdentification.Type)
	assert.Equal(t, "2016-04-07", *custom.ReferenceDateString)
	assert.Equal(t, "1239K23092", custom.AlternatePartyIdentification.Value)
	assert.Equal(t, "TIJUANA", custom.NameAndAddress.Name)
}

func InternalLogisticUnits(t *testing.T, logistic *detallista10.LogisticUnits) {
	assert.NotNil(t, logistic)
	assert.Equal(t, "SRV", logistic.SerialShippingContainerCode.Type)
	assert.Equal(t, "17717121128", logistic.SerialShippingContainerCode.Value)

}

func InternalPalletInformation(t *testing.T, pallet *detallista10.PalletInformation) {
	assert.NotNil(t, pallet)
	assert.Equal(t, "2", pallet.PalletQuantity)
	assert.Equal(t, "EXCHANGE_PALLETS", pallet.Description.Type)
	assert.Equal(t, "Empaque", pallet.Description.Value)
	assert.Equal(t, "PAID_BY_BUYER", pallet.MethodPayment)
}

func InternalExtendedAttributes(t *testing.T, extended *detallista10.ExtendedAttributes) {
	assert.NotNil(t, extended)
	lots := extended.LotNumber
	assert.Equal(t, len(lots), 1)
	lot := lots[0]
	assert.Equal(t, "2016-04-07", lot.ProductionDateString)
	assert.Equal(t, "1123", lot.Value)
}

func InternalAllowanceChargeLine(t *testing.T, allowance *[]detallista10.AllowanceChargeLineItem) {
	assert.NotNil(t, allowance)
	assert.Equal(t, len(*allowance), 1)
	allowed := (*allowance)[0]
	assert.Equal(t, "ALLOWANCE_GLOBAL", allowed.AllowanceChargeType)
	assert.Equal(t, "OFF_INVOICE", allowed.SettlementType)
	assert.Equal(t, "1", *allowed.SequenceNumber)
	assert.Equal(t, "5%", allowed.MonetaryAmountOrPercentage.PercentagePerUnit)
	assert.Equal(t, "3.00", allowed.MonetaryAmountOrPercentage.RatePerUnit)
}
func InternalTradeItemTaxInformation(t *testing.T, trades *[]detallista10.TradeItemTaxInformation) {
	assert.NotNil(t, trades)
	assert.Equal(t, len(*trades), 1)
	trade := (*trades)[0]
	assert.Equal(t, "GST", trade.TaxTypeDescription)
	assert.Equal(t, "112312", *trade.ReferenceNumber)
	assert.Equal(t, "TRANSFERIDO", *trade.TaxCategory)
	assert.Equal(t, 1.00, trade.TradeItemTaxAmount.TaxAmount)
	assert.Equal(t, 1.00, trade.TradeItemTaxAmount.TaxPercentage)
}

func InternalTotalLineAmount(t *testing.T, total detallista10.TotalLineAmount) {
	assert.NotNil(t, total)
	assert.Equal(t, 1.0, *total.NetAmount)
	assert.Equal(t, 1.0, total.GrossAmount)
}

func InternalTotalAmount(t *testing.T, total *detallista10.TotalAmount) {
	assert.NotNil(t, total)
	assert.Equal(t, 1.0, total.Amount)

}
func InternalTotalAllowanceCharge(t *testing.T, totales *[]detallista10.TotalAllowanceCharge) {
	assert.NotNil(t, totales)
	assert.Equal(t, len(*totales), 1)
	total := (*totales)[0]
	assert.Equal(t, 1.0, *total.Amount)
	assert.Equal(t, "CHARGE", total.AllowanceOrChargeType)
	assert.Equal(t, "ADS", *total.SpecialServicesType)

}

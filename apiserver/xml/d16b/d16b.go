package d16b

import "encoding/xml"

type CrossIndustryInvoice struct {
	XMLName                     xml.Name                    `xml:"CrossIndustryInvoice"`
	SupplyChainTradeTransaction SupplyChainTradeTransaction `xml:"SupplyChainTradeTransaction"`
}

type SupplyChainTradeTransaction struct {
	ApplicableHeaderTradeAgreement  HeaderTradeAgreementType  `xml:"ApplicableHeaderTradeAgreement"`
	ApplicableHeaderTradeSettlement HeaderTradeSettlementType `xml:"ApplicableHeaderTradeSettlement"`
}

type HeaderTradeAgreementType struct {
	SellerTradeParty TradePartyType `xml:"SellerTradeParty"`
	BuyerTradeParty  TradePartyType `xml:"BuyerTradeParty"`
}

type TradePartyType struct {
	Name string `xml:"Name"`
}

type HeaderTradeSettlementType struct {
	SpecifiedTradeSettlementHeaderMonetarySummation TradeSettlementHeaderMonetarySummationType `xml:"SpecifiedTradeSettlementHeaderMonetarySummation"`
}

type TradeSettlementHeaderMonetarySummationType struct {
	LineTotalAmount AmountType `xml:"LineTotalAmount"`
}

type AmountType struct {
	Value string `xml:",innerxml"`
}

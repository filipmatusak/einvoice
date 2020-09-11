package ubl21

import "encoding/xml"

type Invoice struct {
	XMLName            xml.Name                `xml:"Invoice"`
	ID                 string                  `xml:"ID"`
	Supplier           AccountingSupplierParty `xml:"AccountingSupplierParty"`
	Customer           AccountingCustomerParty `xml:"AccountingCustomerParty"`
	LegalMonetaryTotal MonetaryTotalType       `xml:"LegalMonetaryTotal"`
}

type AccountingSupplierParty struct {
	Party Party `xml:"Party"`
}

type AccountingCustomerParty struct {
	Party Party `xml:"Party"`
}

type Party struct {
	Name PartyName `xml:"PartyName"`
}

type PartyName struct {
	Name string `xml:"Name"`
}

type MonetaryTotalType struct {
	PayableAmount PayableAmountType `xml:"PayableAmount"`
}

type PayableAmountType struct {
	Value string `xml:",innerxml"`
}

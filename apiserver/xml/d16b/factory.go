package d16b

import (
	"encoding/xml"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"strconv"
)

func Create(value string) (error, *invoice.Meta) {
	inv := &CrossIndustryInvoice{}
	err := xml.Unmarshal([]byte(value), &inv)
	if err != nil {
		return err, nil
	}

	price, _ := strconv.ParseFloat(inv.SupplyChainTradeTransaction.ApplicableHeaderTradeSettlement.SpecifiedTradeSettlementHeaderMonetarySummation.LineTotalAmount.Value, 64)
	return nil, &invoice.Meta{
		Sender:   inv.SupplyChainTradeTransaction.ApplicableHeaderTradeAgreement.SellerTradeParty.Name,
		Receiver: inv.SupplyChainTradeTransaction.ApplicableHeaderTradeAgreement.BuyerTradeParty.Name,
		Format:   invoice.D16bFormat,
		Price:    price,
	}
}

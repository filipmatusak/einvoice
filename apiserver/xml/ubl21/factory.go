package ubl21

import (
	"encoding/xml"
	"github.com/slovak-egov/einvoice/apiserver/invoice"
	"strconv"
)

func Create(value string) (error, *invoice.Meta) {
	inv := &Invoice{}
	err := xml.Unmarshal([]byte(value), &inv)
	if err != nil {
		return err, nil
	}

	price, _ := strconv.ParseFloat(inv.LegalMonetaryTotal.PayableAmount.Value, 64)
	return nil, &invoice.Meta{
		Sender:   inv.Supplier.Party.Name.Name,
		Receiver: inv.Customer.Party.Name.Name,
		Format:   invoice.UblFormat,
		Price:    price,
	}
}

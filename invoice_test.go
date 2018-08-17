package fivaldi_test

import (
	"fmt"
	"testing"
	"time"

	fixedwidth "github.com/omniboost/go-fixedwidth"
	fivaldi "github.com/tim-online/go-visma-fivaldi"
)

func TestReslasMarshalling(t *testing.T) {
	reslas := fivaldi.Reslas{
		CompanyID:               12,
		Type:                    "RESLAS",
		CustomerNumber:          23,
		InvoiceNumber:           34,
		RefundableInvoiceNumber: 45,
		AccountingPeriod:        fivaldi.Period{time.Now()},
		InvoiceDate:             fivaldi.Date{time.Now()},
		PaymentTerms:            56,
		CashDiscountDate:        fivaldi.Date{time.Now()},
		CashDiscountPercentage:  99,
		InvoiceDueDate:          fivaldi.Date{time.Now().AddDate(0, 0, 14)},
		GrossAmount:             fivaldi.Amount(66666.6),
		Sign:                    '-',
		Currency:                "EUR",
	}

	data, err := fixedwidth.Marshal(reslas)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
}

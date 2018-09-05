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
		Currency:                "EUR",
	}

	data, err := fixedwidth.Marshal(reslas)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
}

func TestInvoiceMarshalling(t *testing.T) {
	invoice := fivaldi.Invoice{
		CompanyID:               12,
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
		Currency:                "EUR",
		InvoiceLines: fivaldi.InvoiceLines{
			{
				SalesLedger: "SL",
				CostCenter1: "CC1",
				CostCenter2: "CC2",
				CostCenter3: "CC3",
				NetAmount:   fivaldi.Amount(123.45),
				FCNetAmount: fivaldi.Amount(-123.45),
				VATCode:     "VC",
				CostCenter4: "CC4",
				VATAmount:   fivaldi.Amount(987.65),
				FCVATAmount: fivaldi.Amount(-987.65),

				VATLedger: "12345678",
			},
		},
	}

	data, err := fixedwidth.Marshal(invoice)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(data))
}

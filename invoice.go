package fivaldi

type Invoice struct {
	Lines InvoiceLines
}

type InvoiceLines []InvoiceLine

type InvoiceLine struct {
}

type Reslas struct {
	CompanyID               int
	Type                    Type // always RESLAS
	CustomerNumber          int  // Customer number in Fivaldi
	InvoiceNumber           int
	RefundableInvoiceNumber int
	AccountingPeriod        Period
	InvoiceDate             Date
	PaymentTerms            int
	CashDiscountDate        Date
	CashDiscountPercentage  float64
	InvoiceDueDate          Date
	GrossAmount             Amount
	Sign                    char
	Currency                string
	FCGrossAmount           Amount
	FCSign                  byte
	FCExchangeRate          float64
	AccountReceivableLedger string
	Reference               string
	Description             string
	CustomerName1           string
	CustomerName2           string
	Address1                string
	Address2                string
	PostalCode              string
	Country                 string // iso3166
}

type Restap struct {
	CompanyID      int
	Type           Type // always RESLAS
	CustomerNumber int  // Customer number in Fivaldi
	InvoiceNumber  int
	SalesLedger    string
	CostCenter1    string
	CostCenter2    string
	CostCenter3    string
	NetAmount      Amount
	Sign           byte
	FCNetAmount    Amount
	FCSign         byte
	VATCode        string
	CostCenter4    string
	VATAmount      Amount
	VATSign        byte
	FCVATAmount    Amount
	FCVATSign      byte
	VATLedger      string
}

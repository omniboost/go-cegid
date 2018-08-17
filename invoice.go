package fivaldi

import (
	"fmt"
	"strconv"
	"time"

	fixedwidth "github.com/omniboost/go-fixedwidth"
)

type Invoice struct {
	// Customer Customer
	Lines InvoiceLines
}

type InvoiceLines []InvoiceLine

type InvoiceLine struct {
}

type Reslas struct {
	// fixed:"{startPos},{endPos}"
	CompanyID               Int     `fixed:"1,6"`
	Type                    Type    `fixed:"7,12"`  // always RESLAS
	CustomerNumber          Int     `fixed:"13,22"` // Customer number in Fivaldi
	InvoiceNumber           Int     `fixed:"23,30"`
	RefundableInvoiceNumber Int     `fixed:"31,38"`
	AccountingPeriod        Period  `fixed:"39,44"`
	InvoiceDate             Date    `fixed:"45,50"`
	PaymentTerms            Int     `fixed:"51,54"`
	CashDiscountDate        Date    `fixed:"55,60"`
	CashDiscountPercentage  Int     `fixed:"61,62"`
	InvoiceDueDate          Date    `fixed:"63,68"`
	GrossAmount             Amount  `fixed:"69,86"`
	Sign                    Sign    `fixed:"87,87"`
	Currency                string  `fixed:"88,90"`
	FCGrossAmount           Amount  `fixed:"91,108"`
	FCSign                  Sign    `fixed:"109,109"`
	FCExchangeRate          Decimal `fixed:"110,125"`
	AccountReceivableLedger string  `fixed:"126,133"`
	Reference               string  `fixed:"134,153"`
	Description             string  `fixed:"154,253"`
	CustomerName1           string  `fixed:"254,293"`
	CustomerName2           string  `fixed:"294,333"`
	Address1                string  `fixed:"334,413"`
	Address2                string  `fixed:"414,493"`
	PostalCode              string  `fixed:"494,495"`
	Country                 string  `fixed:"496,497"` // iso3166
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
	Sign           Sign
	FCNetAmount    Amount
	FCSign         Sign
	VATCode        string
	CostCenter4    string
	VATAmount      Amount
	VATSign        Sign
	FCVATAmount    Amount
	FCVATSign      Sign
	VATLedger      string
}

type Int int

func (i Int) String() string {
	return fmt.Sprint(int(i))
}

func (i Int) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	length := strconv.Itoa(spec.EndPos - spec.StartPos + 1)
	format := "%0" + length + "v"
	padded := fmt.Sprintf(format, i.String())
	return []byte(padded), nil
}

type Decimal float64

func (d Decimal) String() string {
	return fmt.Sprint(float64(d))
}

func (d Decimal) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	length := strconv.Itoa(spec.EndPos - spec.StartPos + 1)
	format := "%0" + length + "v"
	padded := fmt.Sprintf(format, d.String())
	return []byte(padded), nil
}

type Type string

func (t Type) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	return []byte(t), nil
}

type Period struct {
	time.Time
}

func (p Period) String() string {
	return p.Time.Format("200601")
}

func (p Period) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	return []byte(p.String()), nil
}

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Time.Format("060102")
}

func (d Date) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	return []byte(d.String()), nil
}

type Amount float64

func (a Amount) String() string {
	return fmt.Sprintf("%.2f", a)
}

func (a Amount) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	length := strconv.Itoa(spec.EndPos - spec.StartPos + 1)
	format := "%0" + length + "v"
	padded := fmt.Sprintf(format, a.String())
	return []byte(padded), nil
}

type Sign byte

func (s Sign) MarshalText() ([]byte, error) {
	return []byte{byte(s)}, nil
}

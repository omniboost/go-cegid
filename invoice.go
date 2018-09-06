package fivaldi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	fixedwidth "github.com/omniboost/go-fixedwidth"
)

type Invoice struct {
	FivaldiCompanyID        int
	CustomerNumber          int
	InvoiceNumber           int
	RefundableInvoiceNumber int
	AccountingPeriod        Period
	InvoiceDate             ShortDate
	PaymentTerms            int
	CashDiscountDate        ShortDate
	CashDiscountPercentage  int
	InvoiceDueDate          ShortDate
	GrossAmount             Amount
	Currency                string
	FCGrossAmount           Amount
	FCExchangeRate          Decimal
	AccountReceivableLedger string
	Reference               string
	Description             string
	CustomerName1           string
	CustomerName2           string
	Address1                string
	Address2                string
	PostalCode              string
	Country                 string
	InvoiceLines            InvoiceLines
}

func (i Invoice) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	bufs := [][]byte{}

	reslas := Reslas{
		FivaldiCompanyID:        Int(i.FivaldiCompanyID),
		Type:                    "RESLAS",
		CustomerNumber:          Int(i.CustomerNumber),
		InvoiceNumber:           Int(i.InvoiceNumber),
		RefundableInvoiceNumber: Int(i.RefundableInvoiceNumber),
		AccountingPeriod:        i.AccountingPeriod,
		InvoiceDate:             i.InvoiceDate,
		PaymentTerms:            Int(i.PaymentTerms),
		CashDiscountDate:        i.CashDiscountDate,
		InvoiceDueDate:          i.InvoiceDueDate,
		GrossAmount:             i.GrossAmount,
		Sign:                    i.GrossAmount.Sign(),
		Currency:                i.Currency,
		FCGrossAmount:           i.FCGrossAmount,
		FCSign:                  i.FCGrossAmount.Sign(),
		FCExchangeRate:          i.FCExchangeRate,
		AccountReceivableLedger: i.AccountReceivableLedger,
		Reference:               i.Reference,
		Description:             i.Description,
		CustomerName1:           i.CustomerName1,
		CustomerName2:           i.CustomerName2,
		Address1:                i.Address1,
		Address2:                i.Address2,
		PostalCode:              i.PostalCode,
		Country:                 i.Country,
	}

	buf, err := fixedwidth.Marshal(reslas)
	if err != nil {
		return buf, err
	}
	bufs = append(bufs, buf)

	for _, l := range i.InvoiceLines {
		restap := i.invoiceLineToRestap(l)

		buf, err := fixedwidth.Marshal(restap)
		if err != nil {
			return buf, err
		}
		bufs = append(bufs, buf)
	}

	buf = bytes.Join(bufs, []byte("\r\n"))
	return buf, nil
}

func (i Invoice) invoiceLineToRestap(l InvoiceLine) Restap {
	restap := Restap{
		FivaldiCompanyID: Int(i.FivaldiCompanyID),
		Type:             "RESTAP",
		CustomerNumber:   Int(i.CustomerNumber),
		InvoiceNumber:    Int(i.InvoiceNumber),
		SalesLedger:      l.SalesLedger,
		CostCenter1:      l.CostCenter1,
		CostCenter2:      l.CostCenter2,
		CostCenter3:      l.CostCenter3,
		NetAmount:        l.NetAmount,
		Sign:             l.NetAmount.Sign(),
		FCNetAmount:      l.FCNetAmount,
		FCSign:           l.FCNetAmount.Sign(),
		VATCode:          l.VATCode,
		CostCenter4:      l.CostCenter4,
		VATAmount:        l.VATAmount,
		VATSign:          l.VATAmount.Sign(),
		FCVATAmount:      l.FCVATAmount,
		FCVATSign:        l.FCVATAmount.Sign(),
		VATLedger:        l.VATLedger,
	}
	return restap
}

type InvoiceLines []InvoiceLine

type InvoiceLine struct {
	SalesLedger string
	CostCenter1 string
	CostCenter2 string
	CostCenter3 string
	NetAmount   Amount
	FCNetAmount Amount
	VATCode     string
	CostCenter4 string
	VATAmount   Amount
	FCVATAmount Amount
	VATLedger   string
}

type Reslas struct {
	// fixed:"{startPos},{endPos}"
	FivaldiCompanyID        Int       `fixed:"1,6"`
	Type                    Type      `fixed:"7,12"`  // always RESLAS
	CustomerNumber          Int       `fixed:"13,22"` // Customer number in Fivaldi
	InvoiceNumber           Int       `fixed:"23,30"`
	RefundableInvoiceNumber Int       `fixed:"31,38"`
	AccountingPeriod        Period    `fixed:"39,44"`
	InvoiceDate             ShortDate `fixed:"45,50"`
	PaymentTerms            Int       `fixed:"51,54"`
	CashDiscountDate        ShortDate `fixed:"55,60"`
	CashDiscountPercentage  Int       `fixed:"61,62"`
	InvoiceDueDate          ShortDate `fixed:"63,68"`
	GrossAmount             Amount    `fixed:"69,86"`
	Sign                    Sign      `fixed:"87,87"`
	Currency                string    `fixed:"88,90"`
	FCGrossAmount           Amount    `fixed:"91,108"`
	FCSign                  Sign      `fixed:"109,109"`
	FCExchangeRate          Decimal   `fixed:"110,125"`
	AccountReceivableLedger string    `fixed:"126,133"`
	Reference               string    `fixed:"134,153"`
	Description             string    `fixed:"154,253"`
	CustomerName1           string    `fixed:"254,293"`
	CustomerName2           string    `fixed:"294,333"`
	Address1                string    `fixed:"334,413"`
	Address2                string    `fixed:"414,493"`
	PostalCode              string    `fixed:"494,495"`
	Country                 string    `fixed:"496,497"` // iso3166
}

type Restap struct {
	// fixed:"{startPos},{endPos}"
	FivaldiCompanyID Int    `fixed:"1,6"`
	Type             Type   `fixed:"7,12"`  // always RESTAP
	CustomerNumber   Int    `fixed:"13,22"` // Customer number in Fivaldi
	InvoiceNumber    Int    `fixed:"23,30"`
	SalesLedger      string `fixed:"31,38"`
	CostCenter1      string `fixed:"39,46"`
	CostCenter2      string `fixed:"47,54"`
	CostCenter3      string `fixed:"55,62"`
	NetAmount        Amount `fixed:"63,80"`
	Sign             Sign   `fixed:"81,81"`
	FCNetAmount      Amount `fixed:"82,99"`
	FCSign           Sign   `fixed:"100,100"`
	VATCode          string `fixed:"101,101"`
	CostCenter4      string `fixed:"102,109"`
	VATAmount        Amount `fixed:"110,127"`
	VATSign          Sign   `fixed:"128,128"`
	FCVATAmount      Amount `fixed:"129,146"`
	FCVATSign        Sign   `fixed:"147,147"`
	VATLedger        string `fixed:"148,156"`
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

func (p *Period) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	p.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try fivaldi date format
	p.Time, err = time.Parse("200601", value)
	return err
}

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Time.Format("20060102")
}

func (d Date) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d Date) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try fivaldi date format
	d.Time, err = time.Parse("20060102", value)
	return err
}

type ShortDate struct {
	time.Time
}

func (d ShortDate) String() string {
	return d.Time.Format("060102")
}

func (d ShortDate) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	return []byte(d.String()), nil
}

func (d ShortDate) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *ShortDate) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try fivaldi date format
	d.Time, err = time.Parse("060102", value)
	return err
}

type Amount float64

func (a Amount) String() string {
	s := fmt.Sprintf("%.2f", a)
	return strings.Replace(s, ".", "", -1)
}

func (a Amount) Abs() Amount {
	aa := math.Abs(float64(a))
	return Amount(aa)
}

func (a Amount) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	aa := a.Abs()
	length := strconv.Itoa(spec.EndPos - spec.StartPos + 1)
	format := "%0" + length + "v"
	padded := fmt.Sprintf(format, aa.String())
	return []byte(padded), nil
}

func (a Amount) Sign() Sign {
	if a < 0 {
		return '-'
	}

	return ' '
}

type Sign byte

func (s Sign) MarshalText() ([]byte, error) {
	return []byte{byte(s)}, nil
}

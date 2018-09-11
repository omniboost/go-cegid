package fivaldi

import (
	"encoding"
	"fmt"
)

type GeneralJournalEntry struct {
	RecordDate       Date   // record date of the voucher, YYYYMMDD
	FivaldiCompanyID Int    // Fivaldi business id
	VoucherType      string // voucher type
	VoucherNumber    string // Voucher number. If no voucher number exists, the date change will change to the next voucher
	Lines            GeneralJournalEntryLines
}

type GeneralJournalEntryLines []GeneralJournalEntryLine

type GeneralJournalEntryLine struct {
	AccountNumber     string // account number
	CostCenter1       string // cost center/follow up code 1
	CostCenter2       string // cost center/follow up code 2
	CostCenter3       string // cost center/follow up code 3
	CostCenter4       string // cost center/follow up code 4
	Amount            Amount // The amount in the currency of the company with its sign (only if “-“)
	AmountInCents     int    // Amount in cents with sign
	Description       string // booking entry description
	VATCode           string // tax rate code
	VATAccountNumber  string // tax account number
	VATAccountNumber2 string // tax account number 2 (counter entry)
	VATAmount         Amount // The amount of tax in the company's currency with its sign
	VATAmountInCents  int    // The amount of tax in cents with sign
	Status            string // Special handling of entry
	NetDate           Date   // net date YYYYMMDD
	CustomerID        string // Customer ID in Fivaldi customer register
	ExportTypeID      string // export type id
	InvoiceNumber     string // invoice number
}

type GeneralJournalCSVLine struct {
	RecordDate        Date   // record date of the voucher, YYYYMMDD
	AccountNumber     string // account number
	VoucherNumber     string // Voucher number. If no voucher number exists, the date change will change to the next voucher
	CostCenter1       string // cost center/follow up code 1
	CostCenter2       string // cost center/follow up code 2
	CostCenter3       string // cost center/follow up code 3
	CostCenter4       string // cost center/follow up code 4
	Amount            Amount // The amount in the currency of the company with its sign (only if “-“)
	AmountInCents     int    // Amount in cents with sign
	Description       string // booking entry description
	VATCode           string // tax rate code
	VATAccountNumber  string // tax account number
	VATAccountNumber2 string // tax account number 2 (counter entry)
	VATAmount         Amount // The amount of tax in the company's currency with its sign
	VATAmountInCents  int    // The amount of tax in cents with sign
	Status            string // Special handling of entry
	NetDate           Date   // net date YYYYMMDD
	CustomerID        string // Customer ID in Fivaldi customer register
	ExportTypeID      string // export type id
	InvoiceNumber     string // invoice number
	FivaldiCompanyID  Int    // Fivaldi business id
	VoucherType       string // voucher type
}

func (l GeneralJournalCSVLine) Headers() []string {
	return []string{
		"PVM",
		"TILI",
		"TOSITE",
		"SK1",
		"SK2",
		"SK3",
		"SK4",
		"SUMMA",
		"SENTIT",
		"SELITE",
		"VEROKANTA",
		"VEROTILI",
		"VEROTILI2",
		"VEROSUMMA",
		"VEROSENTIT",
		"STATUS",
		"NETTOPVM",
		"ASIAKASTUNNUS",
		"VIENTILAJI",
		"LASKUNO",
		"YT",
		"LAJI",
	}
}

func (l GeneralJournalCSVLine) ToStrings() []string {
	m := l.ToMap()
	return []string{
		m["PVM"],
		m["TILI"],
		m["TOSITE"],
		m["SK1"],
		m["SK2"],
		m["SK3"],
		m["SK4"],
		m["SUMMA"],
		m["SENTIT"],
		m["SELITE"],
		m["VEROKANTA"],
		m["VEROTILI"],
		m["VEROTILI2"],
		m["VEROSUMMA"],
		m["VEROSENTIT"],
		m["STATUS"],
		m["NETTOPVM"],
		m["ASIAKASTUNNUS"],
		m["VIENTILAJI"],
		m["LASKUNO"],
		m["YT"],
		m["LAJI"],
	}
}

func (l GeneralJournalCSVLine) ToMap() map[string]string {
	mi := map[string]interface{}{
		"PVM":           l.RecordDate,
		"TILI":          l.AccountNumber,
		"TOSITE":        l.VoucherNumber,
		"SK1":           l.CostCenter1,
		"SK2":           l.CostCenter2,
		"SK3":           l.CostCenter3,
		"SK4":           l.CostCenter4,
		"SUMMA":         l.Amount,
		"SENTIT":        l.AmountInCents,
		"SELITE":        l.Description,
		"VEROKANTA":     l.VATCode,
		"VEROTILI":      l.VATAccountNumber,
		"VEROTILI2":     l.VATAccountNumber2,
		"VEROSUMMA":     l.VATAmount,
		"VEROSENTIT":    l.VATAmountInCents,
		"STATUS":        l.Status,
		"NETTOPVM":      l.NetDate,
		"ASIAKASTUNNUS": l.CustomerID,
		"VIENTILAJI":    l.ExportTypeID,
		"LASKUNO":       l.InvoiceNumber,
		"YT":            l.FivaldiCompanyID,
		"LAJI":          l.VoucherType,
	}

	m := map[string]string{}
	for k, v := range mi {
		marshaller, ok := v.(encoding.TextMarshaler)
		if ok {
			b, err := marshaller.MarshalText()
			if err == nil {
				m[k] = string(b)
				continue
			}
		}

		m[k] = fmt.Sprint(v)
	}

	if m["NETTOPVM"] == "00010101" {
		m["NETTOPVM"] = ""
	}

	if m["SENTIT"] == "0" {
		m["SENTIT"] = ""
	}

	if m["VEROSUMA"] == "0" {
		m["VEROSUMA"] = ""
	}

	if m["VEROSENTIT"] == "0" {
		m["VEROSENTIT"] = ""
	}

	return m
}

type GeneralJournalCSVLines []GeneralJournalCSVLine

func GeneralJournalEntryToCSVLines(entry GeneralJournalEntry) (GeneralJournalCSVLines, error) {
	csvLines := GeneralJournalCSVLines{}
	for _, l := range entry.Lines {
		csvLine := GeneralJournalCSVLine{
			RecordDate:        entry.RecordDate,
			AccountNumber:     l.AccountNumber,
			VoucherNumber:     entry.VoucherNumber,
			CostCenter1:       l.CostCenter1,
			CostCenter2:       l.CostCenter2,
			CostCenter3:       l.CostCenter3,
			CostCenter4:       l.CostCenter4,
			Amount:            l.Amount,
			AmountInCents:     l.AmountInCents,
			Description:       l.Description,
			VATCode:           l.VATCode,
			VATAccountNumber:  l.VATAccountNumber,
			VATAccountNumber2: l.VATAccountNumber2,
			VATAmount:         l.VATAmount,
			VATAmountInCents:  l.VATAmountInCents,
			Status:            l.Status,
			NetDate:           l.NetDate,
			CustomerID:        l.CustomerID,
			ExportTypeID:      l.ExportTypeID,
			InvoiceNumber:     l.InvoiceNumber,
			FivaldiCompanyID:  entry.FivaldiCompanyID,
			VoucherType:       entry.VoucherType,
		}
		csvLines = append(csvLines, csvLine)
	}
	return csvLines, nil
}

package main

import (
	"time"
	"github.com/shopspring/decimal"
	"regexp"
	"sort"
)

type Party struct {
	Name string
	Nip string
	Regon string
	Address string
	Address2 string
	BankAccount string
	Version time.Time
}

type InvoiceEntry struct {
	Description string
	Pkwiu string
	Quantity string
	QuantityUnit string
	PriceNet string
	Vat string
}

type Invoice struct {
	Number string
	Seller string
	Buyer string
	IssueDate string
	IssuePlace string
	SellDate string
	DueDate string
	InvoiceEntries []InvoiceEntry
}

type Data struct {
	Name string
	Nip string
	Parties map[string]Party
	Invoices map[string][]Invoice
}

type Val struct {
	v decimal.Decimal
} 

type summary struct {
	Tax string
	NetValue Val
	TaxValue Val
	GrossValue Val
}

func (v Val) Add(addon Val) Val {
	return Val{v.v.Add(addon.v)}
}

func Parse(val string) Val {
	value, _ := decimal.NewFromString(val)
	return Val{value}
}

var dot, _ = regexp.Compile("[.]")

func PolishNum(v string) string {
	result := ""
	length := len(v)
	for i, r := range v {
		result += string(r)
		if (length - i - 1) % 3 == 0 && i < length - 6 {
			result += " "
		}
	}

	return dot.ReplaceAllString(result, ",")
}

func FormatDecimal(v string) string {
	val, _ := decimal.NewFromString(v)
	return PolishNum(val.StringFixed(2))
}

func (v Val) String() string {
	return PolishNum(v.v.StringFixed(2))
}

func (invoice Invoice) NetValueGrouped() []summary {
	var groups map[string]summary = make(map[string]summary)
	for _, entry := range invoice.InvoiceEntries {
		if _, ok := groups[entry.Vat]; !ok {
			groups[entry.Vat] = summary {
				entry.Vat,
				Val{zero},
				Val{zero},
				Val{zero},
			}
		}
		sum := groups[entry.Vat]
		groups[entry.Vat] = summary {
			entry.Vat,
			sum.NetValue.Add(entry.NetValue()),
			sum.TaxValue.Add(entry.TaxValue()),
			sum.GrossValue.Add(entry.NetValue()).Add(entry.TaxValue()),
		}
	}
	var keys []string
	for k := range groups {
	    keys = append(keys, k)
	}
	sort.Strings(keys)

	var values []summary
	for _, k := range keys {
	    values = append(values, groups[k])
	}
	return values
}

func (invoice Invoice) NetValueSum() Val {
	result := Val{zero}
	for _, entry := range invoice.InvoiceEntries {
		entryValue := entry.NetValue()
		result = result.Add(entryValue)
	}
	return result
}

func (entry InvoiceEntry) NetValue() Val {
	quantity, _ := decimal.NewFromString(entry.Quantity)
	price, _ := decimal.NewFromString(entry.PriceNet)
	return Val{quantity.Mul(price)}
}

func (invoice Invoice) GrossValueSum() Val {
	result := Val{zero}
	for _, entry := range invoice.InvoiceEntries {
		entryValue := entry.GrossValue()
		result = result.Add(entryValue)
	}
	return result
}

func (entry InvoiceEntry) GrossValue() Val {
	return entry.NetValue().Add(entry.TaxValue())
}

func (invoice Invoice) TaxValueSum() Val {
	result := Val{zero}
	for _, entry := range invoice.InvoiceEntries {
		entryValue := entry.TaxValue()
		result = result.Add(entryValue)
	}
	return result
}

func (entry InvoiceEntry) TaxValue() Val {
	net := entry.NetValue()
	tax, _ := decimal.NewFromString(entry.Vat)
	return Val{net.v.Mul(tax).Shift(-2)}
}

func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}


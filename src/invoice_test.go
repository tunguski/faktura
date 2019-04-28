package main

import (
	"fmt"
	"testing"
)

// "buyer" "Code of the buyer"
// "seller" "Name of the seller"
// "issuanceDate" "Date of issuance"
// "issuancePlace" "Place of issuance"
// "sellDate" "Sell date"
// "positionFormat" "Defines input format of the positions"
// "positions" "Positions declared on invoice"

func TestValidateAddInvoice(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"buyer":         "pzu",
		"issuanceDate":  "2019-04-04",
		"issuancePlace": "Warszawa",
		"positions":     "none;0"})

	// when
	err := validateAddInvoice(flags)

	// then
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}

func TestAddInvoice(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"buyer":         "pzu",
		"issuanceDate":  "2019-04-04",
		"issuancePlace": "Warszawa",
		"positions":     "zadanie 1;30\nzadanie 2;75\nzadanie 3; 100"})
	configfile = "../.out.toml"

	// when
	data, err := addInvoiceNoStore(flags)
	//fmt.Printf("Data: %v", data)

	// then
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	invoice := data.Invoices["test"][0]
	if len(invoice.InvoiceEntries) != 3 {
		fmt.Printf("Invoice: %v", invoice)
		t.Errorf("Expected 2 invoice positions but got %d", len(invoice.InvoiceEntries))
	}
}

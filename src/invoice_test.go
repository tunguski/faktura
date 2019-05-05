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
	configfile = "../.out.toml"
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

func TestValidateAddInvoice_sellerNotDefined(t *testing.T) {
	// given
	configfile = "../.out.toml"
	flags := paramsFromMap(map[string]string{
		"buyer":         "pzu",
		"seller":        "test",
		"issuanceDate":  "2019-04-04",
		"issuancePlace": "Warszawa",
		"positions":     "none;0"})

	// when
	err := validateAddInvoice(flags)

	// then
	if err == nil {
		t.Errorf("Expected error - no seller defined")
	}
}

func TestAddInvoice_defaultSeller(t *testing.T) {
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
		return
	}

	invoices := data.Invoices["szamoni"]
	invoice := invoices[len(invoices)-1]
	if len(invoice.InvoiceEntries) != 3 {
		fmt.Printf("Invoice: %v", invoice)
		t.Errorf("Expected 2 invoice positions but got %d", len(invoice.InvoiceEntries))
	}
}

func TestAddInvoice(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"buyer":         "pzu",
		"seller":        "tortoise",
		"issuanceDate":  "2019-04-04",
		"issuancePlace": "Warszawa",
		"positions": `zadanie 1;30
				zadanie 2;75;7
				zadanie 3;10;100;7

				// empty lines are ignored
				# same as comments (both // and #)

				zadanie 4;07.29.11.1;1000;9;23
				zadanie 5;07.29.11.1;1000;ton;9;23`})
	configfile = "../.out.toml"

	// when
	data, err := addInvoiceNoStore(flags)
	//fmt.Printf("Data: %v", data)

	// then
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
		return
	}

	invoice := data.Invoices["tortoise"][0]
	if len(invoice.InvoiceEntries) != 5 {
		fmt.Printf("Invoice: %v", invoice)
		t.Errorf("Expected 2 invoice positions but got %d", len(invoice.InvoiceEntries))
	}
	if invoice.Number != "FV/2019/04/01" {
		t.Errorf("Expected invoice number 'FV/2019/04/01' but got '%s'", invoice.Number)
	}
}

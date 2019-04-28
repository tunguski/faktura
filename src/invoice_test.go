package main

import (
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
		"positions":     "none;0"})
	configfile = "../.out.toml"

	// when
	data, err := addInvoiceNoStore(flags)

	// then
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if data == nil {
		t.Errorf("No data returned")
	}
}

package main

import "testing"

func TestValidatePrintInvoice(t *testing.T) {
	// given
	configfile = "../.out.toml"
	flags := paramsFromMap(map[string]string{
		"party": "pzu",
		"last":  "true"})

	// when
	err := validatePrintInvoice(flags)

	// then
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
}

func TestValidatePrintInvoice_noParty(t *testing.T) {
	// given
	configfile = "../.out.toml"
	flags := paramsFromMap(map[string]string{
		"last": "true"})

	// when
	err := validatePrintInvoice(flags)

	// then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestValidatePrintInvoice_noSpecifier(t *testing.T) {
	// given
	configfile = "../.out.toml"
	flags := paramsFromMap(map[string]string{
		"party": "pzu"})

	// when
	err := validatePrintInvoice(flags)

	// then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestPrintInvoice(t *testing.T) {
	// given
	configfile = "../.out.toml"
	tmpFolder = "../tmp"
	fakturaTex = "../" + fakturaTex
	flags := paramsFromMap(map[string]string{
		"party": "szamoni",
		"last":  "true"})

	// when
	err := printInvoice(flags)

	// then
	if err != nil {
		t.Errorf("Expected success but got: %v", err)
	}
}

package main

import (
	"testing"
)

func TestValidateAddParty_missingName(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"code":    "test-code",
		"address": "test-address",
		"nip":     "test-nip"})

	// when
	err := validateAddParty(flags)

	// then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestValidateAddParty_missingCode(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"name":    "test-name",
		"address": "test-address",
		"nip":     "test-nip"})

	// when
	err := validateAddParty(flags)

	// then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestValidateAddParty_missingNip(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"name":    "test-name",
		"address": "test-address",
		"code":    "test-code"})

	// when
	err := validateAddParty(flags)

	// then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestValidateAddParty_missingAddress(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"code": "test-code",
		"name": "test-name",
		"nip":  "test-nip"})

	// when
	err := validateAddParty(flags)

	// then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestValidateAddParty(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"name":    "test-name",
		"code":    "test-code",
		"address": "test-address",
		"nip":     "test-nip"})

	// when
	err := validateAddParty(flags)

	// then
	if err != nil {
		t.Errorf("Expected nil but got %d", err)
	}
}

func TestAddParty(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"name":    "test-name",
		"code":    "test-code",
		"address": "test-address",
		"nip":     "test-nip"})
	configfile = "../.out.toml"

	// when
	data, err := addPartyNoStore(flags)

	// then
	if err != nil {
		t.Errorf("Expected nil but got: %v", err)
	} else if length := len(data.Parties); length != 4 {
		t.Errorf("Expected 4 paries but got %d", length)
	}
}

func TestValidateModifyParty_missingCode(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{})

	// when
	err := validateModifyParty(flags)

	// then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestValidateModifyParty(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"code": "test-code"})

	// when
	err := validateModifyParty(flags)

	// then
	if err != nil {
		t.Errorf("Expected nil but got %d", err)
	}
}

func TestModifyParty(t *testing.T) {
	// given
	flags := paramsFromMap(map[string]string{
		"code":              "szamoni",
		"name":              "test-name",
		"nip":               "test-nip",
		"regon":             "test-regon",
		"address":           "test-address",
		"address2":          "test-address2",
		"bankAccount":       "test-bankAccount",
		"numbering-pattern": "test-invoiceNumberingPattern",
	})
	configfile = "../.out.toml"

	// when
	data, err := modifyPartyNoStore(flags)

	// then
	if err != nil {
		t.Errorf("Expected nil but got: %v", err)
	} else if length := len(data.Parties); length != 3 {
		t.Errorf("Expected 3 paries but got %d", length)
	} else {
		party, _ := data.Parties["szamoni"]
		if party.Name != "test-name" {
			t.Errorf("Name missmatch: %s", party.Name)
		}
		if party.Nip != "test-nip" {
			t.Errorf("Nip missmatch: %s", party.Nip)
		}
		if party.Regon != "test-regon" {
			t.Errorf("Regon missmatch: %s", party.Regon)
		}
		if party.Address != "test-address" {
			t.Errorf("Address missmatch: %s", party.Address)
		}
		if party.Address2 != "test-address2" {
			t.Errorf("Address2 missmatch: %s", party.Address2)
		}
		if party.BankAccount != "test-bankAccount" {
			t.Errorf("BankAccount missmatch: %s", party.BankAccount)
		}
		if party.InvoiceNumberingPattern != "test-invoiceNumberingPattern" {
			t.Errorf("InvoiceNumberingPattern missmatch: %s", party.InvoiceNumberingPattern)
		}
	}
}

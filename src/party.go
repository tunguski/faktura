package main

import (
	"errors"
	"fmt"
	"time"
)

func validateAddParty(c paramsAccessor) error {
	if c.String("name") == "" ||
		c.String("code") == "" ||
		c.String("address") == "" ||
		c.String("nip") == "" {
		fmt.Printf("Code, name, address and nip are required: %s; %s; %s; %s\n",
			c.String("name"),
			c.String("code"),
			c.String("address"),
			c.String("nip"))
		return errors.New("missing_data")
	}
	return nil
}

func addPartyNoStore(c paramsAccessor) (*Data, error) {
	code := c.String("code")
	party := &Party{
		c.String("name"),
		c.String("nip"),
		c.String("regon"),
		c.String("address"),
		c.String("address2"),
		c.String("bankAccount"),
		"",
		time.Now()}

	fmt.Println(code)

	var data = readConfig()

	if _, ok := data.Parties[code]; ok {
		return nil, errors.New("" + code + " already defined")
	}

	data.Parties[code] = *party

	return data, nil
}

func addParty(c paramsAccessor) error {
	data, err := addPartyNoStore(c)

	storeData(data)

	return err
}

func validateModifyParty(c paramsAccessor) error {
	if c.String("code") == "" {
		fmt.Println("Code is required")
		return errors.New("missing_data")
	}
	return nil
}

func modifyPartyNoStore(c paramsAccessor) (*Data, error) {
	code := c.String("code")
	//party := &Party{
	//	c.String("name"),
	//	c.String("nip"),
	//	c.String("regon"),
	//	c.String("address"),
	//	c.String("address2"),
	//	c.String("bankAccount"),
	//	"",
	//	time.Now()}

	var data = readConfig()

	if party, ok := data.Parties[code]; ok {
		fmt.Println(code)

		if c.IsSet("name") {
			party.Name = c.String("name")
		}
		if c.IsSet("nip") {
			party.Nip = c.String("nip")
		}
		if c.IsSet("regon") {
			party.Regon = c.String("regon")
		}
		if c.IsSet("address") {
			party.Address = c.String("address")
		}
		if c.IsSet("address2") {
			party.Address2 = c.String("address2")
		}
		if c.IsSet("bankAccount") {
			party.BankAccount = c.String("bankAccount")
		}
		if c.IsSet("numbering-pattern") {
			party.InvoiceNumberingPattern = c.String("numbering-pattern")
		}

		data.Parties[code] = party
		return data, nil
	} else {
		return nil, errors.New("" + code + " not found")
	}
}

func modifyParty(c paramsAccessor) error {
	data, err := modifyPartyNoStore(c)

	storeData(data)

	return err
}

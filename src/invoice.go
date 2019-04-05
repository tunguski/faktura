package main

import (
	"os"
	"fmt"
	"errors"
	"encoding/json"
	"github.com/urfave/cli"
)

func validateAddInvoice(c *cli.Context) error {
	if c.String("buyer") == "" || 
		c.String("issuanceDate") == "" || 
		c.String("issuancePlace") == "" || 
		c.String("positions") == "" {
    	fmt.Println("Buyer, issuanceDate, issuancePlace and positions are required")
		cli.ShowCommandHelp(c, "add")
		return errors.New("missing_param")
	}
	return nil
}

func addInvoice(c *cli.Context) error {
	var positions []InvoiceEntry
		//c.String("positions")

	issuanceDate := c.String("issuanceDate")
	if issuanceDate == "" {
		fmt.Println("Issuance date not defined")
		os.Exit(1)
	}

	sellDate := c.String("sellDate")
	if sellDate == "" {
		sellDate = issuanceDate
	}

	dueDate := c.String("dueDate")
	if issuanceDate == "" {
		fmt.Println("Due date not define")
		os.Exit(1)
	}

	seller := c.String("seller")
	// TODO: set default seller if empty

	buyer := c.String("buyer")
	if issuanceDate == "" {
		fmt.Println("Buyer not defined")
		os.Exit(1)
	}

	invoice := Invoice{
		"invoice-number",
		seller,
		buyer,
		issuanceDate,
		c.String("issuancePlace"),
		sellDate,
		dueDate,
		positions,
	}

	out, _ := json.Marshal(invoice)
	fmt.Println(string(out))

	var data = readConfig()

	if _, ok := data.Invoices["test"]; ok {
		fmt.Println("" + "test" + " already defined")
		os.Exit(1)
	}

	data.Invoices["test"] = append(data.Invoices["test"], invoice)

	storeData(".out.toml", data)

	return nil
}


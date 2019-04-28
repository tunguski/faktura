package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateAddInvoice(c paramsAccessor) error {
	if c.String("buyer") == "" ||
		c.String("issuanceDate") == "" ||
		c.String("issuancePlace") == "" ||
		c.String("positions") == "" {
		fmt.Println("Buyer, issuanceDate, issuancePlace and positions are required")
		return errors.New("missing_param")
	}
	return nil
}

func parsePositions(text string) ([]InvoiceEntry, error) {
	lines := strings.Split(text, "\n")
	length := len(lines)
	split := make([][]string, length)

	for index, element := range lines {
		split[index] = strings.Split(element, ";")
	}

	for _, element := range split {
		if len(split[0]) != len(element) {
			return nil, errors.New("Unparseable positions")
		}
	}

	positions := make([]InvoiceEntry, length)

	//for index, element := range split {

	//}

	return positions, nil
}

func addInvoiceNoStore(c paramsAccessor) (*Data, error) {
	positions, err := parsePositions(c.String("positions"))

	if err != nil {
		return nil, err
	}

	issuanceDate := c.String("issuanceDate")
	sellDate := c.String("sellDate")
	dueDate := c.String("dueDate")
	// TODO: set default seller if empty
	seller := c.String("seller")
	buyer := c.String("buyer")

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

	data := readConfig()
	invoices, ok := data.Invoices["test"]
	if !ok {
		invoices = []Invoice{}
	}

	data.Invoices["test"] = append(invoices, invoice)

	return data, nil
}

func addInvoice(c paramsAccessor) error {
	data, err := addInvoiceNoStore(c)

	if err != nil {
		return err
	}

	storeData(data)

	return nil

}

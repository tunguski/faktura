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

	positions := make([]InvoiceEntry, 0)

	for index, element := range split {
		size := len(element)
		position := InvoiceEntry{
			"",
			"",
			"1",
			"szt.",
			"",
			"23",
		}

		if size == 2 {
			position.Description = element[0]
			position.PriceNet = element[1]
		} else {
			return nil, errors.New(fmt.Sprintf("Could not parse position %d: %s", index, lines[index]))
		}
		positions = append(positions, position)
	}

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
	if seller == "" {
		seller = "test"
	}

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
	invoices, ok := data.Invoices[seller]
	if !ok {
		invoices = []Invoice{}
	}

	data.Invoices[seller] = append(invoices, invoice)

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

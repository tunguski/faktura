package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"
)

func validateAddInvoice(c paramsAccessor) error {
	if c.String("buyer") == "" ||
		c.String("issuanceDate") == "" ||
		c.String("issuancePlace") == "" ||
		c.String("positions") == "" {
		fmt.Println("Buyer, issuanceDate, issuancePlace and positions are required")
		return errors.New("missing_param")
	}

	data := readConfig()

	sellerName := c.String("seller")
	if sellerName == "" {
		sellerName = data.DefaultParty
	}

	_, ok := data.Parties[sellerName]
	if !ok {
		return errors.New("No party with code " + sellerName + " defined")
	}

	return nil
}

func parsePositions(text string) ([]InvoiceEntry, error) {
	lines := strings.Split(text, "\n")
	split := make([][]string, 0)

	for _, element := range lines {
		trimmed := strings.TrimSpace(element)
		if trimmed != "" &&
			!strings.HasPrefix(trimmed, "//") &&
			!strings.HasPrefix(trimmed, "#") {
			split = append(split, strings.Split(element, ";"))
		}
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
		} else if size == 3 {
			position.Description = element[0]
			position.PriceNet = element[1]
			position.Vat = element[2]
		} else if size == 4 {
			position.Description = element[0]
			position.Quantity = element[1]
			position.PriceNet = element[2]
			position.Vat = element[3]
		} else if size == 5 {
			position.Description = element[0]
			position.Pkwiu = element[1]
			position.Quantity = element[2]
			position.PriceNet = element[3]
			position.Vat = element[4]
		} else if size == 6 {
			position.Description = element[0]
			position.Pkwiu = element[1]
			position.Quantity = element[2]
			position.QuantityUnit = element[3]
			position.PriceNet = element[4]
			position.Vat = element[5]
		} else {
			return nil, fmt.Errorf(
				"Could not parse position %d: %s", index, lines[index])
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

	data := readConfig()

	issuanceDate := c.String("issuanceDate")
	sellDate := c.String("sellDate")
	dueDate := c.String("dueDate")
	// TODO: set default seller if empty
	sellerName := c.String("seller")
	if sellerName == "" {
		sellerName = data.DefaultParty
	}

	buyer := c.String("buyer")

	invoice := Invoice{
		"<nil>",
		sellerName,
		buyer,
		issuanceDate,
		c.String("issuancePlace"),
		sellDate,
		dueDate,
		positions,
	}
	err = generateInvoiceNumber(&invoice, data)
	if err != nil {
		return nil, err
	}

	invoices, ok := data.Invoices[sellerName]
	if !ok {
		invoices = []Invoice{}
	}

	data.Invoices[sellerName] = append(invoices, invoice)

	return data, nil
}

type generateInvoiceNumberData struct {
	Year   string
	Month  string
	Number string
}

func createInvoiceNumberData(invoice *Invoice, data *Data) generateInvoiceNumberData {
	year := "2010"
	month := "00"
	number := "00"

	if split := strings.Split(invoice.IssueDate, "-"); len(split) > 2 {
		year = split[0]
		month = split[1]
	} else if split = strings.Split(invoice.IssueDate, "/"); len(split) > 2 {
		year = split[0]
		month = split[1]
	}

	return generateInvoiceNumberData{
		Year:   year,
		Month:  month,
		Number: number,
	}
}

func generateInvoiceNumber(invoice *Invoice, data *Data) error {
	var numberTemplate = "FV/{{.Year}}/{{.Month}}/{{.Number}}"
	tmpl, err := template.
		New("invoice").
		//Funcs(templateFunctions()).
		Parse(numberTemplate)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	err = tmpl.Execute(writer, createInvoiceNumberData(invoice, data))
	if err != nil {
		return err
	}
	writer.Flush()

	invoice.Number = b.String()

	return nil
}

func addInvoice(c paramsAccessor) error {
	data, err := addInvoiceNoStore(c)

	if err != nil {
		return err
	}

	storeData(data)

	return nil

}

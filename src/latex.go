package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type invoicePrintData struct {
	Invoice *Invoice
	Seller  Party
	Buyer   Party
}

func validatePrintInvoice(c paramsAccessor) error {
	if c.String("party") == "" ||
		c.String("last") == "" {
		fmt.Printf("Party code and invoice specification is required\n")
		return errors.New("missing_param")
	}
	return nil
}

func templateFunctions() map[string]interface{} {
	return map[string]interface{}{
		"FormatDecimal": FormatDecimal,
		"inc": func(i int) int {
			return i + 1
		},
		"valToPolishText": valToPolishText,
	}
}

func templateData(invoice *Invoice, data *Data) invoicePrintData {
	return invoicePrintData{
		Invoice: invoice,
		Seller:  data.Parties[invoice.Seller],
		Buyer:   data.Parties[invoice.Buyer],
	}
}

func printInvoice(c paramsAccessor) error {

	data := readConfig()

	invoice, err := findInvoice(data, c)
	if err != nil {
		return err
	}

	err = writeFilledTemplate(templateData(invoice, data))
	if err != nil {
		return err
	}

	return executeGeneration()
}

func findInvoice(data *Data, c paramsAccessor) (*Invoice, error) {
	party := c.String("party")
	last := c.Bool("last")

	fmt.Println("generate pdf {} {}", party, last)

	invoices, ok := data.Invoices[party]
	if !ok {
		return nil, errors.New("Party does not have invoices")
	}
	if len(invoices) == 0 {
		return nil, errors.New("Party does not have invoices")
	}
	invoice := invoices[0]

	return &invoice, nil
}

func writeFilledTemplate(templateData invoicePrintData) error {
	dat := loadFile(fakturaTex)
	file, err := os.Create(getTmpFolder() + "invoice_source")

	if err != nil {
		return err
	}

	tmpl, err := template.
		New("invoice").
		Funcs(templateFunctions()).
		Parse(string(dat))
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	err = tmpl.Execute(writer, templateData)
	if err != nil {
		return err
	}
	writer.Flush()
	file.Close()

	return nil
}

func executeGeneration() error {
	cmd := exec.Command("pdflatex", "invoice_source") //, "--output-directory=data/pdf")
	cmd.Dir = getTmpFolder()
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	return cmd.Run()
}

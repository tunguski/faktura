package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

type InvoicePrintData struct {
	Invoice Invoice
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

func printInvoice(c paramsAccessor) error {
	party := c.String("party")
	last := c.Bool("last")

	fmt.Println("generate pdf {} {}", party, last)

	dat := loadFile(fakturaTex)
	data := readConfig()

	invoice := data.Invoices["test"][0]

	templateData := InvoicePrintData{
		Invoice: invoice,
		Seller:  data.Parties[invoice.Seller],
		Buyer:   data.Parties[invoice.Buyer],
	}

	file, err := os.Create("tmp/invoice_source")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("invoice").Funcs(map[string]interface{}{
		"FormatDecimal": FormatDecimal,
		"inc": func(i int) int {
			return i + 1
		},
		"valToPolishText": valToPolishText,
	}).Parse(string(dat))
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

	cmd := exec.Command("pdflatex", "invoice_source") //, "--output-directory=data/pdf")
	cmd.Dir = "tmp"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}

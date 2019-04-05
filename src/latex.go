package main

import (
	"log"
	"fmt"
	"bufio"
	"os"
	"errors"
	"os/exec"
	"text/template"
	"github.com/urfave/cli"
)

type InvoicePrintData struct {
	Invoice Invoice
	Seller Party
	Buyer Party
}

func validatePrintInvoice(c *cli.Context) error {
	if c.String("party") == "" || 
		c.String("last") == "" {
		fmt.Println("Party code and invoice specification is required\n")
		cli.ShowCommandHelp(c, "add")
		return errors.New("missing_param")
	}
	return nil
}

func printInvoice(c *cli.Context) error {
	party := c.String("party")
	last := c.Bool("last")

	fmt.Println("generate pdf {} {}", party, last)

	dat := loadFile(fakturaTex)
	data := readConfig()

	invoice := data.Invoices["test"][0]

	templateData := InvoicePrintData{
		Invoice: invoice,
		Seller: data.Parties[invoice.Seller],
		Buyer: data.Parties[invoice.Buyer],
	}

	file, err := os.Create("tmp/invoice_source")
	if err != nil { panic(err) }

	tmpl, err := template.New("invoice").Funcs(map[string]interface{}{
			"FormatDecimal": FormatDecimal,
			"inc": func (i int) int {
				return i + 1
			},
			"valToPolishText": valToPolishText,
		}).Parse(string(dat))
	if err != nil { panic(err) }
	writer := bufio.NewWriter(file)
	err = tmpl.Execute(writer, templateData)
	if err != nil { panic(err) }
	writer.Flush()
	file.Close()

	cmd := exec.Command("pdflatex", "invoice_source")//, "--output-directory=data/pdf")
	cmd.Dir = "tmp"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	return nil
}


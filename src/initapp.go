package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

type validate func(c paramsAccessor) error
type action func(c paramsAccessor) error

type paramsAccessor interface {
	String(name string) string
	Bool(name string) bool
	IsSet(name string) bool
}

type mapBasedParams struct {
	data map[string]string
}

func (v mapBasedParams) String(key string) string {
	return v.data[key]
}

func (v mapBasedParams) Bool(key string) bool {
	value := v.data[key]
	return value == "true"
}

func (v mapBasedParams) IsSet(key string) bool {
	_, ok := v.data[key]
	return ok
}

func paramsFromMap(data map[string]string) paramsAccessor {
	return mapBasedParams{data: data}
}

func process(c *cli.Context, v validate, a action) error {
	if err := v(c); err != nil {
		return err
	} else {
		return a(c)
	}
}

func stringFlag(name string, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Usage: usage,
	}
}

func boolFlag(name string, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func initApp() *cli.App {
	return doInitApp(os.Args)
}

func doInitApp(args []string) *cli.App {
	app := createApp()

	fmt.Println(args)

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}

	return app
}

func createApp() *cli.App {
	app := cli.NewApp()
	app.Name = "faktura"
	app.Usage = "Create polish invoices in latex style"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "party",
			Usage: "Operations on parties",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "Add new party",
					Action: func(c *cli.Context) error {
						return process(c, validateAddParty, addParty)
					},
					Flags: []cli.Flag{
						stringFlag("code", "Code of the party"),
						stringFlag("name", "Name of the party"),
						stringFlag("address", "Address of the party"),
						stringFlag("address2", "Second address line of the party"),
						stringFlag("nip", "NIP of the party"),
						stringFlag("regon", "Regon of the party"),
					},
				},
				{
					Name:  "modify",
					Usage: "Modify party",
					Action: func(c *cli.Context) error {
						return process(c, validateModifyParty, modifyParty)
					},
					Flags: []cli.Flag{
						stringFlag("code", "Code of the party"),
						stringFlag("name", "Name of the party"),
						stringFlag("address", "Address of the party"),
						stringFlag("address2", "Second address line of the party"),
						stringFlag("nip", "NIP of the party"),
						stringFlag("regon", "Regon of the party"),
						stringFlag("numbering-pattern", "Invoice numbering pattern"),
						stringFlag("active-from", "Start date from which this version is actual"),
					},
				},
			},
		},
		{
			Name:  "generate",
			Usage: "Operations for generating print",
			Subcommands: []cli.Command{
				{
					Name:  "invoice",
					Usage: "Print invoice",
					Action: func(c *cli.Context) error {
						return process(c, validatePrintInvoice, printInvoice)
					},
					Flags: []cli.Flag{
						stringFlag("party", "Code of the party"),
						boolFlag("last", "Generate last invoice created for the party"),
					},
				},
			},
		},
		{
			Name:  "invoice",
			Usage: "Operations on invoices",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "Add new invoice",
					Action: func(c *cli.Context) error {
						return process(c, validateAddInvoice, addInvoice)
					},
					Flags: []cli.Flag{
						stringFlag("buyer", "Code of the buyer"),
						stringFlag("seller", "Name of the seller"),
						stringFlag("issuanceDate", "Date of issuance"),
						stringFlag("issuancePlace", "Place of issuance"),
						stringFlag("sellDate", "Sell date"),
						stringFlag("positionFormat", "Defines input format of the positions"),
						cli.StringSliceFlag{
							Name:  "positions",
							Usage: "Positions declared on invoice",
						},
					},
				},
			},
		},
	}

	// initialize after parsing
	app.Action = func(c *cli.Context) error {
		fmt.Println("Cmd params initialized")
		return nil
	}

	// command not found
	app.CommandNotFound = func(c *cli.Context, s string) {
		fmt.Println("Command not reckognized")
		cli.ShowAppHelpAndExit(c, 1)
	}

	return app
}

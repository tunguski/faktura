package main

import (
	"os"
	"log"
	"fmt"
	"github.com/urfave/cli"
)

type validate func(c *cli.Context) error
type action func(c *cli.Context) error

func process(c *cli.Context, v validate, a action) error {
	if err := v(c); err != nil {
		return err
	} else {
		return a(c)
	}
}

func stringFlag(name string, usage string) cli.StringFlag {
	return cli.StringFlag {
		Name: name,
		Usage: usage,
	}
}

func boolFlag(name string, usage string) cli.BoolFlag {
	return cli.BoolFlag {
		Name: name,
		Usage: usage,
	}
}

func initApp() *cli.App {
	app := cli.NewApp()
	app.Name = "faktura"
	app.Usage = "Create polish invoices in latex style"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Commands = []cli.Command {
		{
			Name: "party",
			Usage: "Operations on parties",
			Subcommands: []cli.Command {
				{
					Name: "add",
					Usage: "Add new party",
					Action: func(c *cli.Context) error {
						return process(c, validateAddParty, addParty)
			        },
					Flags: []cli.Flag {
						stringFlag("code", "Code of the party"),
						stringFlag("name", "Name of the party"),
						stringFlag("address", "Address of the party"),
						stringFlag("address2", "Second address line of the party"),
						stringFlag("nip", "NIP of the party"),
						stringFlag("regon", "Regon of the party"),
					},
				},
			},
		},
		{
			Name: "generate",
			Usage: "Operations for generating print",
			Subcommands: []cli.Command {
				{
					Name: "invoice",
					Usage: "Print invoice",
					Action: func(c *cli.Context) error {
						return process(c, validatePrintInvoice, printInvoice)
			        },
					Flags: []cli.Flag {
						stringFlag("party", "Code of the party"),
						boolFlag("last", "Generate last invoice created for the party"),
					},
				},
			},
		},
		{
			Name: "invoice",
			Usage: "Operations on invoices",
			Subcommands: []cli.Command {
				{
					Name: "add",
					Usage: "Add new invoice",
					Action: func(c *cli.Context) error {
						return process(c, validateAddInvoice, addInvoice)
			        },
					Flags: []cli.Flag {
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

	err := app.Run(os.Args)
	check(err)
	fmt.Println(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	return app
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

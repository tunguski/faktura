package main

import (
	"os"
	"fmt"
	"errors"
	"time"
	"encoding/json"
	"github.com/urfave/cli"
)

func validateAddParty(c *cli.Context) error {
	if c.String("name") == "" || 
		c.String("code") == "" || 
		c.String("address") == "" || 
		c.String("nip") == "" {
		fmt.Println("Code, name, address and nip are required\n")
		cli.ShowCommandHelp(c, "add")
		return errors.New("missing_data")
	}
	return nil
}

func addParty(c *cli.Context) error {
	code := c.String("code")
	party := &Party{
		c.String("name"),
		c.String("nip"),
		c.String("regon"),
		c.String("address"),
		c.String("address2"),
		c.String("bankAccount"),
		time.Now()}

	fmt.Println(code)
	out, _ := json.Marshal(party)
	fmt.Println(string(out))

	var data = readConfig()

	if _, ok := data.Parties[code]; ok {
		fmt.Println("" + code + " already defined")
		os.Exit(1)
	}

	data.Parties[code] = *party

	storeData(".out.toml", data)

	return nil
}

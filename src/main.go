package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	baseDir    string = "."
	configfile string = "config.toml"
	fakturaTex string = "faktura.tex"
	dataFolder string = "data"
	tmpFolder  string = "tmp"
)

func getTmpFolder() string {
	return baseDir + "/" + tmpFolder + "/"
}

func main() {
	app := initApp()
	fmt.Println(app.Version)
}

func loadFile(filename string) []byte {
	_, err := os.Stat(filename)
	var dat = []byte{}
	if err != nil {
		log.Fatal("Config file is missing: ", filename)
	} else {
		dat, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal("Could not read file")
			//} else {
			//	fmt.Println(string(dat))
		}
	}
	return dat
}

// Reads info from config file
func readConfig() *Data {
	loadFile(configfile)

	var config Data
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	if config.Parties == nil {
		config.Parties = make(map[string]Party)
	}

	if config.Invoices == nil {
		config.Invoices = make(map[string][]Invoice)
	}

	return &config
}

func storeData(data *Data) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	message := []byte(buf.String())
	err := ioutil.WriteFile(configfile, message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

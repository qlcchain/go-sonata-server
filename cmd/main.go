/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/qlcchain/go-sonata-server/restapi"
	"github.com/qlcchain/go-sonata-server/restapi/operations"
)

var (
	version string
	date    string
	commit  string
)

const defaultPort = 9999

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		logrus.Fatal(err)
	}

	api := operations.NewSonataAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "MEF LSO Sonata"
	parser.LongDescription = "\nA set of APIs based on the LSO Reference Architecture for\nServiceability (Address Validation, Site Queries, Product Offering Qualification) |\nQuoting | Product Inventory | Ordering | Trouble-ticketing Billing | Contract & Catalog\n\n"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	// set default port
	if server.Port == 0 {
		server.Port = defaultPort
	}
	server.EnabledListeners = []string{"http"}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		logrus.Fatal(err)
	}
}

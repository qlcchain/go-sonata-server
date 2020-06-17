/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/client/product"

	"github.com/qlcchain/go-sonata-server/client"

	httptransport "github.com/go-openapi/runtime/client"

	gs "github.com/qlcchain/go-sonata-server/client/geographic_site"
	po "github.com/qlcchain/go-sonata-server/client/product_order"
	"github.com/qlcchain/go-sonata-server/util"
)

var (
	token = httptransport.BearerToken(`eyJhbGciOiJFUzUxMiIsInR5cCI6IkpXVCJ9.eyJyb2xlcyI6WyJhZG1pbiJdLCJhdWQiOiJRTENDaGFpbiBCb3QiLCJleHAiOjE2MjA3OTQzMTIsImp0aSI6IjE3Y2E4MWU1LTA2YzYtNDZhMi1iMTU2LTdkMzJlMGRjZmIyNSIsImlhdCI6MTU4OTI1ODMxMiwiaXNzIjoiUUxDQ2hhaW4gQm90Iiwic3ViIjoic29uYXRhIn0.AN01SFxQWJ--LOwie9xQSb5IUZUXt6HFuF6kUSSvtUhH0Q5tG7hUU0tc0I-vvJZoRyo8humQ2KwU3-qbp6rbMAS5AHz5tGkja-JI4MT9MYMABLQZcPCy6n9VjKPuxhv2rP74HfJ8fEBC7Vo6JHDMXYOzrfO13i_3oNqQ6aX6miW3tFrg`)
)

func main() {
	var opts struct {
		EndPoint string `short:"e" long:"url" description:"sonata server url" default:"sonata.qlcchain.online"`
		City     string `long:"city" description:"geographicAddress.city" default:"Hong Kong"`
		Country  string `long:"country" description:"geographicAddress.country" default:"China"`
		Schema   string `short:"s" long:"schema" description:"http server schema, http or https" default:"https"`
	}

	if _, err := flags.Parse(&opts); err != nil {
		log.Fatal(err)
	}

	log.SetLevel(log.DebugLevel)

	c := client.New(httptransport.New(opts.EndPoint, "/api/mef", []string{opts.Schema}), strfmt.Default)

	findParams := gs.NewGeographicSiteFindParams()
	findParams.WithGeographicAddressCountry(swag.String(opts.Country)).WithGeographicAddressCity(swag.String(opts.City))
	log.Debugf("STEP1: find product by country(%s) and city(%s)\n", opts.Country, opts.City)
	if resp, err := c.GeographicSite.GeographicSiteFind(findParams, token); err != nil {
		log.Error(err)
	} else {
		sites := resp.Payload
		if len(sites) > 0 {
			log.Infof("find %d site, use the 1st one to search product", len(sites))
			site := sites[0]
			log.Debugf(util.ToIndentString(site))
			productFindParams := product.NewProductFindParams()
			productFindParams.WithGeographicalSiteID(swag.String(site.ID))
			if resp, err := c.Product.ProductFind(productFindParams, token); err == nil {
				summaries := resp.Payload
				if len(summaries) > 0 {
					summary := summaries[0]
					getParams := product.NewProductGetParams()
					getParams.WithProductID(*summary.ID)
					if resp, err := c.Product.ProductGet(getParams, token); err == nil {
						log.Infof("find %d products: use the 1st one, %s", len(sites),
							util.ToIndentString(resp.Payload))
					}
				} else {
					log.Fatal("can not find any product")
				}
			} else {
				log.Fatal(err)
			}
		}
	}

	//getParams := gs.NewGeographicSiteGetParams()
	//getParams.WithSiteID("")
	//if resp, err := c.GeographicSite.GeographicSiteGet(getParams, token); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(resp)
	//}

	params := po.NewProductOrderFindParams()
	if resp, err := c.ProductOrder.ProductOrderFind(params, token); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}

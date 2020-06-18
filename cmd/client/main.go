/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

import (
	"time"

	"github.com/brianvoe/gofakeit/v5"

	"github.com/qlcchain/go-sonata-server/client/quote"
	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"

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
			billingAccount := util.NewIDPtr()
			site := sites[0]
			log.Debugf("GeographicSite payload: %s", util.ToIndentString(site))
			log.Infof("find %d site(s), use the 1st one to search product by geographicalSiteId(%s)", len(sites), site.ID)
			time.Sleep(time.Second)
			productFindParams := product.NewProductFindParams()
			productFindParams.WithGeographicalSiteID(swag.String(site.ID))
			if resp, err := c.Product.ProductFind(productFindParams, token); err == nil {
				summaries := resp.Payload
				if len(summaries) > 0 {
					summary := summaries[0]
					getParams := product.NewProductGetParams()
					getParams.WithProductID(*summary.ID)
					if resp, err := c.Product.ProductGet(getParams, token); err == nil {
						p := resp.Payload
						log.Infof("find %d products: use the 1st one to place order", len(summaries))
						log.Debugf("Product payload: %s", util.ToIndentString(p))
						time.Sleep(time.Second)
						log.Debugf("STEP2: create quote for product %s", swag.StringValue(p.ID))
						quoteCreateParams := quote.NewQuoteCreateParams()
						quoteCreateParams.WithQuote(&models.QuoteCreate{
							//AtBaseType:       "",
							//AtSchemaLocation: "",
							//AtType:           "",
							Agreement: []*models.AgreementRef{
								{
									Href: p.Agreement[0].Path,
									ID:   p.Agreement[0].ID,
									Name: swag.String(p.Agreement[0].Name),
									Path: swag.String(p.Agreement[0].Path),
								},
							},
							Description:                  "auto generated quote",
							ExpectedFulfillmentStartDate: strfmt.Date(time.Now().AddDate(0, 0, 2)),
							ExternalID:                   util.NewID(),
							InstantSyncQuoting:           swag.Bool(false),
							Note: []*models.Note{
								{
									Author: swag.String("Order bot"),
									Date:   handler.NewDatetime(time.Now()),
									Text:   swag.String("auto generated"),
								},
							},
							ProjectID: util.NewID(),
							QuoteItem: []*models.QuoteItemCreate{
								{
									//AtSchemaLocation:       "",
									//AtType:                 "",
									Action: models.ProductActionTypeAdd,
									ID:     util.NewIDPtr(),
									//Note:                   nil,
									Product:         p,
									ProductOffering: &models.ProductOfferingRef{ID: util.NewIDPtr()},
									Qualification: &models.ProductOfferingQualificationRef{
										//AtReferredType:    "",
										//Href:              "",
										ID:                util.NewIDPtr(),
										QualificationItem: util.NewIDPtr(),
									},
									//QuoteItemRelationship:  nil,
									//RelatedParty:           nil,
									RequestedQuoteItemTerm: &models.ItemTerm{
										//Description: "",
										Duration: &models.Duration{
											Unit:  models.DurationUnitMONTH,
											Value: swag.Int32(6),
										},
										//Name:        "",
									},
								},
							},
							QuoteLevel: models.QuoteLevelBUDGETARY,
							RelatedParty: []*models.RelatedParty{
								p.RelatedParty[0],
								{
									AtReferredType:  "Organization",
									EmailAddress:    "hi@pccwglobal.com",
									ID:              util.NewIDPtr(),
									Name:            swag.String("PCCW Global"),
									Number:          gofakeit.Phone(),
									NumberExtension: "",
									Role:            []string{"Buyer"},
								},
							},
							RequestedQuoteCompletionDate: handler.NewDatetime(time.Now().AddDate(0, 0, 1)),
						})
						if resp, err := c.Quote.QuoteCreate(quoteCreateParams, token); err == nil {
							q := resp.Payload
							q.QuoteItem[0].State = models.QuoteItemStateTypeREADY
							q.QuoteItem[0].Product.StatusChange = []*models.StatusChange{}
							q.State = models.QuoteStateTypeACCEPTED
							log.Infof("create quote: id(%s); status(%s)", q.ID, q.State)
							log.Debugf("Quote payload: %s", util.ToIndentString(q))
							time.Sleep(time.Second)

							log.Debugf("STEP3: place order of product(%s), quote(%s)", swag.StringValue(p.ID), q.ID)
							orderCreateParams := po.NewProductOrderCreateParams()
							orderCreateParams.WithProductOrder(&models.ProductOrderCreate{
								//AtBaseType: "",
								//AtSchemaLocation:        "",
								//AtType:                  "",
								BillingAccount: &models.BillingAccountRef{
									ID: billingAccount,
								},
								BuyerRequestDate: handler.NewDatetime(time.Now().Add(time.Hour * 12)),
								DesiredResponse:  models.DesiredOrderResponsesConfirmationAndEngineeringDesign,
								ExpeditePriority: false,
								ExternalID:       util.NewIDPtr(),
								//Note:             nil,
								OrderActivity: models.OrderActivityInstall,
								OrderItem: []*models.ProductOrderItemCreate{
									{
										//AtSchemaLocation:      "",
										//AtType:                "",
										Action: models.ProductActionTypeAdd,
										BillingAccount: &models.BillingAccountRef{
											ID: billingAccount,
										},
										ID:             util.NewIDPtr(),
										OrderItemPrice: nil,
										//OrderItemRelationship: nil,
										PricingMethod: models.PricingMethodTariff,
										//PricingReference:      "",
										PricingTerm:     swag.Int32(6),
										Product:         p,
										ProductOffering: &models.ProductOfferingRef{ID: util.NewIDPtr()},
										//Qualification:   nil,
										Quote: &models.QuoteRef{
											//AtReferredType: "",
											//Href:           "",
											ID:        util.NewIDPtr(),
											QuoteItem: q.ID,
										},
										RelatedParty: q.RelatedParty,
									},
								},
								OrderVersion: swag.String("01"),
								//PricingMethod:    models.PricingMethodTariff,
								//PricingReference: util.NewID(),
								//PricingTerm:      0,
								//Priority:         0,
								ProjectID: util.NewID(),
								//RelatedBuyerPON:         "",
								RelatedParty:            q.RelatedParty,
								RequestedCompletionDate: handler.NewDatetime(time.Now().AddDate(0, 0, 1)),
								RequestedStartDate:      strfmt.DateTime(time.Now()),
								//TspRestorationPriority:  "",
							})
							if resp, err := c.ProductOrder.ProductOrderCreate(orderCreateParams, token); err == nil {
								order := resp.Payload
								log.Infof("create order, id(%s)", swag.StringValue(order.ID))
								order.StateChange = append(order.StateChange, &models.StateChange{
									ChangeDate:   strfmt.DateTime(time.Now()),
									ChangeReason: "placed",
									State:        models.ProductOfferingQualificationStateTypeDone,
								})
								log.Debugf("ProductOrder payload: %s", util.ToIndentString(order))
								time.Sleep(time.Second)
							} else {
								log.Fatal(err)
							}
						} else {
							log.Fatal(err)
						}
					}
				} else {
					log.Fatal("can not find any product")
				}
			} else {
				log.Fatal(err)
			}
		}
	}
}

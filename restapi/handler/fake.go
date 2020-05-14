/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package handler

import (
	"fmt"
	"time"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-sonata-server/models"
)

func FieldAddress() *models.FieldedAddress {
	return &models.FieldedAddress{
		City:    gofakeit.City(),
		Country: gofakeit.Country(),
		GeographicSubAddress: []*models.GeographicSubAddress{
			{
				//AtSchemaLocation:    "",
				//AtType:              "",
				BuildingName:        "Tower01",
				ID:                  util.NewID(),
				LevelNumber:         "BASEMENT 1",
				LevelType:           "1",
				PrivateStreetName:   "university",
				PrivateStreetNumber: "01",
				SubUnit: []*models.SubUnit{
					{
						SubUnitIdentifier: swag.String(fmt.Sprintf("%d", gofakeit.Uint32())),
						SubUnitType:       swag.String("UNIT"),
					},
				},
			},
		},
		ID:                 util.NewID(),
		Locality:           "Locality",
		PostCodeExtension:  "",
		Postcode:           gofakeit.Zip(),
		StateOrProvince:    gofakeit.State(),
		StreetName:         gofakeit.StreetName(),
		StreetNr:           gofakeit.StreetNumber(),
		StreetNrLast:       "",
		StreetNrLastSuffix: "",
		StreetNrSuffix:     "",
		StreetSuffix:       gofakeit.StreetSuffix(),
		StreetType:         "Alley",
	}
}

func FormattedAddress(a *models.FieldedAddress) *models.FormattedAddress {
	if a == nil {
		a = FieldAddress()
	}
	return &models.FormattedAddress{
		AddrLine1: swag.String(fmt.Sprintf("%s %s %s %s, %s %s", a.StreetNr, a.StreetNrLast, a.StreetSuffix,
			a.StreetNrLastSuffix, a.StreetName, a.StreetSuffix)),
		AddrLine2:         "",
		City:              a.City,
		Country:           a.Country,
		Locality:          a.Locality,
		PostCodeExtension: a.PostCodeExtension,
		Postcode:          a.Postcode,
		StateOrProvince:   a.StateOrProvince,
	}
}

func GeographicSiteAndAddress() (*models.GeographicAddress, *models.GeographicSite) {
	a := GeographicAddress()
	s := GeographicSite()
	s.FieldedAddress = a.FieldedAddress
	s.FormattedAddress = a.FormattedAddress
	s.ReferencedAddress = a.ReferencedAddress
	return a, s
}

func GeographicLocation() *models.GeographicLocation {
	return &models.GeographicLocation{
		ID: util.NewID(),
		GeographicPoint: []*models.GeographicPoint{
			{
				ID:        util.NewID(),
				Latitude:  swag.String(fmt.Sprintf("%f", gofakeit.Latitude())),
				Longitude: swag.String(fmt.Sprintf("%f", gofakeit.Longitude())),
			},
		},
		SpatialRef: swag.String(gofakeit.CountryAbr()),
	}
}

func ReferencedAddress(refID *string) *models.ReferencedAddress {
	return &models.ReferencedAddress{
		ID:            util.NewID(),
		ReferenceID:   util.NewOrDefaultPtr(refID),
		ReferenceType: swag.String("refer"),
	}
}

func GeographicAddress() *models.GeographicAddress {
	a := FieldAddress()
	return &models.GeographicAddress{
		//AtSchemaLocation: "",
		//AtType:           "",
		AllowsNewSite:      false,
		FieldedAddress:     a,
		FormattedAddress:   FormattedAddress(a),
		GeographicLocation: GeographicLocation(),
		HasPublicSite:      false,
		ID:                 util.NewID(),
		ReferencedAddress:  ReferencedAddress(nil),
	}
}

func Note() *models.Note {
	return &models.Note{
		Author: swag.String(gofakeit.Name()),
		Date:   NewDatetime(time.Now()),
		Text:   swag.String(gofakeit.Phrase()),
	}
}

func ItemTerm() *models.ItemTerm {
	return &models.ItemTerm{
		Description: gofakeit.JobDescriptor(),
		Duration: &models.Duration{
			Unit:  models.DurationUnitDAY,
			Value: swag.Int32(gofakeit.Int32()),
		},
		Name: gofakeit.Name(),
	}
}

func RelatedParty() *models.RelatedParty {
	return &models.RelatedParty{
		AtReferredType:  "Organization",
		EmailAddress:    gofakeit.Email(),
		ID:              util.NewIDPtr(),
		Name:            swag.String(gofakeit.Name()),
		Number:          gofakeit.Phone(),
		NumberExtension: "",
		Role: []string{
			"Buyer", "Seller", "Site Contact",
		},
	}
}

func Describing() *models.Describing {
	return &models.Describing{
		AtSchemaLocation: swag.String("{MEF Product Spec Storage Location}/ElineSpec_v2.json"),
		AtType:           swag.String("E-Line Spec"),
	}
}

func Agreement() *models.Agreement {
	return &models.Agreement{
		ID:   util.NewIDPtr(),
		Name: gofakeit.Name(),
		Path: gofakeit.URL(),
	}
}

func ProductPrice() *models.ProductPrice {
	return &models.ProductPrice{
		//AtType:      "",
		Description: gofakeit.Phrase(),
		Name:        swag.String("Subscription price"),
		Price: &models.Price{
			//AtType:            "",
			DutyFreeAmount: &models.Money{
				Unit:  swag.String("USD"),
				Value: swag.Float32(gofakeit.Float32()),
			},
			TaxIncludedAmount: &models.Money{
				Unit:  swag.String("USD"),
				Value: swag.Float32(gofakeit.Float32()),
			},
			TaxRate: 0.01,
		},
		PriceType:             models.PriceTypeRecurring,
		RecurringChargePeriod: models.ChargePeriodMonth,
		UnitOfMeasure:         "Gb",
	}
}

func Product() *models.Product {
	return &models.Product{
		//AtBaseType:       "",
		//AtSchemaLocation: "",
		AtType: "E-Line Spec",
		Agreement: []*models.Agreement{
			Agreement(),
		},
		BillingAccount: []*models.BillingAccountRef{
			{
				ID: util.NewIDPtr(),
			},
		},
		BuyerProductID: util.NewID(),
		Href:           gofakeit.URL(),
		ID:             util.NewIDPtr(),
		LastUpdateDate: strfmt.DateTime(time.Now()),
		ProductOffering: &models.ProductOfferingRef{
			ID: util.NewIDPtr(),
		},
		ProductOrder: []*models.ProductOrderRef{
			{
				Href:        gofakeit.URL(),
				ID:          util.NewIDPtr(),
				OrderItemID: swag.String("01"),
			},
		},
		ProductPrice: []*models.ProductPrice{
			ProductPrice(),
		},
		ProductRelationship: []*models.ProductRelationship{
			{
				Product: &models.ProductRef{
					BuyerProductID: util.NewID(),
					Href:           gofakeit.URL(),
					ID:             util.NewIDPtr(),
				},
				Type: swag.String("bundled"),
			},
		},
		ProductSpecification: &models.ProductSpecificationRef{
			Describing: Describing(),
			ID:         util.NewIDPtr(),
		},
		ProductTerm: []*models.ProductTerm{
			{
				Description: gofakeit.JobDescriptor(),
				Duration: &models.Quantity{
					Amount: gofakeit.Float32(),
					Units:  "day",
				},
				Name:     gofakeit.Name(),
				ValidFor: TimePeriod(),
			},
		},
		RelatedParty: []*models.RelatedParty{
			RelatedParty(),
		},
		Site: []*models.GeographicSite{
			GeographicSite(),
		},
		StartDate: NewDatetime(time.Now().AddDate(-1, 0, 0)),
		Status:    models.ProductStatusActive,
		StatusChange: []*models.StatusChange{
			StatusChange(),
		},
		TerminationDate: strfmt.DateTime(time.Now().AddDate(10, 0, 0)),
	}
}

func GeographicSite() *models.GeographicSite {
	fa := FieldAddress()
	return &models.GeographicSite{
		AtSchemaLocation:          "",
		AtType:                    "",
		AdditionalSiteInformation: gofakeit.BuzzWord(),
		Description:               gofakeit.JobDescriptor(),
		FieldedAddress:            fa,
		FormattedAddress:          FormattedAddress(fa),
		GeographicLocation:        GeographicLocation(),
		ID:                        util.NewID(),
		ReferencedAddress: &models.ReferencedAddress{
			ID:            util.NewID(),
			ReferenceID:   util.NewIDPtr(),
			ReferenceType: swag.String(gofakeit.RandomString([]string{"01", "02"})),
		},
		RelatedParty: []*models.RelatedParty{
			RelatedParty(),
		},
		SiteCompanyName:  gofakeit.Company(),
		SiteCustomerName: gofakeit.Name(),
		SiteName:         gofakeit.Name(),
		SiteType:         "PUBLIC",
		Status:           models.StatusPlanned,
	}
}

func TimePeriod() *models.TimePeriod {
	t := gofakeit.Date()
	return &models.TimePeriod{
		EndDateTime:   strfmt.DateTime(t.AddDate(1, 0, 0)),
		StartDateTime: strfmt.DateTime(t),
	}
}

func StatusChange() *models.StatusChange {
	return &models.StatusChange{
		ChangeDate:   strfmt.DateTime(gofakeit.Date()),
		ChangeReason: gofakeit.Quote(),
	}
}

func QuoteCreate() *models.QuoteCreate {
	return &models.QuoteCreate{
		//AtBaseType:                   "",
		//AtSchemaLocation:             "",
		AtType: "",
		Agreement: []*models.AgreementRef{
			{
				Href: gofakeit.URL(),
				ID:   util.NewIDPtr(),
				Name: swag.String(gofakeit.Name()),
				Path: swag.String(gofakeit.URL()),
			},
		},
		Description:                  "demo create quote",
		ExpectedFulfillmentStartDate: strfmt.Date(time.Now().AddDate(0, 0, 5)),
		ExternalID:                   util.NewID(),
		InstantSyncQuoting:           swag.Bool(false),
		Note:                         []*models.Note{Note()},
		ProjectID:                    util.NewID(),
		QuoteItem: []*models.QuoteItemCreate{
			{
				//AtSchemaLocation:       "",
				//AtType:                 "",
				Action: models.ProductActionTypeAdd,
				ID:     util.NewIDPtr(),
				Note: []*models.Note{
					Note(),
				},
				Product: Product(),
				ProductOffering: &models.ProductOfferingRef{
					ID: util.NewIDPtr(),
				},
				Qualification: &models.ProductOfferingQualificationRef{
					//AtReferredType:    "",
					Href:              gofakeit.URL(),
					ID:                util.NewIDPtr(),
					QualificationItem: util.NewIDPtr(),
				},
				QuoteItemRelationship: []*models.QuoteItemRelationship{
					{
						ID:   util.NewIDPtr(),
						Type: models.RelationshipTypeBundled,
					},
				},
				RelatedParty: []*models.RelatedParty{
					RelatedParty(),
				},
				RequestedQuoteItemTerm: ItemTerm(),
			},
		},
		QuoteLevel: models.QuoteLevelFIRM,
		RelatedParty: []*models.RelatedParty{
			RelatedParty(),
		},
		RequestedQuoteCompletionDate: NewDatetime(time.Now().AddDate(0, 0, 3)),
	}
}

func ProductOfferingQualificationCreate() *models.ProductOfferingQualificationCreate {
	return &models.ProductOfferingQualificationCreate{
		AtSchemaLocation:         "",
		AtType:                   "",
		InstantSyncQualification: swag.Bool(false),
		ProductOfferingQualificationItem: []*models.ProductOfferingQualificationItemCreate{
			{
				AtType: "E-Line Spec",
				Action: models.ProductActionTypeAdd,
				ID:     util.NewIDPtr(),
				Product: &models.Product{
					AtType: "E-Line Spec",
					Agreement: []*models.Agreement{
						Agreement(),
					},
					BillingAccount: []*models.BillingAccountRef{
						{
							ID: util.NewIDPtr(),
						},
					},
					BuyerProductID: util.NewID(),
					Href:           gofakeit.URL(),
					ID:             util.NewIDPtr(),
					LastUpdateDate: strfmt.DateTime(time.Now()),
					ProductOffering: &models.ProductOfferingRef{
						ID: util.NewIDPtr(),
					},
					//TODO: fill it
					//ProductOrder: []*models.ProductOrderRef{
					//	{
					//		Href:        gofakeit.URL(),
					//		ID:          handler.NewID(),
					//		OrderItemID: swag.String("01"),
					//	},
					//},
					//ProductPrice: []*models.ProductPrice{
					//	ProductPrice(),
					//},
					ProductRelationship: []*models.ProductRelationship{
						{
							Product: &models.ProductRef{
								BuyerProductID: "",
								Href:           gofakeit.URL(),
								ID:             util.NewIDPtr(),
							},
							Type: swag.String("bundled"),
						},
					},
					ProductSpecification: &models.ProductSpecificationRef{
						Describing: Describing(),
						ID:         util.NewIDPtr(),
					},
					ProductTerm: []*models.ProductTerm{
						{
							Description: gofakeit.JobDescriptor(),
							Duration: &models.Quantity{
								Amount: gofakeit.Float32(),
								Units:  "day",
							},
							Name:     gofakeit.Name(),
							ValidFor: TimePeriod(),
						},
					},
					RelatedParty: []*models.RelatedParty{
						RelatedParty(),
					},
					Site:      []*models.GeographicSite{},
					StartDate: NewDatetime(time.Now().AddDate(-1, 0, 0)),
					Status:    models.ProductStatusActive,
					StatusChange: []*models.StatusChange{
						StatusChange(),
					},
					TerminationDate: strfmt.DateTime(time.Now().AddDate(10, 0, 0)),
				},
				ProductOffering: &models.ProductOfferingRef{
					ID: util.NewIDPtr(),
				},
				ProductOfferingQualificationItemRelationship: []*models.ProductOfferingQualificationItemRelationship{
					{
						ID:   util.NewIDPtr(),
						Type: models.RelationshipTypeBundled,
					},
				},
				RelatedParty: []*models.RelatedParty{
					RelatedParty(),
				},
			},
		},
		ProjectID:          util.NewID(),
		ProvideAlternative: swag.Bool(false),
		RelatedParty: []*models.RelatedParty{
			RelatedParty(),
		},
		RequestedResponseDate: strfmt.Date(time.Now().AddDate(0, 0, 2)),
	}
}

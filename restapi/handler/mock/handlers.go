/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"encoding/json"

	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"

	"github.com/qlcchain/go-sonata-server/restapi/operations"
	"github.com/qlcchain/go-sonata-server/restapi/operations/cancel_product_order"
	ga "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address"
	gav "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address_validation"
	gs "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	"github.com/qlcchain/go-sonata-server/restapi/operations/notification"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product"
	poq "github.com/qlcchain/go-sonata-server/restapi/operations/product_offering_qualification"
	po "github.com/qlcchain/go-sonata-server/restapi/operations/product_order"
	"github.com/qlcchain/go-sonata-server/restapi/operations/quote"
)

const (
	addresses = `[
  {
    "fieldedAddress": {
      "city": "West Beahan",
      "country": "Netherlands Antilles",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcq50",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "1237770961",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcq5g",
      "locality": "Locality",
      "postcode": "50604",
      "stateOrProvince": "Idaho",
      "streetName": "Causeway",
      "streetNr": "841",
      "streetSuffix": "fort",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "841  fort , Causeway fort",
      "city": "West Beahan",
      "country": "Netherlands Antilles",
      "locality": "Locality",
      "postcode": "50604",
      "stateOrProvince": "Idaho"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcq6g",
          "latitude": "32.235241",
          "longitude": "-101.320902"
        }
      ],
      "id": "bqucsgug10l34acgcq60",
      "spatialRef": "JO"
    },
    "id": "bqucsgug10l34acgcq70",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcq7g",
      "referenceId": "bqucsgug10l34acgcq80",
      "referenceType": "refer"
    }
  },
  {
    "fieldedAddress": {
      "city": "West Spencer",
      "country": "Niue",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcqcg",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "377280272",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcqd0",
      "locality": "Locality",
      "postcode": "74736",
      "stateOrProvince": "Massachusetts",
      "streetName": "Orchard",
      "streetNr": "73193",
      "streetSuffix": "burgh",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "73193  burgh , Orchard burgh",
      "city": "West Spencer",
      "country": "Niue",
      "locality": "Locality",
      "postcode": "74736",
      "stateOrProvince": "Massachusetts"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcqe0",
          "latitude": "-74.606309",
          "longitude": "61.047107"
        }
      ],
      "id": "bqucsgug10l34acgcqdg",
      "spatialRef": "SI"
    },
    "id": "bqucsgug10l34acgcqeg",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcqf0",
      "referenceId": "bqucsgug10l34acgcqfg",
      "referenceType": "refer"
    }
  },
  {
    "fieldedAddress": {
      "city": "South Ullrich",
      "country": "Slovakia (Slovak Republic)",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcqk0",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "1658235342",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcqkg",
      "locality": "Locality",
      "postcode": "26906",
      "stateOrProvince": "Alaska",
      "streetName": "Locks",
      "streetNr": "1445",
      "streetSuffix": "land",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "1445  land , Locks land",
      "city": "South Ullrich",
      "country": "Slovakia (Slovak Republic)",
      "locality": "Locality",
      "postcode": "26906",
      "stateOrProvince": "Alaska"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcqlg",
          "latitude": "50.946295",
          "longitude": "-38.429641"
        }
      ],
      "id": "bqucsgug10l34acgcql0",
      "spatialRef": "TZ"
    },
    "id": "bqucsgug10l34acgcqm0",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcqmg",
      "referenceId": "bqucsgug10l34acgcqn0",
      "referenceType": "refer"
    }
  },
  {
    "fieldedAddress": {
      "city": "West Beier",
      "country": "Saint Martin",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcqrg",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "473656281",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcqs0",
      "locality": "Locality",
      "postcode": "13886",
      "stateOrProvince": "Maryland",
      "streetName": "Station",
      "streetNr": "2513",
      "streetSuffix": "shire",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "2513  shire , Station shire",
      "city": "West Beier",
      "country": "Saint Martin",
      "locality": "Locality",
      "postcode": "13886",
      "stateOrProvince": "Maryland"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcqt0",
          "latitude": "26.106995",
          "longitude": "-25.672166"
        }
      ],
      "id": "bqucsgug10l34acgcqsg",
      "spatialRef": "GL"
    },
    "id": "bqucsgug10l34acgcqtg",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcqu0",
      "referenceId": "bqucsgug10l34acgcqug",
      "referenceType": "refer"
    }
  },
  {
    "fieldedAddress": {
      "city": "Winfieldshire",
      "country": "San Marino",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcr30",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "487793581",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcr3g",
      "locality": "Locality",
      "postcode": "22171",
      "stateOrProvince": "North Carolina",
      "streetName": "Brook",
      "streetNr": "911",
      "streetSuffix": "port",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "911  port , Brook port",
      "city": "Winfieldshire",
      "country": "San Marino",
      "locality": "Locality",
      "postcode": "22171",
      "stateOrProvince": "North Carolina"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcr4g",
          "latitude": "-27.564502",
          "longitude": "-102.724663"
        }
      ],
      "id": "bqucsgug10l34acgcr40",
      "spatialRef": "ZA"
    },
    "id": "bqucsgug10l34acgcr50",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcr5g",
      "referenceId": "bqucsgug10l34acgcr60",
      "referenceType": "refer"
    }
  }
]`
	sites = `[
  {
    "additionalSiteInformation": "Phased",
    "description": "Legacy",
    "fieldedAddress": {
      "city": "West Beahan",
      "country": "Netherlands Antilles",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcq50",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "1237770961",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcq5g",
      "locality": "Locality",
      "postcode": "50604",
      "stateOrProvince": "Idaho",
      "streetName": "Causeway",
      "streetNr": "841",
      "streetSuffix": "fort",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "841  fort , Causeway fort",
      "city": "West Beahan",
      "country": "Netherlands Antilles",
      "locality": "Locality",
      "postcode": "50604",
      "stateOrProvince": "Idaho"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcqa0",
          "latitude": "-39.868629",
          "longitude": "-27.665208"
        }
      ],
      "id": "bqucsgug10l34acgcq9g",
      "spatialRef": "AU"
    },
    "id": "bqucsgug10l34acgcqag",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcq7g",
      "referenceId": "bqucsgug10l34acgcq80",
      "referenceType": "refer"
    },
    "relatedParty": [
      {
        "@referredType": "Organization",
        "emailAddress": "gilbertokuneva@hartmann.io",
        "id": "bqucsgug10l34acgcqc0",
        "name": "Brody Walker",
        "number": "9058835105",
        "role": [
          "Buyer",
          "Seller",
          "Site Contact"
        ]
      }
    ],
    "siteCompanyName": "Verdafero",
    "siteCustomerName": "Toy Lueilwitz",
    "siteName": "Brandy Grady",
    "siteType": "PUBLIC",
    "status": "planned"
  },
  {
    "additionalSiteInformation": "modular",
    "description": "Customer",
    "fieldedAddress": {
      "city": "West Spencer",
      "country": "Niue",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcqcg",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "377280272",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcqd0",
      "locality": "Locality",
      "postcode": "74736",
      "stateOrProvince": "Massachusetts",
      "streetName": "Orchard",
      "streetNr": "73193",
      "streetSuffix": "burgh",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "73193  burgh , Orchard burgh",
      "city": "West Spencer",
      "country": "Niue",
      "locality": "Locality",
      "postcode": "74736",
      "stateOrProvince": "Massachusetts"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcqhg",
          "latitude": "-84.659170",
          "longitude": "-179.314598"
        }
      ],
      "id": "bqucsgug10l34acgcqh0",
      "spatialRef": "MV"
    },
    "id": "bqucsgug10l34acgcqi0",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcqf0",
      "referenceId": "bqucsgug10l34acgcqfg",
      "referenceType": "refer"
    },
    "relatedParty": [
      {
        "@referredType": "Organization",
        "emailAddress": "santospowlowski@daugherty.name",
        "id": "bqucsgug10l34acgcqjg",
        "name": "Wallace Bins",
        "number": "1960369125",
        "role": [
          "Buyer",
          "Seller",
          "Site Contact"
        ]
      }
    ],
    "siteCompanyName": "Quertle",
    "siteCustomerName": "Godfrey Emmerich",
    "siteName": "Yvette Hansen",
    "siteType": "PUBLIC",
    "status": "planned"
  },
  {
    "additionalSiteInformation": "radical",
    "description": "International",
    "fieldedAddress": {
      "city": "South Ullrich",
      "country": "Slovakia (Slovak Republic)",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcqk0",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "1658235342",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcqkg",
      "locality": "Locality",
      "postcode": "26906",
      "stateOrProvince": "Alaska",
      "streetName": "Locks",
      "streetNr": "1445",
      "streetSuffix": "land",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "1445  land , Locks land",
      "city": "South Ullrich",
      "country": "Slovakia (Slovak Republic)",
      "locality": "Locality",
      "postcode": "26906",
      "stateOrProvince": "Alaska"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcqp0",
          "latitude": "-87.691337",
          "longitude": "-168.954410"
        }
      ],
      "id": "bqucsgug10l34acgcqog",
      "spatialRef": "MU"
    },
    "id": "bqucsgug10l34acgcqpg",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcqmg",
      "referenceId": "bqucsgug10l34acgcqn0",
      "referenceType": "refer"
    },
    "relatedParty": [
      {
        "@referredType": "Organization",
        "emailAddress": "missourimoen@aufderhar.net",
        "id": "bqucsgug10l34acgcqr0",
        "name": "Ford Blanda",
        "number": "5290865606",
        "role": [
          "Buyer",
          "Seller",
          "Site Contact"
        ]
      }
    ],
    "siteCompanyName": "DataLogix",
    "siteCustomerName": "Sadie Douglas",
    "siteName": "Dax Kiehn",
    "siteType": "PUBLIC",
    "status": "planned"
  },
  {
    "additionalSiteInformation": "Ameliorated",
    "description": "Future",
    "fieldedAddress": {
      "city": "West Beier",
      "country": "Saint Martin",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcqrg",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "473656281",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcqs0",
      "locality": "Locality",
      "postcode": "13886",
      "stateOrProvince": "Maryland",
      "streetName": "Station",
      "streetNr": "2513",
      "streetSuffix": "shire",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "2513  shire , Station shire",
      "city": "West Beier",
      "country": "Saint Martin",
      "locality": "Locality",
      "postcode": "13886",
      "stateOrProvince": "Maryland"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcr0g",
          "latitude": "15.090253",
          "longitude": "-155.247770"
        }
      ],
      "id": "bqucsgug10l34acgcr00",
      "spatialRef": "GB"
    },
    "id": "bqucsgug10l34acgcr10",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcqu0",
      "referenceId": "bqucsgug10l34acgcqug",
      "referenceType": "refer"
    },
    "relatedParty": [
      {
        "@referredType": "Organization",
        "emailAddress": "kristianpowlowski@christiansen.info",
        "id": "bqucsgug10l34acgcr2g",
        "name": "Breanna Windler",
        "number": "1435310490",
        "role": [
          "Buyer",
          "Seller",
          "Site Contact"
        ]
      }
    ],
    "siteCompanyName": "VisualDoD, LLC",
    "siteCustomerName": "Florida Towne",
    "siteName": "Herta Kirlin",
    "siteType": "PUBLIC",
    "status": "planned"
  },
  {
    "additionalSiteInformation": "Polarised",
    "description": "National",
    "fieldedAddress": {
      "city": "Winfieldshire",
      "country": "San Marino",
      "geographicSubAddress": [
        {
          "buildingName": "Tower01",
          "id": "bqucsgug10l34acgcr30",
          "levelNumber": "BASEMENT 1",
          "levelType": "1",
          "privateStreetName": "university",
          "privateStreetNumber": "01",
          "subUnit": [
            {
              "subUnitIdentifier": "487793581",
              "subUnitType": "UNIT"
            }
          ]
        }
      ],
      "id": "bqucsgug10l34acgcr3g",
      "locality": "Locality",
      "postcode": "22171",
      "stateOrProvince": "North Carolina",
      "streetName": "Brook",
      "streetNr": "911",
      "streetSuffix": "port",
      "streetType": "Alley"
    },
    "formattedAddress": {
      "addrLine1": "911  port , Brook port",
      "city": "Winfieldshire",
      "country": "San Marino",
      "locality": "Locality",
      "postcode": "22171",
      "stateOrProvince": "North Carolina"
    },
    "geographicLocation": {
      "geographicPoint": [
        {
          "id": "bqucsgug10l34acgcr80",
          "latitude": "25.355667",
          "longitude": "163.173151"
        }
      ],
      "id": "bqucsgug10l34acgcr7g",
      "spatialRef": "MV"
    },
    "id": "bqucsgug10l34acgcr8g",
    "referencedAddress": {
      "id": "bqucsgug10l34acgcr5g",
      "referenceId": "bqucsgug10l34acgcr60",
      "referenceType": "refer"
    },
    "relatedParty": [
      {
        "@referredType": "Organization",
        "emailAddress": "alberthahane@pagac.net",
        "id": "bqucsgug10l34acgcra0",
        "name": "Trystan Ebert",
        "number": "6726152987",
        "role": [
          "Buyer",
          "Seller",
          "Site Contact"
        ]
      }
    ],
    "siteCompanyName": "TuvaLabs",
    "siteCustomerName": "Elnora Walter",
    "siteName": "Mariela Armstrong",
    "siteType": "PUBLIC",
    "status": "planned"
  }
]
`
)

var (
	Store *gorm.DB
)

func init() {
	var err error
	// prepare db
	//dir := config.DBDir()
	//Store, err = gorm.Open(sqlite.Open("file:mockdb?mode=memory&cache=shared"), &gorm.Config{
	//	SkipDefaultTransaction: false,
	//	NamingStrategy:         nil,
	//	Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//		SlowThreshold: 100 * time.Millisecond,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	}),
	//})
	//file := path.Join(dir, "sonata.db")
	file := "file:mockdb?mode=memory&cache=shared"
	Store, err = gorm.Open("sqlite3", file)
	if err != nil {
		log.Fatalln(err)
	}
	Store.LogMode(true)
	if err := Store.AutoMigrate(
		// user management
		&schema.User{}, &schema.HubSubscriber{},
		// GeographicAddress
		&schema.GeographicAddress{}, &schema.FieldedAddress{}, &schema.FormattedAddress{},
		&schema.GeographicLocation{}, &schema.GeographicAddress{}, &schema.GeographicSubAddress{}, &schema.GeographicPoint{},
		&schema.SubUnit{}, &models.ReferencedAddress{},
		// GeographicSite
		&schema.GeographicSite{}, &schema.FormattedAddress{}, &schema.FieldedAddress{},
		&schema.GeographicSubAddress{}, &schema.SubUnit{}, &schema.GeographicLocation{}, &schema.GeographicPoint{},
		&models.ReferencedAddress{}, &schema.GeographicSite{},
		// Product
		&schema.Product{}, &models.Agreement{}, &schema.ProductRelationship{}, &models.ProductRef{}, &schema.ProductSpecificationRef{},
		&schema.ProductTerm{}, &models.ProductOfferingRef{}, &models.ProductOfferingQualificationItemRelationship{}, &schema.ProductOrderRef{},
		&schema.ProductPrice{}, &models.TerminationError{}, &models.ProductOfferingQualificationRef{},
		// Order
		// POQ
		&schema.ProductOfferingQualification{}, &schema.ProductOfferingQualificationItem{}, &schema.StatusChange{}, &schema.StateChange{},
		&models.BillingAccountRef{},
		&schema.RelatedParty{}, &schema.AlternateProductProposal{},
		// Quote
		&schema.Quote{}, &models.AgreementRef{}, &schema.Note{}, &schema.QuoteItem{}, &models.QuoteItemRelationship{},
	).Error; err != nil {
		log.Fatalln(err)
	}
}

func mockData() error {
	//mock address
	var address []*schema.GeographicAddress
	if err := json.Unmarshal([]byte(addresses), &address); err != nil {
		return err
	} else {
		for _, geographicAddress := range address {
			if err := Store.Create(geographicAddress).Error; err != nil {
				log.Error(err)
			}
		}
	}
	//mock site
	var site []*schema.GeographicSite
	if err := json.Unmarshal([]byte(sites), &site); err != nil {
		return err
	} else {
		for _, geographicSite := range site {
			if err := Store.Create(geographicSite).Error; err != nil {
				log.Error(err)
			}
		}
	}

	return nil
}

func Bind(api *operations.SonataAPI) {
	if err := mockData(); err != nil {
		log.Fatal(err)
	}

	// GeographicAddress
	api.GeographicAddressGeographicAddressGetHandler = ga.GeographicAddressGetHandlerFunc(GeographicAddressGeographicAddressGetHandler)
	api.GeographicAddressValidationGeographicAddressValidationCreateHandler = gav.GeographicAddressValidationCreateHandlerFunc(GeographicAddressValidationGeographicAddressValidationCreateHandler)

	// GeographicSite
	api.GeographicSiteGeographicSiteFindHandler = gs.GeographicSiteFindHandlerFunc(GeographicSiteGeographicSiteFindHandler)
	api.GeographicSiteGeographicSiteGetHandler = gs.GeographicSiteGetHandlerFunc(GeographicSiteGeographicSiteGetHandler)

	api.CancelProductOrderCancelProductOrderCreateHandler = cancel_product_order.CancelProductOrderCreateHandlerFunc(CancelProductOrderCancelProductOrderCreateHandler)
	api.CancelProductOrderCancelProductOrderFindHandler = cancel_product_order.CancelProductOrderFindHandlerFunc(CancelProductOrderCancelProductOrderFindHandler)
	api.CancelProductOrderCancelProductOrderGetHandler = cancel_product_order.CancelProductOrderGetHandlerFunc(CancelProductOrderCancelProductOrderGetHandler)

	// ProductOrder
	api.ProductOrderProductOrderCreateHandler = po.ProductOrderCreateHandlerFunc(ProductOrderProductOrderCreateHandler)
	api.ProductOrderProductOrderFindHandler = po.ProductOrderFindHandlerFunc(ProductOrderProductOrderFindHandler)
	api.ProductOrderProductOrderGetHandler = po.ProductOrderGetHandlerFunc(ProductOrderProductOrderGetHandler)
	// ProductOrder hub
	api.HubProductOrderManagementHubCreateHandler = hub.ProductOrderManagementHubCreateHandlerFunc(HubProductOrderManagementHubCreateHandler)
	api.HubProductOrderManagementHubDeleteHandler = hub.ProductOrderManagementHubDeleteHandlerFunc(HubProductOrderManagementHubDeleteHandler)
	api.HubProductOrderManagementHubFindHandler = hub.ProductOrderManagementHubFindHandlerFunc(HubProductOrderManagementHubFindHandler)
	// ProductOrder Notification
	api.NotificationNotificationProductOrderAttributeValueChangeNotificationHandler = notification.NotificationProductOrderAttributeValueChangeNotificationHandlerFunc(NotificationNotificationProductOrderAttributeValueChangeNotificationHandler)
	api.NotificationNotificationProductOrderCreationNotificationHandler = notification.NotificationProductOrderCreationNotificationHandlerFunc(NotificationNotificationProductOrderCreationNotificationHandler)
	api.NotificationNotificationProductOrderInformationRequiredNotificationHandler = notification.NotificationProductOrderInformationRequiredNotificationHandlerFunc(NotificationNotificationProductOrderInformationRequiredNotificationHandler)
	api.NotificationNotificationProductOrderStateChangeNotificationHandler = notification.NotificationProductOrderStateChangeNotificationHandlerFunc(NotificationNotificationProductOrderStateChangeNotificationHandler)

	// Quote
	api.QuoteQuoteCreateHandler = quote.QuoteCreateHandlerFunc(QuoteQuoteCreateHandler)
	api.QuoteQuoteFindHandler = quote.QuoteFindHandlerFunc(QuoteQuoteFindHandler)
	api.QuoteQuoteGetHandler = quote.QuoteGetHandlerFunc(QuoteQuoteGetHandler)
	api.QuoteQuoteRequestStateChangeHandler = quote.QuoteRequestStateChangeHandlerFunc(QuoteQuoteRequestStateChangeHandler)
	api.HubQuoteManagementHubDeleteHandler = hub.QuoteManagementHubDeleteHandlerFunc(HubQuoteManagementHubDeleteHandler)
	// Quote Hub
	api.HubQuoteManagementHubCreateHandler = hub.QuoteManagementHubCreateHandlerFunc(HubQuoteManagementHubCreateHandler)
	api.HubQuoteManagementHubFindHandler = hub.QuoteManagementHubFindHandlerFunc(HubQuoteManagementHubFindHandler)
	// Quote Notification
	api.NotificationNotificationQuoteAttributeValueChangeNotificationHandler = notification.NotificationQuoteAttributeValueChangeNotificationHandlerFunc(NotificationNotificationQuoteAttributeValueChangeNotificationHandler)
	api.NotificationNotificationQuoteCreationNotificationHandler = notification.NotificationQuoteCreationNotificationHandlerFunc(NotificationNotificationQuoteCreationNotificationHandler)
	api.NotificationNotificationQuoteInformationRequiredNotificationHandler = notification.NotificationQuoteInformationRequiredNotificationHandlerFunc(NotificationNotificationQuoteInformationRequiredNotificationHandler)
	api.NotificationNotificationQuoteStateChangeNotificationHandler = notification.NotificationQuoteStateChangeNotificationHandlerFunc(NotificationNotificationQuoteStateChangeNotificationHandler)

	// Product
	api.ProductProductFindHandler = product.ProductFindHandlerFunc(ProductProductFindHandler)
	api.ProductProductGetHandler = product.ProductGetHandlerFunc(ProductProductGetHandler)

	// ProductOfferingQualification
	api.ProductOfferingQualificationProductOfferingQualificationCreateHandler = poq.ProductOfferingQualificationCreateHandlerFunc(ProductOfferingQualificationProductOfferingQualificationCreateHandler)
	api.ProductOfferingQualificationProductOfferingQualificationFindHandler = poq.ProductOfferingQualificationFindHandlerFunc(ProductOfferingQualificationProductOfferingQualificationFindHandler)
	api.ProductOfferingQualificationProductOfferingQualificationGetHandler = poq.ProductOfferingQualificationGetHandlerFunc(ProductOfferingQualificationProductOfferingQualificationGetHandler)
	// ProductOfferingQualification Hub
	api.HubProductOfferingQualificationManagementHubDeleteHandler = hub.ProductOfferingQualificationManagementHubDeleteHandlerFunc(HubProductOfferingQualificationManagementHubDeleteHandler)
	api.HubProductOfferingQualificationManagementHubGetHandler = hub.ProductOfferingQualificationManagementHubGetHandlerFunc(HubProductOfferingQualificationManagementHubGetHandler)
	api.HubProductOfferingQualificationManagementHubPostHandler = hub.ProductOfferingQualificationManagementHubPostHandlerFunc(HubProductOfferingQualificationManagementHubPostHandler)
	// ProductOfferingQualification Notification
	api.NotificationNotificationProductOfferingQualificationCreationNotificationHandler = notification.NotificationProductOfferingQualificationCreationNotificationHandlerFunc(NotificationProductOfferingQualificationCreationNotificationHandler)
	api.NotificationNotificationProductOfferingQualificationStateChangeNotificationHandler = notification.NotificationProductOfferingQualificationStateChangeNotificationHandlerFunc(NotificationProductOfferingQualificationStateChangeNotificationHandler)
}

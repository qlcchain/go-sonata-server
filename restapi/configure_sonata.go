// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"path"

	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/qlcchain/go-sonata-server/restapi/handler"

	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-openapi/swag"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/config"

	"github.com/qlcchain/go-sonata-server/restapi/handler/mock"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"github.com/qlcchain/go-sonata-server/auth"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/operations"
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

func init() {
	dir := config.LogDir()
	fn := path.Join(dir, "sonata.log")

	lw, _ := rotatelogs.New(
		fn+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fn),
	)

	lh := lfshook.NewHook(
		lw,
		&log.JSONFormatter{},
	)
	log.AddHook(lh)
	//log.SetFormatter(&log.TextFormatter{
	//	FullTimestamp: true,
	//	DisableColors: true,
	//})
	log.SetReportCaller(true)
	log.SetLevel(log.ErrorLevel)
}

//go:generate swagger generate server --target ../../go-sonata-server --name Sonata --spec ../spec/all.yaml --principal models.Principal

func configureFlags(api *operations.SonataAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "jwt",
			LongDescription:  "jwt options",
			Options:          config.Cfg.Jwt,
		}, {
			ShortDescription: "debug",
			LongDescription:  "debug options",
			Options:          config.Cfg.Debug,
		}, {
			ShortDescription: "server",
			LongDescription:  "server options",
			Options:          config.Cfg.Server,
		},
	}
}

func configureAPI(api *operations.SonataAPI) http.Handler {
	cfg := config.Cfg

	var (
		err  error
		file string
	)
	if config.Cfg.Debug.IsFile {
		dir := config.DBDir()
		file = path.Join(dir, "sonata.db")
	} else {
		file = "file:mockdb?mode=memory&cache=shared"
	}
	//Store, err = gorm.Open(sqlite.Open("file:mockdb?mode=memory&cache=shared"), &gorm.Config{
	//	SkipDefaultTransaction: false,
	//	NamingStrategy:         nil,
	//	Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//		SlowThreshold: 100 * time.Millisecond,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	}),
	//})
	db.Store, err = gorm.Open("sqlite3", file)
	if err != nil {
		log.Fatalln(err)
	}

	if cfg.Debug.Verbose {
		log.SetLevel(log.DebugLevel)
		db.Store.LogMode(true)
	}

	if err := db.CreateTables(); err != nil {
		log.Fatalln(err)
	}

	if cfg.Debug.IsMock {
		if err := mockData(); err != nil {
			log.Fatal(err)
		} else {
			log.Debug("insert mock data successful")
		}
	}

	jwtCfg := cfg.Jwt
	if jwtCfg.Key != "" {
		var err error
		if jwtCfg.PrivateKey, err = auth.FromBase64(jwtCfg.Key); err != nil {
			log.Fatal(err)
		}
	}

	if jwtCfg.PrivateKey == nil {
		log.Debug("use default key...")
		jwtCfg.PrivateKey = auth.Decode(auth.DefaultKey)
		jwtCfg.PublicKey = &jwtCfg.PrivateKey.PublicKey
	}

	// configure the api here
	api.ServeError = errors.ServeError
	api.Logger = log.Printf
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.BearerAuth = func(token string, scopes []string) (*models.Principal, error) {
		// TODO: verify scopes???
		if claims, err := auth.ParseAndCheckToken(token, jwtCfg.PublicKey); err == nil {
			return &models.Principal{
				Code:   0,
				Reason: "",
				Roles:  claims.Roles,
			}, nil
		} else {
			switch vErr, ok := err.(*jwt.ValidationError); ok {
			case vErr.Errors&jwt.ValidationErrorMalformed != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorUnverifiable != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorSignatureInvalid != 0:
				fallthrough
			case vErr.Errors&jwt.ValidationErrorClaimsInvalid != 0:
				return &models.Principal{
					Code: 41, Reason: "Invalid credentials",
				}, nil
			case vErr.Errors&jwt.ValidationErrorExpired != 0:
				return &models.Principal{
					Code: 42, Reason: "Expired credentials",
				}, nil
			default:
				return &models.Principal{
					Code: 400, Reason: err.Error(),
				}, nil
			}
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	binder := &mock.MockBinder{}
	binder.Bind(api)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		// clear all subscriber
		db.Store.Delete(&schema.HubSubscriber{})
		if db.Store != nil {
			if err := db.Store.Close(); err != nil {
				log.Error(err)
			}
		}
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	writer := &logger{}
	cfg := config.Cfg.Server
	return alice.New(handlers.CORS(handlers.AllowedOrigins(cfg.AllowedOrigins), handlers.AllowCredentials()),
		func(handler http.Handler) http.Handler {
			return handlers.LoggingHandler(writer, handler)
		},
		handlers.ProxyHeaders,
		handlers.CompressHandler,
		handlers.RecoveryHandler(
			handlers.RecoveryLogger(writer),
			handlers.PrintRecoveryStack(true),
		)).Then(handler)
}

type logger struct {
}

func (l *logger) Write(p []byte) (n int, err error) {
	log.Debugln(string(p))
	return len(p), nil
}

func (l *logger) Println(args ...interface{}) {
	log.Error(args...)
}

func mockData() error {
	//mock address
	var address []*schema.GeographicAddress
	if err := json.Unmarshal([]byte(addresses), &address); err != nil {
		return err
	} else {
		for _, geographicAddress := range address {
			if err := db.Store.Create(geographicAddress).Error; err != nil {
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
			if err := db.Store.Create(geographicSite).Error; err != nil {
				log.Error(err)
			}
		}
	}
	//mock product
	for i := 0; i < 10; i++ {
		from := handler.Product()
		to := schema.FromProduct(from)
		if err := db.Store.Create(to).Error; err != nil {
			log.Error(err)
		}
	}

	return nil
}

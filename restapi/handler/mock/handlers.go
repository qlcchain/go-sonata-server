package mock

import (
	"fmt"
	"path"

	"github.com/qlcchain/go-sonata-server/config"
	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/go-openapi/swag"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"

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

var (
	Store *gorm.DB
)

func init() {
	// prepare db
	dir := config.DBDir()
	var err error
	//Store, err = gorm.Open(sqlite.Open("file:mockdb?mode=memory&cache=shared"), &gorm.Config{
	//	SkipDefaultTransaction: false,
	//	NamingStrategy:         nil,
	//	Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//		SlowThreshold: 100 * time.Millisecond,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	}),
	//})
	file := path.Join(dir, "sonata.db")
	//file := "file:mockdb?mode=memory&cache=shared"
	Store, err = gorm.Open("sqlite3", file)
	if err != nil {
		logrus.Fatalln(err)
	}
	Store.LogMode(true)
	if err := Store.AutoMigrate(
		// user management
		&schema.User{},
		// GeographicAddress
		&schema.GeographicAddress{}, &schema.FieldedAddress{}, &schema.FormattedAddress{},
		&schema.GeographicLocation{}, &schema.GeographicAddress{}, &schema.GeographicSubAddress{}, &schema.GeographicPoint{},
		&models.SubUnit{}, &models.ReferencedAddress{},
		// GeographicSite
		// Product
		// Order
	).Error; err != nil {
		logrus.Fatalln(err)
	}
}

func mockGeographicAddress(size int) {
	for i := 0; i < size; i++ {
		a := &schema.FieldedAddress{
			City:    gofakeit.City(),
			Country: gofakeit.Country(),
			GeographicSubAddress: []*schema.GeographicSubAddress{
				{
					//AtSchemaLocation:    "",
					//AtType:              "",
					BuildingName:        "Tower01",
					ID:                  xid.New().String(),
					LevelNumber:         "BASEMENT 1",
					LevelType:           "1",
					PrivateStreetName:   "university",
					PrivateStreetNumber: "01",
					SubUnit: []*models.SubUnit{
						{
							SubUnitIdentifier: swag.String(xid.New().String()),
							SubUnitType:       swag.String("BERTH"),
						},
					},
				},
			},
			ID:                 xid.New().String(),
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

		address := &schema.GeographicAddress{
			//AtSchemaLocation: "",
			//AtType:           "",
			AllowsNewSite:  false,
			FieldedAddress: a,
			FormattedAddress: &schema.FormattedAddress{
				AddrLine1: swag.String(fmt.Sprintf("%s %s %s %s, %s %s", a.StreetNr, a.StreetNrLast, a.StreetSuffix,
					a.StreetNrLastSuffix, a.StreetName, a.StreetSuffix)),
				AddrLine2:         "",
				City:              a.City,
				Country:           a.Country,
				Locality:          a.Locality,
				PostCodeExtension: a.PostCodeExtension,
				Postcode:          a.Postcode,
				StateOrProvince:   a.StateOrProvince,
			},
			GeographicLocation: &schema.GeographicLocation{
				ID: xid.New().String(),
				GeographicPoint: []*schema.GeographicPoint{
					{
						ID:        xid.New().String(),
						Latitude:  swag.String(fmt.Sprintf("%f", gofakeit.Latitude())),
						Longitude: swag.String(fmt.Sprintf("%f", gofakeit.Longitude())),
					},
				},
				SpatialRef: swag.String(gofakeit.CountryAbr()),
			},
			HasPublicSite: false,
			ID:            xid.New().String(),
			ReferencedAddress: &models.ReferencedAddress{
				ID:            xid.New().String(),
				ReferenceID:   swag.String("ReferenceID"),
				ReferenceType: swag.String("ReferenceType"),
			},
		}

		if err := Store.Create(address).Error; err != nil {
			logrus.Error(err)
		}
	}
}

func mockData() error {
	mockGeographicAddress(10)
	return nil
}

func Bind(api *operations.SonataAPI) {
	if err := mockData(); err != nil {
		logrus.Fatal(err)
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
}

package mock

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/util"

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
	DB *gorm.DB
)

func init() {
	// prepare db
	//dir := util.DBDir()
	var err error
	//DB, err = gorm.Open(sqlite.Open("file:mockdb?mode=memory&cache=shared"), &gorm.Config{
	//	SkipDefaultTransaction: false,
	//	NamingStrategy:         nil,
	//	Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//		SlowThreshold: 100 * time.Millisecond,
	//		LogLevel:      logger.Info,
	//		Colorful:      true,
	//	}),
	//})
	DB, err = gorm.Open("sqlite3", "file:mockdb?mode=memory&cache=shared")
	if err != nil {
		logrus.Fatalln(err)
	}
	DB.LogMode(true)
	if err := DB.AutoMigrate(&db.User{}, &db.GeographicAddressModel{}, &db.FieldedAddressModel{}).Error; err != nil {
		logrus.Fatalln(err)
	}
}

func mockGeographicAddress() {
	for i := 0; i < 10; i++ {
		a := &db.FieldedAddressModel{
			City:    gofakeit.City(),
			Country: gofakeit.Country(),
			//GeographicSubAdress: []*models.GeographicSubAddress{},
			//ID:                 xid.New().String(),
			Locality:           "",
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

		address := &db.GeographicAddressModel{
			//AtSchemaLocation: "",
			//AtType:           "",
			AllowsNewSite:  false,
			FieldedAddress: a,
			//FormattedAddress: &models.FormattedAddress{
			//	AddrLine1: swag.String(fmt.Sprintf("%s %s %s %s, %s %s", a.StreetNr, a.StreetNrLast, a.StreetSuffix,
			//		a.StreetNrLastSuffix, a.StreetName, a.StreetSuffix)),
			//	AddrLine2:         "",
			//	City:              a.City,
			//	Country:           a.Country,
			//	ID:                xid.New().String(),
			//	Locality:          a.Locality,
			//	PostCodeExtension: a.PostCodeExtension,
			//	Postcode:          a.Postcode,
			//	StateOrProvince:   a.StateOrProvince,
			//},
			//GeographicLocation: &models.GeographicLocation{
			//	GeographicPoint: []*models.GeographicPoint{{
			//		Latitude:  swag.String(fmt.Sprintf("%f", gofakeit.Latitude())),
			//		Longitude: swag.String(fmt.Sprintf("%f", gofakeit.Longitude())),
			//	}},
			//	ID:         xid.New().String(),
			//	SpatialRef: swag.String(gofakeit.CountryAbr()),
			//},
			HasPublicSite: false,
			ID:            xid.New().String(),
			//ReferencedAddress: &models.ReferencedAddress{
			//	ID:            "",
			//	ReferenceID:   nil,
			//	ReferenceType: nil,
			//},
		}

		fmt.Println(util.ToIndentString(address))
		if err := DB.Create(address).Error; err != nil {
			logrus.Error(err)
		} else {
			count := 0
			var aa []db.GeographicAddressModel
			if err := DB.Find(&aa).Count(&count).Error; err != nil {
				logrus.Error(err)
			} else {
				logrus.Debug(count)
			}
		}
	}
}

func mockData() error {
	mockGeographicAddress()
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

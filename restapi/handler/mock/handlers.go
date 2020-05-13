/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
		logrus.Fatalln(err)
	}
	Store.LogMode(true)
	if err := Store.AutoMigrate(
		// user management
		&schema.User{}, &schema.HubSubscriber{},
		// GeographicAddress
		&schema.GeographicAddress{}, &schema.FieldedAddress{}, &schema.FormattedAddress{},
		&schema.GeographicLocation{}, &schema.GeographicAddress{}, &schema.GeographicSubAddress{}, &schema.GeographicPoint{},
		&models.SubUnit{}, &models.ReferencedAddress{},
		// GeographicSite
		&schema.GeographicSite{},
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
		logrus.Fatalln(err)
	}
}

func mockGeographicAddress(size int) {
	//for i := 0; i < size; i++ {
	//	address := GeographicAddress()
	//
	//	if err := Store.Create(address).Error; err != nil {
	//		logrus.Error(err)
	//	}
	//}
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
	// ProductOfferingQualification Notification
	api.NotificationNotificationProductOfferingQualificationCreationNotificationHandler = notification.NotificationProductOfferingQualificationCreationNotificationHandlerFunc(NotificationProductOfferingQualificationCreationNotificationHandler)
	api.NotificationNotificationProductOfferingQualificationStateChangeNotificationHandler = notification.NotificationProductOfferingQualificationStateChangeNotificationHandlerFunc(NotificationProductOfferingQualificationStateChangeNotificationHandler)
}

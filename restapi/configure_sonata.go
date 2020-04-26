// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/restapi/operations"
	"github.com/qlcchain/go-sonata-server/restapi/operations/cancel_product_order"
	"github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address"
	"github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address_validation"
	"github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	"github.com/qlcchain/go-sonata-server/restapi/operations/notification"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product_offering_qualification"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product_order"
	"github.com/qlcchain/go-sonata-server/restapi/operations/quote"
)

//go:generate swagger generate server --target ../../go-sonata-server --name Sonata --spec ../spec/all.yaml

func configureFlags(api *operations.SonataAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SonataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CancelProductOrderCancelProductOrderCreateHandler == nil {
		api.CancelProductOrderCancelProductOrderCreateHandler = cancel_product_order.CancelProductOrderCreateHandlerFunc(func(params cancel_product_order.CancelProductOrderCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderCreate has not yet been implemented")
		})
	}
	if api.CancelProductOrderCancelProductOrderFindHandler == nil {
		api.CancelProductOrderCancelProductOrderFindHandler = cancel_product_order.CancelProductOrderFindHandlerFunc(func(params cancel_product_order.CancelProductOrderFindParams) middleware.Responder {
			return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderFind has not yet been implemented")
		})
	}
	if api.CancelProductOrderCancelProductOrderGetHandler == nil {
		api.CancelProductOrderCancelProductOrderGetHandler = cancel_product_order.CancelProductOrderGetHandlerFunc(func(params cancel_product_order.CancelProductOrderGetParams) middleware.Responder {
			return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderGet has not yet been implemented")
		})
	}
	if api.GeographicAddressGeographicAddressGetHandler == nil {
		api.GeographicAddressGeographicAddressGetHandler = geographic_address.GeographicAddressGetHandlerFunc(func(params geographic_address.GeographicAddressGetParams) middleware.Responder {
			return middleware.NotImplemented("operation geographic_address.GeographicAddressGet has not yet been implemented")
		})
	}
	if api.GeographicAddressValidationGeographicAddressValidationCreateHandler == nil {
		api.GeographicAddressValidationGeographicAddressValidationCreateHandler = geographic_address_validation.GeographicAddressValidationCreateHandlerFunc(func(params geographic_address_validation.GeographicAddressValidationCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation geographic_address_validation.GeographicAddressValidationCreate has not yet been implemented")
		})
	}
	if api.GeographicSiteGeographicSiteFindHandler == nil {
		api.GeographicSiteGeographicSiteFindHandler = geographic_site.GeographicSiteFindHandlerFunc(func(params geographic_site.GeographicSiteFindParams) middleware.Responder {
			return middleware.NotImplemented("operation geographic_site.GeographicSiteFind has not yet been implemented")
		})
	}
	if api.GeographicSiteGeographicSiteGetHandler == nil {
		api.GeographicSiteGeographicSiteGetHandler = geographic_site.GeographicSiteGetHandlerFunc(func(params geographic_site.GeographicSiteGetParams) middleware.Responder {
			return middleware.NotImplemented("operation geographic_site.GeographicSiteGet has not yet been implemented")
		})
	}
	if api.NotificationNotificationProductOrderAttributeValueChangeNotificationHandler == nil {
		api.NotificationNotificationProductOrderAttributeValueChangeNotificationHandler = notification.NotificationProductOrderAttributeValueChangeNotificationHandlerFunc(func(params notification.NotificationProductOrderAttributeValueChangeNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationProductOrderAttributeValueChangeNotification has not yet been implemented")
		})
	}
	if api.NotificationNotificationProductOrderCreationNotificationHandler == nil {
		api.NotificationNotificationProductOrderCreationNotificationHandler = notification.NotificationProductOrderCreationNotificationHandlerFunc(func(params notification.NotificationProductOrderCreationNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationProductOrderCreationNotification has not yet been implemented")
		})
	}
	if api.NotificationNotificationProductOrderInformationRequiredNotificationHandler == nil {
		api.NotificationNotificationProductOrderInformationRequiredNotificationHandler = notification.NotificationProductOrderInformationRequiredNotificationHandlerFunc(func(params notification.NotificationProductOrderInformationRequiredNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationProductOrderInformationRequiredNotification has not yet been implemented")
		})
	}
	if api.NotificationNotificationProductOrderStateChangeNotificationHandler == nil {
		api.NotificationNotificationProductOrderStateChangeNotificationHandler = notification.NotificationProductOrderStateChangeNotificationHandlerFunc(func(params notification.NotificationProductOrderStateChangeNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationProductOrderStateChangeNotification has not yet been implemented")
		})
	}
	if api.NotificationNotificationQuoteAttributeValueChangeNotificationHandler == nil {
		api.NotificationNotificationQuoteAttributeValueChangeNotificationHandler = notification.NotificationQuoteAttributeValueChangeNotificationHandlerFunc(func(params notification.NotificationQuoteAttributeValueChangeNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationQuoteAttributeValueChangeNotification has not yet been implemented")
		})
	}
	if api.NotificationNotificationQuoteCreationNotificationHandler == nil {
		api.NotificationNotificationQuoteCreationNotificationHandler = notification.NotificationQuoteCreationNotificationHandlerFunc(func(params notification.NotificationQuoteCreationNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationQuoteCreationNotification has not yet been implemented")
		})
	}
	if api.NotificationNotificationQuoteInformationRequiredNotificationHandler == nil {
		api.NotificationNotificationQuoteInformationRequiredNotificationHandler = notification.NotificationQuoteInformationRequiredNotificationHandlerFunc(func(params notification.NotificationQuoteInformationRequiredNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationQuoteInformationRequiredNotification has not yet been implemented")
		})
	}
	if api.NotificationNotificationQuoteStateChangeNotificationHandler == nil {
		api.NotificationNotificationQuoteStateChangeNotificationHandler = notification.NotificationQuoteStateChangeNotificationHandlerFunc(func(params notification.NotificationQuoteStateChangeNotificationParams) middleware.Responder {
			return middleware.NotImplemented("operation notification.NotificationQuoteStateChangeNotification has not yet been implemented")
		})
	}
	if api.ProductProductFindHandler == nil {
		api.ProductProductFindHandler = product.ProductFindHandlerFunc(func(params product.ProductFindParams) middleware.Responder {
			return middleware.NotImplemented("operation product.ProductFind has not yet been implemented")
		})
	}
	if api.ProductProductGetHandler == nil {
		api.ProductProductGetHandler = product.ProductGetHandlerFunc(func(params product.ProductGetParams) middleware.Responder {
			return middleware.NotImplemented("operation product.ProductGet has not yet been implemented")
		})
	}
	if api.ProductOfferingQualificationProductOfferingQualificationCreateHandler == nil {
		api.ProductOfferingQualificationProductOfferingQualificationCreateHandler = product_offering_qualification.ProductOfferingQualificationCreateHandlerFunc(func(params product_offering_qualification.ProductOfferingQualificationCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation product_offering_qualification.ProductOfferingQualificationCreate has not yet been implemented")
		})
	}
	if api.ProductOfferingQualificationProductOfferingQualificationFindHandler == nil {
		api.ProductOfferingQualificationProductOfferingQualificationFindHandler = product_offering_qualification.ProductOfferingQualificationFindHandlerFunc(func(params product_offering_qualification.ProductOfferingQualificationFindParams) middleware.Responder {
			return middleware.NotImplemented("operation product_offering_qualification.ProductOfferingQualificationFind has not yet been implemented")
		})
	}
	if api.ProductOfferingQualificationProductOfferingQualificationGetHandler == nil {
		api.ProductOfferingQualificationProductOfferingQualificationGetHandler = product_offering_qualification.ProductOfferingQualificationGetHandlerFunc(func(params product_offering_qualification.ProductOfferingQualificationGetParams) middleware.Responder {
			return middleware.NotImplemented("operation product_offering_qualification.ProductOfferingQualificationGet has not yet been implemented")
		})
	}
	if api.HubProductOfferingQualificationManagementHubDeleteHandler == nil {
		api.HubProductOfferingQualificationManagementHubDeleteHandler = hub.ProductOfferingQualificationManagementHubDeleteHandlerFunc(func(params hub.ProductOfferingQualificationManagementHubDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubDelete has not yet been implemented")
		})
	}
	if api.HubProductOfferingQualificationManagementHubGetHandler == nil {
		api.HubProductOfferingQualificationManagementHubGetHandler = hub.ProductOfferingQualificationManagementHubGetHandlerFunc(func(params hub.ProductOfferingQualificationManagementHubGetParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubGet has not yet been implemented")
		})
	}
	if api.HubProductOfferingQualificationManagementHubPostHandler == nil {
		api.HubProductOfferingQualificationManagementHubPostHandler = hub.ProductOfferingQualificationManagementHubPostHandlerFunc(func(params hub.ProductOfferingQualificationManagementHubPostParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubPost has not yet been implemented")
		})
	}
	if api.ProductOrderProductOrderCreateHandler == nil {
		api.ProductOrderProductOrderCreateHandler = product_order.ProductOrderCreateHandlerFunc(func(params product_order.ProductOrderCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation product_order.ProductOrderCreate has not yet been implemented")
		})
	}
	if api.ProductOrderProductOrderFindHandler == nil {
		api.ProductOrderProductOrderFindHandler = product_order.ProductOrderFindHandlerFunc(func(params product_order.ProductOrderFindParams) middleware.Responder {
			return middleware.NotImplemented("operation product_order.ProductOrderFind has not yet been implemented")
		})
	}
	if api.ProductOrderProductOrderGetHandler == nil {
		api.ProductOrderProductOrderGetHandler = product_order.ProductOrderGetHandlerFunc(func(params product_order.ProductOrderGetParams) middleware.Responder {
			return middleware.NotImplemented("operation product_order.ProductOrderGet has not yet been implemented")
		})
	}
	if api.HubProductOrderManagementHubCreateHandler == nil {
		api.HubProductOrderManagementHubCreateHandler = hub.ProductOrderManagementHubCreateHandlerFunc(func(params hub.ProductOrderManagementHubCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.ProductOrderManagementHubCreate has not yet been implemented")
		})
	}
	if api.HubProductOrderManagementHubDeleteHandler == nil {
		api.HubProductOrderManagementHubDeleteHandler = hub.ProductOrderManagementHubDeleteHandlerFunc(func(params hub.ProductOrderManagementHubDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.ProductOrderManagementHubDelete has not yet been implemented")
		})
	}
	if api.HubProductOrderManagementHubFindHandler == nil {
		api.HubProductOrderManagementHubFindHandler = hub.ProductOrderManagementHubFindHandlerFunc(func(params hub.ProductOrderManagementHubFindParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.ProductOrderManagementHubFind has not yet been implemented")
		})
	}
	if api.QuoteQuoteCreateHandler == nil {
		api.QuoteQuoteCreateHandler = quote.QuoteCreateHandlerFunc(func(params quote.QuoteCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation quote.QuoteCreate has not yet been implemented")
		})
	}
	if api.QuoteQuoteFindHandler == nil {
		api.QuoteQuoteFindHandler = quote.QuoteFindHandlerFunc(func(params quote.QuoteFindParams) middleware.Responder {
			return middleware.NotImplemented("operation quote.QuoteFind has not yet been implemented")
		})
	}
	if api.QuoteQuoteGetHandler == nil {
		api.QuoteQuoteGetHandler = quote.QuoteGetHandlerFunc(func(params quote.QuoteGetParams) middleware.Responder {
			return middleware.NotImplemented("operation quote.QuoteGet has not yet been implemented")
		})
	}
	if api.HubQuoteManagementHubCreateHandler == nil {
		api.HubQuoteManagementHubCreateHandler = hub.QuoteManagementHubCreateHandlerFunc(func(params hub.QuoteManagementHubCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.QuoteManagementHubCreate has not yet been implemented")
		})
	}
	if api.HubQuoteManagementHubDeleteHandler == nil {
		api.HubQuoteManagementHubDeleteHandler = hub.QuoteManagementHubDeleteHandlerFunc(func(params hub.QuoteManagementHubDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.QuoteManagementHubDelete has not yet been implemented")
		})
	}
	if api.HubQuoteManagementHubFindHandler == nil {
		api.HubQuoteManagementHubFindHandler = hub.QuoteManagementHubFindHandlerFunc(func(params hub.QuoteManagementHubFindParams) middleware.Responder {
			return middleware.NotImplemented("operation hub.QuoteManagementHubFind has not yet been implemented")
		})
	}
	if api.QuoteQuoteRequestStateChangeHandler == nil {
		api.QuoteQuoteRequestStateChangeHandler = quote.QuoteRequestStateChangeHandlerFunc(func(params quote.QuoteRequestStateChangeParams) middleware.Responder {
			return middleware.NotImplemented("operation quote.QuoteRequestStateChange has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

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
	return handler
}

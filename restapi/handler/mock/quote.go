package mock

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	"github.com/qlcchain/go-sonata-server/restapi/operations/notification"
	"github.com/qlcchain/go-sonata-server/restapi/operations/quote"
)

func QuoteQuoteRequestStateChangeHandler(params quote.QuoteRequestStateChangeParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation quote.QuoteRequestStateChange has not yet been implemented")
}

func QuoteQuoteCreateHandler(params quote.QuoteCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation quote.QuoteCreate has not yet been implemented")
}

func QuoteQuoteFindHandler(params quote.QuoteFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation quote.QuoteFind has not yet been implemented")
}

func QuoteQuoteGetHandler(params quote.QuoteGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation quote.QuoteGet has not yet been implemented")
}

func HubQuoteManagementHubCreateHandler(params hub.QuoteManagementHubCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.QuoteManagementHubCreate has not yet been implemented")
}

func HubQuoteManagementHubDeleteHandler(params hub.QuoteManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.QuoteManagementHubDelete has not yet been implemented")
}

func HubQuoteManagementHubFindHandler(params hub.QuoteManagementHubFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.QuoteManagementHubFind has not yet been implemented")
}

func NotificationNotificationQuoteAttributeValueChangeNotificationHandler(params notification.NotificationQuoteAttributeValueChangeNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationQuoteAttributeValueChangeNotification has not yet been implemented")
}

func NotificationNotificationQuoteCreationNotificationHandler(params notification.NotificationQuoteCreationNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationQuoteCreationNotification has not yet been implemented")
}

func NotificationNotificationQuoteInformationRequiredNotificationHandler(params notification.NotificationQuoteInformationRequiredNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationQuoteInformationRequiredNotification has not yet been implemented")
}

func NotificationNotificationQuoteStateChangeNotificationHandler(params notification.NotificationQuoteStateChangeNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationQuoteStateChangeNotification has not yet been implemented")
}

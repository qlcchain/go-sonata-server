package mock

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/operations/cancel_product_order"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	"github.com/qlcchain/go-sonata-server/restapi/operations/notification"
	po "github.com/qlcchain/go-sonata-server/restapi/operations/product_order"
)

func CancelProductOrderCancelProductOrderCreateHandler(params cancel_product_order.CancelProductOrderCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderCreate has not yet been implemented")
}

func CancelProductOrderCancelProductOrderFindHandler(params cancel_product_order.CancelProductOrderFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderFind has not yet been implemented")
}

func CancelProductOrderCancelProductOrderGetHandler(params cancel_product_order.CancelProductOrderGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderGet has not yet been implemented")
}

func ProductOrderProductOrderCreateHandler(params po.ProductOrderCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation product_order.ProductOrderCreate has not yet been implemented")
}

func ProductOrderProductOrderFindHandler(params po.ProductOrderFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation product_order.ProductOrderFind has not yet been implemented")
}

func ProductOrderProductOrderGetHandler(params po.ProductOrderGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation product_order.ProductOrderGet has not yet been implemented")
}

func HubProductOrderManagementHubCreateHandler(params hub.ProductOrderManagementHubCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.ProductOrderManagementHubCreate has not yet been implemented")
}

func HubProductOrderManagementHubDeleteHandler(params hub.ProductOrderManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.ProductOrderManagementHubDelete has not yet been implemented")
}

func HubProductOrderManagementHubFindHandler(params hub.ProductOrderManagementHubFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.ProductOrderManagementHubFind has not yet been implemented")
}

func NotificationNotificationProductOrderAttributeValueChangeNotificationHandler(params notification.NotificationProductOrderAttributeValueChangeNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationProductOrderAttributeValueChangeNotification has not yet been implemented")
}

func NotificationNotificationProductOrderCreationNotificationHandler(params notification.NotificationProductOrderCreationNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationProductOrderCreationNotification has not yet been implemented")
}

func NotificationNotificationProductOrderInformationRequiredNotificationHandler(params notification.NotificationProductOrderInformationRequiredNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationProductOrderInformationRequiredNotification has not yet been implemented")
}

func NotificationNotificationProductOrderStateChangeNotificationHandler(params notification.NotificationProductOrderStateChangeNotificationParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation notification.NotificationProductOrderStateChangeNotification has not yet been implemented")
}

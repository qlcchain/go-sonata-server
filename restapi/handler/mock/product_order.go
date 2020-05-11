/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
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
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return po.NewProductOrderGetBadRequest().WithPayload(payload)
	}

	if o, err := db.GetProductOrder(Store, params.ProductOrderID); err == nil {
		return po.NewProductOrderGetOK().WithPayload(o)
	} else if err == gorm.ErrRecordNotFound {
		return po.NewProductOrderGetNotFound()
	} else {
		return po.NewProductOrderGetInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
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

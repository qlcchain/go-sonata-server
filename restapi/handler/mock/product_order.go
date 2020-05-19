/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/qlcchain/go-sonata-server/event"
	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/restapi/operations/cancel_product_order"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	"github.com/qlcchain/go-sonata-server/restapi/operations/notification"
	po "github.com/qlcchain/go-sonata-server/restapi/operations/product_order"
)

const (
	productOrderTopic = "productOrder"
)

var poBus = event.GetEventBus(productOrderTopic)

// create/get/find product order
func ProductOrderProductOrderCreateHandler(params po.ProductOrderCreateParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return po.NewProductOrderGetUnauthorized().WithPayload(payload)
	}

	if order, err := db.ProductOrderCreateToProductOrder(params.ProductOrder); err == nil {
		if err := db.Store.Save(order).Error; err == nil {
			ev := models.PoEventPlus{
				PoEvent: models.PoEvent{
					Event:     order.ProductOrderEvent(),
					EventID:   util.NewIDPtr(),
					EventTime: handler.NewDatetime(time.Now()),
					EventType: models.ProductOrderEventTypeProductOrderCreationNotification,
				},
				FieldPath:    []string{},
				ResourcePath: swag.String(""),
			}
			poBus.Publish(string(models.ProductOrderEventTypeProductOrderCreationNotification), ev)
			return po.NewProductOrderCreateCreated().WithPayload(order.To())
		} else {
			return po.NewProductOrderCreateInternalServerError().WithPayload(&models.ErrorRepresentation{
				Reason: swag.String(err.Error()),
			})
		}
	} else {
		return po.NewProductOrderCreateInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func ProductOrderProductOrderFindHandler(params po.ProductOrderFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return po.NewProductOrderFindUnauthorized().WithPayload(payload)
	}

	if r, err := db.GetProductOrderByParams(&params); err == nil {
		return po.NewProductOrderFindOK().WithPayload(r)
	} else if err == gorm.ErrRecordNotFound {
		return po.NewProductOrderFindNotFound()
	} else {
		return po.NewProductOrderFindInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func ProductOrderProductOrderGetHandler(params po.ProductOrderGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return po.NewProductOrderGetUnauthorized().WithPayload(payload)
	}

	if o, err := db.GetProductOrder(params.ProductOrderID); err == nil {
		return po.NewProductOrderGetOK().WithPayload(o)
	} else if err == gorm.ErrRecordNotFound {
		return po.NewProductOrderGetNotFound()
	} else {
		return po.NewProductOrderGetInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

// cancel product order
func CancelProductOrderCancelProductOrderCreateHandler(params cancel_product_order.CancelProductOrderCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderCreate has not yet been implemented")
}

func CancelProductOrderCancelProductOrderFindHandler(params cancel_product_order.CancelProductOrderFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderFind has not yet been implemented")
}

func CancelProductOrderCancelProductOrderGetHandler(params cancel_product_order.CancelProductOrderGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation cancel_product_order.CancelProductOrderGet has not yet been implemented")
}

func HubProductOrderManagementHubCreateHandler(params hub.ProductOrderManagementHubCreateParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewProductOrderManagementHubCreateUnauthorized().WithPayload(payload)
	}

	// verify input
	input := params.Hub
	if input == nil {
		return hub.NewProductOrderManagementHubCreateBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(21),
			Reason: swag.String("Missing body"),
		})
	}

	if input.Query == nil || input.Callback == nil {
		return hub.NewProductOrderManagementHubCreateBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(23),
			Reason: swag.String("Missing body field"),
		})
	}

	if *input.Query == "" {
		return hub.NewProductOrderManagementHubCreateBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(24),
			Reason: swag.String("Invalid body field"),
		})
	}

	if payload := handler.VerifyCallback(input.Callback); payload != nil {
		return hub.NewProductOrderManagementHubCreateBadRequest().WithPayload(payload)
	}

	if id, err := poBus.Subscribe(handler.ParseType(*input.Query), commonCallbackHandler, &event.CallbackOption{}); err != nil {
		return hub.NewProductOrderManagementHubCreateInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	} else {
		if err := db.AddSubscriber(&schema.HubSubscriber{
			ID:       id,
			Type:     productOrderTopic,
			Query:    *input.Query,
			Callback: *input.Callback,
		}); err == nil {
			return hub.NewProductOrderManagementHubCreateCreated().WithPayload(&models.Hub{
				Callback: input.Callback,
				ID:       swag.String(id),
				Query:    input.Query,
			})
		} else {
			return hub.NewProductOrderManagementHubCreateInternalServerError().WithPayload(&models.ErrorRepresentation{
				Reason: swag.String(err.Error()),
			})
		}
	}
}

func HubProductOrderManagementHubDeleteHandler(params hub.ProductOrderManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewProductOrderManagementHubDeleteUnauthorized().WithPayload(payload)
	}
	// verify id
	id := params.HubID

	if s, err := db.FindSubscriber(id); err == nil {
		if err := poBus.Unsubscribe(handler.ParseType(s.Query), id); err == nil {
			if err := db.DeleteSubscriber(id); err == nil {
				return hub.NewProductOrderManagementHubDeleteNoContent()
			} else {
				return hub.NewProductOrderManagementHubDeleteBadRequest().WithPayload(&models.ErrorRepresentation{
					Reason: swag.String(err.Error()),
				})
			}
		} else {
			return hub.NewProductOrderManagementHubDeleteBadRequest().WithPayload(&models.ErrorRepresentation{
				Reason: swag.String(err.Error()),
			})
		}
	} else if err == gorm.ErrRecordNotFound {
		return hub.NewProductOrderManagementHubDeleteNotFound()
	} else {
		return hub.NewProductOrderManagementHubDeleteBadRequest().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func HubProductOrderManagementHubFindHandler(params hub.ProductOrderManagementHubFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewProductOrderManagementHubFindUnauthorized().WithPayload(payload)
	}

	if subscribers, err := db.ListSubscribers(productOrderTopic); err == nil {
		var payload []*models.Hub
		for _, s := range subscribers {
			payload = append(payload, &models.Hub{
				Callback: swag.String(s.Callback),
				ID:       swag.String(s.ID),
				Query:    swag.String(s.Query),
			})
		}
		return hub.NewProductOrderManagementHubFindOK().WithPayload(payload)
	} else if err == gorm.ErrRecordNotFound {
		return hub.NewProductOrderManagementHubFindNotFound()
	} else {
		return hub.NewProductOrderManagementHubFindInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func NotificationNotificationProductOrderAttributeValueChangeNotificationHandler(params notification.NotificationProductOrderAttributeValueChangeNotificationParams) middleware.Responder {
	log.Debugf("got NotificationProductOrderAttributeValueChangeNotificationParams: %s", util.ToString(params.ProductOrderAttributeValueChange))
	return notification.NewNotificationProductOrderAttributeValueChangeNotificationNoContent()
}

func NotificationNotificationProductOrderCreationNotificationHandler(params notification.NotificationProductOrderCreationNotificationParams) middleware.Responder {
	log.Debugf("got NotificationProductOrderCreationNotificationParams: %s", util.ToString(params.ProductOrderCreationNotification))
	return notification.NewNotificationProductOfferingQualificationCreationNotificationNoContent()
}

func NotificationNotificationProductOrderInformationRequiredNotificationHandler(params notification.NotificationProductOrderInformationRequiredNotificationParams) middleware.Responder {
	log.Debugf("got NotificationProductOrderInformationRequiredNotificationParams:%s", util.ToString(params.ProductOrderInformationRequired))
	return notification.NewNotificationProductOrderInformationRequiredNotificationNoContent()
}

func NotificationNotificationProductOrderStateChangeNotificationHandler(params notification.NotificationProductOrderStateChangeNotificationParams) middleware.Responder {
	log.Debugf("got NotificationProductOrderStateChangeNotificationParams:%s", util.ToString(params.ProductOrderStateChange))
	return notification.NewNotificationProductOrderStateChangeNotificationNoContent()
}

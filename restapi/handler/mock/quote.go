/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"time"

	"github.com/qlcchain/go-sonata-server/restapi/operations/quote"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/event"
	"github.com/qlcchain/go-sonata-server/schema"
	"github.com/qlcchain/go-sonata-server/util"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	"github.com/qlcchain/go-sonata-server/restapi/operations/notification"
)

const (
	quoteTopic = "quote"
)

var (
	quoteBus = event.GetEventBus(quoteTopic)
)

func QuoteQuoteRequestStateChangeHandler(params quote.QuoteRequestStateChangeParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return quote.NewQuoteRequestStateChangeUnauthorized().WithPayload(payload)
	}
	input := params.ChangeQuoteStateRequest
	if input == nil {
		return quote.NewQuoteRequestStateChangeBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(21),
			Reason: swag.String("Missing body"),
		})
	}
	now := strfmt.DateTime(time.Now())
	state := &schema.Quote{
		ID:         *input.ID,
		ExternalID: input.ExternalID,
		State:      models.QuoteStateType(input.State),
	}

	if err := db.Store.Model(&schema.Quote{}).Update(state).Error; err == nil {
		if q, err := db.GetQuote(*input.ID); err != nil {
			//FIXME: calculate file path and resource path
			ev := models.QuoteEventPlus{
				QuoteEvent: models.QuoteEvent{
					Event:     q.ToQuoteSummaryView(),
					EventID:   util.NewIDPtr(),
					EventTime: &now,
					EventType: models.QuoteEventTypeQUOTESTATECHANGENOTIFICATION,
				},
				FieldPath:    []string{},
				ResourcePath: nil,
			}
			quoteBus.Publish(string(models.QuoteEventTypeQUOTESTATECHANGENOTIFICATION), ev)
		}
		return quote.NewQuoteRequestStateChangeOK().WithPayload(&models.ChangeQuoteStateResponse{
			ExternalID:                    input.ExternalID,
			ID:                            input.ID,
			QuoteEffectiveStateChangeDate: &now,
			State:                         models.QuoteStateType(input.State),
		})
	} else if err == gorm.ErrRecordNotFound {
		return quote.NewQuoteRequestStateChangeNotFound()
	} else {
		return quote.NewQuoteRequestStateChangeInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func QuoteQuoteCreateHandler(params quote.QuoteCreateParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return quote.NewQuoteCreateUnauthorized().WithPayload(payload)
	}
	input := params.Quote
	id := util.NewID()
	now := strfmt.DateTime(time.Now())
	var items []*schema.QuoteItem
	for _, item := range input.QuoteItem {
		items = append(items, &schema.QuoteItem{
			AtSchemaLocation: item.AtSchemaLocation,
			AtType:           item.AtType,
			Action:           item.Action,
			ID:               util.NewOrDefaultPtr(item.ID),
			Note:             schema.FromNotes(item.Note),
			Product:          schema.FromProduct(item.Product),
			ProductOffering:  item.ProductOffering,
			Qualification:    []*models.ProductOfferingQualificationRef{item.Qualification},
			//QuoteItemPrice:         item.,
			QuoteItemRelationship: item.QuoteItemRelationship,
			//QuoteItemTerm:          item.q,
			RelatedParty:           schema.FromRelatedParty(item.RelatedParty),
			RequestedQuoteItemTerm: item.RequestedQuoteItemTerm,
			State:                  models.QuoteItemStateTypeINPROGRESS,
		})
	}

	for _, a := range input.Agreement {
		a.ID = util.NewOrDefaultPtr(a.ID)
	}

	q := &schema.Quote{
		ID:                           util.NewID(),
		AtBaseType:                   input.AtBaseType,
		AtSchemaLocation:             input.AtSchemaLocation,
		AtType:                       input.AtType,
		Agreement:                    input.Agreement,
		Description:                  input.Description,
		ExpectedFulfillmentStartDate: input.ExpectedFulfillmentStartDate,
		//TODO: maybe config
		ExpectedQuoteCompletionDate:  strfmt.Date(time.Now().AddDate(0, 0, 5)),
		EffectiveQuoteCompletionDate: strfmt.NewDateTime(),
		ExternalID:                   input.ExternalID,
		InstantSyncQuoting:           swag.BoolValue(input.InstantSyncQuoting),
		Note:                         schema.FromNotes(input.Note),
		ProjectID:                    input.ProjectID,
		QuoteItem:                    items,
		QuoteLevel:                   input.QuoteLevel,
		RelatedParty:                 handler.ConvertRelatedParty(input.RelatedParty),
		RequestedQuoteCompletionDate: input.RequestedQuoteCompletionDate,
		Href:                         handler.HrefToID("", id),
	}
	if err := db.Store.Save(q).Error; err == nil {
		//FIXME: calculate file path and resource path
		ev := models.QuoteEventPlus{
			QuoteEvent: models.QuoteEvent{
				Event:     q.ToQuoteSummaryView(),
				EventID:   util.NewIDPtr(),
				EventTime: &now,
				EventType: models.QuoteEventTypeQUOTECREATIONNOTIFICATION,
			},
			FieldPath:    []string{},
			ResourcePath: swag.String(""),
		}
		quoteBus.Publish(string(models.QuoteEventTypeQUOTECREATIONNOTIFICATION), ev)
		return quote.NewQuoteCreateCreated().WithPayload(q.ToQuote())
	} else {
		return quote.NewQuoteCreateInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func QuoteQuoteFindHandler(params quote.QuoteFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return quote.NewQuoteFindUnauthorized().WithPayload(payload)
	}

	// build query sql
	tx := db.GetTx()
	if v, b := handler.VerifyField(params.ExternalID); b {
		tx = tx.Where("externalId=?", v)
	}

	if v, b := handler.VerifyField(params.ProjectID); b {
		tx = tx.Where("projectId=?", v)
	}

	if v, b := handler.VerifyField(params.QuoteLevel); b {
		tx = tx.Where("quoteLevel=?", v)
	}

	if v, b := handler.VerifyField(params.State); b {
		tx = tx.Where("state=?", v)
	}

	if v, b := handler.VerifyField(params.QuoteCompletionDateGt); b {
		tx = tx.Where("expectedQuoteCompletionDate>=?", v)
	}

	if v, b := handler.VerifyField(params.QuoteCompletionDateLt); b {
		tx = tx.Where("expectedQuoteCompletionDate<=?", v)
	}

	if v, b := handler.VerifyField(params.QuoteDateGt); b {
		tx = tx.Where("quoteDate>=?", v)
	}

	if v, b := handler.VerifyField(params.QuoteDateLt); b {
		tx = tx.Where("quoteDate<=?", v)
	}

	if v, b := handler.VerifyField(params.Limit); b {
		tx = tx.Limit(v)
	}

	if v, b := handler.VerifyField(params.Offset); b {
		tx = tx.Offset(v)
	}
	var quotes []*schema.Quote
	if err := tx.Find(&quotes).Error; err == nil {
		var payload []*models.QuoteFind
		for _, q := range quotes {
			payload = append(payload, q.ToQuoteFind())
		}
		return quote.NewQuoteFindOK().WithPayload(payload)
	} else if err == gorm.ErrRecordNotFound {
		return quote.NewQuoteFindNotFound()
	} else {
		return quote.NewQuoteFindInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func QuoteQuoteGetHandler(params quote.QuoteGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return quote.NewQuoteGetUnauthorized().WithPayload(payload)
	}
	if q, err := db.GetQuote(params.ID); err == nil {
		return quote.NewQuoteGetOK().WithPayload(q.ToQuote())
	} else if err == gorm.ErrRecordNotFound {
		return quote.NewQuoteGetNotFound()
	} else {
		return quote.NewQuoteGetInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func HubQuoteManagementHubCreateHandler(params hub.QuoteManagementHubCreateParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewQuoteManagementHubCreateUnauthorized().WithPayload(payload)
	}

	// verify input
	input := params.Hub
	if input == nil {
		return hub.NewQuoteManagementHubCreateBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(21),
			Reason: swag.String("Missing body"),
		})
	}

	if input.Query == nil || input.Callback == nil {
		return hub.NewQuoteManagementHubCreateBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(23),
			Reason: swag.String("Missing body field"),
		})
	}

	if *input.Query == "" {
		return hub.NewQuoteManagementHubCreateBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(24),
			Reason: swag.String("Invalid body field"),
		})
	}

	if payload := handler.VerifyCallback(input.Callback); payload != nil {
		return hub.NewQuoteManagementHubCreateBadRequest().WithPayload(payload)
	}

	if id, err := quoteBus.Subscribe(handler.ParseType(*input.Query), commonCallbackHandler, &event.CallbackOption{}); err != nil {
		return hub.NewQuoteManagementHubCreateInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	} else {
		if err := db.AddSubscriber(&schema.HubSubscriber{
			ID:       id,
			Type:     quoteTopic,
			Query:    *input.Query,
			Callback: *input.Callback,
		}); err == nil {
			return hub.NewQuoteManagementHubCreateCreated().WithPayload(&models.Hub{
				Callback: input.Callback,
				ID:       swag.String(id),
				Query:    input.Query,
			})
		} else {
			return hub.NewQuoteManagementHubCreateInternalServerError().WithPayload(&models.ErrorRepresentation{
				Reason: swag.String(err.Error()),
			})
		}
	}
}

func HubQuoteManagementHubDeleteHandler(params hub.QuoteManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewQuoteManagementHubDeleteUnauthorized().WithPayload(payload)
	}
	// verify id
	id := params.HubID

	if s, err := db.FindSubscriber(id); err == nil {
		if err := quoteBus.Unsubscribe(handler.ParseType(s.Query), id); err == nil {
			if err := db.DeleteSubscriber(id); err == nil {
				return hub.NewProductOfferingQualificationManagementHubDeleteNoContent()
			} else {
				return hub.NewQuoteManagementHubDeleteInternalServerError().WithPayload(&models.ErrorRepresentation{
					Reason: swag.String(err.Error()),
				})
			}
		} else {
			return hub.NewQuoteManagementHubDeleteInternalServerError().WithPayload(&models.ErrorRepresentation{
				Reason: swag.String(err.Error()),
			})
		}
	} else if err == gorm.ErrRecordNotFound {
		return hub.NewQuoteManagementHubDeleteNotFound()
	} else {
		return hub.NewQuoteManagementHubDeleteInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func HubQuoteManagementHubFindHandler(params hub.QuoteManagementHubFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewQuoteManagementHubFindUnauthorized().WithPayload(payload)
	}

	if subscribers, err := db.ListSubscribers(quoteTopic); err == nil {
		var payload []*models.Hub
		for _, s := range subscribers {
			payload = append(payload, &models.Hub{
				Callback: swag.String(s.Callback),
				ID:       swag.String(s.ID),
				Query:    swag.String(s.Query),
			})
		}
		return hub.NewQuoteManagementHubFindOK().WithPayload(payload)
	} else if err == gorm.ErrRecordNotFound {
		return hub.NewQuoteManagementHubFindNotFound()
	} else {
		return hub.NewQuoteManagementHubFindInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func NotificationNotificationQuoteAttributeValueChangeNotificationHandler(params notification.NotificationQuoteAttributeValueChangeNotificationParams) middleware.Responder {
	log.Debugf("got NotificationQuoteAttributeValueChangeNotificationParams: %s", util.ToString(params.QuoteAttributeValueChangeNotification))
	return notification.NewNotificationQuoteAttributeValueChangeNotificationNoContent()
}

func NotificationNotificationQuoteCreationNotificationHandler(params notification.NotificationQuoteCreationNotificationParams) middleware.Responder {
	log.Debugf("got NotificationQuoteCreationNotificationParams:%s", util.ToString(params.QuoteCreationNotification))
	return notification.NewNotificationQuoteCreationNotificationNoContent()
}

func NotificationNotificationQuoteInformationRequiredNotificationHandler(params notification.NotificationQuoteInformationRequiredNotificationParams) middleware.Responder {
	log.Debugf("got NotificationQuoteInformationRequiredNotificationParams: %s", util.ToString(params.QuoteInformationRequiredNotification))
	return notification.NewNotificationProductOrderInformationRequiredNotificationNoContent()
}

func NotificationNotificationQuoteStateChangeNotificationHandler(params notification.NotificationQuoteStateChangeNotificationParams) middleware.Responder {
	log.Debugf("got NotificationQuoteStateChangeNotificationParams: %s", util.ToString(params.QuoteStateChangeNotification))
	return notification.NewNotificationQuoteStateChangeNotificationNoContent()
}

func commonCallbackHandler(option *event.CallbackOption, v interface{}) {
	if subscriber, err := db.FindSubscriber(option.ID); err == nil {
		if v != nil {
			if content, err := handler.HttpPost(subscriber.Callback, v); err != nil {
				log.Error(err)
			} else {
				log.Debug(content)
			}
		}
	} else {
		log.Error(err)
	}
}

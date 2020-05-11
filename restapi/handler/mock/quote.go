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
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"

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
		IDField:    *input.ID,
		StateField: models.QuoteStateType(input.State),
	}

	if err := Store.Save(state).Error; err == nil {
		// TODO: publish event
		quoteBus.Publish(string(models.QuoteEventTypeQUOTESTATECHANGENOTIFICATION), models.QuoteEventPlus{
			FieldPath:    nil,
			ResourcePath: nil,
		})
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
	id := xid.New().String()
	now := strfmt.DateTime(time.Now())
	var items []*schema.QuoteItem
	for _, item := range input.QuoteItem {
		to := &schema.QuoteItem{}
		if err := util.Convert(item, to); err == nil {
			to.ID = swag.String(xid.New().String())
			items = append(items, to)
		} else {
			logrus.Error(err)
		}
	}
	q := &schema.Quote{
		AtBaseTypeField:                   input.AtBaseType,
		AtSchemaLocationField:             input.AtSchemaLocation,
		Type:                              input.AtType,
		AgreementField:                    input.Agreement,
		DescriptionField:                  input.Description,
		ExpectedFulfillmentStartDateField: input.ExpectedFulfillmentStartDate,
		//TODO: maybe config
		ExpectedQuoteCompletionDateField:  strfmt.Date(time.Now().AddDate(0, 0, 5)),
		EffectiveQuoteCompletionDateField: strfmt.NewDateTime(),
		ExternalIDField:                   input.ExternalID,
		InstantSyncQuotingField:           input.InstantSyncQuoting,
		//NoteField:                         notes,
		ProjectIDField:                    input.ProjectID,
		QuoteItemField:                    items,
		QuoteLevelField:                   input.QuoteLevel,
		RelatedPartyField:                 input.RelatedParty,
		RequestedQuoteCompletionDateField: input.RequestedQuoteCompletionDate,
		HrefField:                         handler.HrefToID("", id),
		IDField:                           id,
		QuoteDateField:                    now,
		StateField:                        models.QuoteStateTypeACCEPTED,
		ValidForField: &models.TimePeriod{
			EndDateTime:   strfmt.DateTime(time.Now().AddDate(1, 1, 10)),
			StartDateTime: now,
		},
	}
	q.SetNote(input.Note)
	if err := Store.Save(q).Error; err == nil {
		//TODO: publish event
		quoteBus.Publish(string(models.QuoteEventTypeQUOTECREATIONNOTIFICATION), &models.QuoteEventPlus{
			FieldPath:    nil,
			ResourcePath: nil,
		})
		return quote.NewQuoteCreateCreated().WithPayload(q)
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
	tx := Store.Set(db.AutoPreLoad, true)
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
		var payload []models.QuoteFind
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
	if q, err := db.GetQuote(Store, params.ID); err == nil {
		return quote.NewQuoteGetOK().WithPayload(q)
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
		if err := db.AddSubscriber(Store, &schema.HubSubscriber{
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

	if s, err := db.FindSubscriber(Store, id); err == nil {
		if err := quoteBus.Unsubscribe(handler.ParseType(s.Query), id); err == nil {
			if err := db.DeleteSubscriber(Store, id); err == nil {
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

	if subscribers, err := db.ListSubscribers(Store, quoteTopic); err == nil {
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

func NotificationNotificationQuoteAttributeValueChangeNotificationHandler(params notification.NotificationQuoteAttributeValueChangeNotificationParams, principal *models.Principal) middleware.Responder {
	logrus.Debugf("got NotificationQuoteAttributeValueChangeNotificationParams: %s", util.ToString(params.QuoteAttributeValueChangeNotification))
	return notification.NewNotificationQuoteAttributeValueChangeNotificationNoContent()
}

func NotificationNotificationQuoteCreationNotificationHandler(params notification.NotificationQuoteCreationNotificationParams, principal *models.Principal) middleware.Responder {
	logrus.Debugf("got NotificationQuoteCreationNotificationParams:%s", util.ToString(params.QuoteCreationNotification))
	return notification.NewNotificationQuoteCreationNotificationNoContent()
}

func NotificationNotificationQuoteInformationRequiredNotificationHandler(params notification.NotificationQuoteInformationRequiredNotificationParams, principal *models.Principal) middleware.Responder {
	logrus.Debugf("got NotificationQuoteInformationRequiredNotificationParams: %s", util.ToString(params.QuoteInformationRequiredNotification))
	return notification.NewNotificationProductOrderInformationRequiredNotificationNoContent()
}

func NotificationNotificationQuoteStateChangeNotificationHandler(params notification.NotificationQuoteStateChangeNotificationParams, principal *models.Principal) middleware.Responder {
	logrus.Debugf("got NotificationQuoteStateChangeNotificationParams: %s", util.ToString(params.QuoteStateChangeNotification))
	return notification.NewNotificationQuoteStateChangeNotificationNoContent()
}

func commonCallbackHandler(option *event.CallbackOption, v interface{}) {
	if subscriber, err := db.FindSubscriber(Store, option.ID); err == nil {
		if v != nil {
			if content, err := handler.HttpPost(subscriber.Callback, v); err != nil {
				logrus.Error(err)
			} else {
				logrus.Debug(content)
			}
		}
	} else {
		logrus.Error(err)
	}
}

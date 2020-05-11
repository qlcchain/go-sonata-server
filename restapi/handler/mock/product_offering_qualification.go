/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"time"

	"github.com/qlcchain/go-sonata-server/event"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	poq "github.com/qlcchain/go-sonata-server/restapi/operations/product_offering_qualification"
	"github.com/qlcchain/go-sonata-server/schema"
	"github.com/qlcchain/go-sonata-server/util"
)

const (
	poqType = "poq"
)

var (
	poqBus = event.GetEventBus(poqType)
)

func ProductOfferingQualificationProductOfferingQualificationCreateHandler(params poq.ProductOfferingQualificationCreateParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return poq.NewProductOfferingQualificationCreateBadRequest().WithPayload(payload)
	}

	input := params.ProductOfferingQualification
	id := xid.New().String()
	var items []*schema.ProductOfferingQualificationItem
	for _, i := range input.ProductOfferingQualificationItem {
		p := &schema.Product{}
		if err := util.Convert(i.Product, p); err == nil {
			items = append(items, &schema.ProductOfferingQualificationItem{
				AtType:                   i.AtType,
				Action:                   i.Action,
				AlternateProductProposal: nil,
				EligibleProductOffering:  nil,
				GuaranteedUntilDate:      strfmt.DateTime(time.Now().AddDate(0, 0, 10)),
				ID:                       swag.String(xid.New().String()),
				InstallationInterval: &models.TimeInterval{
					Amount:   swag.Int32(15),
					TimeUnit: models.TimeUnitBusinessDays,
				},
				Product:         p,
				ProductOffering: i.ProductOffering,
				ProductOfferingQualificationItemRelationship: i.ProductOfferingQualificationItemRelationship,
				RelatedParty:             i.RelatedParty,
				ServiceConfidenceReason:  "green",
				ServiceabilityConfidence: models.ServiceabilityColorGreen,
				State:                    models.ProductOfferingQualificationItemStateTypeInProgress,
				StateChange: []*schema.StateChange{
					{
						ID:           swag.String(xid.New().String()),
						ChangeDate:   strfmt.DateTime(time.Now()),
						ChangeReason: "creat",
						State:        models.ProductOfferingQualificationStateTypeInProgress,
					},
				},
				TerminationError: nil,
			})
		} else {
			logrus.Error(err)
		}
	}

	qualification := &schema.ProductOfferingQualification{
		AtSchemaLocation:                     input.AtSchemaLocation,
		AtType:                               input.AtType,
		EffectiveQualificationCompletionDate: strfmt.DateTime(time.Now().AddDate(0, 0, 5)),
		ExpectedResponseDate:                 strfmt.DateTime{},
		Href:                                 id,
		ID:                                   swag.String(id),
		InstantSyncQualification:             input.InstantSyncQualification,
		ProductOfferingQualificationItem:     items,
		ProjectID:                            input.ProjectID,
		ProvideAlternative:                   false,
		RelatedParty:                         input.RelatedParty,
		RequestedResponseDate:                strfmt.DateTime(input.RequestedResponseDate),
		State:                                models.ProductOfferingQualificationStateTypeInProgress,
		StateChange: []*schema.StateChange{
			{
				ID:           swag.String(xid.New().String()),
				ChangeDate:   strfmt.DateTime(time.Now()),
				ChangeReason: "creat",
				State:        models.ProductOfferingQualificationStateTypeInProgress,
			},
		},
	}

	var err error
	if err = Store.Create(qualification).Error; err == nil {
		payload := &models.ProductOfferingQualification{}
		if err = util.Convert(qualification, payload); err == nil {
			now := strfmt.DateTime(time.Now())
			t := models.PoqEventTypeProductOfferingQualificationCreateEventNotification
			poqBus.Publish(string(t), &models.PoQEventContainer{
				Event: &models.PoqEvent{
					Href: handler.HrefToID("", swag.StringValue(payload.ID)),
					ID:   swag.StringValue(payload.ID),
				},
				EventID:   swag.String(xid.New().String()),
				EventTime: &now,
				EventType: t,
			})
			return poq.NewProductOfferingQualificationCreateCreated().WithPayload(payload)
		}
	}

	return poq.NewProductOfferingQualificationCreateServiceUnavailable().WithPayload(&models.ErrorRepresentation{
		Reason: swag.String(err.Error()),
	})
}

func ProductOfferingQualificationProductOfferingQualificationFindHandler(params poq.ProductOfferingQualificationFindParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return poq.NewProductOfferingQualificationFindBadRequest().WithPayload(payload)
	}
	var poqs []schema.ProductOfferingQualification
	tx := Store.Set(db.AutoPreLoad, true)
	if v, b := handler.VerifyField(params.ProjectID); b {
		tx = tx.Where("projectId=?", v)
	}
	if v, b := handler.VerifyField(params.State); b {
		tx = tx.Where("state=?", v)
	}

	if v, b := handler.VerifyField(params.RequestedResponseDateGt); b {
		tx = tx.Where("requestedResponseDate>=?", v)
	}

	if v, b := handler.VerifyField(params.RequestedResponseDateLt); b {
		tx = tx.Where("requestedResponseDate<=?", v)
	}

	if v, b := handler.VerifyField(params.Limit); b {
		tx = tx.Limit(v)
	}

	if v, b := handler.VerifyField(params.Offset); b {
		tx = tx.Offset(v)
	}

	if err := tx.Find(&poqs).Error; err == nil {
		var result []*models.ProductOfferingQualificationFind
		for _, p := range poqs {
			result = append(result, &models.ProductOfferingQualificationFind{
				ID:                    *p.ID,
				ProjectID:             p.ProjectID,
				RequestedResponseDate: strfmt.Date(p.RequestedResponseDate),
				State:                 p.State,
			})
		}
		return poq.NewProductOfferingQualificationFindOK().WithPayload(result)
	} else if err == gorm.ErrRecordNotFound {
		return poq.NewProductOfferingQualificationFindNotFound()
	} else {
		return poq.NewProductOfferingQualificationFindServiceUnavailable().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func ProductOfferingQualificationProductOfferingQualificationGetHandler(params poq.ProductOfferingQualificationGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return poq.NewProductOfferingQualificationGetUnauthorized().WithPayload(payload)
	}

	if p, err := db.GetProductOfferingQualification(Store, params.ProductOfferingQualificationID); err == nil {
		return poq.NewProductOfferingQualificationGetOK().WithPayload(p)
	} else if err == gorm.ErrRecordNotFound {
		return poq.NewProductOfferingQualificationGetNotFound()
	} else {
		return poq.NewProductOfferingQualificationGetServiceUnavailable().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func HubProductOfferingQualificationManagementHubDeleteHandler(params hub.ProductOfferingQualificationManagementHubDeleteParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewProductOfferingQualificationManagementHubDeleteUnauthorized().WithPayload(payload)
	}
	// verify id
	id := params.HubID

	if s, err := db.FindSubscriber(Store, id); err == nil {
		if err := poqBus.Unsubscribe(handler.ParseType(s.Query), id); err == nil {
			if err := db.DeleteSubscriber(Store, id); err == nil {
				return hub.NewProductOfferingQualificationManagementHubDeleteNoContent()
			} else {
				return hub.NewProductOrderManagementHubDeleteInternalServerError().WithPayload(&models.ErrorRepresentation{
					Reason: swag.String(err.Error()),
				})
			}
		} else {
			return hub.NewProductOrderManagementHubDeleteInternalServerError().WithPayload(&models.ErrorRepresentation{
				Reason: swag.String(err.Error()),
			})
		}
	} else if err == gorm.ErrRecordNotFound {
		return hub.NewProductOrderManagementHubDeleteNotFound()
	} else {
		return hub.NewProductOrderManagementHubDeleteInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func HubProductOfferingQualificationManagementHubGetHandler(params hub.ProductOfferingQualificationManagementHubGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewProductOfferingQualificationManagementHubGetUnauthorized().WithPayload(payload)
	}
	if subscribers, err := db.ListSubscribers(Store, poqType); err == nil {
		var payload []*models.Hub
		for _, s := range subscribers {
			payload = append(payload, &models.Hub{
				Callback: swag.String(s.Callback),
				ID:       swag.String(s.ID),
				Query:    swag.String(s.Query),
			})
		}
		return hub.NewProductOfferingQualificationManagementHubGetOK().WithPayload(payload)
	} else if err == gorm.ErrRecordNotFound {
		return hub.NewProductOfferingQualificationManagementHubGetNotFound()
	} else {
		return hub.NewProductOfferingQualificationManagementHubGetServiceUnavailable().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func HubProductOfferingQualificationManagementHubPostHandler(params hub.ProductOfferingQualificationManagementHubPostParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return hub.NewProductOfferingQualificationManagementHubPostUnauthorized().WithPayload(payload)
	}

	// verify input
	input := params.Hub
	if input == nil {
		return hub.NewProductOfferingQualificationManagementHubPostBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(21),
			Reason: swag.String("Missing body"),
		})
	}

	if input.Query == nil || input.Callback == nil {
		return hub.NewProductOfferingQualificationManagementHubPostBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(23),
			Reason: swag.String("Missing body field"),
		})
	}

	if *input.Query == "" {
		return hub.NewProductOfferingQualificationManagementHubPostBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(24),
			Reason: swag.String("Invalid body field"),
		})
	}

	if payload := handler.VerifyCallback(input.Callback); payload != nil {
		return hub.NewProductOfferingQualificationManagementHubPostBadRequest().WithPayload(payload)
	}

	if id, err := poqBus.Subscribe(handler.ParseType(*input.Query), commonCallbackHandler, &event.CallbackOption{}); err != nil {
		return hub.NewProductOfferingQualificationManagementHubPostServiceUnavailable().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	} else {
		if err := db.AddSubscriber(Store, &schema.HubSubscriber{
			ID:       id,
			Type:     poqType,
			Query:    *input.Query,
			Callback: *input.Callback,
		}); err == nil {
			return hub.NewProductOfferingQualificationManagementHubPostCreated().WithPayload(&models.Hub{
				Callback: input.Callback,
				ID:       swag.String(id),
				Query:    input.Query,
			})
		} else {
			return hub.NewProductOfferingQualificationManagementHubPostServiceUnavailable().WithPayload(&models.ErrorRepresentation{
				Reason: swag.String(err.Error()),
			})
		}
	}
}

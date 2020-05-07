package mock

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/restapi/operations/hub"
	poq "github.com/qlcchain/go-sonata-server/restapi/operations/product_offering_qualification"
	"github.com/qlcchain/go-sonata-server/schema"
	"github.com/qlcchain/go-sonata-server/util"
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
		var payload *models.ProductOfferingQualification
		if err = util.Convert(qualification, payload); err == nil {
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
	if err := Store.Set(db.AutoPreLoad, true).Where("", params).Find(&poqs).Error; err == nil {
		var result []*models.ProductOfferingQualificationFind
		for _, p := range poqs {
			result = append(result, &models.ProductOfferingQualificationFind{
				ID:                    *p.ID,
				ProjectID:             p.ProjectID,
				RequestedResponseDate: strfmt.Date(time.Now()),
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
		return poq.NewProductOfferingQualificationGetBadRequest().WithPayload(payload)
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
	return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubDelete has not yet been implemented")
}

func HubProductOfferingQualificationManagementHubGetHandler(params hub.ProductOfferingQualificationManagementHubGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubGet has not yet been implemented")
}

func HubProductOfferingQualificationManagementHubPostHandler(params hub.ProductOfferingQualificationManagementHubPostParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation hub.ProductOfferingQualificationManagementHubPost has not yet been implemented")
}

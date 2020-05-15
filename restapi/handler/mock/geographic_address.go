/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/util"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/restapi/handler"

	"github.com/qlcchain/go-sonata-server/models"
	ga "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address"
	gav "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address_validation"
)

func GeographicAddressGeographicAddressGetHandler(params ga.GeographicAddressGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return ga.NewGeographicAddressGetUnauthorized().WithPayload(payload)
	}
	if address, err := db.GetGeographicAddress(Store, params.GeographicAddressID); err == nil {
		return ga.NewGeographicAddressGetOK().WithPayload(address)
	} else if err == gorm.ErrRecordNotFound {
		return ga.NewGeographicAddressGetNotFound()
	} else {
		return ga.NewGeographicAddressGetInternalServerError().WithPayload(&models.ErrorRepresentation{
			Reason: swag.String(err.Error()),
		})
	}
}

func GeographicAddressValidationGeographicAddressValidationCreateHandler(params gav.GeographicAddressValidationCreateParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return gav.NewGeographicAddressValidationCreateUnauthorized().WithPayload(payload)
	}
	input := params.AddressValidationRequest.RequestedAddress

	var verifiedAddress []*models.GeographicAddressSellerResponse
	filter := make(map[string]interface{})
	result := models.ValidationResultFails

	if address, err := db.GetGeographicAddressByFieldedAddress(Store, input.FieldedAddress); err == nil {
		if len(address) > 0 {
			result = models.ValidationResultPartial
		}

		for _, a := range address {
			if _, ok := filter[a.ID]; !ok {
				verifiedAddress = append(verifiedAddress, &models.GeographicAddressSellerResponse{
					AtSchemaLocation: a.AtSchemaLocation,
					AtType:           a.AtType,
					AllowsNewSite:    a.AllowsNewSite,
					FieldedAddress:   a.FieldedAddress.To(),
					FormattedAddress: a.FormattedAddress.To(),
					HasPublicSite:    a.HasPublicSite,
					ID:               a.ID,
					IsBestMatch:      swag.Bool(true),
				})
				filter[a.ID] = struct{}{}
			}
		}
	} else {
		log.Error(err)
	}

	if address, err := db.GetGeographicAddressByFormattedAddress(Store, input.FormattedAddress); err == nil {
		if len(address) > 0 {
			if result == models.ValidationResultPartial {
				result = models.ValidationResultSuccess
			} else {
				result = models.ValidationResultPartial
			}
		}

		for _, a := range address {
			if _, ok := filter[a.ID]; !ok {
				verifiedAddress = append(verifiedAddress, &models.GeographicAddressSellerResponse{
					AtSchemaLocation: a.AtSchemaLocation,
					AtType:           a.AtType,
					AllowsNewSite:    a.AllowsNewSite,
					FieldedAddress:   a.FieldedAddress.To(),
					FormattedAddress: a.FormattedAddress.To(),
					HasPublicSite:    a.HasPublicSite,
					ID:               a.ID,
					IsBestMatch:      swag.Bool(true),
				})
				filter[a.ID] = struct{}{}
			}
		}
	} else {
		log.Error(err)
	}

	// TODO: fuzzy query when can not fetch any records by best match
	if len(verifiedAddress) == 0 {
		return gav.NewGeographicAddressValidationCreateNotFound()
	} else {
		return gav.NewGeographicAddressValidationCreateCreated().WithPayload(&models.GeographicAddressValidation{
			ID: util.NewID(),
			RequestedAddress: &models.GeographicAddressRequestBuyerInput{
				FieldedAddress:   input.FieldedAddress,
				FormattedAddress: input.FormattedAddress,
			},
			ValidationDate:   strfmt.DateTime(time.Now()),
			ValidationResult: result,
			VerifiedAddress:  verifiedAddress,
		})

	}

}

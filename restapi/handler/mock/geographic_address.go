/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"time"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"

	"github.com/qlcchain/go-sonata-server/restapi/handler"

	"github.com/qlcchain/go-sonata-server/models"
	ga "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address"
	gav "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address_validation"
)

func GeographicAddressGeographicAddressGetHandler(params ga.GeographicAddressGetParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return ga.NewGeographicAddressGetBadRequest().WithPayload(payload)
	}
	address := &models.GeographicAddress{}
	if err := Store.Set(db.AutoPreLoad, true).First(address, params.GeographicAddressID).Error; err == gorm.ErrRecordNotFound {
		return ga.NewGeographicAddressGetNotFound()
	}

	return ga.NewGeographicAddressGetOK().WithPayload(address)
}

func GeographicAddressValidationGeographicAddressValidationCreateHandler(params gav.GeographicAddressValidationCreateParams, principal *models.Principal) middleware.Responder {
	if payload := handler.ToErrorRepresentation(principal); payload != nil {
		return gav.NewGeographicAddressValidationCreateBadRequest().WithPayload(payload)
	}
	input := params.AddressValidationRequest.RequestedAddress

	var verifiedAddress []*models.GeographicAddressSellerResponse
	filter := make(map[string]interface{})
	result := models.ValidationResultFails

	if fa := handler.FieldedAddressRequest2FieldedAddress(input.FieldedAddress); fa != nil {
		var address []models.GeographicAddress
		if err := Store.Set(db.AutoPreLoad, true).Model(fa).Related(&address, "ID"); err == nil {
			if len(address) > 0 {
				result = models.ValidationResultPartial
			}
		}
		for _, a := range address {
			if _, ok := filter[a.ID]; !ok {
				verifiedAddress = append(verifiedAddress, &models.GeographicAddressSellerResponse{
					AtSchemaLocation: a.AtSchemaLocation,
					AtType:           a.AtType,
					AllowsNewSite:    a.AllowsNewSite,
					FieldedAddress:   a.FieldedAddress,
					FormattedAddress: a.FormattedAddress,
					HasPublicSite:    a.HasPublicSite,
					ID:               a.ID,
					IsBestMatch:      swag.Bool(true),
				})
				filter[a.ID] = struct{}{}
			}
		}
	}

	if fa := handler.FormattedAddressRequest2FormattedAddress(input.FormattedAddress); fa != nil {
		var address []models.GeographicAddress
		if err := Store.Set(db.AutoPreLoad, true).Model(fa).Related(&address, "ID"); err == nil {
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
						FieldedAddress:   a.FieldedAddress,
						FormattedAddress: a.FormattedAddress,
						HasPublicSite:    a.HasPublicSite,
						ID:               a.ID,
						IsBestMatch:      swag.Bool(true),
					})
					filter[a.ID] = struct{}{}
				}
			}
		}
	}

	return gav.NewGeographicAddressValidationCreateCreated().WithPayload(&models.GeographicAddressValidation{
		ID: xid.New().String(),
		RequestedAddress: &models.GeographicAddressRequestBuyerInput{
			FieldedAddress:   input.FieldedAddress,
			FormattedAddress: input.FormattedAddress,
		},
		ValidationDate:   strfmt.DateTime(time.Now()),
		ValidationResult: result,
		VerifiedAddress:  verifiedAddress,
	})
}

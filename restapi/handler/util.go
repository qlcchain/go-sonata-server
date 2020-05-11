package handler

import (
	"net/url"

	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-sonata-server/models"
)

func ToErrorRepresentation(principal *models.Principal) *models.ErrorRepresentation {
	if principal.Code != 0 {
		return &models.ErrorRepresentation{
			Code:   swag.Int32(principal.Code),
			Reason: swag.String(principal.Reason),
		}
	}
	return nil
}

func VerifyCallback(callback *string) *models.ErrorRepresentation {
	if _, err := url.Parse(*callback); err != nil {
		return &models.ErrorRepresentation{
			Code:   swag.Int32(24),
			Reason: swag.String("Invalid  body field"),
		}
	}
	return nil
}

func FieldedAddressRequest2FieldedAddress(req *models.FieldedAddressRequest) *models.FieldedAddress {
	if req == nil {
		return nil
	}

	var a []*models.GeographicSubAddress
	sub := req.GeographicSubAddress
	if sub != nil {
		a = append(a, &models.GeographicSubAddress{
			AtSchemaLocation:    sub.AtSchemaLocation,
			AtType:              sub.AtType,
			BuildingName:        sub.BuildingName,
			LevelNumber:         sub.LevelNumber,
			LevelType:           sub.LevelType,
			PrivateStreetName:   sub.PrivateStreetName,
			PrivateStreetNumber: sub.PrivateStreetNumber,
			SubUnit:             sub.SubUnit,
		})
	}

	return &models.FieldedAddress{
		City:                 req.City,
		Country:              req.Country,
		GeographicSubAddress: a,
		Locality:             req.Locality,
		PostCodeExtension:    req.PostCodeExtension,
		Postcode:             req.Postcode,
		StateOrProvince:      req.StateOrProvince,
		StreetName:           req.StreetName,
		StreetNr:             req.StreetNr,
		StreetNrLast:         req.StreetNrLast,
		StreetNrLastSuffix:   req.StreetNrLastSuffix,
		StreetNrSuffix:       req.StreetNrSuffix,
		StreetSuffix:         req.StreetSuffix,
		StreetType:           req.StreetType,
	}
}

func FormattedAddressRequest2FormattedAddress(req *models.FormattedAddressRequest) *models.FormattedAddress {
	if req == nil {
		return nil
	}

	return &models.FormattedAddress{
		AddrLine1:         req.AddrLine1,
		AddrLine2:         req.AddrLine2,
		City:              req.City,
		Country:           req.Country,
		Locality:          req.Locality,
		PostCodeExtension: req.PostCodeExtension,
		Postcode:          req.Postcode,
		StateOrProvince:   req.StateOrProvince,
	}
}

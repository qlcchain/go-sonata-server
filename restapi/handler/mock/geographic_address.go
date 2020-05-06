package mock

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/qlcchain/go-sonata-server/models"
	ga "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address"
	gav "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address_validation"
)

func GeographicAddressGeographicAddressGetHandler(params ga.GeographicAddressGetParams, principal *models.Principal) middleware.Responder {
	if principal.Code != 0 {
		return ga.NewGeographicAddressGetBadRequest().WithPayload(&models.ErrorRepresentation{
			Code:   swag.Int32(principal.Code),
			Reason: swag.String(principal.Reason),
		})
	}
	address := &db.GeographicAddressModel{}
	if err := DB.First(address, params.GeographicAddressID).Error; err == gorm.ErrRecordNotFound {
		return ga.NewGeographicAddressGetNotFound()
	}

	return ga.NewGeographicAddressGetOK().WithPayload(&models.GeographicAddress{})
}

func GeographicAddressValidationGeographicAddressValidationCreateHandler(params gav.GeographicAddressValidationCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation geographic_address_validation.GeographicAddressValidationCreate has not yet been implemented")
}

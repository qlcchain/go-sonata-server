package mock

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
	ga "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address"
	gav "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_address_validation"
)

func GeographicAddressGeographicAddressGetHandler(params ga.GeographicAddressGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation geographic_address.GeographicAddressGet has not yet been implemented")
}

func GeographicAddressValidationGeographicAddressValidationCreateHandler(params gav.GeographicAddressValidationCreateParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation geographic_address_validation.GeographicAddressValidationCreate has not yet been implemented")
}

package mock

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/operations/product"
)

func ProductProductFindHandler(params product.ProductFindParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation product.ProductFind has not yet been implemented")
}

func ProductProductGetHandler(params product.ProductGetParams, principal *models.Principal) middleware.Responder {
	return middleware.NotImplemented("operation product.ProductGet has not yet been implemented")
}

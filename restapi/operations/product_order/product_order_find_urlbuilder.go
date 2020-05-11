// Code generated by go-swagger; DO NOT EDIT.

package product_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/strfmt"
)

// ProductOrderFindURL generates an URL for the product order find operation
type ProductOrderFindURL struct {
	BuyerID                   *string
	BuyerRequestedDateGt      *strfmt.DateTime
	BuyerRequestedDateLt      *strfmt.DateTime
	CompletionDateGt          *strfmt.DateTime
	CompletionDateLt          *strfmt.DateTime
	ExpectedCompletionDateGt  *strfmt.DateTime
	ExpectedCompletionDateLt  *strfmt.DateTime
	ExternalID                *string
	Limit                     *string
	Offset                    *string
	OrderCancellationDateGt   *strfmt.DateTime
	OrderCancellationDateLt   *strfmt.DateTime
	OrderDateGt               *strfmt.DateTime
	OrderDateLt               *strfmt.DateTime
	ProjectID                 *string
	RequestedCompletionDateGt *strfmt.DateTime
	RequestedCompletionDateLt *strfmt.DateTime
	RequestedStartDateGt      *strfmt.DateTime
	RequestedStartDateLt      *strfmt.DateTime
	SellerID                  *string
	SiteCompanyName           *string
	SiteCustomerName          *string
	SiteName                  *string
	State                     *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ProductOrderFindURL) WithBasePath(bp string) *ProductOrderFindURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ProductOrderFindURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ProductOrderFindURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/productOfferingQualificationNotification/v3/productOrderManagement/v3/productOrder"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/mef"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var buyerIDQ string
	if o.BuyerID != nil {
		buyerIDQ = *o.BuyerID
	}
	if buyerIDQ != "" {
		qs.Set("buyerId", buyerIDQ)
	}

	var buyerRequestedDateGtQ string
	if o.BuyerRequestedDateGt != nil {
		buyerRequestedDateGtQ = o.BuyerRequestedDateGt.String()
	}
	if buyerRequestedDateGtQ != "" {
		qs.Set("buyerRequestedDate.gt", buyerRequestedDateGtQ)
	}

	var buyerRequestedDateLtQ string
	if o.BuyerRequestedDateLt != nil {
		buyerRequestedDateLtQ = o.BuyerRequestedDateLt.String()
	}
	if buyerRequestedDateLtQ != "" {
		qs.Set("buyerRequestedDate.lt", buyerRequestedDateLtQ)
	}

	var completionDateGtQ string
	if o.CompletionDateGt != nil {
		completionDateGtQ = o.CompletionDateGt.String()
	}
	if completionDateGtQ != "" {
		qs.Set("completionDate.gt", completionDateGtQ)
	}

	var completionDateLtQ string
	if o.CompletionDateLt != nil {
		completionDateLtQ = o.CompletionDateLt.String()
	}
	if completionDateLtQ != "" {
		qs.Set("completionDate.lt", completionDateLtQ)
	}

	var expectedCompletionDateGtQ string
	if o.ExpectedCompletionDateGt != nil {
		expectedCompletionDateGtQ = o.ExpectedCompletionDateGt.String()
	}
	if expectedCompletionDateGtQ != "" {
		qs.Set("expectedCompletionDate.gt", expectedCompletionDateGtQ)
	}

	var expectedCompletionDateLtQ string
	if o.ExpectedCompletionDateLt != nil {
		expectedCompletionDateLtQ = o.ExpectedCompletionDateLt.String()
	}
	if expectedCompletionDateLtQ != "" {
		qs.Set("expectedCompletionDate.lt", expectedCompletionDateLtQ)
	}

	var externalIDQ string
	if o.ExternalID != nil {
		externalIDQ = *o.ExternalID
	}
	if externalIDQ != "" {
		qs.Set("externalId", externalIDQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = *o.Limit
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = *o.Offset
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var orderCancellationDateGtQ string
	if o.OrderCancellationDateGt != nil {
		orderCancellationDateGtQ = o.OrderCancellationDateGt.String()
	}
	if orderCancellationDateGtQ != "" {
		qs.Set("orderCancellationDate.gt", orderCancellationDateGtQ)
	}

	var orderCancellationDateLtQ string
	if o.OrderCancellationDateLt != nil {
		orderCancellationDateLtQ = o.OrderCancellationDateLt.String()
	}
	if orderCancellationDateLtQ != "" {
		qs.Set("orderCancellationDate.lt", orderCancellationDateLtQ)
	}

	var orderDateGtQ string
	if o.OrderDateGt != nil {
		orderDateGtQ = o.OrderDateGt.String()
	}
	if orderDateGtQ != "" {
		qs.Set("orderDate.gt", orderDateGtQ)
	}

	var orderDateLtQ string
	if o.OrderDateLt != nil {
		orderDateLtQ = o.OrderDateLt.String()
	}
	if orderDateLtQ != "" {
		qs.Set("orderDate.lt", orderDateLtQ)
	}

	var projectIDQ string
	if o.ProjectID != nil {
		projectIDQ = *o.ProjectID
	}
	if projectIDQ != "" {
		qs.Set("projectId", projectIDQ)
	}

	var requestedCompletionDateGtQ string
	if o.RequestedCompletionDateGt != nil {
		requestedCompletionDateGtQ = o.RequestedCompletionDateGt.String()
	}
	if requestedCompletionDateGtQ != "" {
		qs.Set("requestedCompletionDate.gt", requestedCompletionDateGtQ)
	}

	var requestedCompletionDateLtQ string
	if o.RequestedCompletionDateLt != nil {
		requestedCompletionDateLtQ = o.RequestedCompletionDateLt.String()
	}
	if requestedCompletionDateLtQ != "" {
		qs.Set("requestedCompletionDate.lt", requestedCompletionDateLtQ)
	}

	var requestedStartDateGtQ string
	if o.RequestedStartDateGt != nil {
		requestedStartDateGtQ = o.RequestedStartDateGt.String()
	}
	if requestedStartDateGtQ != "" {
		qs.Set("requestedStartDate.gt", requestedStartDateGtQ)
	}

	var requestedStartDateLtQ string
	if o.RequestedStartDateLt != nil {
		requestedStartDateLtQ = o.RequestedStartDateLt.String()
	}
	if requestedStartDateLtQ != "" {
		qs.Set("requestedStartDate.lt", requestedStartDateLtQ)
	}

	var sellerIDQ string
	if o.SellerID != nil {
		sellerIDQ = *o.SellerID
	}
	if sellerIDQ != "" {
		qs.Set("sellerId", sellerIDQ)
	}

	var siteCompanyNameQ string
	if o.SiteCompanyName != nil {
		siteCompanyNameQ = *o.SiteCompanyName
	}
	if siteCompanyNameQ != "" {
		qs.Set("siteCompanyName", siteCompanyNameQ)
	}

	var siteCustomerNameQ string
	if o.SiteCustomerName != nil {
		siteCustomerNameQ = *o.SiteCustomerName
	}
	if siteCustomerNameQ != "" {
		qs.Set("siteCustomerName", siteCustomerNameQ)
	}

	var siteNameQ string
	if o.SiteName != nil {
		siteNameQ = *o.SiteName
	}
	if siteNameQ != "" {
		qs.Set("siteName", siteNameQ)
	}

	var stateQ string
	if o.State != nil {
		stateQ = *o.State
	}
	if stateQ != "" {
		qs.Set("state", stateQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ProductOrderFindURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ProductOrderFindURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ProductOrderFindURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ProductOrderFindURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ProductOrderFindURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ProductOrderFindURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}

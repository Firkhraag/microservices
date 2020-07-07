// Package classification Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta
package handlers

import (
	"github.com/firkhraag/product-api/internal/model"
)

// This API endpoint returns a list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []model.Product
}

// This API endpoint returns a product
// swagger:response productResponse
type productResponseWrapper struct {
	// All products in the system
	// in: body
	Body model.Product
}

// This API endpoint returns no content
// swagger:response noContentResponse
type noContentResponseWrapper struct{}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// swagger:parameters deleteProduct getProduct
type productIDParameterWrapper struct {
	// The id of the product
	// in: path
	// required: true
	ID int `json:"id"`
}

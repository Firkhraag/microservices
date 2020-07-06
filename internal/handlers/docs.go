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
	"Projects/microservices/internal/model"
)

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []model.Product
}

// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

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

// swagger:parameters deleteProduct
type productDeleteIDParameterWrapper struct {
	// The id of the product to delete from database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updateProduct
type productUpdateIDParameterWrapper struct {
	// The id of the product to update in database
	// in: path
	// required: true
	ID int `json:"id"`
}

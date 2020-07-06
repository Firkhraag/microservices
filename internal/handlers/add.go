package handlers

import (
	"Projects/microservices/internal/model"
	"net/http"
)

// swagger:route POST /products products addProduct
// Adds a product
// responses:
//    200: noContentResponse
//  400: errorValidation
//  422: errorResponse

// AddProduct adds a product to the database
func (ph *Products) AddProduct(w http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProduct{}).(model.Product)

	ph.l.Println("[DEBUG] adding a new product", prod.Name)
	model.AddProduct(&prod)
}

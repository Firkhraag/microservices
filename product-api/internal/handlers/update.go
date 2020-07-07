package handlers

import (
	"net/http"

	"github.com/firkhraag/product-api/internal/model"
)

// swagger:route PUT /products products updateProduct
// Updates a product
// responses:
//    200: noContentResponse
//  400: errorValidation
//  404: errorResponse
//  422: errorValidation

// UpdateProduct updates a product in the database
func (ph *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	ph.l.Println("[DEBUG] updating the product")

	prod := r.Context().Value(KeyProduct{}).(model.Product)
	if err := model.UpdateProduct(&prod); err != nil {
		if err == model.ErrProductNotFound {
			ph.l.Println("[ERROR] Product not found", err)
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		ph.l.Println("[ERROR] Product not found", err)
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

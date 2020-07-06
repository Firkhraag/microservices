package handlers

import (
	"Projects/microservices/internal/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProduct
// Updates a product
// responses:
//    201: noContentResponse
//  400: errorValidation
//  404: errorResponse
//  422: errorValidation

// UpdateProduct updates a product in the database
func (ph *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		ph.l.Println("[ERROR] Id is not a number")
		http.Error(w, "Id is not a number", http.StatusNotFound)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(model.Product)
	prod.ID = id
	ph.l.Println("[DEBUG] updating record id", prod.ID)

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

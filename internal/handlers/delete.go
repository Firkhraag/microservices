package handlers

import (
	"Projects/microservices/internal/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//    200: noContentResponse
//  404: errorResponse

// DeleteProduct deletes a product from the database
func (ph *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		ph.l.Println("[ERROR] id is not a number", err)
		http.Error(w, "Id is not a number", http.StatusNotFound)
		return
	}

	ph.l.Println("[DEBUG] deleting record id", id)

	if err := model.DeleteProduct(id); err != nil {
		ph.l.Println("[ERROR] Product not found", err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
}

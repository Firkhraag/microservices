package handlers

import (
	"Projects/microservices/internal/model"
	"Projects/microservices/internal/util"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsResponse

// GetProducts responds with a list of products
func (ph *Products) GetProducts(w http.ResponseWriter, r *http.Request) {

	ph.l.Println("[DEBUG] getting all products")

	// fetch the products from the datastore
	lp := model.GetProducts()

	// serialize the list to JSON
	if err := util.ToJSON(lp, w); err != nil {
		ph.l.Println("[ERROR] unable to marshal json", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

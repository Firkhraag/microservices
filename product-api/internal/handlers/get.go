package handlers

import (
	"net/http"

	"github.com/firkhraag/product-api/internal/util"

	"github.com/firkhraag/product-api/internal/model"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//    200: productsResponse
//  500: errorResponse

// GetProducts responds with a list of products
func (ph *Products) GetProducts(w http.ResponseWriter, r *http.Request) {

	ph.l.Println("[DEBUG] getting all products")

	// fetch the products from the datastore
	lp := model.GetProducts()

	w.Header().Add("Content-Type", "application/json")

	// serialize the list to JSON
	if err := util.ToJSON(lp, w); err != nil {
		ph.l.Println("[ERROR] unable to marshal json", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /{id} products getProduct
// Returns a product
//
// responses:
//    200: productResponse
//  404: errorResponse
//  500: errorResponse

// GetProduct not found returns a product
func (ph *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	ph.l.Println("[DEBUG] getting a product")

	id := getProductIDFromURI(r)

	p, err := model.GetProductByID(id)
	if err != nil {
		ph.l.Println("[ERROR] product not found", err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	if err := util.ToJSON(p, w); err != nil {
		ph.l.Println("[ERROR] unable to marshal json", err)
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

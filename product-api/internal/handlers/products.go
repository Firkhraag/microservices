package handlers

import (
	"log"
	"net/http"

	"github.com/firkhraag/product-api/internal/model"
	"github.com/firkhraag/product-api/internal/util"
)

// Products handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//---------------------------- GET ----------------------------

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

//---------------------------- POST ----------------------------

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

//---------------------------- PUT ----------------------------

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

//---------------------------- DELETE ----------------------------

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//    200: noContentResponse
//  404: errorResponse

// DeleteProduct deletes a product from the database
func (ph *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	ph.l.Println("[DEBUG] deleting the product")

	id := getProductIDFromURI(r)
	if err := model.DeleteProduct(id); err != nil {
		ph.l.Println("[ERROR] Product not found", err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
}

package handlers

import (
	"net/http"

	"github.com/firkhraag/product-api/internal/model"
)

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

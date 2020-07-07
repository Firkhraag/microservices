package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/firkhraag/product-api/internal/model"
	"github.com/firkhraag/product-api/internal/util"
)

// KeyProduct key of Product in Context
type KeyProduct struct{}

// MiddlewareProductValidation validates a product
func (ph *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := model.Product{}
		if err := util.FromJSON(prod, r.Body); err != nil {
			ph.l.Println("[ERROR] Unable to unmarshal json", err)
			http.Error(w, "Unable to unmarshal json", http.StatusUnprocessableEntity)
			return
		}
		v := model.NewValidation()
		if err := v.Validate(prod); err != nil {
			ph.l.Println("[ERROR] Failed validation", err)
			http.Error(w, fmt.Sprintf("Failed validation %s", err), http.StatusBadRequest)
			return
		}

		req := r.WithContext(context.WithValue(r.Context(), KeyProduct{}, prod))
		next.ServeHTTP(w, req)
	})
}

package main

import (
	"testing"

	"github.com/firkhraag/product-api/sdk/client"
	"github.com/firkhraag/product-api/sdk/client/products"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	_, err := c.Products.ListProducts(params)
	assert.NoError(t, err)
}

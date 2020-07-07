package model

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the user
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// the name for the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for the product
	//
	// required: false
	// max length: 10000
	Description string `json:"decription"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}

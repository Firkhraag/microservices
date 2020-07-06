package model

// Products defines the list of Product
type Products []*Product

// GetProducts returns all products
func GetProducts() Products {
	return productList
}

// AddProduct adds product to the list
func AddProduct(p *Product) {
	p.ID = len(productList) + 1
	productList = append(productList, p)
}

// UpdateProduct updates the product in the list
func UpdateProduct(p *Product) error {
	id := findIndexByProductID(p.ID)
	if id == -1 {
		return ErrProductNotFound
	}

	productList[id] = p
	return nil
}

// DeleteProduct deletes the product in the list
func DeleteProduct(id int) error {
	if i := findIndexByProductID(id); i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:id], productList[id+1])
	return nil
}

// findIndexByProductID finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc-abc-abc",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd-fgh-jut",
	},
}

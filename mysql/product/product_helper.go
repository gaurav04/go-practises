package product

import (
	"fmt"

	"github.com/gaurav04/go-practises/mysql/database"
)

func getProductList() ([]Product, error) {
	var products []Product
	results, err := database.DbConn.Query("SELECT * FROM products LIMIT 10")
	if err != nil {
		fmt.Println(err)
	}

	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)
	}
	return products, nil
}

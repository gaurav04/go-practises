package product

import (
	"fmt"
	"log"

	"github.com/gaurav04/go-practises/mysql/database"
)

func getProductList() ([]Product, error) {
	var products []Product
	results, err := database.DbConn.Query("SELECT * FROM products ORDER BY productid DESC LIMIT 10")
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

func removeProduct(prodid int) error {
	_, err := database.DbConn.Query("DELETE FROM products where productId = ?", prodid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func insertProduct(product Product) (int, error) {
	stmt, err := database.DbConn.Prepare("INSERT INTO products(manufacturer,sku,upc,pricePerUnit,quantityOnHand,productName) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	result, err := stmt.Exec(product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName)

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertID), nil

}

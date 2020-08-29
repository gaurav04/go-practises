package product

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gaurav04/go-practises/inventroy_mgmt/database"
)

func getProductList() ([]Product, error) {
	var products []Product
	fmt.Println(database.DbConn)
	results, err := database.DbConn.Query("SELECT * FROM products ORDER BY productid DESC LIMIT 10")

	if err != nil {
		fmt.Println(err)
	}

	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)
	}
	fmt.Println(products)
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
	stmt, err := database.DbConn.Prepare("INSERT INTO products(manufacturer,pricePerUnit,quantityOnHand,productName) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	result, err := stmt.Exec(product.Manufacturer,
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

func getProduct(productID int) (*Product, error) {
	row := database.DbConn.QueryRow(`SELECT 
	productId, 
	manufacturer, 
	pricePerUnit, 
	quantityOnHand, 
	productName 
	FROM products 
	WHERE productId = ?`, productID)

	product := &Product{}
	err := row.Scan(
		&product.ProductID,
		&product.Manufacturer,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.ProductName,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}
func updateProduct(product Product) error {
	_, err := database.DbConn.Query(`UPDATE products SET 
		manufacturer=?, 
		pricePerUnit=?, 
		quantityOnHand=?, 
		productName=?
		WHERE productId=?`,
		product.Manufacturer,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

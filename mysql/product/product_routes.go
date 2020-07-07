package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func productshandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList, err := getProductList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.MarshalIndent(productList, "", "   ")
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPost:
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		productID, err := insertProduct(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"productId":%d}`, productID)))

	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	fmt.Println(urlPathSegments)
	productID, _ := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])

	switch r.Method {
	case http.MethodDelete:
		err := removeProduct(productID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}

func SetupRoutes() {
	http.HandleFunc("/products", productshandler)
	http.HandleFunc("/products/", productHandler)
}

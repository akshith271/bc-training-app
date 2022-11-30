package Products

import (
	"encoding/json"
	"net/http"
	db "sample/git/connection"
	"sample/git/models"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	products := []models.Product{}
	db.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	product := []models.Product{}
	params := mux.Vars(r)["productid"]
	db.Find(&product).First(&product, params)
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	product := models.Product{}
	json.NewDecoder(r.Body).Decode(&product)
	db.Save(&product)
	json.NewEncoder(w).Encode(product)
}

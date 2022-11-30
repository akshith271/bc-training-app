package Products

import (
	"encoding/json"
	"net/http"
	"sample/git/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	Db *gorm.DB
}

func (s Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{}
	s.Db.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (s Server) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	product := []models.Product{}
	params := mux.Vars(r)["productid"]
	s.Db.Find(&product).First(&product, params)
	json.NewEncoder(w).Encode(product)
}

func (s Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	product := models.Product{}
	json.NewDecoder(r.Body).Decode(&product)
	s.Db.Save(&product)
	json.NewEncoder(w).Encode(product)
}

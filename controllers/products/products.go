package Products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample/git/database"
	"sample/git/models"
)

type Server struct {
	Db database.DbProducts
}

type JsonResponse struct {
	Type string
	Data []models.Product
}

func (s Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := s.Db.GetProducts()
	if err != nil {
		fmt.Printf("Thers an error")
	}
	response := JsonResponse{Type: "success", Data: result}

	json.NewEncoder(w).Encode(response)
}

// func (s Server) GetProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	product := []models.Product{}
// 	params := mux.Vars(r)["productid"]
// 	s.Db.Find(&product).First(&product, params)
// 	json.NewEncoder(w).Encode(product)
// }

// func (s Server) CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	product := models.Product{}
// 	json.NewDecoder(r.Body).Decode(&product)
// 	s.Db.Save(&product)
// 	json.NewEncoder(w).Encode(product)
// }

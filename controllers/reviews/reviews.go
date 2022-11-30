package reviews

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

func (s Server) CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	review := models.Rating{}
	json.NewDecoder(r.Body).Decode(&review)
	s.Db.Save(&review)
	json.NewEncoder(w).Encode(review)
}

func (s Server) GetReviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reviews := []models.Rating{}
	params := mux.Vars(r)["productid"]
	s.Db.Where("product_id=?", params).Find(&reviews)
	result := []string{}
	for _, value := range reviews {
		result = append(result, value.Review)
	}
	json.NewEncoder(w).Encode(result)
}

func (s Server) DeleteReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ratings := []models.Rating{}
	productid := mux.Vars(r)["product_id"]
	ratingid := mux.Vars(r)["name"]
	s.Db.Where("name = ? AND product_id = ?", ratingid, productid).Delete(&ratings)
	json.NewEncoder(w).Encode(&ratings)

}

package reviews

import (
	"encoding/json"
	"net/http"
	db "sample/git/connection"
	"sample/git/models"

	"github.com/gorilla/mux"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	review := models.Rating{}
	json.NewDecoder(r.Body).Decode(&review)
	db.Save(&review)
	json.NewEncoder(w).Encode(review)
}

func GetReviews(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	reviews := []models.Rating{}
	params := mux.Vars(r)["productid"]
	db.Where("product_id=?", params).Find(&reviews)
	result := []string{}
	for _, value := range reviews {
		result = append(result, value.Review)
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	ratings := []models.Rating{}
	productid := mux.Vars(r)["product_id"]
	ratingid := mux.Vars(r)["name"]
	db.Where("name = ? AND product_id = ?", ratingid, productid).Delete(&ratings)
	json.NewEncoder(w).Encode(&ratings)

}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sample/git/authorization"
	DB "sample/git/connection"
	product "sample/git/controllers/products"
	database "sample/git/database"
	token "sample/git/token-generator"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	token.GenerateJWt()
	db_name := flag.String("db", "sample", "give db name")
	db, err := gorm.Open("postgres", fmt.Sprintf("user=postgres password=root dbname=%v sslmode=disable", *db_name))
	DB.CheckErr(err)
	defer db.Close()
	var product = product.Server{
		Db: database.DbClient{
			Db: db},
	}
	// var review reviews.Server = reviews.Server{db}
	r := mux.NewRouter() //router init

	r.Handle("/products", authorization.IsAuthorized(product.GetProducts)).Methods("GET")
	// r.HandleFunc("/products/{productid}", product.GetProduct).Methods("GET")
	// r.HandleFunc("/products/create", product.CreateProduct).Methods("POST")
	// r.HandleFunc("/products/{productid}/reviews", review.GetReviews).Methods("GET")
	// r.HandleFunc("/products/{productid}/reviews/create", review.CreateReview).Methods("POST")

	fmt.Printf("starting port at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	DB "sample/git/connection"
	products "sample/git/controllers/products"
	reviews "sample/git/controllers/reviews"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db_name := flag.String("db", "product", "give db name")
	db, err := gorm.Open("postgres", fmt.Sprintf("user=postgres password=root dbname=%v sslmode=disable", *db_name))
	DB.CheckErr(err)

	var product products.Server = products.Server{db}
	var review reviews.Server = reviews.Server{db}
	r := mux.NewRouter() //router init

	r.HandleFunc("/products", product.GetProducts).Methods("GET")
	r.HandleFunc("/products/{productid}", product.GetProduct).Methods("GET")
	r.HandleFunc("/products/create", product.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{productid}/reviews", review.GetReviews).Methods("GET")
	r.HandleFunc("/products/{productid}/reviews/create", review.CreateReview).Methods("POST")

	fmt.Printf("starting port at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

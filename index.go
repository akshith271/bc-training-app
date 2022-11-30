package main

import (
	"fmt"
	"log"
	"net/http"
	products "sample/git/controllers"
	reviews "sample/git/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	r := mux.NewRouter() //router init

	r.HandleFunc("/products", products.GetProducts).Methods("GET")
	r.HandleFunc("/products/{productid}", products.GetProduct).Methods("GET")
	r.HandleFunc("/products/create", products.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{productid}/reviews", reviews.GetReviews).Methods("GET")
	r.HandleFunc("/products/{productid}/reviews/create", reviews.CreateReview).Methods("POST")

	fmt.Printf("starting port at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

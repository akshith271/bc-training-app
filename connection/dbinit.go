package db

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres dbname=sample password=root  sslmode=disable")
	f := flag.String("dbname", "sample", "created by akshith")
	CheckErr(err)
	// defer db.Close()
	// db.DropTableIfExists(&Product{})
	// db.CreateTable(&Product{})
	// db.DropTableIfExists(&Rating{})
	// db.CreateTable(&Rating{})
	flag.Parse()
	fmt.Printf(*f)
	// db.Save(&Product{
	// 	ProductName:        "Product 1",
	// 	ProductDescription: "This is product 1",
	// 	ProductCategory:    "Category 1",
	// 	ProductQuantity:    "Quantity 1",
	// 	ProductPrice:       "23423",
	// 	ProductImage:       "This is image 1",
	// 	ProductRating: &Rating{
	// 		ProductID: 1,
	// 		Name:      "Reviewer 1",
	// 		Review:    "This is review 1",
	// 	},
	// })
	// db.Save(&Product{
	// 	ProductName:        "Product 2",
	// 	ProductDescription: "This is product 2",
	// 	ProductCategory:    "Category 2",
	// 	ProductQuantity:    "Quantity 2",
	// 	ProductPrice:       "3434534",
	// 	ProductImage:       "This is image 2",
	// 	ProductRating: &Rating{
	// 		Name:   "Reviewer 2",
	// 		Review: "This is review 2",
	// 	},
	// })
	fmt.Printf("success")
	return db
}
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

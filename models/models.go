package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Product struct {
	gorm.Model
	ProductName        string  `json:"productname"`
	ProductDescription string  `json:"productdescription"`
	ProductCategory    string  `json:"productcategory"`
	ProductQuantity    string  `json:"productquantity"`
	ProductPrice       string  `json:"productprice"`
	ProductImage       string  `json:"productimage"`
	ProductRating      *Rating `json:"rating"`
}

type Rating struct {
	ProductID int    `json:"reviewProduct"`
	Name      string `json:"name"`
	Review    string `json:"review"`
	Rating    int    `json:"rating"`
}

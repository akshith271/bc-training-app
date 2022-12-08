package Products

import (
	// model "ecommerce/db_model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	mocks "sample/git/mocks"
	"sample/git/models"
	"testing"

	"github.com/golang/mock/gomock"
)

func (*Server) CheckErr(err error) {
	if err != nil {
		fmt.Printf("Thers an error")
	}
}
func TestGetProducts(t *testing.T) {

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockcntrl := gomock.NewController(t)
	defer mockcntrl.Finish()

	mockProd := mocks.NewMockDbProducts(mockcntrl)
	testProd := Server{Db: mockProd}

	product1 := models.Product{
		ProductName:        "Asus Zenbook 11",
		ProductDescription: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate",
		ProductCategory:    "Laptop",
		ProductQuantity:    "2",
		ProductPrice:       "88000",
		ProductImage:       "image.jpg",
		ProductRating: &models.Rating{
			Name:   "Akshith",
			Review: "This laptop is good",
			Rating: 5,
		},
	}
	mockProducts := []models.Product{product1}

	mockProd.EXPECT().GetProducts().Return(mockProducts, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testProd.GetProducts)
	handler.ServeHTTP(rr, req)

	// Checking status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checking body
	var got JsonResponse
	json.NewDecoder(rr.Body).Decode(&got)
	testProd.CheckErr(err)
	product2 := models.Product{
		ProductName:        "Asus Zenbook 11",
		ProductDescription: "This Laptop is with Intel i7 12th gen processor and it has 120hz High refresh rate",
		ProductCategory:    "Laptop",
		ProductQuantity:    "2",
		ProductPrice:       "88000",
		ProductImage:       "image.jpg",
		ProductRating: &models.Rating{
			Name:   "Akshith",
			Review: "This laptop is good",
			Rating: 5,
		},
	}

	mockProducts2 := []models.Product{product2}
	var mock = JsonResponse{Type: "success", Data: mockProducts2}

	if !reflect.DeepEqual(got, mock) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, mock)
	}
}

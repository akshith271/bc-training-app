package Products

import (
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

	request, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockcontroller := gomock.NewController(t)
	defer mockcontroller.Finish()

	mockProd := mocks.NewMockDbProducts(mockcontroller)
	testProd := Server{Db: mockProd}

	product1 := models.Product{
		ProductName:        "Asus Zenbook 11",
		ProductDescription: "This is a good laptop",
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
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(testProd.GetProducts)
	handler.ServeHTTP(recorder, request)

	// Checking status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checking body
	var got JsonResponse
	json.NewDecoder(recorder.Body).Decode(&got)
	testProd.CheckErr(err)
	product2 := models.Product{
		ProductName:        "Asus Zenbook 11",
		ProductDescription: "This is a good laptop",
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

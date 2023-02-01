package main_test

import (
	"Haioo_6/api/controller"
	"Haioo_6/api/dto"
	"Haioo_6/api/injection"
	"Haioo_6/api/service"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	db             *gorm.DB                  = injection.CreateDatabase()
	CartService    service.CartService       = service.NewCartService(db)
	CartController controller.CartController = controller.NewCartController(CartService)
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetCartWithoutFilter(t *testing.T) {
	r := SetUpRouter()
	r.POST("/cart/", CartController.GetCart)

	productName := ""

	params := dto.ListCartInput{Filter: &dto.FilterCartInput{
		NamaProduk: &productName,
		Kuantitas:  nil,
	}}

	jsonValue, _ := json.Marshal(&params)
	req, _ := http.NewRequest("POST", "/cart/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCartWithNameFilter(t *testing.T) {
	r := SetUpRouter()
	r.POST("/cart/", CartController.GetCart)

	productName := "Tahu Tok"

	params := dto.ListCartInput{Filter: &dto.FilterCartInput{
		NamaProduk: &productName,
		Kuantitas:  nil,
	}}

	jsonValue, _ := json.Marshal(&params)
	req, _ := http.NewRequest("POST", "/cart/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCartWithQuantityFilter(t *testing.T) {
	r := SetUpRouter()
	r.POST("/cart/", CartController.GetCart)

	productName := ""
	quantity := int64(12)

	params := dto.ListCartInput{Filter: &dto.FilterCartInput{
		NamaProduk: &productName,
		Kuantitas:  &quantity,
	}}

	jsonValue, _ := json.Marshal(&params)
	req, _ := http.NewRequest("POST", "/cart/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCartWithNameandQuantityFilter(t *testing.T) {
	r := SetUpRouter()
	r.POST("/cart/", CartController.GetCart)

	productName := "Tahu Tok"
	quantity := int64(12)

	params := dto.ListCartInput{Filter: &dto.FilterCartInput{
		NamaProduk: &productName,
		Kuantitas:  &quantity,
	}}

	jsonValue, _ := json.Marshal(&params)
	req, _ := http.NewRequest("POST", "/cart/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInsertCartWithNewProduct(t *testing.T) {
	r := SetUpRouter()
	r.POST("/cart/add", CartController.CreateCart)

	params := dto.CreateCart{
		KodeProduk: "12345",
		NamaProduk: "Kacang Hijau",
		Kuantitas:  30,
	}

	jsonValue, _ := json.Marshal(&params)
	req, _ := http.NewRequest("POST", "/cart/add", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInsertCartWithOldProduct(t *testing.T) {
	r := SetUpRouter()
	r.POST("/cart/add", CartController.CreateCart)

	params := dto.CreateCart{
		KodeProduk: "12345",
		NamaProduk: "Kacang Hijau",
		Kuantitas:  20,
	}

	jsonValue, _ := json.Marshal(&params)
	req, _ := http.NewRequest("POST", "/cart/add", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteProduct(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/cart/delete", CartController.DeleteCart)

	req, _ := http.NewRequest("DELETE", "/cart/delete", nil)
	q := req.URL.Query()
	q.Add("kodeProduk", "12345")

	req.URL.RawQuery = q.Encode()

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

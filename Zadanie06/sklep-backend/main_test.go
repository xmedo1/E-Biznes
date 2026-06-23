package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsHandler(t *testing.T) {
	// GET /products
	req := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	productsHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization", w.Header().Get("Access-Control-Allow-Headers"))

	var products []Product
	err := json.NewDecoder(w.Body).Decode(&products)
	assert.NoError(t, err)
	assert.Len(t, products, 2)

	assert.Equal(t, 1, products[0].ID)
	assert.Equal(t, "Jablko", products[0].Name)
	assert.Equal(t, 1.2, products[0].Price)

	assert.Equal(t, 2, products[1].ID)
	assert.Equal(t, "Banan", products[1].Name)
	assert.Equal(t, 1.5, products[1].Price)

	// OPTIONS /products
	reqOpt := httptest.NewRequest("OPTIONS", "/products", nil)
	wOpt := httptest.NewRecorder()
	productsHandler(wOpt, reqOpt)
	assert.Equal(t, http.StatusOK, wOpt.Code)
	assert.Equal(t, "*", wOpt.Header().Get("Access-Control-Allow-Origin"))
}

func TestPaymentsHandler(t *testing.T) {
	// POST /payments
	payload := []byte(`{"amount": 100.5}`)
	req := httptest.NewRequest("POST", "/payments", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	paymentsHandler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))

	// OPTIONS /payments
	reqOpt := httptest.NewRequest("OPTIONS", "/payments", nil)
	wOpt := httptest.NewRecorder()
	paymentsHandler(wOpt, reqOpt)
	assert.Equal(t, http.StatusOK, wOpt.Code)
	assert.Equal(t, "*", wOpt.Header().Get("Access-Control-Allow-Origin"))

	// GET /payments 
	reqGet := httptest.NewRequest("GET", "/payments", nil)
	wGet := httptest.NewRecorder()
	paymentsHandler(wGet, reqGet)
	assert.Equal(t, http.StatusOK, wGet.Code)
	assert.Equal(t, "*", wGet.Header().Get("Access-Control-Allow-Origin"))
}

func TestCartHandler(t *testing.T) {
	// POST /cart
	payload := []byte(`[{"id": 1, "name": "Test"}]`)
	req := httptest.NewRequest("POST", "/cart", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	cartHandler(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))

	// OPTIONS /cart
	reqOpt := httptest.NewRequest("OPTIONS", "/cart", nil)
	wOpt := httptest.NewRecorder()
	cartHandler(wOpt, reqOpt)
	assert.Equal(t, http.StatusOK, wOpt.Code)
	assert.Equal(t, "*", wOpt.Header().Get("Access-Control-Allow-Origin"))

	// GET /cart
	reqGet := httptest.NewRequest("GET", "/cart", nil)
	wGet := httptest.NewRecorder()
	cartHandler(wGet, reqGet)
	assert.Equal(t, http.StatusOK, wGet.Code)
	assert.Equal(t, "*", wGet.Header().Get("Access-Control-Allow-Origin"))
}

func TestEnableCors(t *testing.T) {
	var rw http.ResponseWriter = httptest.NewRecorder()
	enableCors(&rw)

	assert.Equal(t, "*", rw.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE", rw.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization", rw.Header().Get("Access-Control-Allow-Headers"))
	assert.NotEmpty(t, rw.Header().Get("Access-Control-Allow-Origin"))
	assert.NotEmpty(t, rw.Header().Get("Access-Control-Allow-Methods"))
	assert.NotEmpty(t, rw.Header().Get("Access-Control-Allow-Headers"))
}

func TestProductModelsAndDataValidations(t *testing.T) {
	p1 := Product{ID: 10, Name: "ProduktA", Price: 9.99}
	assert.Equal(t, 10, p1.ID)
	assert.Equal(t, "ProduktA", p1.Name)
	assert.Equal(t, 9.99, p1.Price)
	assert.NotNil(t, p1)
	assert.True(t, p1.ID > 0)
	assert.True(t, p1.Price > 0)

	p2 := Product{ID: 20, Name: "ProduktB", Price: 15.50}
	assert.Equal(t, 20, p2.ID)
	assert.Equal(t, "ProduktB", p2.Name)
	assert.Equal(t, 15.50, p2.Price)
	assert.NotNil(t, p2)
	assert.True(t, p2.ID > 0)
	assert.True(t, p2.Price > 0)

	p3 := Product{ID: 30, Name: "ProduktC", Price: 20.00}
	assert.Equal(t, 30, p3.ID)
	assert.Equal(t, "ProduktC", p3.Name)
	assert.Equal(t, 20.00, p3.Price)
	assert.NotNil(t, p3)
	assert.True(t, p3.ID > 0)
	assert.True(t, p3.Price > 0)
}
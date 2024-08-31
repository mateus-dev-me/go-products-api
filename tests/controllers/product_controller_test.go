package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-products/interfaces/api"
	"go-products/interfaces/api/controllers"
	"go-products/internal/infrastructure/db"
	"go-products/tests/helpers"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestController_Save(t *testing.T) {
	gin.SetMode(gin.TestMode)

	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)
	controller := controllers.NewProductController(repo)

	router := api.SetupRouter(controller)

	product := helpers.GenerateFakeProduct()
	payload, err := json.Marshal(product)

	require.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestController_GetById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)
	controller := controllers.NewProductController(repo)

	router := api.SetupRouter(controller)

	err := helpers.CreateProduct(connectionDB)

	require.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestController_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	connectionDB := helpers.SetupTestDB(t)
	defer connectionDB.Close()
	defer helpers.TeardownTestDB(t)

	repo := db.NewProductRepositoryDB(connectionDB)
	controller := controllers.NewProductController(repo)

	router := api.SetupRouter(controller)

	err := helpers.CreateProducts(connectionDB)

	require.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

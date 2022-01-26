package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/zenklot/catatan-dapur/controller"
	"github.com/zenklot/catatan-dapur/model/domain"
	"github.com/zenklot/catatan-dapur/repository"
	"github.com/zenklot/catatan-dapur/routes"
	"github.com/zenklot/catatan-dapur/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=catatan_dapur_test port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&domain.Bahan{}, &domain.ResepDetail{}, &domain.Resep{}, &domain.Kategori{})
	fmt.Println("Database Migrated")
	return DB
}

func setupRouter(db *gorm.DB) *httprouter.Router {
	validate := validator.New()
	bahanRepository := repository.NewBahanRepository()
	bahanService := service.NewBahanService(bahanRepository, db, validate)
	bahanController := controller.NewBahanController(bahanService)

	return routes.BahanRouter(bahanController)

}

func TestCreateBahanSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"bahan" : "Terigu"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/bahan", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "CREATED", responseBody["status"])
}

func TestCreateBahanFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"bahan" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/bahan", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestDeleteBahanSuccess(t *testing.T) {
	db := setupTestDB()

	tx := db.Begin()
	bahanRepository := repository.NewBahanRepository()
	bahan := bahanRepository.Save(tx, domain.Bahan{
		Bahan: "Terigu",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/bahan/"+strconv.Itoa(bahan.Id), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

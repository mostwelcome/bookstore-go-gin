package main

import (
	"bookstore-go-gin/controllers"
	"bookstore-go-gin/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomepageHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to the Bookstore listing API with Golang"}`
	r := SetUpRouter()
	r.GET("/", controllers.HomePageHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNewBookHandler(t *testing.T) {
	models.ConnectDatabase()
	defer models.CloseDatabase()

	r := SetUpRouter()
	r.POST("/books", controllers.CreateBook)
	book := models.Book{
		Title:  "The Infinite Game",
		Author: "Simon Sinek",
	}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	models.ClearTable()
}

func TestGetBooksHandler(t *testing.T) {
	a := assert.New(t)
	models.ConnectDatabase()
	defer models.CloseDatabase()

	_, err := models.InsertTestBook()
	if err != nil {
		a.Error(err)
	}

	r := SetUpRouter()
	r.GET("/books", controllers.FindBooks)
	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var book []models.Book
	json.Unmarshal(w.Body.Bytes(), &book)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, book)
	models.ClearTable()
}

func TestUpdateBookHandler(t *testing.T) {
	a := assert.New(t)
	models.ConnectDatabase()
	defer models.CloseDatabase()

	_, err := models.InsertTestBook()
	if err != nil {
		a.Error(err)
	}

	r := SetUpRouter()
	r.PATCH("/books/:id", controllers.UpdateBook)
	book := models.Book{
		ID:    1,
		Title: "The Infinite Game",
	}
	jsonValue, _ := json.Marshal(book)
	var id string = strconv.FormatUint(uint64(book.ID), 10)
	reqFound, _ := http.NewRequest("PATCH", "/books/"+id, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PATCH", "/books/12", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	models.ClearTable()
}

func TestDeleteBookHandler(t *testing.T) {
	a := assert.New(t)
	models.ConnectDatabase()
	defer models.CloseDatabase()

	_, err := models.InsertTestBook()
	if err != nil {
		a.Error(err)
	}

	r := SetUpRouter()
	r.DELETE("/books/:id", controllers.UpdateBook)
	book := models.Book{
		ID:    1,
		Title: "The Infinite Game",
	}
	jsonValue, _ := json.Marshal(book)
	var id string = strconv.FormatUint(uint64(book.ID), 10)
	reqFound, _ := http.NewRequest("DELETE", "/books/"+id, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("DELETE", "/books/12", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	models.ClearTable()
}

func TestGetBookByIDHandler(t *testing.T) {
	a := assert.New(t)
	models.ConnectDatabase()
	defer models.CloseDatabase()

	_, err := models.InsertTestBook()
	if err != nil {
		a.Error(err)
	}

	r := SetUpRouter()
	r.GET("/books/:id", controllers.FindBook)
	req, _ := http.NewRequest("GET", "/books/"+"1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var book models.Book
	json.Unmarshal(w.Body.Bytes(), &book)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, book)
	models.ClearTable()
}

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
	models.ConnectTestDatabase()
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
}

func TestUpdateBookHandler(t *testing.T) {
	models.ConnectTestDatabase()
	defer models.CloseDatabase()

	bookToInsert, _ := models.InsertTestBook()

	r := SetUpRouter()
	r.PATCH("/books/:id", controllers.UpdateBook)
	book := models.Book{
		ID:    bookToInsert.ID,
		Title: "The Infinite Game Updated",
	}
	jsonValue, _ := json.Marshal(book)
	var id string = strconv.FormatUint(uint64(book.ID), 10)
	reqFound, _ := http.NewRequest("PATCH", "/books/"+id, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	// Make a GET request to verify the data is updated
	r.GET("/books/:id", controllers.FindBook)
	req, _ := http.NewRequest("GET", "/books/"+id, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var updatedBook models.Book
	err := json.Unmarshal(w.Body.Bytes(), &updatedBook)
	if err != nil {
		return
	}
	assert.Equal(t, "The Infinite Game Updated", updatedBook.Title)
}

func TestDeleteBookHandler(t *testing.T) {
	models.ConnectTestDatabase()
	defer models.CloseDatabase()

	bookToInsert, _ := models.InsertTestBook()

	r := SetUpRouter()
	r.DELETE("/books/:id", controllers.DeleteBook)
	var id string = strconv.FormatUint(uint64(bookToInsert.ID), 10)
	reqFound, _ := http.NewRequest("DELETE", "/books/"+id, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	// Make a GET request to verify the data is deleted
	r.GET("/books/:id", controllers.FindBook)
	req, _ := http.NewRequest("GET", "/books/"+id, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

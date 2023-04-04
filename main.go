package main

import (
	"bookstore-go-gin/controllers"
	"bookstore-go-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)         // get list of books
	r.POST("/books", controllers.CreateBook)       // add new book
	r.GET("/books/:id", controllers.FindBook)      // get book by id
	r.PATCH("/books/:id", controllers.UpdateBook)  //update book by id
	r.DELETE("/books/:id", controllers.DeleteBook) // delete book by id
	r.Run()
}

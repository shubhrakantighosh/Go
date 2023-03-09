package main

import (
	_ "BookAPI/docs"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const portNumber = ":8080"

type book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{Id: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of books.
// @Tags			/books
// @Success			200 {books} response.Response{}
// @Router			/books [get]
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook book
	err := c.BindJSON(&newBook)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func getBookById(id string) (*book, error) {
	for i, j := range books {
		if j.Id == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func searchBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func updateBook(c *gin.Context) (*book, error) {
	id, ok := c.GetQuery("id")
	if ok == false {
		return nil, nil
	}
	book, err := getBookById(id)

	if err != nil {
		return nil, nil
	}

	if book.Quantity <= 0 {
		return nil, nil
	}

	return book, nil
}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagsRequest true  "Update tags"
// @Tags			/books
// @Success			200 {object} response.Response{}
// @Router			/books/{id} [patch]
func checkOutBook(c *gin.Context) {
	//id, ok := c.GetQuery("id")
	//if ok == false {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
	//	return
	//}
	//book, err := getBookById(id)
	//
	//if err != nil {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
	//	return
	//}
	//
	//if book.Quantity <= 0 {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not available"})
	//	return
	//}
	book, err := updateBook(c)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
	}

	book.Quantity -= 1

	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	//id, ok := c.GetQuery("id")
	//if ok == false {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
	//	return
	//}
	//book, err := getBookById(id)
	//
	//if err != nil {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
	//	return
	//}
	//
	//if book.Quantity <= 0 {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not available"})
	//	return
	//}
	book, err := updateBook(c)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found"})
	}

	book.Quantity += 1

	c.IndentedJSON(http.StatusOK, book)
}

// @title 	book Service API
// @version	1.0
// @description A book service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath /
func main() {

	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.GET("/books/:id", searchBookById)
	router.PATCH("/books/checkout", checkOutBook)
	router.PATCH("/books/return", returnBook)
	router.Run(portNumber)

}

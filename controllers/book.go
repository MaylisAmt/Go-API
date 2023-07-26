package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/golang/models"
)

var Library []models.Book
var Counter int

func InitDataBase() {
	Counter = 1

	book1 := models.Book{
		Id:     1,
		Title:  "Jusqu'à ce que la mort nous sépare",
		Author: "Lisa Gardner",
	}
	Library = append(Library, book1)
}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Library})
}

func CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Counter++

	book := models.Book{Id: Counter, Title: input.Title, Author: input.Author}
	Library = append(Library, book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func removeItem(book models.Book, bookSlice []models.Book) []models.Book {
	for idx, v := range bookSlice {
		if v == book {
			return append(bookSlice[0:idx], bookSlice[idx+1:]...)
		}
	}
	return bookSlice
}

func DeleteBook(c *gin.Context) {
	bookFound := false

	var bookFind models.Book
	for _, book := range Library {
		if c.Param("id") == strconv.Itoa(book.Id) {
			bookFound = true
			bookFind = book
		}
	}

	if !bookFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	Library = removeItem(bookFind, Library)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

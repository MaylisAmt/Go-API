package main

import (
	"github.com/gin-gonic/gin"
	"main.go/golang/controllers"
)

func main() {
	controllers.InitDataBase()
	r := gin.Default()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}

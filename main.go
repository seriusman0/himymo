package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:tmDevFlats0987^(@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	router := gin.Default()

	if err != nil {
		log.Fatal("DB Connection Error")
		log.Fatal("Oke")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	bookRequest := book.BookRequest{
		Title: "Gundam",
		Price: "200000",
	}

	bookService.Create(bookRequest)

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8888")

	//main
	//handler
	//service
	//repository
	//db
	//mysql

}

//  02:53:24 - --:--:-- Perbaiki service.go untuk memetakan book

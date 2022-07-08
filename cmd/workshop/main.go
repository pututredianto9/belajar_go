package main

import (
	"belajar/book"
	"belajar/book/handler"
	"belajar/package/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadConfig("./package/config")
	if err != nil {
		log.Fatal("Can't load config: ", err)
	}
	router := gin.Default()
	v1 := router.Group("/v1")

	db, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}

	fmt.Println("DB connection succeeded")

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandle := handler.NewBookHandler(bookService)

	v1.POST("/book", bookHandle.CreateBooks)
	v1.GET("/books", bookHandle.GetBooks)
	v1.GET("/book/:id", bookHandle.GetByIdBook)
	v1.PUT("/book/:id", bookHandle.UpdateBook)
	v1.DELETE("/book/:id", bookHandle.DeleteBook)

	router.Run(config.ServerAddress)
}

// main
// handler
// service
// repository
// db
// mysql

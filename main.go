package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"himymo/config"
	"himymo/handler"
	"himymo/middleware"
	"himymo/repository"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db, err := config.DBConn()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := &repository.Repository{DB: db}
	if err := repo.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	h := handler.NewHandler(db)

	router := gin.Default()

	router.Use(middleware.AuthMiddleware)

	v1 := router.Group("/v1")
	{
		v1.POST("/transactions", h.CreateTransaction)
		v1.GET("/transactions/:id", h.GetTransactionById)
		v1.GET("/transactions", h.GetTransactions)
		v1.PUT("/transactions/:id", h.UpdateTransaction)
		v1.DELETE("/transactions/:id", h.DeleteTransactionById)
	}

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("failed to start server: %v\n", err)
	}
}

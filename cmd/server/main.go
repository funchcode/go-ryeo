package main

import (
	"Goryeo/internal/receipt"
	"Goryeo/internal/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := repository.NewDB()

	router := gin.Default()
	v1 := router.Group("/v1")

	receipt.RegisterHandlers(v1, receipt.NewService(db))
	router.Run()
}

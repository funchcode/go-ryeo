package main

import (
	"Goryeo/internal/receipt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	receipt.RegisterHandlers(v1)
	router.Run()
}
